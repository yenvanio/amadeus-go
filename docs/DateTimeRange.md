# DateTimeRange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | **string** | Dates are specified in the [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) YYYY-MM-DD format, e.g. 2018-12-25 | 
**Time** | Pointer to **string** | Local time. hh:mm:ss format, e.g 10:30:00 | [optional] 
**DateWindow** | Pointer to **string** | Either 1, 2 or 3 extra days around the local date (IxD for +/- x days - Ex: I3D), Either 1 to 3 days after the local date (PxD for +x days - Ex: P3D), or 1 to 3 days before the local date (MxD for -x days - Ex: M3D)  Can not be combined with \&quot;originRadius\&quot; or \&quot;destinationRadius\&quot;.  | [optional] 
**TimeWindow** | Pointer to **string** | 1 to 12 hours around (both +and -) the local time. Possibly limited by the number of extra days when specified, i.e.  in some situations, it may not be used to exceed the maximum date range. [1-12]H format, e.g. 6H  Can not be combined with \&quot;originRadius\&quot; or \&quot;destinationRadius\&quot;.  | [optional] 

## Methods

### NewDateTimeRange

`func NewDateTimeRange(date string, ) *DateTimeRange`

NewDateTimeRange instantiates a new DateTimeRange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDateTimeRangeWithDefaults

`func NewDateTimeRangeWithDefaults() *DateTimeRange`

NewDateTimeRangeWithDefaults instantiates a new DateTimeRange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *DateTimeRange) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *DateTimeRange) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *DateTimeRange) SetDate(v string)`

SetDate sets Date field to given value.


### GetTime

`func (o *DateTimeRange) GetTime() string`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *DateTimeRange) GetTimeOk() (*string, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *DateTimeRange) SetTime(v string)`

SetTime sets Time field to given value.

### HasTime

`func (o *DateTimeRange) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetDateWindow

`func (o *DateTimeRange) GetDateWindow() string`

GetDateWindow returns the DateWindow field if non-nil, zero value otherwise.

### GetDateWindowOk

`func (o *DateTimeRange) GetDateWindowOk() (*string, bool)`

GetDateWindowOk returns a tuple with the DateWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateWindow

`func (o *DateTimeRange) SetDateWindow(v string)`

SetDateWindow sets DateWindow field to given value.

### HasDateWindow

`func (o *DateTimeRange) HasDateWindow() bool`

HasDateWindow returns a boolean if a field has been set.

### GetTimeWindow

`func (o *DateTimeRange) GetTimeWindow() string`

GetTimeWindow returns the TimeWindow field if non-nil, zero value otherwise.

### GetTimeWindowOk

`func (o *DateTimeRange) GetTimeWindowOk() (*string, bool)`

GetTimeWindowOk returns a tuple with the TimeWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeWindow

`func (o *DateTimeRange) SetTimeWindow(v string)`

SetTimeWindow sets TimeWindow field to given value.

### HasTimeWindow

`func (o *DateTimeRange) HasTimeWindow() bool`

HasTimeWindow returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


