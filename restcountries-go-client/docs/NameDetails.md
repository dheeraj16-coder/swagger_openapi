# NameDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Common** | Pointer to **string** |  | [optional] 
**Official** | Pointer to **string** |  | [optional] 
**NativeName** | Pointer to [**map[string]NativeNameDetail**](NativeNameDetail.md) |  | [optional] 

## Methods

### NewNameDetails

`func NewNameDetails() *NameDetails`

NewNameDetails instantiates a new NameDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNameDetailsWithDefaults

`func NewNameDetailsWithDefaults() *NameDetails`

NewNameDetailsWithDefaults instantiates a new NameDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommon

`func (o *NameDetails) GetCommon() string`

GetCommon returns the Common field if non-nil, zero value otherwise.

### GetCommonOk

`func (o *NameDetails) GetCommonOk() (*string, bool)`

GetCommonOk returns a tuple with the Common field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommon

`func (o *NameDetails) SetCommon(v string)`

SetCommon sets Common field to given value.

### HasCommon

`func (o *NameDetails) HasCommon() bool`

HasCommon returns a boolean if a field has been set.

### GetOfficial

`func (o *NameDetails) GetOfficial() string`

GetOfficial returns the Official field if non-nil, zero value otherwise.

### GetOfficialOk

`func (o *NameDetails) GetOfficialOk() (*string, bool)`

GetOfficialOk returns a tuple with the Official field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfficial

`func (o *NameDetails) SetOfficial(v string)`

SetOfficial sets Official field to given value.

### HasOfficial

`func (o *NameDetails) HasOfficial() bool`

HasOfficial returns a boolean if a field has been set.

### GetNativeName

`func (o *NameDetails) GetNativeName() map[string]NativeNameDetail`

GetNativeName returns the NativeName field if non-nil, zero value otherwise.

### GetNativeNameOk

`func (o *NameDetails) GetNativeNameOk() (*map[string]NativeNameDetail, bool)`

GetNativeNameOk returns a tuple with the NativeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNativeName

`func (o *NameDetails) SetNativeName(v map[string]NativeNameDetail)`

SetNativeName sets NativeName field to given value.

### HasNativeName

`func (o *NameDetails) HasNativeName() bool`

HasNativeName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


