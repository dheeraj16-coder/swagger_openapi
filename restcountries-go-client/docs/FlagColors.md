# FlagColors

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dominant** | Pointer to **string** |  | [optional] 
**Prominent** | Pointer to **string** |  | [optional] 
**Palette** | Pointer to [**[]FlagColorsPaletteInner**](FlagColorsPaletteInner.md) |  | [optional] 
**Swatches** | Pointer to [**FlagColorsSwatches**](FlagColorsSwatches.md) |  | [optional] 

## Methods

### NewFlagColors

`func NewFlagColors() *FlagColors`

NewFlagColors instantiates a new FlagColors object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlagColorsWithDefaults

`func NewFlagColorsWithDefaults() *FlagColors`

NewFlagColorsWithDefaults instantiates a new FlagColors object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDominant

`func (o *FlagColors) GetDominant() string`

GetDominant returns the Dominant field if non-nil, zero value otherwise.

### GetDominantOk

`func (o *FlagColors) GetDominantOk() (*string, bool)`

GetDominantOk returns a tuple with the Dominant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDominant

`func (o *FlagColors) SetDominant(v string)`

SetDominant sets Dominant field to given value.

### HasDominant

`func (o *FlagColors) HasDominant() bool`

HasDominant returns a boolean if a field has been set.

### GetProminent

`func (o *FlagColors) GetProminent() string`

GetProminent returns the Prominent field if non-nil, zero value otherwise.

### GetProminentOk

`func (o *FlagColors) GetProminentOk() (*string, bool)`

GetProminentOk returns a tuple with the Prominent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProminent

`func (o *FlagColors) SetProminent(v string)`

SetProminent sets Prominent field to given value.

### HasProminent

`func (o *FlagColors) HasProminent() bool`

HasProminent returns a boolean if a field has been set.

### GetPalette

`func (o *FlagColors) GetPalette() []FlagColorsPaletteInner`

GetPalette returns the Palette field if non-nil, zero value otherwise.

### GetPaletteOk

`func (o *FlagColors) GetPaletteOk() (*[]FlagColorsPaletteInner, bool)`

GetPaletteOk returns a tuple with the Palette field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPalette

`func (o *FlagColors) SetPalette(v []FlagColorsPaletteInner)`

SetPalette sets Palette field to given value.

### HasPalette

`func (o *FlagColors) HasPalette() bool`

HasPalette returns a boolean if a field has been set.

### GetSwatches

`func (o *FlagColors) GetSwatches() FlagColorsSwatches`

GetSwatches returns the Swatches field if non-nil, zero value otherwise.

### GetSwatchesOk

`func (o *FlagColors) GetSwatchesOk() (*FlagColorsSwatches, bool)`

GetSwatchesOk returns a tuple with the Swatches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwatches

`func (o *FlagColors) SetSwatches(v FlagColorsSwatches)`

SetSwatches sets Swatches field to given value.

### HasSwatches

`func (o *FlagColors) HasSwatches() bool`

HasSwatches returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


