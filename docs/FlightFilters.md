# FlightFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CrossBorderAllowed** | Pointer to **bool** | Allows to search a location outside the borders when a radius around a location is specified. Default is false. | [optional] 
**MoreOvernightsAllowed** | Pointer to **bool** | This flag enables/disables the possibility to have more overnight flights in Low Fare Search | [optional] 
**ReturnToDepartureAirport** | Pointer to **bool** | This option force to retrieve flight-offer with a departure and a return in the same airport | [optional] 
**RailSegmentAllowed** | Pointer to **bool** | This flag enable/disable filtering of rail segment (TGV AIR, RAIL ...) | [optional] 
**BusSegmentAllowed** | Pointer to **bool** | This flag enable/disable filtering of bus segment | [optional] 
**MaxFlightTime** | Pointer to **float32** | Maximum flight time as a percentage relative to the shortest flight time available for the itinerary | [optional] 
**CarrierRestrictions** | Pointer to [**CarrierRestrictions**](CarrierRestrictions.md) |  | [optional] 
**CabinRestrictions** | Pointer to [**[]ExtendedCabinRestriction**](ExtendedCabinRestriction.md) | Restriction towards cabins. | [optional] 
**ConnectionRestriction** | Pointer to [**ConnectionRestriction**](ConnectionRestriction.md) |  | [optional] 

## Methods

### NewFlightFilters

`func NewFlightFilters() *FlightFilters`

NewFlightFilters instantiates a new FlightFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlightFiltersWithDefaults

`func NewFlightFiltersWithDefaults() *FlightFilters`

NewFlightFiltersWithDefaults instantiates a new FlightFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCrossBorderAllowed

`func (o *FlightFilters) GetCrossBorderAllowed() bool`

GetCrossBorderAllowed returns the CrossBorderAllowed field if non-nil, zero value otherwise.

### GetCrossBorderAllowedOk

`func (o *FlightFilters) GetCrossBorderAllowedOk() (*bool, bool)`

GetCrossBorderAllowedOk returns a tuple with the CrossBorderAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrossBorderAllowed

`func (o *FlightFilters) SetCrossBorderAllowed(v bool)`

SetCrossBorderAllowed sets CrossBorderAllowed field to given value.

### HasCrossBorderAllowed

`func (o *FlightFilters) HasCrossBorderAllowed() bool`

HasCrossBorderAllowed returns a boolean if a field has been set.

### GetMoreOvernightsAllowed

`func (o *FlightFilters) GetMoreOvernightsAllowed() bool`

GetMoreOvernightsAllowed returns the MoreOvernightsAllowed field if non-nil, zero value otherwise.

### GetMoreOvernightsAllowedOk

`func (o *FlightFilters) GetMoreOvernightsAllowedOk() (*bool, bool)`

GetMoreOvernightsAllowedOk returns a tuple with the MoreOvernightsAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoreOvernightsAllowed

`func (o *FlightFilters) SetMoreOvernightsAllowed(v bool)`

SetMoreOvernightsAllowed sets MoreOvernightsAllowed field to given value.

### HasMoreOvernightsAllowed

`func (o *FlightFilters) HasMoreOvernightsAllowed() bool`

HasMoreOvernightsAllowed returns a boolean if a field has been set.

### GetReturnToDepartureAirport

`func (o *FlightFilters) GetReturnToDepartureAirport() bool`

GetReturnToDepartureAirport returns the ReturnToDepartureAirport field if non-nil, zero value otherwise.

### GetReturnToDepartureAirportOk

`func (o *FlightFilters) GetReturnToDepartureAirportOk() (*bool, bool)`

GetReturnToDepartureAirportOk returns a tuple with the ReturnToDepartureAirport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnToDepartureAirport

`func (o *FlightFilters) SetReturnToDepartureAirport(v bool)`

SetReturnToDepartureAirport sets ReturnToDepartureAirport field to given value.

### HasReturnToDepartureAirport

`func (o *FlightFilters) HasReturnToDepartureAirport() bool`

HasReturnToDepartureAirport returns a boolean if a field has been set.

### GetRailSegmentAllowed

`func (o *FlightFilters) GetRailSegmentAllowed() bool`

GetRailSegmentAllowed returns the RailSegmentAllowed field if non-nil, zero value otherwise.

### GetRailSegmentAllowedOk

`func (o *FlightFilters) GetRailSegmentAllowedOk() (*bool, bool)`

GetRailSegmentAllowedOk returns a tuple with the RailSegmentAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRailSegmentAllowed

`func (o *FlightFilters) SetRailSegmentAllowed(v bool)`

SetRailSegmentAllowed sets RailSegmentAllowed field to given value.

### HasRailSegmentAllowed

`func (o *FlightFilters) HasRailSegmentAllowed() bool`

HasRailSegmentAllowed returns a boolean if a field has been set.

### GetBusSegmentAllowed

`func (o *FlightFilters) GetBusSegmentAllowed() bool`

GetBusSegmentAllowed returns the BusSegmentAllowed field if non-nil, zero value otherwise.

### GetBusSegmentAllowedOk

`func (o *FlightFilters) GetBusSegmentAllowedOk() (*bool, bool)`

GetBusSegmentAllowedOk returns a tuple with the BusSegmentAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusSegmentAllowed

`func (o *FlightFilters) SetBusSegmentAllowed(v bool)`

SetBusSegmentAllowed sets BusSegmentAllowed field to given value.

### HasBusSegmentAllowed

`func (o *FlightFilters) HasBusSegmentAllowed() bool`

HasBusSegmentAllowed returns a boolean if a field has been set.

### GetMaxFlightTime

`func (o *FlightFilters) GetMaxFlightTime() float32`

GetMaxFlightTime returns the MaxFlightTime field if non-nil, zero value otherwise.

### GetMaxFlightTimeOk

`func (o *FlightFilters) GetMaxFlightTimeOk() (*float32, bool)`

GetMaxFlightTimeOk returns a tuple with the MaxFlightTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxFlightTime

`func (o *FlightFilters) SetMaxFlightTime(v float32)`

SetMaxFlightTime sets MaxFlightTime field to given value.

### HasMaxFlightTime

`func (o *FlightFilters) HasMaxFlightTime() bool`

HasMaxFlightTime returns a boolean if a field has been set.

### GetCarrierRestrictions

`func (o *FlightFilters) GetCarrierRestrictions() CarrierRestrictions`

GetCarrierRestrictions returns the CarrierRestrictions field if non-nil, zero value otherwise.

### GetCarrierRestrictionsOk

`func (o *FlightFilters) GetCarrierRestrictionsOk() (*CarrierRestrictions, bool)`

GetCarrierRestrictionsOk returns a tuple with the CarrierRestrictions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierRestrictions

`func (o *FlightFilters) SetCarrierRestrictions(v CarrierRestrictions)`

SetCarrierRestrictions sets CarrierRestrictions field to given value.

### HasCarrierRestrictions

`func (o *FlightFilters) HasCarrierRestrictions() bool`

HasCarrierRestrictions returns a boolean if a field has been set.

### GetCabinRestrictions

`func (o *FlightFilters) GetCabinRestrictions() []ExtendedCabinRestriction`

GetCabinRestrictions returns the CabinRestrictions field if non-nil, zero value otherwise.

### GetCabinRestrictionsOk

`func (o *FlightFilters) GetCabinRestrictionsOk() (*[]ExtendedCabinRestriction, bool)`

GetCabinRestrictionsOk returns a tuple with the CabinRestrictions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCabinRestrictions

`func (o *FlightFilters) SetCabinRestrictions(v []ExtendedCabinRestriction)`

SetCabinRestrictions sets CabinRestrictions field to given value.

### HasCabinRestrictions

`func (o *FlightFilters) HasCabinRestrictions() bool`

HasCabinRestrictions returns a boolean if a field has been set.

### GetConnectionRestriction

`func (o *FlightFilters) GetConnectionRestriction() ConnectionRestriction`

GetConnectionRestriction returns the ConnectionRestriction field if non-nil, zero value otherwise.

### GetConnectionRestrictionOk

`func (o *FlightFilters) GetConnectionRestrictionOk() (*ConnectionRestriction, bool)`

GetConnectionRestrictionOk returns a tuple with the ConnectionRestriction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionRestriction

`func (o *FlightFilters) SetConnectionRestriction(v ConnectionRestriction)`

SetConnectionRestriction sets ConnectionRestriction field to given value.

### HasConnectionRestriction

`func (o *FlightFilters) HasConnectionRestriction() bool`

HasConnectionRestriction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


