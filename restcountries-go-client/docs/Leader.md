# Leader

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Assets** | Pointer to **map[string]interface{}** |  | [optional] 
**Attributes** | Pointer to **map[string]interface{}** |  | [optional] 
**Links** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewLeader

`func NewLeader() *Leader`

NewLeader instantiates a new Leader object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLeaderWithDefaults

`func NewLeaderWithDefaults() *Leader`

NewLeaderWithDefaults instantiates a new Leader object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Leader) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Leader) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Leader) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Leader) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTitle

`func (o *Leader) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Leader) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Leader) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Leader) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetAssets

`func (o *Leader) GetAssets() map[string]interface{}`

GetAssets returns the Assets field if non-nil, zero value otherwise.

### GetAssetsOk

`func (o *Leader) GetAssetsOk() (*map[string]interface{}, bool)`

GetAssetsOk returns a tuple with the Assets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssets

`func (o *Leader) SetAssets(v map[string]interface{})`

SetAssets sets Assets field to given value.

### HasAssets

`func (o *Leader) HasAssets() bool`

HasAssets returns a boolean if a field has been set.

### GetAttributes

`func (o *Leader) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *Leader) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *Leader) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *Leader) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetLinks

`func (o *Leader) GetLinks() map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Leader) GetLinksOk() (*map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Leader) SetLinks(v map[string]interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *Leader) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


