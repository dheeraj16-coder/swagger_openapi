package openapi

import (
	"context"
	"fmt"
	"net/http"
	restcountries "restcountries"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

var countryCache = NewCache(1 * time.Hour)

type CountriesAPI struct {
	client *restcountries.APIClient
}

func newClient() *restcountries.APIClient {
	return restcountries.NewAPIClient(restcountries.NewConfiguration())
}

func handleExternalError(c *gin.Context, err error, resource string) {
	errStr := err.Error()
	switch {
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

func (api *CountriesAPI) GetAllCountries(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}

func (api *CountriesAPI) GetCountryByCapital(c *gin.Context) {
	capital := c.Param("capital")
	fields := c.Query("fields")
	cacheKey := fmt.Sprintf("capital:%s:%s", capital, fields)
	if cached, ok := countryCache.Get(cacheKey); ok {
		c.JSON(http.StatusOK, cached)
		return
	}
	client := newClient()
	helper := client.CountriesAPI.GetCountryByCapital(context.Background(), capital)
	if fields != "" {
		helper = helper.Fields(strings.Split(fields, ","))
	}
	countries, _, err := helper.Execute()
	if err != nil {
		handleExternalError(c, err, fmt.Sprintf("capital '%s'", capital))
		return
	}
	if len(countries) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("No countries found for capital '%s'", capital)})
		return
	}
	countryCache.Set(cacheKey, countries)
	c.JSON(http.StatusOK, countries)
}

func (api *CountriesAPI) GetCountryByCode(c *gin.Context) {
	code := c.Param("code")
	fields := c.Query("fields")
	cacheKey := fmt.Sprintf("code:%s:%s", code, fields)
	if cached, ok := countryCache.Get(cacheKey); ok {
		c.JSON(http.StatusOK, cached)
		return
	}
	client := newClient()
	helper := client.CountriesAPI.GetCountryByCode(context.Background(), code)
	if fields != "" {
		helper = helper.Fields(strings.Split(fields, ","))
	}
	countries, _, err := helper.Execute()
	if err != nil {
		handleExternalError(c, err, fmt.Sprintf("code '%s'", code))
		return
	}
	if len(countries) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("No countries found for code '%s'", code)})
		return
	}
	countryCache.Set(cacheKey, countries)
	c.JSON(http.StatusOK, countries)
}

func (api *CountriesAPI) GetCountryByCurrency(c *gin.Context) {
	currency := c.Param("currency")
	fields := c.Query("fields")
	cacheKey := fmt.Sprintf("currency:%s:%s", currency, fields)
	if cached, ok := countryCache.Get(cacheKey); ok {
		c.JSON(http.StatusOK, cached)
		return
	}
	client := newClient()
	helper := client.CountriesAPI.GetCountryByCurrency(context.Background(), currency)
	if fields != "" {
		helper = helper.Fields(strings.Split(fields, ","))
	}
	countries, _, err := helper.Execute()
	if err != nil {
		handleExternalError(c, err, fmt.Sprintf("currency '%s'", currency))
		return
	}
	if len(countries) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("No countries found for currency '%s'", currency)})
		return
	}
	countryCache.Set(cacheKey, countries)
	c.JSON(http.StatusOK, countries)
}

func (api *CountriesAPI) GetCountryByLanguage(c *gin.Context) {
	language := c.Param("language")
	fields := c.Query("fields")
	cacheKey := fmt.Sprintf("lang:%s:%s", language, fields)
	if cached, ok := countryCache.Get(cacheKey); ok {
		c.JSON(http.StatusOK, cached)
		return
	}
	client := newClient()
	helper := client.CountriesAPI.GetCountryByLanguage(context.Background(), language)
	if fields != "" {
		helper = helper.Fields(strings.Split(fields, ","))
	}
	countries, _, err := helper.Execute()
	if err != nil {
		handleExternalError(c, err, fmt.Sprintf("language '%s'", language))
		return
	}
	if len(countries) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("No countries found for language '%s'", language)})
		return
	}
	countryCache.Set(cacheKey, countries)
	c.JSON(http.StatusOK, countries)
}

func (api *CountriesAPI) GetCountryByName(c *gin.Context) {
	name := c.Param("name")
	fields := c.Query("fields")
	cacheKey := fmt.Sprintf("name:%s:%s", name, fields)
	if cached, ok := countryCache.Get(cacheKey); ok {
		c.JSON(http.StatusOK, cached)
		return
	}
	client := newClient()
	helper := client.CountriesAPI.GetCountryByName(context.Background(), name)
	if fields != "" {
		helper = helper.Fields(strings.Split(fields, ","))
	}
	countries, _, err := helper.Execute()
	if err != nil {
		handleExternalError(c, err, fmt.Sprintf("name '%s'", name))
		return
	}
	if len(countries) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("No countries found for name '%s'", name)})
		return
	}
	countryCache.Set(cacheKey, countries)
	c.JSON(http.StatusOK, countries)
}

func (api *CountriesAPI) GetCountryByRegion(c *gin.Context) {
	region := c.Param("region")
	fields := c.Query("fields")
	cacheKey := fmt.Sprintf("region:%s:%s", region, fields)
	if cached, ok := countryCache.Get(cacheKey); ok {
		c.JSON(http.StatusOK, cached)
		return
	}
	client := newClient()
	helper := client.CountriesAPI.GetCountryByRegion(context.Background(), region)
	if fields != "" {
		helper = helper.Fields(strings.Split(fields, ","))
	}
	countries, _, err := helper.Execute()
	if err != nil {
		handleExternalError(c, err, fmt.Sprintf("region '%s'", region))
		return
	}
	if len(countries) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("No countries found for region '%s'", region)})
		return
	}
	countryCache.Set(cacheKey, countries)
	c.JSON(http.StatusOK, countries)
}

func (api *CountriesAPI) GetCountryBySubregion(c *gin.Context) {
	subregion := c.Param("subregion")
	fields := c.Query("fields")
	cacheKey := fmt.Sprintf("subregion:%s:%s", subregion, fields)
	if cached, ok := countryCache.Get(cacheKey); ok {
		c.JSON(http.StatusOK, cached)
		return
	}
	client := newClient()
	helper := client.CountriesAPI.GetCountryBySubregion(context.Background(), subregion)
	if fields != "" {
		helper = helper.Fields(strings.Split(fields, ","))
	}
	countries, _, err := helper.Execute()
	if err != nil {
		handleExternalError(c, err, fmt.Sprintf("subregion '%s'", subregion))
		return
	}
	if len(countries) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("No countries found for subregion '%s'", subregion)})
		return
	}
	countryCache.Set(cacheKey, countries)
	c.JSON(http.StatusOK, countries)
}

func (api *CountriesAPI) GetCountryByTranslation(c *gin.Context) {
	translation := c.Param("translation")
	fields := c.Query("fields")
	cacheKey := fmt.Sprintf("translation:%s:%s", translation, fields)
	if cached, ok := countryCache.Get(cacheKey); ok {
		c.JSON(http.StatusOK, cached)
		return
	}
	client := newClient()
	helper := client.CountriesAPI.GetCountryByTranslation(context.Background(), translation)
	if fields != "" {
		helper = helper.Fields(strings.Split(fields, ","))
	}
	countries, _, err := helper.Execute()
	if err != nil {
		handleExternalError(c, err, fmt.Sprintf("translation '%s'", translation))
		return
	}
	if len(countries) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("No countries found for translation '%s'", translation)})
		return
	}
	countryCache.Set(cacheKey, countries)
	c.JSON(http.StatusOK, countries)
}

func (api *CountriesAPI) GetIndependentCountries(c *gin.Context) {
	status := c.Query("status")
	fields := c.Query("fields")
	cacheKey := fmt.Sprintf("independent:%s:%s", status, fields)
	if cached, ok := countryCache.Get(cacheKey); ok {
		c.JSON(http.StatusOK, cached)
		return
	}
	client := newClient()
	helper := client.CountriesAPI.GetIndependentCountries(context.Background())
	helper = helper.Status(status == "true")
	if fields != "" {
		helper = helper.Fields(strings.Split(fields, ","))
	}
	countries, _, err := helper.Execute()
	if err != nil {
		handleExternalError(c, err, "independent countries")
		return
	}
	if len(countries) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No independent countries found"})
		return
	}
	countryCache.Set(cacheKey, countries)
	c.JSON(http.StatusOK, countries)
}

func HealthCheck(c *gin.Context) {
	client := http.Client{
		Timeout: 2 * time.Second,
	}
	resp, err := client.Get("https://restcountries.com/v3.1/alpha/us")
	if err != nil || resp.StatusCode != http.StatusOK {
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
