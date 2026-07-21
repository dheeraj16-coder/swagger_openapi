# Classification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sovereign** | Pointer to **bool** | Replaces v3.1&#39;s &#x60;independent&#x60; (verified against a live v5 record). | [optional] 
**UnMember** | Pointer to **bool** |  | [optional] 
**UnObserver** | Pointer to **bool** |  | [optional] 
**Disputed** | Pointer to **bool** |  | [optional] 
**Dependency** | Pointer to **bool** |  | [optional] 
**DependencyType** | Pointer to **string** |  | [optional] 
**IsoStatus** | Pointer to **string** |  | [optional] 

## Methods

### NewClassification

`func NewClassification() *Classification`

NewClassification instantiates a new Classification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClassificationWithDefaults

`func NewClassificationWithDefaults() *Classification`

NewClassificationWithDefaults instantiates a new Classification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSovereign

`func (o *Classification) GetSovereign() bool`

GetSovereign returns the Sovereign field if non-nil, zero value otherwise.

### GetSovereignOk

`func (o *Classification) GetSovereignOk() (*bool, bool)`

GetSovereignOk returns a tuple with the Sovereign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSovereign

`func (o *Classification) SetSovereign(v bool)`

SetSovereign sets Sovereign field to given value.

### HasSovereign

`func (o *Classification) HasSovereign() bool`

HasSovereign returns a boolean if a field has been set.

### GetUnMember

`func (o *Classification) GetUnMember() bool`

GetUnMember returns the UnMember field if non-nil, zero value otherwise.

### GetUnMemberOk

`func (o *Classification) GetUnMemberOk() (*bool, bool)`

GetUnMemberOk returns a tuple with the UnMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnMember

`func (o *Classification) SetUnMember(v bool)`

SetUnMember sets UnMember field to given value.

### HasUnMember

`func (o *Classification) HasUnMember() bool`

HasUnMember returns a boolean if a field has been set.

### GetUnObserver

`func (o *Classification) GetUnObserver() bool`

GetUnObserver returns the UnObserver field if non-nil, zero value otherwise.

### GetUnObserverOk

`func (o *Classification) GetUnObserverOk() (*bool, bool)`

GetUnObserverOk returns a tuple with the UnObserver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnObserver

`func (o *Classification) SetUnObserver(v bool)`

SetUnObserver sets UnObserver field to given value.

### HasUnObserver

`func (o *Classification) HasUnObserver() bool`

HasUnObserver returns a boolean if a field has been set.

### GetDisputed

`func (o *Classification) GetDisputed() bool`

GetDisputed returns the Disputed field if non-nil, zero value otherwise.

### GetDisputedOk

`func (o *Classification) GetDisputedOk() (*bool, bool)`

GetDisputedOk returns a tuple with the Disputed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisputed

`func (o *Classification) SetDisputed(v bool)`

SetDisputed sets Disputed field to given value.

### HasDisputed

`func (o *Classification) HasDisputed() bool`

HasDisputed returns a boolean if a field has been set.

### GetDependency

`func (o *Classification) GetDependency() bool`

GetDependency returns the Dependency field if non-nil, zero value otherwise.

### GetDependencyOk

`func (o *Classification) GetDependencyOk() (*bool, bool)`

GetDependencyOk returns a tuple with the Dependency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependency

`func (o *Classification) SetDependency(v bool)`

SetDependency sets Dependency field to given value.

### HasDependency

`func (o *Classification) HasDependency() bool`

HasDependency returns a boolean if a field has been set.

### GetDependencyType

`func (o *Classification) GetDependencyType() string`

GetDependencyType returns the DependencyType field if non-nil, zero value otherwise.

### GetDependencyTypeOk

`func (o *Classification) GetDependencyTypeOk() (*string, bool)`

GetDependencyTypeOk returns a tuple with the DependencyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependencyType

`func (o *Classification) SetDependencyType(v string)`

SetDependencyType sets DependencyType field to given value.

### HasDependencyType

`func (o *Classification) HasDependencyType() bool`

HasDependencyType returns a boolean if a field has been set.

### GetIsoStatus

`func (o *Classification) GetIsoStatus() string`

GetIsoStatus returns the IsoStatus field if non-nil, zero value otherwise.

### GetIsoStatusOk

`func (o *Classification) GetIsoStatusOk() (*string, bool)`

GetIsoStatusOk returns a tuple with the IsoStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsoStatus

`func (o *Classification) SetIsoStatus(v string)`

SetIsoStatus sets IsoStatus field to given value.

### HasIsoStatus

`func (o *Classification) HasIsoStatus() bool`

HasIsoStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


