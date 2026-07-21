# CountryListResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Objects** | Pointer to [**[]Country**](Country.md) |  | [optional] 
**Meta** | Pointer to [**ResponseMeta**](ResponseMeta.md) |  | [optional] 

## Methods

### NewCountryListResponseData

`func NewCountryListResponseData() *CountryListResponseData`

NewCountryListResponseData instantiates a new CountryListResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCountryListResponseDataWithDefaults

`func NewCountryListResponseDataWithDefaults() *CountryListResponseData`

NewCountryListResponseDataWithDefaults instantiates a new CountryListResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjects

`func (o *CountryListResponseData) GetObjects() []Country`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *CountryListResponseData) GetObjectsOk() (*[]Country, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *CountryListResponseData) SetObjects(v []Country)`

SetObjects sets Objects field to given value.

### HasObjects

`func (o *CountryListResponseData) HasObjects() bool`

HasObjects returns a boolean if a field has been set.

### GetMeta

`func (o *CountryListResponseData) GetMeta() ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *CountryListResponseData) GetMetaOk() (*ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *CountryListResponseData) SetMeta(v ResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *CountryListResponseData) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


