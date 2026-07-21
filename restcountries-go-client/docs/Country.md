# Country

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Names** | Pointer to [**Names**](Names.md) |  | [optional] 
**Capitals** | Pointer to [**[]Capital**](Capital.md) |  | [optional] 
**Demonyms** | Pointer to [**map[string]Demonym**](Demonym.md) | Keyed by ISO 639-3 language code. | [optional] 
**Codes** | Pointer to [**Codes**](Codes.md) |  | [optional] 
**Flag** | Pointer to [**Flag**](Flag.md) |  | [optional] 
**Region** | Pointer to **string** | Africa, Americas, Asia, Europe, Oceania. | [optional] 
**Subregion** | Pointer to **string** |  | [optional] 
**Continents** | Pointer to **[]string** |  | [optional] 
**Landlocked** | Pointer to **bool** |  | [optional] 
**Borders** | Pointer to **[]string** | ISO alpha-3 codes of land-border neighbours. | [optional] 
**Area** | Pointer to [**CountryArea**](CountryArea.md) |  | [optional] 
**Coordinates** | Pointer to [**Coordinates**](Coordinates.md) |  | [optional] 
**Timezones** | Pointer to **[]string** | UTC offsets observed (e.g. UTC+05:30). | [optional] 
**Population** | Pointer to **int64** | Top-level integer (verified against a live v5 record). Dynamic, 4-hour sync. | [optional] 
**Economy** | Pointer to [**CountryEconomy**](CountryEconomy.md) |  | [optional] 
**Languages** | Pointer to [**[]Language**](Language.md) | Languages used in the country. &#x60;name&#x60; (English name) verified as a live lookup path; the identifier key names below follow the docs&#39; description (ISO 639 identifiers, BCP 47 tag, native name) and should be confirmed against one live record before relying on them (x-unverified).  | [optional] 
**Currencies** | Pointer to [**[]Currency**](Currency.md) | Officially supported currencies. &#x60;code&#x60; verified as a live lookup path; &#x60;name&#x60;/&#x60;symbol&#x60; follow convention and should be confirmed against one live record (x-unverified).  | [optional] 
**CallingCodes** | Pointer to **[]string** |  | [optional] 
**Tlds** | Pointer to **[]string** |  | [optional] 
**Cars** | Pointer to [**CountryCars**](CountryCars.md) |  | [optional] 
**PostalCode** | Pointer to [**CountryPostalCode**](CountryPostalCode.md) |  | [optional] 
**Date** | Pointer to [**DateConventions**](DateConventions.md) |  | [optional] 
**NumberFormat** | Pointer to [**CountryNumberFormat**](CountryNumberFormat.md) |  | [optional] 
**Units** | Pointer to [**CountryUnits**](CountryUnits.md) |  | [optional] 
**Classification** | Pointer to [**Classification**](Classification.md) |  | [optional] 
**Parent** | Pointer to [**CountryParent**](CountryParent.md) |  | [optional] 
**Memberships** | Pointer to [**Memberships**](Memberships.md) |  | [optional] 
**GovernmentType** | Pointer to **string** | Dynamic field. | [optional] 
**Leaders** | Pointer to [**[]Leader**](Leader.md) | Premium dynamic field (Personal plans and above). On the free plan the array carries a single upgrade-notice object instead of leader records, so all leaf fields must stay optional.  | [optional] 
**Links** | Pointer to [**CountryLinks**](CountryLinks.md) |  | [optional] 

## Methods

### NewCountry

`func NewCountry() *Country`

NewCountry instantiates a new Country object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCountryWithDefaults

`func NewCountryWithDefaults() *Country`

NewCountryWithDefaults instantiates a new Country object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNames

`func (o *Country) GetNames() Names`

GetNames returns the Names field if non-nil, zero value otherwise.

### GetNamesOk

`func (o *Country) GetNamesOk() (*Names, bool)`

GetNamesOk returns a tuple with the Names field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNames

`func (o *Country) SetNames(v Names)`

SetNames sets Names field to given value.

### HasNames

`func (o *Country) HasNames() bool`

HasNames returns a boolean if a field has been set.

### GetCapitals

`func (o *Country) GetCapitals() []Capital`

GetCapitals returns the Capitals field if non-nil, zero value otherwise.

### GetCapitalsOk

`func (o *Country) GetCapitalsOk() (*[]Capital, bool)`

GetCapitalsOk returns a tuple with the Capitals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapitals

`func (o *Country) SetCapitals(v []Capital)`

SetCapitals sets Capitals field to given value.

### HasCapitals

`func (o *Country) HasCapitals() bool`

HasCapitals returns a boolean if a field has been set.

### GetDemonyms

`func (o *Country) GetDemonyms() map[string]Demonym`

GetDemonyms returns the Demonyms field if non-nil, zero value otherwise.

### GetDemonymsOk

`func (o *Country) GetDemonymsOk() (*map[string]Demonym, bool)`

GetDemonymsOk returns a tuple with the Demonyms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDemonyms

`func (o *Country) SetDemonyms(v map[string]Demonym)`

SetDemonyms sets Demonyms field to given value.

### HasDemonyms

`func (o *Country) HasDemonyms() bool`

HasDemonyms returns a boolean if a field has been set.

### GetCodes

`func (o *Country) GetCodes() Codes`

GetCodes returns the Codes field if non-nil, zero value otherwise.

### GetCodesOk

`func (o *Country) GetCodesOk() (*Codes, bool)`

GetCodesOk returns a tuple with the Codes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodes

`func (o *Country) SetCodes(v Codes)`

SetCodes sets Codes field to given value.

### HasCodes

`func (o *Country) HasCodes() bool`

HasCodes returns a boolean if a field has been set.

### GetFlag

`func (o *Country) GetFlag() Flag`

GetFlag returns the Flag field if non-nil, zero value otherwise.

### GetFlagOk

`func (o *Country) GetFlagOk() (*Flag, bool)`

GetFlagOk returns a tuple with the Flag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlag

`func (o *Country) SetFlag(v Flag)`

SetFlag sets Flag field to given value.

### HasFlag

`func (o *Country) HasFlag() bool`

HasFlag returns a boolean if a field has been set.

### GetRegion

`func (o *Country) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *Country) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *Country) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *Country) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetSubregion

`func (o *Country) GetSubregion() string`

GetSubregion returns the Subregion field if non-nil, zero value otherwise.

### GetSubregionOk

`func (o *Country) GetSubregionOk() (*string, bool)`

GetSubregionOk returns a tuple with the Subregion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubregion

`func (o *Country) SetSubregion(v string)`

SetSubregion sets Subregion field to given value.

### HasSubregion

`func (o *Country) HasSubregion() bool`

HasSubregion returns a boolean if a field has been set.

### GetContinents

`func (o *Country) GetContinents() []string`

GetContinents returns the Continents field if non-nil, zero value otherwise.

### GetContinentsOk

`func (o *Country) GetContinentsOk() (*[]string, bool)`

GetContinentsOk returns a tuple with the Continents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinents

`func (o *Country) SetContinents(v []string)`

SetContinents sets Continents field to given value.

### HasContinents

`func (o *Country) HasContinents() bool`

HasContinents returns a boolean if a field has been set.

### GetLandlocked

`func (o *Country) GetLandlocked() bool`

GetLandlocked returns the Landlocked field if non-nil, zero value otherwise.

### GetLandlockedOk

`func (o *Country) GetLandlockedOk() (*bool, bool)`

GetLandlockedOk returns a tuple with the Landlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLandlocked

`func (o *Country) SetLandlocked(v bool)`

SetLandlocked sets Landlocked field to given value.

### HasLandlocked

`func (o *Country) HasLandlocked() bool`

HasLandlocked returns a boolean if a field has been set.

### GetBorders

`func (o *Country) GetBorders() []string`

GetBorders returns the Borders field if non-nil, zero value otherwise.

### GetBordersOk

`func (o *Country) GetBordersOk() (*[]string, bool)`

GetBordersOk returns a tuple with the Borders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBorders

`func (o *Country) SetBorders(v []string)`

SetBorders sets Borders field to given value.

### HasBorders

`func (o *Country) HasBorders() bool`

HasBorders returns a boolean if a field has been set.

### GetArea

`func (o *Country) GetArea() CountryArea`

GetArea returns the Area field if non-nil, zero value otherwise.

### GetAreaOk

`func (o *Country) GetAreaOk() (*CountryArea, bool)`

GetAreaOk returns a tuple with the Area field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArea

`func (o *Country) SetArea(v CountryArea)`

SetArea sets Area field to given value.

### HasArea

`func (o *Country) HasArea() bool`

HasArea returns a boolean if a field has been set.

### GetCoordinates

`func (o *Country) GetCoordinates() Coordinates`

GetCoordinates returns the Coordinates field if non-nil, zero value otherwise.

### GetCoordinatesOk

`func (o *Country) GetCoordinatesOk() (*Coordinates, bool)`

GetCoordinatesOk returns a tuple with the Coordinates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoordinates

`func (o *Country) SetCoordinates(v Coordinates)`

SetCoordinates sets Coordinates field to given value.

### HasCoordinates

`func (o *Country) HasCoordinates() bool`

HasCoordinates returns a boolean if a field has been set.

### GetTimezones

`func (o *Country) GetTimezones() []string`

GetTimezones returns the Timezones field if non-nil, zero value otherwise.

### GetTimezonesOk

`func (o *Country) GetTimezonesOk() (*[]string, bool)`

GetTimezonesOk returns a tuple with the Timezones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezones

`func (o *Country) SetTimezones(v []string)`

SetTimezones sets Timezones field to given value.

### HasTimezones

`func (o *Country) HasTimezones() bool`

HasTimezones returns a boolean if a field has been set.

### GetPopulation

`func (o *Country) GetPopulation() int64`

GetPopulation returns the Population field if non-nil, zero value otherwise.

### GetPopulationOk

`func (o *Country) GetPopulationOk() (*int64, bool)`

GetPopulationOk returns a tuple with the Population field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPopulation

`func (o *Country) SetPopulation(v int64)`

SetPopulation sets Population field to given value.

### HasPopulation

`func (o *Country) HasPopulation() bool`

HasPopulation returns a boolean if a field has been set.

### GetEconomy

`func (o *Country) GetEconomy() CountryEconomy`

GetEconomy returns the Economy field if non-nil, zero value otherwise.

### GetEconomyOk

`func (o *Country) GetEconomyOk() (*CountryEconomy, bool)`

GetEconomyOk returns a tuple with the Economy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEconomy

`func (o *Country) SetEconomy(v CountryEconomy)`

SetEconomy sets Economy field to given value.

### HasEconomy

`func (o *Country) HasEconomy() bool`

HasEconomy returns a boolean if a field has been set.

### GetLanguages

`func (o *Country) GetLanguages() []Language`

GetLanguages returns the Languages field if non-nil, zero value otherwise.

### GetLanguagesOk

`func (o *Country) GetLanguagesOk() (*[]Language, bool)`

GetLanguagesOk returns a tuple with the Languages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguages

`func (o *Country) SetLanguages(v []Language)`

SetLanguages sets Languages field to given value.

### HasLanguages

`func (o *Country) HasLanguages() bool`

HasLanguages returns a boolean if a field has been set.

### GetCurrencies

`func (o *Country) GetCurrencies() []Currency`

GetCurrencies returns the Currencies field if non-nil, zero value otherwise.

### GetCurrenciesOk

`func (o *Country) GetCurrenciesOk() (*[]Currency, bool)`

GetCurrenciesOk returns a tuple with the Currencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencies

`func (o *Country) SetCurrencies(v []Currency)`

SetCurrencies sets Currencies field to given value.

### HasCurrencies

`func (o *Country) HasCurrencies() bool`

HasCurrencies returns a boolean if a field has been set.

### GetCallingCodes

`func (o *Country) GetCallingCodes() []string`

GetCallingCodes returns the CallingCodes field if non-nil, zero value otherwise.

### GetCallingCodesOk

`func (o *Country) GetCallingCodesOk() (*[]string, bool)`

GetCallingCodesOk returns a tuple with the CallingCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallingCodes

`func (o *Country) SetCallingCodes(v []string)`

SetCallingCodes sets CallingCodes field to given value.

### HasCallingCodes

`func (o *Country) HasCallingCodes() bool`

HasCallingCodes returns a boolean if a field has been set.

### GetTlds

`func (o *Country) GetTlds() []string`

GetTlds returns the Tlds field if non-nil, zero value otherwise.

### GetTldsOk

`func (o *Country) GetTldsOk() (*[]string, bool)`

GetTldsOk returns a tuple with the Tlds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlds

`func (o *Country) SetTlds(v []string)`

SetTlds sets Tlds field to given value.

### HasTlds

`func (o *Country) HasTlds() bool`

HasTlds returns a boolean if a field has been set.

### GetCars

`func (o *Country) GetCars() CountryCars`

GetCars returns the Cars field if non-nil, zero value otherwise.

### GetCarsOk

`func (o *Country) GetCarsOk() (*CountryCars, bool)`

GetCarsOk returns a tuple with the Cars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCars

`func (o *Country) SetCars(v CountryCars)`

SetCars sets Cars field to given value.

### HasCars

`func (o *Country) HasCars() bool`

HasCars returns a boolean if a field has been set.

### GetPostalCode

`func (o *Country) GetPostalCode() CountryPostalCode`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *Country) GetPostalCodeOk() (*CountryPostalCode, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *Country) SetPostalCode(v CountryPostalCode)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *Country) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### GetDate

`func (o *Country) GetDate() DateConventions`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *Country) GetDateOk() (*DateConventions, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *Country) SetDate(v DateConventions)`

SetDate sets Date field to given value.

### HasDate

`func (o *Country) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetNumberFormat

`func (o *Country) GetNumberFormat() CountryNumberFormat`

GetNumberFormat returns the NumberFormat field if non-nil, zero value otherwise.

### GetNumberFormatOk

`func (o *Country) GetNumberFormatOk() (*CountryNumberFormat, bool)`

GetNumberFormatOk returns a tuple with the NumberFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberFormat

`func (o *Country) SetNumberFormat(v CountryNumberFormat)`

SetNumberFormat sets NumberFormat field to given value.

### HasNumberFormat

`func (o *Country) HasNumberFormat() bool`

HasNumberFormat returns a boolean if a field has been set.

### GetUnits

`func (o *Country) GetUnits() CountryUnits`

GetUnits returns the Units field if non-nil, zero value otherwise.

### GetUnitsOk

`func (o *Country) GetUnitsOk() (*CountryUnits, bool)`

GetUnitsOk returns a tuple with the Units field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnits

`func (o *Country) SetUnits(v CountryUnits)`

SetUnits sets Units field to given value.

### HasUnits

`func (o *Country) HasUnits() bool`

HasUnits returns a boolean if a field has been set.

### GetClassification

`func (o *Country) GetClassification() Classification`

GetClassification returns the Classification field if non-nil, zero value otherwise.

### GetClassificationOk

`func (o *Country) GetClassificationOk() (*Classification, bool)`

GetClassificationOk returns a tuple with the Classification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassification

`func (o *Country) SetClassification(v Classification)`

SetClassification sets Classification field to given value.

### HasClassification

`func (o *Country) HasClassification() bool`

HasClassification returns a boolean if a field has been set.

### GetParent

`func (o *Country) GetParent() CountryParent`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *Country) GetParentOk() (*CountryParent, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *Country) SetParent(v CountryParent)`

SetParent sets Parent field to given value.

### HasParent

`func (o *Country) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetMemberships

`func (o *Country) GetMemberships() Memberships`

GetMemberships returns the Memberships field if non-nil, zero value otherwise.

### GetMembershipsOk

`func (o *Country) GetMembershipsOk() (*Memberships, bool)`

GetMembershipsOk returns a tuple with the Memberships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberships

`func (o *Country) SetMemberships(v Memberships)`

SetMemberships sets Memberships field to given value.

### HasMemberships

`func (o *Country) HasMemberships() bool`

HasMemberships returns a boolean if a field has been set.

### GetGovernmentType

`func (o *Country) GetGovernmentType() string`

GetGovernmentType returns the GovernmentType field if non-nil, zero value otherwise.

### GetGovernmentTypeOk

`func (o *Country) GetGovernmentTypeOk() (*string, bool)`

GetGovernmentTypeOk returns a tuple with the GovernmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGovernmentType

`func (o *Country) SetGovernmentType(v string)`

SetGovernmentType sets GovernmentType field to given value.

### HasGovernmentType

`func (o *Country) HasGovernmentType() bool`

HasGovernmentType returns a boolean if a field has been set.

### GetLeaders

`func (o *Country) GetLeaders() []Leader`

GetLeaders returns the Leaders field if non-nil, zero value otherwise.

### GetLeadersOk

`func (o *Country) GetLeadersOk() (*[]Leader, bool)`

GetLeadersOk returns a tuple with the Leaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaders

`func (o *Country) SetLeaders(v []Leader)`

SetLeaders sets Leaders field to given value.

### HasLeaders

`func (o *Country) HasLeaders() bool`

HasLeaders returns a boolean if a field has been set.

### GetLinks

`func (o *Country) GetLinks() CountryLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Country) GetLinksOk() (*CountryLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Country) SetLinks(v CountryLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *Country) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


