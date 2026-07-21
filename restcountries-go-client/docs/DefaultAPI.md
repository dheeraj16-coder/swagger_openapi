# \DefaultAPI

All URIs are relative to *https://api.restcountries.com/countries/v5*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCountriesByPropertyValue**](DefaultAPI.md#GetCountriesByPropertyValue) | **Get** /{property}/{value} | Read countries by exact property value
[**ListCountries**](DefaultAPI.md#ListCountries) | **Get** / | List all countries (paginated), optionally narrowed by free-text search and property filters
[**SearchCountriesByProperty**](DefaultAPI.md#SearchCountriesByProperty) | **Get** /{property} | Substring search within a single searchable property or a named aggregate



## GetCountriesByPropertyValue

> CountryListResponse GetCountriesByPropertyValue(ctx, property, value).Q(q).Limit(limit).Offset(offset).ResponseFields(responseFields).ResponseFieldsOmit(responseFieldsOmit).Pretty(pretty).Execute()

Read countries by exact property value



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
	property := "codes.alpha_3" // string | A readable property dot-path.
	value := "CAN" // string | Exact value to match (case-insensitive).
	q := "q_example" // string | Free-text substring search across searchable properties. Explicitly empty `q` returns 400. (optional)
	limit := int32(56) // int32 | Page size. 1–100 on the free plan (up to 500 on paid plans). Defaults to 25. (optional) (default to 25)
	offset := int32(56) // int32 | Records to skip before returning. Out-of-bounds offsets return an empty objects array with status 200. (optional) (default to 0)
	responseFields := "names.common,codes.alpha_2,flag.emoji" // string | Comma-separated allowlist of dot-path properties to include in each returned object; when set, objects carry only those keys.  (optional)
	responseFieldsOmit := "names.translations" // string | Comma-separated blocklist of dot-path properties to drop from each returned object. Composes with response_fields; omit wins on conflict.  (optional)
	pretty := true // bool | Pretty-print the JSON response. Bare flag or truthy values enable. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetCountriesByPropertyValue(context.Background(), property, value).Q(q).Limit(limit).Offset(offset).ResponseFields(responseFields).ResponseFieldsOmit(responseFieldsOmit).Pretty(pretty).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetCountriesByPropertyValue``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCountriesByPropertyValue`: CountryListResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetCountriesByPropertyValue`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**property** | **string** | A readable property dot-path. | 
**value** | **string** | Exact value to match (case-insensitive). | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCountriesByPropertyValueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **q** | **string** | Free-text substring search across searchable properties. Explicitly empty &#x60;q&#x60; returns 400. | 
 **limit** | **int32** | Page size. 1–100 on the free plan (up to 500 on paid plans). Defaults to 25. | [default to 25]
 **offset** | **int32** | Records to skip before returning. Out-of-bounds offsets return an empty objects array with status 200. | [default to 0]
 **responseFields** | **string** | Comma-separated allowlist of dot-path properties to include in each returned object; when set, objects carry only those keys.  | 
 **responseFieldsOmit** | **string** | Comma-separated blocklist of dot-path properties to drop from each returned object. Composes with response_fields; omit wins on conflict.  | 
 **pretty** | **bool** | Pretty-print the JSON response. Bare flag or truthy values enable. | 

### Return type

[**CountryListResponse**](CountryListResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCountries

> CountryListResponse ListCountries(ctx).Q(q).Limit(limit).Offset(offset).ResponseFields(responseFields).ResponseFieldsOmit(responseFieldsOmit).Pretty(pretty).Execute()

List all countries (paginated), optionally narrowed by free-text search and property filters



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
	q := "q_example" // string | Free-text substring search across searchable properties. Explicitly empty `q` returns 400. (optional)
	limit := int32(56) // int32 | Page size. 1–100 on the free plan (up to 500 on paid plans). Defaults to 25. (optional) (default to 25)
	offset := int32(56) // int32 | Records to skip before returning. Out-of-bounds offsets return an empty objects array with status 200. (optional) (default to 0)
	responseFields := "names.common,codes.alpha_2,flag.emoji" // string | Comma-separated allowlist of dot-path properties to include in each returned object; when set, objects carry only those keys.  (optional)
	responseFieldsOmit := "names.translations" // string | Comma-separated blocklist of dot-path properties to drop from each returned object. Composes with response_fields; omit wins on conflict.  (optional)
	pretty := true // bool | Pretty-print the JSON response. Bare flag or truthy values enable. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ListCountries(context.Background()).Q(q).Limit(limit).Offset(offset).ResponseFields(responseFields).ResponseFieldsOmit(responseFieldsOmit).Pretty(pretty).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ListCountries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCountries`: CountryListResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ListCountries`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCountriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string** | Free-text substring search across searchable properties. Explicitly empty &#x60;q&#x60; returns 400. | 
 **limit** | **int32** | Page size. 1–100 on the free plan (up to 500 on paid plans). Defaults to 25. | [default to 25]
 **offset** | **int32** | Records to skip before returning. Out-of-bounds offsets return an empty objects array with status 200. | [default to 0]
 **responseFields** | **string** | Comma-separated allowlist of dot-path properties to include in each returned object; when set, objects carry only those keys.  | 
 **responseFieldsOmit** | **string** | Comma-separated blocklist of dot-path properties to drop from each returned object. Composes with response_fields; omit wins on conflict.  | 
 **pretty** | **bool** | Pretty-print the JSON response. Bare flag or truthy values enable. | 

### Return type

[**CountryListResponse**](CountryListResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchCountriesByProperty

> CountryListResponse SearchCountriesByProperty(ctx, property).Q(q).Limit(limit).Offset(offset).ResponseFields(responseFields).ResponseFieldsOmit(responseFieldsOmit).Pretty(pretty).Execute()

Substring search within a single searchable property or a named aggregate



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
	property := "names.common" // string | A searchable property dot-path (e.g. `names.common`, `capitals`, `currencies`) or an aggregate name (`name`, `code`). 
	q := "q_example" // string | Search term (substring, case-insensitive). Required; empty returns 400.
	limit := int32(56) // int32 | Page size. 1–100 on the free plan (up to 500 on paid plans). Defaults to 25. (optional) (default to 25)
	offset := int32(56) // int32 | Records to skip before returning. Out-of-bounds offsets return an empty objects array with status 200. (optional) (default to 0)
	responseFields := "names.common,codes.alpha_2,flag.emoji" // string | Comma-separated allowlist of dot-path properties to include in each returned object; when set, objects carry only those keys.  (optional)
	responseFieldsOmit := "names.translations" // string | Comma-separated blocklist of dot-path properties to drop from each returned object. Composes with response_fields; omit wins on conflict.  (optional)
	pretty := true // bool | Pretty-print the JSON response. Bare flag or truthy values enable. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.SearchCountriesByProperty(context.Background(), property).Q(q).Limit(limit).Offset(offset).ResponseFields(responseFields).ResponseFieldsOmit(responseFieldsOmit).Pretty(pretty).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.SearchCountriesByProperty``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchCountriesByProperty`: CountryListResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.SearchCountriesByProperty`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**property** | **string** | A searchable property dot-path (e.g. &#x60;names.common&#x60;, &#x60;capitals&#x60;, &#x60;currencies&#x60;) or an aggregate name (&#x60;name&#x60;, &#x60;code&#x60;).  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchCountriesByPropertyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **q** | **string** | Search term (substring, case-insensitive). Required; empty returns 400. | 
 **limit** | **int32** | Page size. 1–100 on the free plan (up to 500 on paid plans). Defaults to 25. | [default to 25]
 **offset** | **int32** | Records to skip before returning. Out-of-bounds offsets return an empty objects array with status 200. | [default to 0]
 **responseFields** | **string** | Comma-separated allowlist of dot-path properties to include in each returned object; when set, objects carry only those keys.  | 
 **responseFieldsOmit** | **string** | Comma-separated blocklist of dot-path properties to drop from each returned object. Composes with response_fields; omit wins on conflict.  | 
 **pretty** | **bool** | Pretty-print the JSON response. Bare flag or truthy values enable. | 

### Return type

[**CountryListResponse**](CountryListResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

