# Names


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**common** | **string** |  | [optional] [default to undefined]
**official** | **string** |  | [optional] [default to undefined]
**alternates** | **Array&lt;string&gt;** |  | [optional] [default to undefined]
**_native** | [**{ [key: string]: NameForms; }**](NameForms.md) | Keyed by ISO 639-3 language code. | [optional] [default to undefined]
**translations** | [**{ [key: string]: NameForms; }**](NameForms.md) | Keyed by ISO 639-3 language code. Heavy branch; consider response_fields_omit&#x3D;names.translations. | [optional] [default to undefined]

## Example

```typescript
import { Names } from './api';

const instance: Names = {
    common,
    official,
    alternates,
    _native,
    translations,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
