# DateConventions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartOfWeek** | Pointer to **string** |  | [optional] 
**AcademicYearStart** | Pointer to [**MonthDay**](MonthDay.md) |  | [optional] 
**FiscalYearStart** | Pointer to [**DateConventionsFiscalYearStart**](DateConventionsFiscalYearStart.md) |  | [optional] 

## Methods

### NewDateConventions

`func NewDateConventions() *DateConventions`

NewDateConventions instantiates a new DateConventions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDateConventionsWithDefaults

`func NewDateConventionsWithDefaults() *DateConventions`

NewDateConventionsWithDefaults instantiates a new DateConventions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartOfWeek

`func (o *DateConventions) GetStartOfWeek() string`

GetStartOfWeek returns the StartOfWeek field if non-nil, zero value otherwise.

### GetStartOfWeekOk

`func (o *DateConventions) GetStartOfWeekOk() (*string, bool)`

GetStartOfWeekOk returns a tuple with the StartOfWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartOfWeek

`func (o *DateConventions) SetStartOfWeek(v string)`

SetStartOfWeek sets StartOfWeek field to given value.

### HasStartOfWeek

`func (o *DateConventions) HasStartOfWeek() bool`

HasStartOfWeek returns a boolean if a field has been set.

### GetAcademicYearStart

`func (o *DateConventions) GetAcademicYearStart() MonthDay`

GetAcademicYearStart returns the AcademicYearStart field if non-nil, zero value otherwise.

### GetAcademicYearStartOk

`func (o *DateConventions) GetAcademicYearStartOk() (*MonthDay, bool)`

GetAcademicYearStartOk returns a tuple with the AcademicYearStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcademicYearStart

`func (o *DateConventions) SetAcademicYearStart(v MonthDay)`

SetAcademicYearStart sets AcademicYearStart field to given value.

### HasAcademicYearStart

`func (o *DateConventions) HasAcademicYearStart() bool`

HasAcademicYearStart returns a boolean if a field has been set.

### GetFiscalYearStart

`func (o *DateConventions) GetFiscalYearStart() DateConventionsFiscalYearStart`

GetFiscalYearStart returns the FiscalYearStart field if non-nil, zero value otherwise.

### GetFiscalYearStartOk

`func (o *DateConventions) GetFiscalYearStartOk() (*DateConventionsFiscalYearStart, bool)`

GetFiscalYearStartOk returns a tuple with the FiscalYearStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiscalYearStart

`func (o *DateConventions) SetFiscalYearStart(v DateConventionsFiscalYearStart)`

SetFiscalYearStart sets FiscalYearStart field to given value.

### HasFiscalYearStart

`func (o *DateConventions) HasFiscalYearStart() bool`

HasFiscalYearStart returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


