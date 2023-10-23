# Segment

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
**Id** | Pointer to **string** | Id of the segment | [optional] 
**NumberOfStops** | Pointer to **int32** | Number of stops | [optional] 
**BlacklistedInEU** | Pointer to **bool** | When the flight has a marketing or/and operating airline that is identified as blacklisted by the European Commission.   To improve travel safety, the European Commission regularly updates the list of the banned carriers from operating in Europe. It allows any Travel Agency located in the European Union to easily identify and hide any travel recommendation based on some unsafe airlines.  The [list of the banned airlines](https://ec.europa.eu/transport/sites/transport/files/air-safety-list_en.pdf) is published in the Official Journal of the European Union, where they are included as annexes A and B to the Commission Regulation. The blacklist of an airline can concern all its flights or some specific aircraft types pertaining to the airline     | [optional] 
**Co2Emissions** | Pointer to [**[]Co2Emission**](Co2Emission.md) | Co2 informations | [optional] 

## Methods

### NewSegment

`func NewSegment() *Segment`

NewSegment instantiates a new Segment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSegmentWithDefaults

`func NewSegmentWithDefaults() *Segment`

NewSegmentWithDefaults instantiates a new Segment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeparture

`func (o *Segment) GetDeparture() FlightEndPoint`

GetDeparture returns the Departure field if non-nil, zero value otherwise.

### GetDepartureOk

`func (o *Segment) GetDepartureOk() (*FlightEndPoint, bool)`

GetDepartureOk returns a tuple with the Departure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeparture

`func (o *Segment) SetDeparture(v FlightEndPoint)`

SetDeparture sets Departure field to given value.

### HasDeparture

`func (o *Segment) HasDeparture() bool`

HasDeparture returns a boolean if a field has been set.

### GetArrival

`func (o *Segment) GetArrival() FlightEndPoint`

GetArrival returns the Arrival field if non-nil, zero value otherwise.

### GetArrivalOk

`func (o *Segment) GetArrivalOk() (*FlightEndPoint, bool)`

GetArrivalOk returns a tuple with the Arrival field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrival

`func (o *Segment) SetArrival(v FlightEndPoint)`

SetArrival sets Arrival field to given value.

### HasArrival

`func (o *Segment) HasArrival() bool`

HasArrival returns a boolean if a field has been set.

### GetCarrierCode

`func (o *Segment) GetCarrierCode() string`

GetCarrierCode returns the CarrierCode field if non-nil, zero value otherwise.

### GetCarrierCodeOk

`func (o *Segment) GetCarrierCodeOk() (*string, bool)`

GetCarrierCodeOk returns a tuple with the CarrierCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierCode

`func (o *Segment) SetCarrierCode(v string)`

SetCarrierCode sets CarrierCode field to given value.

### HasCarrierCode

`func (o *Segment) HasCarrierCode() bool`

HasCarrierCode returns a boolean if a field has been set.

### GetNumber

`func (o *Segment) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *Segment) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *Segment) SetNumber(v string)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *Segment) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetAircraft

`func (o *Segment) GetAircraft() AircraftEquipment`

GetAircraft returns the Aircraft field if non-nil, zero value otherwise.

### GetAircraftOk

`func (o *Segment) GetAircraftOk() (*AircraftEquipment, bool)`

GetAircraftOk returns a tuple with the Aircraft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAircraft

`func (o *Segment) SetAircraft(v AircraftEquipment)`

SetAircraft sets Aircraft field to given value.

### HasAircraft

`func (o *Segment) HasAircraft() bool`

HasAircraft returns a boolean if a field has been set.

### GetOperating

`func (o *Segment) GetOperating() OperatingFlight`

GetOperating returns the Operating field if non-nil, zero value otherwise.

### GetOperatingOk

`func (o *Segment) GetOperatingOk() (*OperatingFlight, bool)`

GetOperatingOk returns a tuple with the Operating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperating

`func (o *Segment) SetOperating(v OperatingFlight)`

SetOperating sets Operating field to given value.

### HasOperating

`func (o *Segment) HasOperating() bool`

HasOperating returns a boolean if a field has been set.

### GetDuration

`func (o *Segment) GetDuration() string`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *Segment) GetDurationOk() (*string, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *Segment) SetDuration(v string)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *Segment) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetStops

`func (o *Segment) GetStops() []FlightStop`

GetStops returns the Stops field if non-nil, zero value otherwise.

### GetStopsOk

`func (o *Segment) GetStopsOk() (*[]FlightStop, bool)`

GetStopsOk returns a tuple with the Stops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStops

`func (o *Segment) SetStops(v []FlightStop)`

SetStops sets Stops field to given value.

### HasStops

`func (o *Segment) HasStops() bool`

HasStops returns a boolean if a field has been set.

### GetId

`func (o *Segment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Segment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Segment) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Segment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNumberOfStops

`func (o *Segment) GetNumberOfStops() int32`

GetNumberOfStops returns the NumberOfStops field if non-nil, zero value otherwise.

### GetNumberOfStopsOk

`func (o *Segment) GetNumberOfStopsOk() (*int32, bool)`

GetNumberOfStopsOk returns a tuple with the NumberOfStops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfStops

`func (o *Segment) SetNumberOfStops(v int32)`

SetNumberOfStops sets NumberOfStops field to given value.

### HasNumberOfStops

`func (o *Segment) HasNumberOfStops() bool`

HasNumberOfStops returns a boolean if a field has been set.

### GetBlacklistedInEU

`func (o *Segment) GetBlacklistedInEU() bool`

GetBlacklistedInEU returns the BlacklistedInEU field if non-nil, zero value otherwise.

### GetBlacklistedInEUOk

`func (o *Segment) GetBlacklistedInEUOk() (*bool, bool)`

GetBlacklistedInEUOk returns a tuple with the BlacklistedInEU field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlacklistedInEU

`func (o *Segment) SetBlacklistedInEU(v bool)`

SetBlacklistedInEU sets BlacklistedInEU field to given value.

### HasBlacklistedInEU

`func (o *Segment) HasBlacklistedInEU() bool`

HasBlacklistedInEU returns a boolean if a field has been set.

### GetCo2Emissions

`func (o *Segment) GetCo2Emissions() []Co2Emission`

GetCo2Emissions returns the Co2Emissions field if non-nil, zero value otherwise.

### GetCo2EmissionsOk

`func (o *Segment) GetCo2EmissionsOk() (*[]Co2Emission, bool)`

GetCo2EmissionsOk returns a tuple with the Co2Emissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCo2Emissions

`func (o *Segment) SetCo2Emissions(v []Co2Emission)`

SetCo2Emissions sets Co2Emissions field to given value.

### HasCo2Emissions

`func (o *Segment) HasCo2Emissions() bool`

HasCo2Emissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


