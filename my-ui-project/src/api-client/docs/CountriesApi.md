# CountriesApi

All URIs are relative to *https://restcountries.com/v3.1*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**getAllCountries**](#getallcountries) | **GET** /all | GET endpoint for retrieving all countries|
|[**getCountryByCapital**](#getcountrybycapital) | **GET** /capital/{capital} | Search by capital city name|
|[**getCountryByCode**](#getcountrybycode) | **GET** /alpha/{code} | Search by country code. It can be the 2-letter or 3-letter code|
|[**getCountryByCurrency**](#getcountrybycurrency) | **GET** /currency/{currency} | Search by currency code|
|[**getCountryByLanguage**](#getcountrybylanguage) | **GET** /lang/{language} | Search by language code|
|[**getCountryByName**](#getcountrybyname) | **GET** /name/{name} | Search by country name. If you want to get an exact match, use the next endpoint. It can be the common or official value|
|[**getCountryByRegion**](#getcountrybyregion) | **GET** /region/{region} | Search by region name|
|[**getCountryBySubregion**](#getcountrybysubregion) | **GET** /subregion/{subregion} | Search by subregion name|
|[**getCountryByTranslation**](#getcountrybytranslation) | **GET** /translation/{translation} | Search by translation name|
|[**getIndependentCountries**](#getindependentcountries) | **GET** /independent | Get independent countries|

# **getAllCountries**
> Array<AllresponsebodyInner> getAllCountries()

Returns a simple example response.

### Example

```typescript
import {
    CountriesApi,
    Configuration
} from './api';

const configuration = new Configuration();
const apiInstance = new CountriesApi(configuration);

let fields: Array<string>; //Comma-separated list of fields to include in the response for filtering (default to undefined)

const { status, data } = await apiInstance.getAllCountries(
    fields
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **fields** | **Array&lt;string&gt;** | Comma-separated list of fields to include in the response for filtering | defaults to undefined|


### Return type

**Array<AllresponsebodyInner>**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Successful response |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **getCountryByCapital**
> Array<AllresponsebodyInner> getCountryByCapital()


### Example

```typescript
import {
    CountriesApi,
    Configuration
} from './api';

const configuration = new Configuration();
const apiInstance = new CountriesApi(configuration);

let capital: string; //The capital city name to search for (default to undefined)
let fields: Array<string>; //Comma-separated list of fields to include in the response for filtering (optional) (default to undefined)

const { status, data } = await apiInstance.getCountryByCapital(
    capital,
    fields
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **capital** | [**string**] | The capital city name to search for | defaults to undefined|
| **fields** | **Array&lt;string&gt;** | Comma-separated list of fields to include in the response for filtering | (optional) defaults to undefined|


### Return type

**Array<AllresponsebodyInner>**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Successful response with country details |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **getCountryByCode**
> Array<AllresponsebodyInner> getCountryByCode()


### Example

```typescript
import {
    CountriesApi,
    Configuration
} from './api';

const configuration = new Configuration();
const apiInstance = new CountriesApi(configuration);

let code: string; //The country code to search ,for example for moldova 2-letter code i.e cca2  = MD, 3-letter code i.e cca3 and cioc= MDA, numeric code  i.e ccn3=498 (default to undefined)
let fields: Array<string>; //Comma-separated list of fields to include in the response for filtering (optional) (default to undefined)

const { status, data } = await apiInstance.getCountryByCode(
    code,
    fields
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **code** | [**string**] | The country code to search ,for example for moldova 2-letter code i.e cca2  &#x3D; MD, 3-letter code i.e cca3 and cioc&#x3D; MDA, numeric code  i.e ccn3&#x3D;498 | defaults to undefined|
| **fields** | **Array&lt;string&gt;** | Comma-separated list of fields to include in the response for filtering | (optional) defaults to undefined|


### Return type

**Array<AllresponsebodyInner>**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Successful response with country details |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **getCountryByCurrency**
> Array<AllresponsebodyInner> getCountryByCurrency()


### Example

```typescript
import {
    CountriesApi,
    Configuration
} from './api';

const configuration = new Configuration();
const apiInstance = new CountriesApi(configuration);

let currency: string; //The currency code to search for (default to undefined)
let fields: Array<string>; //Comma-separated list of fields to include in the response for filtering (optional) (default to undefined)

const { status, data } = await apiInstance.getCountryByCurrency(
    currency,
    fields
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **currency** | [**string**] | The currency code to search for | defaults to undefined|
| **fields** | **Array&lt;string&gt;** | Comma-separated list of fields to include in the response for filtering | (optional) defaults to undefined|


### Return type

**Array<AllresponsebodyInner>**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Successful response with country details |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **getCountryByLanguage**
> Array<AllresponsebodyInner> getCountryByLanguage()


### Example

```typescript
import {
    CountriesApi,
    Configuration
} from './api';

const configuration = new Configuration();
const apiInstance = new CountriesApi(configuration);

let language: string; //The language code to search for (default to undefined)
let fields: Array<string>; //Comma-separated list of fields to include in the response for filtering (optional) (default to undefined)

const { status, data } = await apiInstance.getCountryByLanguage(
    language,
    fields
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **language** | [**string**] | The language code to search for | defaults to undefined|
| **fields** | **Array&lt;string&gt;** | Comma-separated list of fields to include in the response for filtering | (optional) defaults to undefined|


### Return type

**Array<AllresponsebodyInner>**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Successful response with country details |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **getCountryByName**
> Array<AllresponsebodyInner> getCountryByName()


### Example

```typescript
import {
    CountriesApi,
    Configuration
} from './api';

const configuration = new Configuration();
const apiInstance = new CountriesApi(configuration);

let name: string; //The name of the country to search for (default to undefined)
let fields: Array<string>; //Comma-separated list of fields to include in the response for filtering (optional) (default to undefined)
let fullText: boolean; //Whether to perform a full-text search (optional) (default to undefined)

const { status, data } = await apiInstance.getCountryByName(
    name,
    fields,
    fullText
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **name** | [**string**] | The name of the country to search for | defaults to undefined|
| **fields** | **Array&lt;string&gt;** | Comma-separated list of fields to include in the response for filtering | (optional) defaults to undefined|
| **fullText** | [**boolean**] | Whether to perform a full-text search | (optional) defaults to undefined|


### Return type

**Array<AllresponsebodyInner>**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Successful response with country details |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **getCountryByRegion**
> Array<AllresponsebodyInner> getCountryByRegion()


### Example

```typescript
import {
    CountriesApi,
    Configuration
} from './api';

const configuration = new Configuration();
const apiInstance = new CountriesApi(configuration);

let region: string; //The region name to search for (default to undefined)
let fields: Array<string>; //Comma-separated list of fields to include in the response for filtering (optional) (default to undefined)

const { status, data } = await apiInstance.getCountryByRegion(
    region,
    fields
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **region** | [**string**] | The region name to search for | defaults to undefined|
| **fields** | **Array&lt;string&gt;** | Comma-separated list of fields to include in the response for filtering | (optional) defaults to undefined|


### Return type

**Array<AllresponsebodyInner>**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Successful response with country details |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **getCountryBySubregion**
> Array<AllresponsebodyInner> getCountryBySubregion()


### Example

```typescript
import {
    CountriesApi,
    Configuration
} from './api';

const configuration = new Configuration();
const apiInstance = new CountriesApi(configuration);

let subregion: string; //The subregion name to search for (default to undefined)
let fields: Array<string>; //Comma-separated list of fields to include in the response for filtering (optional) (default to undefined)

const { status, data } = await apiInstance.getCountryBySubregion(
    subregion,
    fields
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **subregion** | [**string**] | The subregion name to search for | defaults to undefined|
| **fields** | **Array&lt;string&gt;** | Comma-separated list of fields to include in the response for filtering | (optional) defaults to undefined|


### Return type

**Array<AllresponsebodyInner>**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Successful response with country details |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **getCountryByTranslation**
> Array<AllresponsebodyInner> getCountryByTranslation()


### Example

```typescript
import {
    CountriesApi,
    Configuration
} from './api';

const configuration = new Configuration();
const apiInstance = new CountriesApi(configuration);

let translation: string; //The translation name to search for (default to undefined)
let fields: Array<string>; //Comma-separated list of fields to include in the response for filtering (optional) (default to undefined)

const { status, data } = await apiInstance.getCountryByTranslation(
    translation,
    fields
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **translation** | [**string**] | The translation name to search for | defaults to undefined|
| **fields** | **Array&lt;string&gt;** | Comma-separated list of fields to include in the response for filtering | (optional) defaults to undefined|


### Return type

**Array<AllresponsebodyInner>**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Successful response with country details |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **getIndependentCountries**
> Array<AllresponsebodyInner> getIndependentCountries()


### Example

```typescript
import {
    CountriesApi,
    Configuration
} from './api';

const configuration = new Configuration();
const apiInstance = new CountriesApi(configuration);

let status: boolean; //Status of the countries to filter by (default to undefined)
let fields: Array<string>; //Comma-separated list of fields to include in the response for filtering (optional) (default to undefined)

const { status, data } = await apiInstance.getIndependentCountries(
    status,
    fields
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **status** | [**boolean**] | Status of the countries to filter by | defaults to undefined|
| **fields** | **Array&lt;string&gt;** | Comma-separated list of fields to include in the response for filtering | (optional) defaults to undefined|


### Return type

**Array<AllresponsebodyInner>**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | Successful response |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

