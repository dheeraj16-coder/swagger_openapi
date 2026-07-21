# Names

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Common** | Pointer to **string** |  | [optional] 
**Official** | Pointer to **string** |  | [optional] 
**Alternates** | Pointer to **[]string** |  | [optional] 
**Native** | Pointer to [**map[string]NameForms**](NameForms.md) | Keyed by ISO 639-3 language code. | [optional] 
**Translations** | Pointer to [**map[string]NameForms**](NameForms.md) | Keyed by ISO 639-3 language code. Heavy branch; consider response_fields_omit&#x3D;names.translations. | [optional] 

## Methods

### NewNames

`func NewNames() *Names`

NewNames instantiates a new Names object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNamesWithDefaults

`func NewNamesWithDefaults() *Names`

NewNamesWithDefaults instantiates a new Names object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommon

`func (o *Names) GetCommon() string`

GetCommon returns the Common field if non-nil, zero value otherwise.

### GetCommonOk

`func (o *Names) GetCommonOk() (*string, bool)`

GetCommonOk returns a tuple with the Common field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommon

`func (o *Names) SetCommon(v string)`

SetCommon sets Common field to given value.

### HasCommon

`func (o *Names) HasCommon() bool`

HasCommon returns a boolean if a field has been set.

### GetOfficial

`func (o *Names) GetOfficial() string`

GetOfficial returns the Official field if non-nil, zero value otherwise.

### GetOfficialOk

`func (o *Names) GetOfficialOk() (*string, bool)`

GetOfficialOk returns a tuple with the Official field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfficial

`func (o *Names) SetOfficial(v string)`

SetOfficial sets Official field to given value.

### HasOfficial

`func (o *Names) HasOfficial() bool`

HasOfficial returns a boolean if a field has been set.

### GetAlternates

`func (o *Names) GetAlternates() []string`

GetAlternates returns the Alternates field if non-nil, zero value otherwise.

### GetAlternatesOk

`func (o *Names) GetAlternatesOk() (*[]string, bool)`

GetAlternatesOk returns a tuple with the Alternates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternates

`func (o *Names) SetAlternates(v []string)`

SetAlternates sets Alternates field to given value.

### HasAlternates

`func (o *Names) HasAlternates() bool`

HasAlternates returns a boolean if a field has been set.

### GetNative

`func (o *Names) GetNative() map[string]NameForms`

GetNative returns the Native field if non-nil, zero value otherwise.

### GetNativeOk

`func (o *Names) GetNativeOk() (*map[string]NameForms, bool)`

GetNativeOk returns a tuple with the Native field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNative

`func (o *Names) SetNative(v map[string]NameForms)`

SetNative sets Native field to given value.

### HasNative

`func (o *Names) HasNative() bool`

HasNative returns a boolean if a field has been set.

### GetTranslations

`func (o *Names) GetTranslations() map[string]NameForms`

GetTranslations returns the Translations field if non-nil, zero value otherwise.

### GetTranslationsOk

`func (o *Names) GetTranslationsOk() (*map[string]NameForms, bool)`

GetTranslationsOk returns a tuple with the Translations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslations

`func (o *Names) SetTranslations(v map[string]NameForms)`

SetTranslations sets Translations field to given value.

### HasTranslations

`func (o *Names) HasTranslations() bool`

HasTranslations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


