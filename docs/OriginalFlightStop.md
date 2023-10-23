# OriginalFlightStop

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IataCode** | Pointer to **string** | [IATA airline codes](http://www.iata.org/publications/Pages/code-search.aspx) | [optional] 
**Duration** | Pointer to **string** | stop duration in [ISO8601](https://en.wikipedia.org/wiki/ISO_8601) PnYnMnDTnHnMnS format, e.g. PT2H10M | [optional] 

## Methods

### NewOriginalFlightStop

`func NewOriginalFlightStop() *OriginalFlightStop`

NewOriginalFlightStop instantiates a new OriginalFlightStop object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOriginalFlightStopWithDefaults

`func NewOriginalFlightStopWithDefaults() *OriginalFlightStop`

NewOriginalFlightStopWithDefaults instantiates a new OriginalFlightStop object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIataCode

`func (o *OriginalFlightStop) GetIataCode() string`

GetIataCode returns the IataCode field if non-nil, zero value otherwise.

### GetIataCodeOk

`func (o *OriginalFlightStop) GetIataCodeOk() (*string, bool)`

GetIataCodeOk returns a tuple with the IataCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIataCode

`func (o *OriginalFlightStop) SetIataCode(v string)`

SetIataCode sets IataCode field to given value.

### HasIataCode

`func (o *OriginalFlightStop) HasIataCode() bool`

HasIataCode returns a boolean if a field has been set.

### GetDuration

`func (o *OriginalFlightStop) GetDuration() string`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *OriginalFlightStop) GetDurationOk() (*string, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *OriginalFlightStop) SetDuration(v string)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *OriginalFlightStop) HasDuration() bool`

HasDuration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


