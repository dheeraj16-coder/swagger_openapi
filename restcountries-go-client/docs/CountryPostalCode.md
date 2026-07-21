# CountryPostalCode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Format** | Pointer to **string** | Human-readable mask, &#x60;#&#x60; for digits (e.g. \&quot;A#A | [optional] 
**Regex** | Pointer to **string** |  | [optional] 

## Methods

### NewCountryPostalCode

`func NewCountryPostalCode() *CountryPostalCode`

NewCountryPostalCode instantiates a new CountryPostalCode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCountryPostalCodeWithDefaults

`func NewCountryPostalCodeWithDefaults() *CountryPostalCode`

NewCountryPostalCodeWithDefaults instantiates a new CountryPostalCode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFormat

`func (o *CountryPostalCode) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *CountryPostalCode) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *CountryPostalCode) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *CountryPostalCode) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetRegex

`func (o *CountryPostalCode) GetRegex() string`

GetRegex returns the Regex field if non-nil, zero value otherwise.

### GetRegexOk

`func (o *CountryPostalCode) GetRegexOk() (*string, bool)`

GetRegexOk returns a tuple with the Regex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegex

`func (o *CountryPostalCode) SetRegex(v string)`

SetRegex sets Regex field to given value.

### HasRegex

`func (o *CountryPostalCode) HasRegex() bool`

HasRegex returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


