# \CountriesAPI

All URIs are relative to *https://restcountries.com/v3.1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAllCountries**](CountriesAPI.md#GetAllCountries) | **Get** /all | GET endpoint for retrieving all countries
[**GetCountryByCapital**](CountriesAPI.md#GetCountryByCapital) | **Get** /capital/{capital} | Search by capital city name
[**GetCountryByCode**](CountriesAPI.md#GetCountryByCode) | **Get** /alpha/{code} | Search by country code. It can be the 2-letter or 3-letter code
[**GetCountryByCurrency**](CountriesAPI.md#GetCountryByCurrency) | **Get** /currency/{currency} | Search by currency code
[**GetCountryByLanguage**](CountriesAPI.md#GetCountryByLanguage) | **Get** /lang/{language} | Search by language code
[**GetCountryByName**](CountriesAPI.md#GetCountryByName) | **Get** /name/{name} | Search by country name. If you want to get an exact match, use the next endpoint. It can be the common or official value
[**GetCountryByRegion**](CountriesAPI.md#GetCountryByRegion) | **Get** /region/{region} | Search by region name
[**GetCountryBySubregion**](CountriesAPI.md#GetCountryBySubregion) | **Get** /subregion/{subregion} | Search by subregion name
[**GetCountryByTranslation**](CountriesAPI.md#GetCountryByTranslation) | **Get** /translation/{translation} | Search by translation name
[**GetIndependentCountries**](CountriesAPI.md#GetIndependentCountries) | **Get** /independent | Get independent countries



## GetAllCountries

> []AllresponsebodyInner GetAllCountries(ctx).Fields(fields).Execute()

GET endpoint for retrieving all countries



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	fields := []string{"name"} // []string | Comma-separated list of fields to include in the response for filtering

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CountriesAPI.GetAllCountries(context.Background()).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CountriesAPI.GetAllCountries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllCountries`: []AllresponsebodyInner
	fmt.Fprintf(os.Stdout, "Response from `CountriesAPI.GetAllCountries`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllCountriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **[]string** | Comma-separated list of fields to include in the response for filtering | 

### Return type

[**[]AllresponsebodyInner**](AllresponsebodyInner.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCountryByCapital

> []AllresponsebodyInner GetCountryByCapital(ctx, capital).Fields(fields).Execute()

Search by capital city name

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	capital := "Washington, D.C." // string | The capital city name to search for
	fields := []string{"name,capital,area,population,flags"} // []string | Comma-separated list of fields to include in the response for filtering (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CountriesAPI.GetCountryByCapital(context.Background(), capital).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CountriesAPI.GetCountryByCapital``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCountryByCapital`: []AllresponsebodyInner
	fmt.Fprintf(os.Stdout, "Response from `CountriesAPI.GetCountryByCapital`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**capital** | **string** | The capital city name to search for | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCountryByCapitalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | Comma-separated list of fields to include in the response for filtering | 

### Return type

[**[]AllresponsebodyInner**](AllresponsebodyInner.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCountryByCode

> []AllresponsebodyInner GetCountryByCode(ctx, code).Fields(fields).Execute()

Search by country code. It can be the 2-letter or 3-letter code

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	code := "code_example" // string | The country code to search ,for example for moldova 2-letter code i.e cca2  = MD, 3-letter code i.e cca3 and cioc= MDA, numeric code  i.e ccn3=498
	fields := []string{"name,capital,area,population,flags"} // []string | Comma-separated list of fields to include in the response for filtering (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CountriesAPI.GetCountryByCode(context.Background(), code).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CountriesAPI.GetCountryByCode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCountryByCode`: []AllresponsebodyInner
	fmt.Fprintf(os.Stdout, "Response from `CountriesAPI.GetCountryByCode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string** | The country code to search ,for example for moldova 2-letter code i.e cca2  &#x3D; MD, 3-letter code i.e cca3 and cioc&#x3D; MDA, numeric code  i.e ccn3&#x3D;498 | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCountryByCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | Comma-separated list of fields to include in the response for filtering | 

### Return type

[**[]AllresponsebodyInner**](AllresponsebodyInner.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCountryByCurrency

> []AllresponsebodyInner GetCountryByCurrency(ctx, currency).Fields(fields).Execute()

Search by currency code

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	currency := "USD" // string | The currency code to search for
	fields := []string{"name,capital,area,population,flags"} // []string | Comma-separated list of fields to include in the response for filtering (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CountriesAPI.GetCountryByCurrency(context.Background(), currency).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CountriesAPI.GetCountryByCurrency``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCountryByCurrency`: []AllresponsebodyInner
	fmt.Fprintf(os.Stdout, "Response from `CountriesAPI.GetCountryByCurrency`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**currency** | **string** | The currency code to search for | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCountryByCurrencyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | Comma-separated list of fields to include in the response for filtering | 

### Return type

[**[]AllresponsebodyInner**](AllresponsebodyInner.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCountryByLanguage

> []AllresponsebodyInner GetCountryByLanguage(ctx, language).Fields(fields).Execute()

Search by language code

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	language := "en" // string | The language code to search for
	fields := []string{"name,capital,area,population,flags"} // []string | Comma-separated list of fields to include in the response for filtering (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CountriesAPI.GetCountryByLanguage(context.Background(), language).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CountriesAPI.GetCountryByLanguage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCountryByLanguage`: []AllresponsebodyInner
	fmt.Fprintf(os.Stdout, "Response from `CountriesAPI.GetCountryByLanguage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**language** | **string** | The language code to search for | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCountryByLanguageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | Comma-separated list of fields to include in the response for filtering | 

### Return type

[**[]AllresponsebodyInner**](AllresponsebodyInner.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCountryByName

> []AllresponsebodyInner GetCountryByName(ctx, name).Fields(fields).FullText(fullText).Execute()

Search by country name. If you want to get an exact match, use the next endpoint. It can be the common or official value

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	name := "United States of America" // string | The name of the country to search for
	fields := []string{"name,capital,area,population,flags"} // []string | Comma-separated list of fields to include in the response for filtering (optional)
	fullText := true // bool | Whether to perform a full-text search (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CountriesAPI.GetCountryByName(context.Background(), name).Fields(fields).FullText(fullText).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CountriesAPI.GetCountryByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCountryByName`: []AllresponsebodyInner
	fmt.Fprintf(os.Stdout, "Response from `CountriesAPI.GetCountryByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The name of the country to search for | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCountryByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | Comma-separated list of fields to include in the response for filtering | 
 **fullText** | **bool** | Whether to perform a full-text search | 

### Return type

[**[]AllresponsebodyInner**](AllresponsebodyInner.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCountryByRegion

> []AllresponsebodyInner GetCountryByRegion(ctx, region).Fields(fields).Execute()

Search by region name

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	region := "Americas" // string | The region name to search for
	fields := []string{"name,capital,area,population,flags"} // []string | Comma-separated list of fields to include in the response for filtering (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CountriesAPI.GetCountryByRegion(context.Background(), region).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CountriesAPI.GetCountryByRegion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCountryByRegion`: []AllresponsebodyInner
	fmt.Fprintf(os.Stdout, "Response from `CountriesAPI.GetCountryByRegion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**region** | **string** | The region name to search for | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCountryByRegionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | Comma-separated list of fields to include in the response for filtering | 

### Return type

[**[]AllresponsebodyInner**](AllresponsebodyInner.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCountryBySubregion

> []AllresponsebodyInner GetCountryBySubregion(ctx, subregion).Fields(fields).Execute()

Search by subregion name

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	subregion := "Northern America" // string | The subregion name to search for
	fields := []string{"name,capital,area,population,flags"} // []string | Comma-separated list of fields to include in the response for filtering (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CountriesAPI.GetCountryBySubregion(context.Background(), subregion).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CountriesAPI.GetCountryBySubregion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCountryBySubregion`: []AllresponsebodyInner
	fmt.Fprintf(os.Stdout, "Response from `CountriesAPI.GetCountryBySubregion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subregion** | **string** | The subregion name to search for | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCountryBySubregionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | Comma-separated list of fields to include in the response for filtering | 

### Return type

[**[]AllresponsebodyInner**](AllresponsebodyInner.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCountryByTranslation

> []AllresponsebodyInner GetCountryByTranslation(ctx, translation).Fields(fields).Execute()

Search by translation name

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	translation := "Estados Unidos" // string | The translation name to search for
	fields := []string{"name,capital,area,population,flags"} // []string | Comma-separated list of fields to include in the response for filtering (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CountriesAPI.GetCountryByTranslation(context.Background(), translation).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CountriesAPI.GetCountryByTranslation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCountryByTranslation`: []AllresponsebodyInner
	fmt.Fprintf(os.Stdout, "Response from `CountriesAPI.GetCountryByTranslation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**translation** | **string** | The translation name to search for | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCountryByTranslationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | Comma-separated list of fields to include in the response for filtering | 

### Return type

[**[]AllresponsebodyInner**](AllresponsebodyInner.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIndependentCountries

> []AllresponsebodyInner GetIndependentCountries(ctx).Status(status).Fields(fields).Execute()

Get independent countries

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	status := false // bool | Status of the countries to filter by
	fields := []string{"name"} // []string | Comma-separated list of fields to include in the response for filtering (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CountriesAPI.GetIndependentCountries(context.Background()).Status(status).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CountriesAPI.GetIndependentCountries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIndependentCountries`: []AllresponsebodyInner
	fmt.Fprintf(os.Stdout, "Response from `CountriesAPI.GetIndependentCountries`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIndependentCountriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | **bool** | Status of the countries to filter by | 
 **fields** | **[]string** | Comma-separated list of fields to include in the response for filtering | 

### Return type

[**[]AllresponsebodyInner**](AllresponsebodyInner.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

