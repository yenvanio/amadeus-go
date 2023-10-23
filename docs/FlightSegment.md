# FlightSegment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Departure** | Pointer to [**FlightEndPoint**](FlightEndPoint.md) |  | [optional] 
**Arrival** | Pointer to [**FlightEndPoint**](FlightEndPoint.md) |  | [optional] 
**CarrierCode** | Pointer to **string** | providing the airline / carrier code | [optional] 
**Number** | Pointer to **string** | the flight number as assigned by the carrier | [optional] 
**Aircraft** | Pointer to [**AircraftEquipment**](AircraftEquipment.md) |  | [optional] 
**Operating** | Pointer to [**OperatingFlight**](OperatingFlight.md) |  | [optional] 
**Duration** | Pointer to **string** | stop duration in [ISO8601](https://en.wikipedia.org/wiki/ISO_8601) PnYnMnDTnHnMnS format, e.g. PT2H10M | [optional] 
**Stops** | Pointer to [**[]FlightStop**](FlightStop.md) | information regarding the different stops composing the flight segment. E.g. technical stop, change of gauge... | [optional] 

## Methods

### NewFlightSegment

`func NewFlightSegment() *FlightSegment`

NewFlightSegment instantiates a new FlightSegment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlightSegmentWithDefaults

`func NewFlightSegmentWithDefaults() *FlightSegment`

NewFlightSegmentWithDefaults instantiates a new FlightSegment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeparture

`func (o *FlightSegment) GetDeparture() FlightEndPoint`

GetDeparture returns the Departure field if non-nil, zero value otherwise.

### GetDepartureOk

`func (o *FlightSegment) GetDepartureOk() (*FlightEndPoint, bool)`

GetDepartureOk returns a tuple with the Departure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeparture

`func (o *FlightSegment) SetDeparture(v FlightEndPoint)`

SetDeparture sets Departure field to given value.

### HasDeparture

`func (o *FlightSegment) HasDeparture() bool`

HasDeparture returns a boolean if a field has been set.

### GetArrival

`func (o *FlightSegment) GetArrival() FlightEndPoint`

GetArrival returns the Arrival field if non-nil, zero value otherwise.

### GetArrivalOk

`func (o *FlightSegment) GetArrivalOk() (*FlightEndPoint, bool)`

GetArrivalOk returns a tuple with the Arrival field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrival

`func (o *FlightSegment) SetArrival(v FlightEndPoint)`

SetArrival sets Arrival field to given value.

### HasArrival

`func (o *FlightSegment) HasArrival() bool`

HasArrival returns a boolean if a field has been set.

### GetCarrierCode

`func (o *FlightSegment) GetCarrierCode() string`

GetCarrierCode returns the CarrierCode field if non-nil, zero value otherwise.

### GetCarrierCodeOk

`func (o *FlightSegment) GetCarrierCodeOk() (*string, bool)`

GetCarrierCodeOk returns a tuple with the CarrierCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierCode

`func (o *FlightSegment) SetCarrierCode(v string)`

SetCarrierCode sets CarrierCode field to given value.

### HasCarrierCode

`func (o *FlightSegment) HasCarrierCode() bool`

HasCarrierCode returns a boolean if a field has been set.

### GetNumber

`func (o *FlightSegment) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *FlightSegment) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *FlightSegment) SetNumber(v string)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *FlightSegment) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetAircraft

`func (o *FlightSegment) GetAircraft() AircraftEquipment`

GetAircraft returns the Aircraft field if non-nil, zero value otherwise.

### GetAircraftOk

`func (o *FlightSegment) GetAircraftOk() (*AircraftEquipment, bool)`

GetAircraftOk returns a tuple with the Aircraft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAircraft

`func (o *FlightSegment) SetAircraft(v AircraftEquipment)`

SetAircraft sets Aircraft field to given value.

### HasAircraft

`func (o *FlightSegment) HasAircraft() bool`

HasAircraft returns a boolean if a field has been set.

### GetOperating

`func (o *FlightSegment) GetOperating() OperatingFlight`

GetOperating returns the Operating field if non-nil, zero value otherwise.

### GetOperatingOk

`func (o *FlightSegment) GetOperatingOk() (*OperatingFlight, bool)`

GetOperatingOk returns a tuple with the Operating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperating

`func (o *FlightSegment) SetOperating(v OperatingFlight)`

SetOperating sets Operating field to given value.

### HasOperating

`func (o *FlightSegment) HasOperating() bool`

HasOperating returns a boolean if a field has been set.

### GetDuration

`func (o *FlightSegment) GetDuration() string`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *FlightSegment) GetDurationOk() (*string, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *FlightSegment) SetDuration(v string)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *FlightSegment) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetStops

`func (o *FlightSegment) GetStops() []FlightStop`

GetStops returns the Stops field if non-nil, zero value otherwise.

### GetStopsOk

`func (o *FlightSegment) GetStopsOk() (*[]FlightStop, bool)`

GetStopsOk returns a tuple with the Stops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStops

`func (o *FlightSegment) SetStops(v []FlightStop)`

SetStops sets Stops field to given value.

### HasStops

`func (o *FlightSegment) HasStops() bool`

HasStops returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


