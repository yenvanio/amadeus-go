# DateTimeType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | **string** | Dates are specified in the [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) YYYY-MM-DD format, e.g. 2018-12-25 | 
**Time** | Pointer to **string** | Local time. hh:mm:ss format, e.g 10:30:00 | [optional] 

## Methods

### NewDateTimeType

`func NewDateTimeType(date string, ) *DateTimeType`

NewDateTimeType instantiates a new DateTimeType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDateTimeTypeWithDefaults

`func NewDateTimeTypeWithDefaults() *DateTimeType`

NewDateTimeTypeWithDefaults instantiates a new DateTimeType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *DateTimeType) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *DateTimeType) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *DateTimeType) SetDate(v string)`

SetDate sets Date field to given value.


### GetTime

`func (o *DateTimeType) GetTime() string`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *DateTimeType) GetTimeOk() (*string, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *DateTimeType) SetTime(v string)`

SetTime sets Time field to given value.

### HasTime

`func (o *DateTimeType) HasTime() bool`

HasTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


