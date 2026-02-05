# AllresponsebodyInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to [**NameDetails**](NameDetails.md) |  | [optional] 
**Tld** | Pointer to **[]string** |  | [optional] 
**Cca2** | Pointer to **string** |  | [optional] 
**Cca3** | Pointer to **string** |  | [optional] 
**Ccn3** | Pointer to **string** |  | [optional] 
**Cioc** | Pointer to **string** |  | [optional] 
**Currencies** | Pointer to [**map[string]CurrenciesDetails**](CurrenciesDetails.md) |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**Independent** | Pointer to **bool** |  | [optional] 
**UnMember** | Pointer to **bool** |  | [optional] 
**Capital** | Pointer to **[]string** |  | [optional] 
**Latlng** | Pointer to **[]float32** |  | [optional] 
**Area** | Pointer to **float32** |  | [optional] 
**Population** | Pointer to **float32** |  | [optional] 
**Flags** | Pointer to [**AllresponsebodyInnerFlags**](AllresponsebodyInnerFlags.md) |  | [optional] 
**Flag** | Pointer to **string** |  | [optional] 
**Landlocked** | Pointer to **bool** |  | [optional] 
**Borders** | Pointer to **[]string** |  | [optional] 
**Translations** | Pointer to [**map[string]NativeNameDetail**](NativeNameDetail.md) |  | [optional] 
**Demonyms** | Pointer to [**map[string]DemonymsDetails**](DemonymsDetails.md) |  | [optional] 
**Timezones** | Pointer to **[]string** |  | [optional] 
**Continents** | Pointer to **[]string** |  | [optional] 
**Fifa** | Pointer to **string** |  | [optional] 
**Car** | Pointer to [**CarDetails**](CarDetails.md) |  | [optional] 
**CoatOfArms** | Pointer to [**ImageDetails**](ImageDetails.md) |  | [optional] 
**StartOfWeek** | Pointer to **string** |  | [optional] 
**CapitalInfo** | Pointer to [**AllresponsebodyInnerCapitalInfo**](AllresponsebodyInnerCapitalInfo.md) |  | [optional] 
**Maps** | Pointer to [**MapsDetails**](MapsDetails.md) |  | [optional] 
**Gini** | Pointer to **map[string]float32** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Subregion** | Pointer to **string** |  | [optional] 
**AltSpellings** | Pointer to **[]string** |  | [optional] 
**Idd** | Pointer to [**IddDetails**](IddDetails.md) |  | [optional] 
**PostalCode** | Pointer to [**PostalCodeDetails**](PostalCodeDetails.md) |  | [optional] 
**Languages** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewAllresponsebodyInner

`func NewAllresponsebodyInner() *AllresponsebodyInner`

NewAllresponsebodyInner instantiates a new AllresponsebodyInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAllresponsebodyInnerWithDefaults

`func NewAllresponsebodyInnerWithDefaults() *AllresponsebodyInner`

NewAllresponsebodyInnerWithDefaults instantiates a new AllresponsebodyInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AllresponsebodyInner) GetName() NameDetails`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AllresponsebodyInner) GetNameOk() (*NameDetails, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AllresponsebodyInner) SetName(v NameDetails)`

SetName sets Name field to given value.

### HasName

`func (o *AllresponsebodyInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTld

`func (o *AllresponsebodyInner) GetTld() []string`

GetTld returns the Tld field if non-nil, zero value otherwise.

### GetTldOk

`func (o *AllresponsebodyInner) GetTldOk() (*[]string, bool)`

GetTldOk returns a tuple with the Tld field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTld

`func (o *AllresponsebodyInner) SetTld(v []string)`

SetTld sets Tld field to given value.

### HasTld

`func (o *AllresponsebodyInner) HasTld() bool`

HasTld returns a boolean if a field has been set.

### GetCca2

`func (o *AllresponsebodyInner) GetCca2() string`

GetCca2 returns the Cca2 field if non-nil, zero value otherwise.

### GetCca2Ok

`func (o *AllresponsebodyInner) GetCca2Ok() (*string, bool)`

GetCca2Ok returns a tuple with the Cca2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCca2

`func (o *AllresponsebodyInner) SetCca2(v string)`

SetCca2 sets Cca2 field to given value.

### HasCca2

`func (o *AllresponsebodyInner) HasCca2() bool`

HasCca2 returns a boolean if a field has been set.

### GetCca3

`func (o *AllresponsebodyInner) GetCca3() string`

GetCca3 returns the Cca3 field if non-nil, zero value otherwise.

### GetCca3Ok

`func (o *AllresponsebodyInner) GetCca3Ok() (*string, bool)`

GetCca3Ok returns a tuple with the Cca3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCca3

`func (o *AllresponsebodyInner) SetCca3(v string)`

SetCca3 sets Cca3 field to given value.

### HasCca3

`func (o *AllresponsebodyInner) HasCca3() bool`

HasCca3 returns a boolean if a field has been set.

### GetCcn3

`func (o *AllresponsebodyInner) GetCcn3() string`

GetCcn3 returns the Ccn3 field if non-nil, zero value otherwise.

### GetCcn3Ok

`func (o *AllresponsebodyInner) GetCcn3Ok() (*string, bool)`

GetCcn3Ok returns a tuple with the Ccn3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcn3

`func (o *AllresponsebodyInner) SetCcn3(v string)`

SetCcn3 sets Ccn3 field to given value.

### HasCcn3

`func (o *AllresponsebodyInner) HasCcn3() bool`

HasCcn3 returns a boolean if a field has been set.

### GetCioc

`func (o *AllresponsebodyInner) GetCioc() string`

GetCioc returns the Cioc field if non-nil, zero value otherwise.

### GetCiocOk

`func (o *AllresponsebodyInner) GetCiocOk() (*string, bool)`

GetCiocOk returns a tuple with the Cioc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCioc

`func (o *AllresponsebodyInner) SetCioc(v string)`

SetCioc sets Cioc field to given value.

### HasCioc

`func (o *AllresponsebodyInner) HasCioc() bool`

HasCioc returns a boolean if a field has been set.

### GetCurrencies

`func (o *AllresponsebodyInner) GetCurrencies() map[string]CurrenciesDetails`

GetCurrencies returns the Currencies field if non-nil, zero value otherwise.

### GetCurrenciesOk

`func (o *AllresponsebodyInner) GetCurrenciesOk() (*map[string]CurrenciesDetails, bool)`

GetCurrenciesOk returns a tuple with the Currencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencies

`func (o *AllresponsebodyInner) SetCurrencies(v map[string]CurrenciesDetails)`

SetCurrencies sets Currencies field to given value.

### HasCurrencies

`func (o *AllresponsebodyInner) HasCurrencies() bool`

HasCurrencies returns a boolean if a field has been set.

### GetRegion

`func (o *AllresponsebodyInner) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *AllresponsebodyInner) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *AllresponsebodyInner) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *AllresponsebodyInner) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetIndependent

`func (o *AllresponsebodyInner) GetIndependent() bool`

GetIndependent returns the Independent field if non-nil, zero value otherwise.

### GetIndependentOk

`func (o *AllresponsebodyInner) GetIndependentOk() (*bool, bool)`

GetIndependentOk returns a tuple with the Independent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndependent

`func (o *AllresponsebodyInner) SetIndependent(v bool)`

SetIndependent sets Independent field to given value.

### HasIndependent

`func (o *AllresponsebodyInner) HasIndependent() bool`

HasIndependent returns a boolean if a field has been set.

### GetUnMember

`func (o *AllresponsebodyInner) GetUnMember() bool`

GetUnMember returns the UnMember field if non-nil, zero value otherwise.

### GetUnMemberOk

`func (o *AllresponsebodyInner) GetUnMemberOk() (*bool, bool)`

GetUnMemberOk returns a tuple with the UnMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnMember

`func (o *AllresponsebodyInner) SetUnMember(v bool)`

SetUnMember sets UnMember field to given value.

### HasUnMember

`func (o *AllresponsebodyInner) HasUnMember() bool`

HasUnMember returns a boolean if a field has been set.

### GetCapital

`func (o *AllresponsebodyInner) GetCapital() []string`

GetCapital returns the Capital field if non-nil, zero value otherwise.

### GetCapitalOk

`func (o *AllresponsebodyInner) GetCapitalOk() (*[]string, bool)`

GetCapitalOk returns a tuple with the Capital field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapital

`func (o *AllresponsebodyInner) SetCapital(v []string)`

SetCapital sets Capital field to given value.

### HasCapital

`func (o *AllresponsebodyInner) HasCapital() bool`

HasCapital returns a boolean if a field has been set.

### GetLatlng

`func (o *AllresponsebodyInner) GetLatlng() []float32`

GetLatlng returns the Latlng field if non-nil, zero value otherwise.

### GetLatlngOk

`func (o *AllresponsebodyInner) GetLatlngOk() (*[]float32, bool)`

GetLatlngOk returns a tuple with the Latlng field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatlng

`func (o *AllresponsebodyInner) SetLatlng(v []float32)`

SetLatlng sets Latlng field to given value.

### HasLatlng

`func (o *AllresponsebodyInner) HasLatlng() bool`

HasLatlng returns a boolean if a field has been set.

### GetArea

`func (o *AllresponsebodyInner) GetArea() float32`

GetArea returns the Area field if non-nil, zero value otherwise.

### GetAreaOk

`func (o *AllresponsebodyInner) GetAreaOk() (*float32, bool)`

GetAreaOk returns a tuple with the Area field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArea

`func (o *AllresponsebodyInner) SetArea(v float32)`

SetArea sets Area field to given value.

### HasArea

`func (o *AllresponsebodyInner) HasArea() bool`

HasArea returns a boolean if a field has been set.

### GetPopulation

`func (o *AllresponsebodyInner) GetPopulation() float32`

GetPopulation returns the Population field if non-nil, zero value otherwise.

### GetPopulationOk

`func (o *AllresponsebodyInner) GetPopulationOk() (*float32, bool)`

GetPopulationOk returns a tuple with the Population field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPopulation

`func (o *AllresponsebodyInner) SetPopulation(v float32)`

SetPopulation sets Population field to given value.

### HasPopulation

`func (o *AllresponsebodyInner) HasPopulation() bool`

HasPopulation returns a boolean if a field has been set.

### GetFlags

`func (o *AllresponsebodyInner) GetFlags() AllresponsebodyInnerFlags`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *AllresponsebodyInner) GetFlagsOk() (*AllresponsebodyInnerFlags, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *AllresponsebodyInner) SetFlags(v AllresponsebodyInnerFlags)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *AllresponsebodyInner) HasFlags() bool`

HasFlags returns a boolean if a field has been set.

### GetFlag

`func (o *AllresponsebodyInner) GetFlag() string`

GetFlag returns the Flag field if non-nil, zero value otherwise.

### GetFlagOk

`func (o *AllresponsebodyInner) GetFlagOk() (*string, bool)`

GetFlagOk returns a tuple with the Flag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlag

`func (o *AllresponsebodyInner) SetFlag(v string)`

SetFlag sets Flag field to given value.

### HasFlag

`func (o *AllresponsebodyInner) HasFlag() bool`

HasFlag returns a boolean if a field has been set.

### GetLandlocked

`func (o *AllresponsebodyInner) GetLandlocked() bool`

GetLandlocked returns the Landlocked field if non-nil, zero value otherwise.

### GetLandlockedOk

`func (o *AllresponsebodyInner) GetLandlockedOk() (*bool, bool)`

GetLandlockedOk returns a tuple with the Landlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLandlocked

`func (o *AllresponsebodyInner) SetLandlocked(v bool)`

SetLandlocked sets Landlocked field to given value.

### HasLandlocked

`func (o *AllresponsebodyInner) HasLandlocked() bool`

HasLandlocked returns a boolean if a field has been set.

### GetBorders

`func (o *AllresponsebodyInner) GetBorders() []string`

GetBorders returns the Borders field if non-nil, zero value otherwise.

### GetBordersOk

`func (o *AllresponsebodyInner) GetBordersOk() (*[]string, bool)`

GetBordersOk returns a tuple with the Borders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBorders

`func (o *AllresponsebodyInner) SetBorders(v []string)`

SetBorders sets Borders field to given value.

### HasBorders

`func (o *AllresponsebodyInner) HasBorders() bool`

HasBorders returns a boolean if a field has been set.

### GetTranslations

`func (o *AllresponsebodyInner) GetTranslations() map[string]NativeNameDetail`

GetTranslations returns the Translations field if non-nil, zero value otherwise.

### GetTranslationsOk

`func (o *AllresponsebodyInner) GetTranslationsOk() (*map[string]NativeNameDetail, bool)`

GetTranslationsOk returns a tuple with the Translations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslations

`func (o *AllresponsebodyInner) SetTranslations(v map[string]NativeNameDetail)`

SetTranslations sets Translations field to given value.

### HasTranslations

`func (o *AllresponsebodyInner) HasTranslations() bool`

HasTranslations returns a boolean if a field has been set.

### GetDemonyms

`func (o *AllresponsebodyInner) GetDemonyms() map[string]DemonymsDetails`

GetDemonyms returns the Demonyms field if non-nil, zero value otherwise.

### GetDemonymsOk

`func (o *AllresponsebodyInner) GetDemonymsOk() (*map[string]DemonymsDetails, bool)`

GetDemonymsOk returns a tuple with the Demonyms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDemonyms

`func (o *AllresponsebodyInner) SetDemonyms(v map[string]DemonymsDetails)`

SetDemonyms sets Demonyms field to given value.

### HasDemonyms

`func (o *AllresponsebodyInner) HasDemonyms() bool`

HasDemonyms returns a boolean if a field has been set.

### GetTimezones

`func (o *AllresponsebodyInner) GetTimezones() []string`

GetTimezones returns the Timezones field if non-nil, zero value otherwise.

### GetTimezonesOk

`func (o *AllresponsebodyInner) GetTimezonesOk() (*[]string, bool)`

GetTimezonesOk returns a tuple with the Timezones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezones

`func (o *AllresponsebodyInner) SetTimezones(v []string)`

SetTimezones sets Timezones field to given value.

### HasTimezones

`func (o *AllresponsebodyInner) HasTimezones() bool`

HasTimezones returns a boolean if a field has been set.

### GetContinents

`func (o *AllresponsebodyInner) GetContinents() []string`

GetContinents returns the Continents field if non-nil, zero value otherwise.

### GetContinentsOk

`func (o *AllresponsebodyInner) GetContinentsOk() (*[]string, bool)`

GetContinentsOk returns a tuple with the Continents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinents

`func (o *AllresponsebodyInner) SetContinents(v []string)`

SetContinents sets Continents field to given value.

### HasContinents

`func (o *AllresponsebodyInner) HasContinents() bool`

HasContinents returns a boolean if a field has been set.

### GetFifa

`func (o *AllresponsebodyInner) GetFifa() string`

GetFifa returns the Fifa field if non-nil, zero value otherwise.

### GetFifaOk

`func (o *AllresponsebodyInner) GetFifaOk() (*string, bool)`

GetFifaOk returns a tuple with the Fifa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFifa

`func (o *AllresponsebodyInner) SetFifa(v string)`

SetFifa sets Fifa field to given value.

### HasFifa

`func (o *AllresponsebodyInner) HasFifa() bool`

HasFifa returns a boolean if a field has been set.

### GetCar

`func (o *AllresponsebodyInner) GetCar() CarDetails`

GetCar returns the Car field if non-nil, zero value otherwise.

### GetCarOk

`func (o *AllresponsebodyInner) GetCarOk() (*CarDetails, bool)`

GetCarOk returns a tuple with the Car field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCar

`func (o *AllresponsebodyInner) SetCar(v CarDetails)`

SetCar sets Car field to given value.

### HasCar

`func (o *AllresponsebodyInner) HasCar() bool`

HasCar returns a boolean if a field has been set.

### GetCoatOfArms

`func (o *AllresponsebodyInner) GetCoatOfArms() ImageDetails`

GetCoatOfArms returns the CoatOfArms field if non-nil, zero value otherwise.

### GetCoatOfArmsOk

`func (o *AllresponsebodyInner) GetCoatOfArmsOk() (*ImageDetails, bool)`

GetCoatOfArmsOk returns a tuple with the CoatOfArms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoatOfArms

`func (o *AllresponsebodyInner) SetCoatOfArms(v ImageDetails)`

SetCoatOfArms sets CoatOfArms field to given value.

### HasCoatOfArms

`func (o *AllresponsebodyInner) HasCoatOfArms() bool`

HasCoatOfArms returns a boolean if a field has been set.

### GetStartOfWeek

`func (o *AllresponsebodyInner) GetStartOfWeek() string`

GetStartOfWeek returns the StartOfWeek field if non-nil, zero value otherwise.

### GetStartOfWeekOk

`func (o *AllresponsebodyInner) GetStartOfWeekOk() (*string, bool)`

GetStartOfWeekOk returns a tuple with the StartOfWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartOfWeek

`func (o *AllresponsebodyInner) SetStartOfWeek(v string)`

SetStartOfWeek sets StartOfWeek field to given value.

### HasStartOfWeek

`func (o *AllresponsebodyInner) HasStartOfWeek() bool`

HasStartOfWeek returns a boolean if a field has been set.

### GetCapitalInfo

`func (o *AllresponsebodyInner) GetCapitalInfo() AllresponsebodyInnerCapitalInfo`

GetCapitalInfo returns the CapitalInfo field if non-nil, zero value otherwise.

### GetCapitalInfoOk

`func (o *AllresponsebodyInner) GetCapitalInfoOk() (*AllresponsebodyInnerCapitalInfo, bool)`

GetCapitalInfoOk returns a tuple with the CapitalInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapitalInfo

`func (o *AllresponsebodyInner) SetCapitalInfo(v AllresponsebodyInnerCapitalInfo)`

SetCapitalInfo sets CapitalInfo field to given value.

### HasCapitalInfo

`func (o *AllresponsebodyInner) HasCapitalInfo() bool`

HasCapitalInfo returns a boolean if a field has been set.

### GetMaps

`func (o *AllresponsebodyInner) GetMaps() MapsDetails`

GetMaps returns the Maps field if non-nil, zero value otherwise.

### GetMapsOk

`func (o *AllresponsebodyInner) GetMapsOk() (*MapsDetails, bool)`

GetMapsOk returns a tuple with the Maps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaps

`func (o *AllresponsebodyInner) SetMaps(v MapsDetails)`

SetMaps sets Maps field to given value.

### HasMaps

`func (o *AllresponsebodyInner) HasMaps() bool`

HasMaps returns a boolean if a field has been set.

### GetGini

`func (o *AllresponsebodyInner) GetGini() map[string]float32`

GetGini returns the Gini field if non-nil, zero value otherwise.

### GetGiniOk

`func (o *AllresponsebodyInner) GetGiniOk() (*map[string]float32, bool)`

GetGiniOk returns a tuple with the Gini field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGini

`func (o *AllresponsebodyInner) SetGini(v map[string]float32)`

SetGini sets Gini field to given value.

### HasGini

`func (o *AllresponsebodyInner) HasGini() bool`

HasGini returns a boolean if a field has been set.

### GetStatus

`func (o *AllresponsebodyInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AllresponsebodyInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AllresponsebodyInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AllresponsebodyInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubregion

`func (o *AllresponsebodyInner) GetSubregion() string`

GetSubregion returns the Subregion field if non-nil, zero value otherwise.

### GetSubregionOk

`func (o *AllresponsebodyInner) GetSubregionOk() (*string, bool)`

GetSubregionOk returns a tuple with the Subregion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubregion

`func (o *AllresponsebodyInner) SetSubregion(v string)`

SetSubregion sets Subregion field to given value.

### HasSubregion

`func (o *AllresponsebodyInner) HasSubregion() bool`

HasSubregion returns a boolean if a field has been set.

### GetAltSpellings

`func (o *AllresponsebodyInner) GetAltSpellings() []string`

GetAltSpellings returns the AltSpellings field if non-nil, zero value otherwise.

### GetAltSpellingsOk

`func (o *AllresponsebodyInner) GetAltSpellingsOk() (*[]string, bool)`

GetAltSpellingsOk returns a tuple with the AltSpellings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAltSpellings

`func (o *AllresponsebodyInner) SetAltSpellings(v []string)`

SetAltSpellings sets AltSpellings field to given value.

### HasAltSpellings

`func (o *AllresponsebodyInner) HasAltSpellings() bool`

HasAltSpellings returns a boolean if a field has been set.

### GetIdd

`func (o *AllresponsebodyInner) GetIdd() IddDetails`

GetIdd returns the Idd field if non-nil, zero value otherwise.

### GetIddOk

`func (o *AllresponsebodyInner) GetIddOk() (*IddDetails, bool)`

GetIddOk returns a tuple with the Idd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdd

`func (o *AllresponsebodyInner) SetIdd(v IddDetails)`

SetIdd sets Idd field to given value.

### HasIdd

`func (o *AllresponsebodyInner) HasIdd() bool`

HasIdd returns a boolean if a field has been set.

### GetPostalCode

`func (o *AllresponsebodyInner) GetPostalCode() PostalCodeDetails`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *AllresponsebodyInner) GetPostalCodeOk() (*PostalCodeDetails, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *AllresponsebodyInner) SetPostalCode(v PostalCodeDetails)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *AllresponsebodyInner) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### GetLanguages

`func (o *AllresponsebodyInner) GetLanguages() map[string]string`

GetLanguages returns the Languages field if non-nil, zero value otherwise.

### GetLanguagesOk

`func (o *AllresponsebodyInner) GetLanguagesOk() (*map[string]string, bool)`

GetLanguagesOk returns a tuple with the Languages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguages

`func (o *AllresponsebodyInner) SetLanguages(v map[string]string)`

SetLanguages sets Languages field to given value.

### HasLanguages

`func (o *AllresponsebodyInner) HasLanguages() bool`

HasLanguages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


