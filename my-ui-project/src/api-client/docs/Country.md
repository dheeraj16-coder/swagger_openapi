# Country

A country record. Every field is optional at the wire level because response_fields / response_fields_omit can project the object to any subset. 

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**names** | [**Names**](Names.md) |  | [optional] [default to undefined]
**capitals** | [**Array&lt;Capital&gt;**](Capital.md) |  | [optional] [default to undefined]
**demonyms** | [**{ [key: string]: Demonym; }**](Demonym.md) | Keyed by ISO 639-3 language code. | [optional] [default to undefined]
**codes** | [**Codes**](Codes.md) |  | [optional] [default to undefined]
**flag** | [**Flag**](Flag.md) |  | [optional] [default to undefined]
**region** | **string** | Africa, Americas, Asia, Europe, Oceania. | [optional] [default to undefined]
**subregion** | **string** |  | [optional] [default to undefined]
**continents** | **Array&lt;string&gt;** |  | [optional] [default to undefined]
**landlocked** | **boolean** |  | [optional] [default to undefined]
**borders** | **Array&lt;string&gt;** | ISO alpha-3 codes of land-border neighbours. | [optional] [default to undefined]
**area** | [**CountryArea**](CountryArea.md) |  | [optional] [default to undefined]
**coordinates** | [**Coordinates**](Coordinates.md) |  | [optional] [default to undefined]
**timezones** | **Array&lt;string&gt;** | UTC offsets observed (e.g. UTC+05:30). | [optional] [default to undefined]
**population** | **number** | Top-level integer (verified against a live v5 record). Dynamic, 4-hour sync. | [optional] [default to undefined]
**economy** | [**CountryEconomy**](CountryEconomy.md) |  | [optional] [default to undefined]
**languages** | [**Array&lt;Language&gt;**](Language.md) | Languages used in the country. &#x60;name&#x60; (English name) verified as a live lookup path; the identifier key names below follow the docs\&#39; description (ISO 639 identifiers, BCP 47 tag, native name) and should be confirmed against one live record before relying on them (x-unverified).  | [optional] [default to undefined]
**currencies** | [**Array&lt;Currency&gt;**](Currency.md) | Officially supported currencies. &#x60;code&#x60; verified as a live lookup path; &#x60;name&#x60;/&#x60;symbol&#x60; follow convention and should be confirmed against one live record (x-unverified).  | [optional] [default to undefined]
**calling_codes** | **Array&lt;string&gt;** |  | [optional] [default to undefined]
**tlds** | **Array&lt;string&gt;** |  | [optional] [default to undefined]
**cars** | [**CountryCars**](CountryCars.md) |  | [optional] [default to undefined]
**postal_code** | [**CountryPostalCode**](CountryPostalCode.md) |  | [optional] [default to undefined]
**date** | [**DateConventions**](DateConventions.md) |  | [optional] [default to undefined]
**number_format** | [**CountryNumberFormat**](CountryNumberFormat.md) |  | [optional] [default to undefined]
**units** | [**CountryUnits**](CountryUnits.md) |  | [optional] [default to undefined]
**classification** | [**Classification**](Classification.md) |  | [optional] [default to undefined]
**parent** | [**CountryParent**](CountryParent.md) |  | [optional] [default to undefined]
**memberships** | [**Memberships**](Memberships.md) |  | [optional] [default to undefined]
**government_type** | **string** | Dynamic field. | [optional] [default to undefined]
**leaders** | [**Array&lt;Leader&gt;**](Leader.md) | Premium dynamic field (Personal plans and above). On the free plan the array carries a single upgrade-notice object instead of leader records, so all leaf fields must stay optional.  | [optional] [default to undefined]
**links** | [**CountryLinks**](CountryLinks.md) |  | [optional] [default to undefined]

## Example

```typescript
import { Country } from './api';

const instance: Country = {
    names,
    capitals,
    demonyms,
    codes,
    flag,
    region,
    subregion,
    continents,
    landlocked,
    borders,
    area,
    coordinates,
    timezones,
    population,
    economy,
    languages,
    currencies,
    calling_codes,
    tlds,
    cars,
    postal_code,
    date,
    number_format,
    units,
    classification,
    parent,
    memberships,
    government_type,
    leaders,
    links,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
