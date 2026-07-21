# Capital

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Coordinates** | Pointer to [**Coordinates**](Coordinates.md) |  | [optional] 
**Attributes** | Pointer to [**CapitalAttributes**](CapitalAttributes.md) |  | [optional] 

## Methods

### NewCapital

`func NewCapital() *Capital`

NewCapital instantiates a new Capital object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapitalWithDefaults

`func NewCapitalWithDefaults() *Capital`

NewCapitalWithDefaults instantiates a new Capital object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Capital) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Capital) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Capital) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Capital) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCoordinates

`func (o *Capital) GetCoordinates() Coordinates`

GetCoordinates returns the Coordinates field if non-nil, zero value otherwise.

### GetCoordinatesOk

`func (o *Capital) GetCoordinatesOk() (*Coordinates, bool)`

GetCoordinatesOk returns a tuple with the Coordinates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoordinates

`func (o *Capital) SetCoordinates(v Coordinates)`

SetCoordinates sets Coordinates field to given value.

### HasCoordinates

`func (o *Capital) HasCoordinates() bool`

HasCoordinates returns a boolean if a field has been set.

### GetAttributes

`func (o *Capital) GetAttributes() CapitalAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *Capital) GetAttributesOk() (*CapitalAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *Capital) SetAttributes(v CapitalAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *Capital) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


