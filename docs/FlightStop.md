# FlightStop

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IataCode** | Pointer to **string** | [IATA airline codes](http://www.iata.org/publications/Pages/code-search.aspx) | [optional] 
**Duration** | Pointer to **string** | stop duration in [ISO8601](https://en.wikipedia.org/wiki/ISO_8601) PnYnMnDTnHnMnS format, e.g. PT2H10M | [optional] 
**ArrivalAt** | Pointer to **time.Time** | arrival at the stop in [ISO8601](https://en.wikipedia.org/wiki/ISO_8601) YYYY-MM-ddThh:mm:ss format, e.g. 2017-02-10T20:40:00 | [optional] 
**DepartureAt** | Pointer to **time.Time** | departure from the stop in [ISO8601](https://en.wikipedia.org/wiki/ISO_8601) YYYY-MM-ddThh:mm:ss format, e.g. 2017-02-10T20:40:00 | [optional] 

## Methods

### NewFlightStop

`func NewFlightStop() *FlightStop`

NewFlightStop instantiates a new FlightStop object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlightStopWithDefaults

`func NewFlightStopWithDefaults() *FlightStop`

NewFlightStopWithDefaults instantiates a new FlightStop object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIataCode

`func (o *FlightStop) GetIataCode() string`

GetIataCode returns the IataCode field if non-nil, zero value otherwise.

### GetIataCodeOk

`func (o *FlightStop) GetIataCodeOk() (*string, bool)`

GetIataCodeOk returns a tuple with the IataCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIataCode

`func (o *FlightStop) SetIataCode(v string)`

SetIataCode sets IataCode field to given value.

### HasIataCode

`func (o *FlightStop) HasIataCode() bool`

HasIataCode returns a boolean if a field has been set.

### GetDuration

`func (o *FlightStop) GetDuration() string`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *FlightStop) GetDurationOk() (*string, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *FlightStop) SetDuration(v string)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *FlightStop) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetArrivalAt

`func (o *FlightStop) GetArrivalAt() time.Time`

GetArrivalAt returns the ArrivalAt field if non-nil, zero value otherwise.

### GetArrivalAtOk

`func (o *FlightStop) GetArrivalAtOk() (*time.Time, bool)`

GetArrivalAtOk returns a tuple with the ArrivalAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrivalAt

`func (o *FlightStop) SetArrivalAt(v time.Time)`

SetArrivalAt sets ArrivalAt field to given value.

### HasArrivalAt

`func (o *FlightStop) HasArrivalAt() bool`

HasArrivalAt returns a boolean if a field has been set.

### GetDepartureAt

`func (o *FlightStop) GetDepartureAt() time.Time`

GetDepartureAt returns the DepartureAt field if non-nil, zero value otherwise.

### GetDepartureAtOk

`func (o *FlightStop) GetDepartureAtOk() (*time.Time, bool)`

GetDepartureAtOk returns a tuple with the DepartureAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartureAt

`func (o *FlightStop) SetDepartureAt(v time.Time)`

SetDepartureAt sets DepartureAt field to given value.

### HasDepartureAt

`func (o *FlightStop) HasDepartureAt() bool`

HasDepartureAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


