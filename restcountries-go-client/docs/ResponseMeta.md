# ResponseMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **int32** | Total matched count after filters. | [optional] 
**Count** | Pointer to **int32** | Size of the returned slice. | [optional] 
**Limit** | Pointer to **int32** |  | [optional] 
**Offset** | Pointer to **int32** |  | [optional] 
**More** | Pointer to **bool** | Whether records exist past this page. | [optional] 
**RequestId** | Pointer to **string** |  | [optional] 

## Methods

### NewResponseMeta

`func NewResponseMeta() *ResponseMeta`

NewResponseMeta instantiates a new ResponseMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseMetaWithDefaults

`func NewResponseMetaWithDefaults() *ResponseMeta`

NewResponseMetaWithDefaults instantiates a new ResponseMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *ResponseMeta) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ResponseMeta) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ResponseMeta) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *ResponseMeta) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetCount

`func (o *ResponseMeta) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ResponseMeta) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ResponseMeta) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *ResponseMeta) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetLimit

`func (o *ResponseMeta) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ResponseMeta) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ResponseMeta) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *ResponseMeta) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *ResponseMeta) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *ResponseMeta) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *ResponseMeta) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *ResponseMeta) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetMore

`func (o *ResponseMeta) GetMore() bool`

GetMore returns the More field if non-nil, zero value otherwise.

### GetMoreOk

`func (o *ResponseMeta) GetMoreOk() (*bool, bool)`

GetMoreOk returns a tuple with the More field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMore

`func (o *ResponseMeta) SetMore(v bool)`

SetMore sets More field to given value.

### HasMore

`func (o *ResponseMeta) HasMore() bool`

HasMore returns a boolean if a field has been set.

### GetRequestId

`func (o *ResponseMeta) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ResponseMeta) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ResponseMeta) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *ResponseMeta) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


