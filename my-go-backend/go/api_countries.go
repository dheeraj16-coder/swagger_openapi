package openapi

import (
	"fmt"
	"net/http"
	"os"
	restcountries "restcountries"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

var countryCache = NewCache(1 * time.Hour)

type CountriesAPI struct {
	client *restcountries.APIClient
}

// NewCountriesAPI builds the handler with ONE reusable v5 client.
// The API key is read once at startup and baked into the client as a
// default Authorization header, so every request reuses the same client
// and no per-request auth work is needed.
func NewCountriesAPI() CountriesAPI {
	cfg := restcountries.NewConfiguration()
	cfg.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("RESTCOUNTRIES_API_KEY"))
	return CountriesAPI{
		client: restcountries.NewAPIClient(cfg),
	}
}

func handleExternalError(c *gin.Context, err error, resource string) {
	errStr := err.Error()
	switch {
	case strings.Contains(errStr, "401"):
		c.JSON(http.StatusBadGateway, gin.H{"error": "upstream authentication failed; check RESTCOUNTRIES_API_KEY"})
	case strings.Contains(errStr, "403"):
		c.JSON(http.StatusBadGateway, gin.H{"error": "upstream quota exceeded or premium field requested"})
	case strings.Contains(errStr, "404"):
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("No countries found for %s", resource)})
	case strings.Contains(errStr, "429"):
		c.JSON(http.StatusTooManyRequests, gin.H{"error": "upstream rate limit reached, please try again later"})
	case strings.Contains(errStr, "503"), strings.Contains(errStr, "connection refused"):
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": "countries service is temporarily unavailable"})
	default:
		c.JSON(http.StatusBadGateway, gin.H{"error": fmt.Sprintf("failed to fetch country data: %s", errStr)})
	}
}

// writeCountries unwraps the v5 envelope (data.objects) and returns the
// bare Country array — the backend's public contract.
func writeCountries(c *gin.Context, cacheKey, resource string, resp *restcountries.CountryListResponse, err error) {
	if err != nil {
		handleExternalError(c, err, resource)
		return
	}
	countries := resp.GetData().Objects
	if len(countries) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("No countries found for %s", resource)})
		return
	}
	countryCache.Set(cacheKey, countries)
	c.JSON(http.StatusOK, countries)
}

func (api CountriesAPI) GetAllCountries(c *gin.Context) {
	fields := c.Query("fields")
	cacheKey := fmt.Sprintf("all:%s", fields)
	if cached, ok := countryCache.Get(cacheKey); ok {
		c.JSON(http.StatusOK, cached)
		return
	}
	req := api.client.DefaultAPI.ListCountries(c.Request.Context()).Limit(100)
	if fields != "" {
		req = req.ResponseFields(fields)
	}
	resp, _, err := req.Execute()
	writeCountries(c, cacheKey, "all countries", resp, err)
}

func (api CountriesAPI) GetCountryByCapital(c *gin.Context) {
	capital := c.Param("capital")
	fields := c.Query("fields")
	cacheKey := fmt.Sprintf("capital:%s:%s", capital, fields)
	if cached, ok := countryCache.Get(cacheKey); ok {
		c.JSON(http.StatusOK, cached)
		return
	}
	req := api.client.DefaultAPI.SearchCountriesByProperty(c.Request.Context(), "capitals").Q(capital)
	if fields != "" {
		req = req.ResponseFields(fields)
	}
	resp, _, err := req.Execute()
	writeCountries(c, cacheKey, fmt.Sprintf("capital '%s'", capital), resp, err)
}

func (api CountriesAPI) GetCountryByCode(c *gin.Context) {
	code := c.Param("code")
	fields := c.Query("fields")
	cacheKey := fmt.Sprintf("code:%s:%s", code, fields)
	if cached, ok := countryCache.Get(cacheKey); ok {
		c.JSON(http.StatusOK, cached)
		return
	}
	// "code" is a v5 aggregate: fans out across alpha_2, alpha_3, ccn3,
	// fips, gec, fifa, cioc — mirrors the old /alpha behavior.
	req := api.client.DefaultAPI.SearchCountriesByProperty(c.Request.Context(), "code").Q(code)
	if fields != "" {
		req = req.ResponseFields(fields)
	}
	resp, _, err := req.Execute()
	writeCountries(c, cacheKey, fmt.Sprintf("code '%s'", code), resp, err)
}

func (api CountriesAPI) GetCountryByCurrency(c *gin.Context) {
	currency := c.Param("currency")
	fields := c.Query("fields")
	cacheKey := fmt.Sprintf("currency:%s:%s", currency, fields)
	if cached, ok := countryCache.Get(cacheKey); ok {
		c.JSON(http.StatusOK, cached)
		return
	}
	// Exact reverse lookup: /currencies/{code} (e.g. USD) is documented v5 behavior.
	req := api.client.DefaultAPI.GetCountriesByPropertyValue(c.Request.Context(), "currencies", currency)
	if fields != "" {
		req = req.ResponseFields(fields)
	}
	resp, _, err := req.Execute()
	writeCountries(c, cacheKey, fmt.Sprintf("currency '%s'", currency), resp, err)
}

func (api CountriesAPI) GetCountryByLanguage(c *gin.Context) {
	language := c.Param("language")
	fields := c.Query("fields")
	cacheKey := fmt.Sprintf("lang:%s:%s", language, fields)
	if cached, ok := countryCache.Get(cacheKey); ok {
		c.JSON(http.StatusOK, cached)
		return
	}
	// Substring search on the language name ("English", "Spanish"),
	// matching what the UI placeholder suggests users type.
	req := api.client.DefaultAPI.SearchCountriesByProperty(c.Request.Context(), "languages").Q(language)
	if fields != "" {
		req = req.ResponseFields(fields)
	}
	resp, _, err := req.Execute()
	writeCountries(c, cacheKey, fmt.Sprintf("language '%s'", language), resp, err)
}

func (api CountriesAPI) GetCountryByName(c *gin.Context) {
	name := c.Param("name")
	fields := c.Query("fields")
	cacheKey := fmt.Sprintf("name:%s:%s", name, fields)
	if cached, ok := countryCache.Get(cacheKey); ok {
		c.JSON(http.StatusOK, cached)
		return
	}
	// "name" aggregate: common, official, alternates, native — substring,
	// mirroring the old fuzzy /name behavior.
	req := api.client.DefaultAPI.SearchCountriesByProperty(c.Request.Context(), "name").Q(name)
	if fields != "" {
		req = req.ResponseFields(fields)
	}
	resp, _, err := req.Execute()
	writeCountries(c, cacheKey, fmt.Sprintf("name '%s'", name), resp, err)
}

func (api CountriesAPI) GetCountryByRegion(c *gin.Context) {
	region := c.Param("region")
	fields := c.Query("fields")
	cacheKey := fmt.Sprintf("region:%s:%s", region, fields)
	if cached, ok := countryCache.Get(cacheKey); ok {
		c.JSON(http.StatusOK, cached)
		return
	}
	req := api.client.DefaultAPI.GetCountriesByPropertyValue(c.Request.Context(), "region", region)
	if fields != "" {
		req = req.ResponseFields(fields)
	}
	resp, _, err := req.Execute()
	writeCountries(c, cacheKey, fmt.Sprintf("region '%s'", region), resp, err)
}

func (api CountriesAPI) GetCountryBySubregion(c *gin.Context) {
	subregion := c.Param("subregion")
	fields := c.Query("fields")
	cacheKey := fmt.Sprintf("subregion:%s:%s", subregion, fields)
	if cached, ok := countryCache.Get(cacheKey); ok {
		c.JSON(http.StatusOK, cached)
		return
	}
	req := api.client.DefaultAPI.GetCountriesByPropertyValue(c.Request.Context(), "subregion", subregion)
	if fields != "" {
		req = req.ResponseFields(fields)
	}
	resp, _, err := req.Execute()
	writeCountries(c, cacheKey, fmt.Sprintf("subregion '%s'", subregion), resp, err)
}

func (api CountriesAPI) GetCountryByTranslation(c *gin.Context) {
	translation := c.Param("translation")
	fields := c.Query("fields")
	cacheKey := fmt.Sprintf("translation:%s:%s", translation, fields)
	if cached, ok := countryCache.Get(cacheKey); ok {
		c.JSON(http.StatusOK, cached)
		return
	}
	req := api.client.DefaultAPI.SearchCountriesByProperty(c.Request.Context(), "names.translations").Q(translation)
	if fields != "" {
		req = req.ResponseFields(fields)
	}
	resp, _, err := req.Execute()
	writeCountries(c, cacheKey, fmt.Sprintf("translation '%s'", translation), resp, err)
}

func (api CountriesAPI) GetIndependentCountries(c *gin.Context) {
	status := c.Query("status")
	fields := c.Query("fields")
	cacheKey := fmt.Sprintf("independent:%s:%s", status, fields)
	if cached, ok := countryCache.Get(cacheKey); ok {
		c.JSON(http.StatusOK, cached)
		return
	}
	// v3.1's "independent" maps to v5's classification.sovereign,
	// read as a boolean property: /classification.sovereign/true
	sovereign := "true"
	if status == "false" {
		sovereign = "false"
	}
	req := api.client.DefaultAPI.GetCountriesByPropertyValue(c.Request.Context(), "classification.sovereign", sovereign)
	if fields != "" {
		req = req.ResponseFields(fields)
	}
	resp, _, err := req.Execute()
	writeCountries(c, cacheKey, "independent countries", resp, err)
}

func HealthCheck(c *gin.Context) {
	client := http.Client{
		Timeout: 2 * time.Second,
	}
	// Unauthenticated probe: a 401 from the v5 API means it is reachable
	// and responding. This burns zero quota (the free tier is 500 req/month,
	// so the health check must never spend an authenticated request).
	resp, err := client.Get("https://api.restcountries.com/countries/v5")
	if err != nil || (resp.StatusCode != http.StatusUnauthorized && resp.StatusCode != http.StatusOK) {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"status":    "unhealthy",
			"reason":    "external_api_unreachable",
			"timestamp": time.Now().UTC().Format(time.RFC3339),
		})
		return
	}

	adotResp, err := client.Get("http://localhost:13133/")
	if err != nil || adotResp.StatusCode != http.StatusOK {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"status":    "unhealthy",
			"reason":    "monitoring_agent_offline",
			"timestamp": time.Now().UTC().Format(time.RFC3339),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":    "healthy",
		"timestamp": time.Now().UTC().Format(time.RFC3339),
	})
}
