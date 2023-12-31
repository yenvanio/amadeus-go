package models

import (
	"encoding/json"
	"github.com/yenvanio/amadeus-go/amadeus/utils"
)

// checks if the FlightFilters type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &FlightFilters{}

// FlightFilters struct for FlightFilters
type FlightFilters struct {
	// Allows to search a location outside the borders when a radius around a location is specified. Default is false.
	CrossBorderAllowed *bool `json:"crossBorderAllowed,omitempty"`
	// This flag enables/disables the possibility to have more overnight flights in Low Fare Search
	MoreOvernightsAllowed *bool `json:"moreOvernightsAllowed,omitempty"`
	// This option force to retrieve flight-offer with a departure and a return in the same airport
	ReturnToDepartureAirport *bool `json:"returnToDepartureAirport,omitempty"`
	// This flag enable/disable filtering of rail segment (TGV AIR, RAIL ...)
	RailSegmentAllowed *bool `json:"railSegmentAllowed,omitempty"`
	// This flag enable/disable filtering of bus segment
	BusSegmentAllowed *bool `json:"busSegmentAllowed,omitempty"`
	// Maximum flight time as a percentage relative to the shortest flight time available for the itinerary
	MaxFlightTime       *float32             `json:"maxFlightTime,omitempty"`
	CarrierRestrictions *CarrierRestrictions `json:"carrierRestrictions,omitempty"`
	// Restriction towards cabins.
	CabinRestrictions     []ExtendedCabinRestriction `json:"cabinRestrictions,omitempty"`
	ConnectionRestriction *ConnectionRestriction     `json:"connectionRestriction,omitempty"`
}

// NewFlightFilters instantiates a new FlightFilters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFlightFilters() *FlightFilters {
	this := FlightFilters{}
	return &this
}

// NewFlightFiltersWithDefaults instantiates a new FlightFilters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFlightFiltersWithDefaults() *FlightFilters {
	this := FlightFilters{}
	return &this
}

// GetCrossBorderAllowed returns the CrossBorderAllowed field value if set, zero value otherwise.
func (o *FlightFilters) GetCrossBorderAllowed() bool {
	if o == nil || utils.IsNil(o.CrossBorderAllowed) {
		var ret bool
		return ret
	}
	return *o.CrossBorderAllowed
}

// GetCrossBorderAllowedOk returns a tuple with the CrossBorderAllowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlightFilters) GetCrossBorderAllowedOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.CrossBorderAllowed) {
		return nil, false
	}
	return o.CrossBorderAllowed, true
}

// HasCrossBorderAllowed returns a boolean if a field has been set.
func (o *FlightFilters) HasCrossBorderAllowed() bool {
	if o != nil && !utils.IsNil(o.CrossBorderAllowed) {
		return true
	}

	return false
}

// SetCrossBorderAllowed gets a reference to the given bool and assigns it to the CrossBorderAllowed field.
func (o *FlightFilters) SetCrossBorderAllowed(v bool) {
	o.CrossBorderAllowed = &v
}

// GetMoreOvernightsAllowed returns the MoreOvernightsAllowed field value if set, zero value otherwise.
func (o *FlightFilters) GetMoreOvernightsAllowed() bool {
	if o == nil || utils.IsNil(o.MoreOvernightsAllowed) {
		var ret bool
		return ret
	}
	return *o.MoreOvernightsAllowed
}

// GetMoreOvernightsAllowedOk returns a tuple with the MoreOvernightsAllowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlightFilters) GetMoreOvernightsAllowedOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.MoreOvernightsAllowed) {
		return nil, false
	}
	return o.MoreOvernightsAllowed, true
}

// HasMoreOvernightsAllowed returns a boolean if a field has been set.
func (o *FlightFilters) HasMoreOvernightsAllowed() bool {
	if o != nil && !utils.IsNil(o.MoreOvernightsAllowed) {
		return true
	}

	return false
}

// SetMoreOvernightsAllowed gets a reference to the given bool and assigns it to the MoreOvernightsAllowed field.
func (o *FlightFilters) SetMoreOvernightsAllowed(v bool) {
	o.MoreOvernightsAllowed = &v
}

// GetReturnToDepartureAirport returns the ReturnToDepartureAirport field value if set, zero value otherwise.
func (o *FlightFilters) GetReturnToDepartureAirport() bool {
	if o == nil || utils.IsNil(o.ReturnToDepartureAirport) {
		var ret bool
		return ret
	}
	return *o.ReturnToDepartureAirport
}

// GetReturnToDepartureAirportOk returns a tuple with the ReturnToDepartureAirport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlightFilters) GetReturnToDepartureAirportOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.ReturnToDepartureAirport) {
		return nil, false
	}
	return o.ReturnToDepartureAirport, true
}

// HasReturnToDepartureAirport returns a boolean if a field has been set.
func (o *FlightFilters) HasReturnToDepartureAirport() bool {
	if o != nil && !utils.IsNil(o.ReturnToDepartureAirport) {
		return true
	}

	return false
}

// SetReturnToDepartureAirport gets a reference to the given bool and assigns it to the ReturnToDepartureAirport field.
func (o *FlightFilters) SetReturnToDepartureAirport(v bool) {
	o.ReturnToDepartureAirport = &v
}

// GetRailSegmentAllowed returns the RailSegmentAllowed field value if set, zero value otherwise.
func (o *FlightFilters) GetRailSegmentAllowed() bool {
	if o == nil || utils.IsNil(o.RailSegmentAllowed) {
		var ret bool
		return ret
	}
	return *o.RailSegmentAllowed
}

// GetRailSegmentAllowedOk returns a tuple with the RailSegmentAllowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlightFilters) GetRailSegmentAllowedOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.RailSegmentAllowed) {
		return nil, false
	}
	return o.RailSegmentAllowed, true
}

// HasRailSegmentAllowed returns a boolean if a field has been set.
func (o *FlightFilters) HasRailSegmentAllowed() bool {
	if o != nil && !utils.IsNil(o.RailSegmentAllowed) {
		return true
	}

	return false
}

// SetRailSegmentAllowed gets a reference to the given bool and assigns it to the RailSegmentAllowed field.
func (o *FlightFilters) SetRailSegmentAllowed(v bool) {
	o.RailSegmentAllowed = &v
}

// GetBusSegmentAllowed returns the BusSegmentAllowed field value if set, zero value otherwise.
func (o *FlightFilters) GetBusSegmentAllowed() bool {
	if o == nil || utils.IsNil(o.BusSegmentAllowed) {
		var ret bool
		return ret
	}
	return *o.BusSegmentAllowed
}

// GetBusSegmentAllowedOk returns a tuple with the BusSegmentAllowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlightFilters) GetBusSegmentAllowedOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.BusSegmentAllowed) {
		return nil, false
	}
	return o.BusSegmentAllowed, true
}

// HasBusSegmentAllowed returns a boolean if a field has been set.
func (o *FlightFilters) HasBusSegmentAllowed() bool {
	if o != nil && !utils.IsNil(o.BusSegmentAllowed) {
		return true
	}

	return false
}

// SetBusSegmentAllowed gets a reference to the given bool and assigns it to the BusSegmentAllowed field.
func (o *FlightFilters) SetBusSegmentAllowed(v bool) {
	o.BusSegmentAllowed = &v
}

// GetMaxFlightTime returns the MaxFlightTime field value if set, zero value otherwise.
func (o *FlightFilters) GetMaxFlightTime() float32 {
	if o == nil || utils.IsNil(o.MaxFlightTime) {
		var ret float32
		return ret
	}
	return *o.MaxFlightTime
}

// GetMaxFlightTimeOk returns a tuple with the MaxFlightTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlightFilters) GetMaxFlightTimeOk() (*float32, bool) {
	if o == nil || utils.IsNil(o.MaxFlightTime) {
		return nil, false
	}
	return o.MaxFlightTime, true
}

// HasMaxFlightTime returns a boolean if a field has been set.
func (o *FlightFilters) HasMaxFlightTime() bool {
	if o != nil && !utils.IsNil(o.MaxFlightTime) {
		return true
	}

	return false
}

// SetMaxFlightTime gets a reference to the given float32 and assigns it to the MaxFlightTime field.
func (o *FlightFilters) SetMaxFlightTime(v float32) {
	o.MaxFlightTime = &v
}

// GetCarrierRestrictions returns the CarrierRestrictions field value if set, zero value otherwise.
func (o *FlightFilters) GetCarrierRestrictions() CarrierRestrictions {
	if o == nil || utils.IsNil(o.CarrierRestrictions) {
		var ret CarrierRestrictions
		return ret
	}
	return *o.CarrierRestrictions
}

// GetCarrierRestrictionsOk returns a tuple with the CarrierRestrictions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlightFilters) GetCarrierRestrictionsOk() (*CarrierRestrictions, bool) {
	if o == nil || utils.IsNil(o.CarrierRestrictions) {
		return nil, false
	}
	return o.CarrierRestrictions, true
}

// HasCarrierRestrictions returns a boolean if a field has been set.
func (o *FlightFilters) HasCarrierRestrictions() bool {
	if o != nil && !utils.IsNil(o.CarrierRestrictions) {
		return true
	}

	return false
}

// SetCarrierRestrictions gets a reference to the given CarrierRestrictions and assigns it to the CarrierRestrictions field.
func (o *FlightFilters) SetCarrierRestrictions(v CarrierRestrictions) {
	o.CarrierRestrictions = &v
}

// GetCabinRestrictions returns the CabinRestrictions field value if set, zero value otherwise.
func (o *FlightFilters) GetCabinRestrictions() []ExtendedCabinRestriction {
	if o == nil || utils.IsNil(o.CabinRestrictions) {
		var ret []ExtendedCabinRestriction
		return ret
	}
	return o.CabinRestrictions
}

// GetCabinRestrictionsOk returns a tuple with the CabinRestrictions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlightFilters) GetCabinRestrictionsOk() ([]ExtendedCabinRestriction, bool) {
	if o == nil || utils.IsNil(o.CabinRestrictions) {
		return nil, false
	}
	return o.CabinRestrictions, true
}

// HasCabinRestrictions returns a boolean if a field has been set.
func (o *FlightFilters) HasCabinRestrictions() bool {
	if o != nil && !utils.IsNil(o.CabinRestrictions) {
		return true
	}

	return false
}

// SetCabinRestrictions gets a reference to the given []ExtendedCabinRestriction and assigns it to the CabinRestrictions field.
func (o *FlightFilters) SetCabinRestrictions(v []ExtendedCabinRestriction) {
	o.CabinRestrictions = v
}

// GetConnectionRestriction returns the ConnectionRestriction field value if set, zero value otherwise.
func (o *FlightFilters) GetConnectionRestriction() ConnectionRestriction {
	if o == nil || utils.IsNil(o.ConnectionRestriction) {
		var ret ConnectionRestriction
		return ret
	}
	return *o.ConnectionRestriction
}

// GetConnectionRestrictionOk returns a tuple with the ConnectionRestriction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlightFilters) GetConnectionRestrictionOk() (*ConnectionRestriction, bool) {
	if o == nil || utils.IsNil(o.ConnectionRestriction) {
		return nil, false
	}
	return o.ConnectionRestriction, true
}

// HasConnectionRestriction returns a boolean if a field has been set.
func (o *FlightFilters) HasConnectionRestriction() bool {
	if o != nil && !utils.IsNil(o.ConnectionRestriction) {
		return true
	}

	return false
}

// SetConnectionRestriction gets a reference to the given ConnectionRestriction and assigns it to the ConnectionRestriction field.
func (o *FlightFilters) SetConnectionRestriction(v ConnectionRestriction) {
	o.ConnectionRestriction = &v
}

func (o FlightFilters) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FlightFilters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.CrossBorderAllowed) {
		toSerialize["crossBorderAllowed"] = o.CrossBorderAllowed
	}
	if !utils.IsNil(o.MoreOvernightsAllowed) {
		toSerialize["moreOvernightsAllowed"] = o.MoreOvernightsAllowed
	}
	if !utils.IsNil(o.ReturnToDepartureAirport) {
		toSerialize["returnToDepartureAirport"] = o.ReturnToDepartureAirport
	}
	if !utils.IsNil(o.RailSegmentAllowed) {
		toSerialize["railSegmentAllowed"] = o.RailSegmentAllowed
	}
	if !utils.IsNil(o.BusSegmentAllowed) {
		toSerialize["busSegmentAllowed"] = o.BusSegmentAllowed
	}
	if !utils.IsNil(o.MaxFlightTime) {
		toSerialize["maxFlightTime"] = o.MaxFlightTime
	}
	if !utils.IsNil(o.CarrierRestrictions) {
		toSerialize["carrierRestrictions"] = o.CarrierRestrictions
	}
	if !utils.IsNil(o.CabinRestrictions) {
		toSerialize["cabinRestrictions"] = o.CabinRestrictions
	}
	if !utils.IsNil(o.ConnectionRestriction) {
		toSerialize["connectionRestriction"] = o.ConnectionRestriction
	}
	return toSerialize, nil
}

type NullableFlightFilters struct {
	value *FlightFilters
	isSet bool
}

func (v NullableFlightFilters) Get() *FlightFilters {
	return v.value
}

func (v *NullableFlightFilters) Set(val *FlightFilters) {
	v.value = val
	v.isSet = true
}

func (v NullableFlightFilters) IsSet() bool {
	return v.isSet
}

func (v *NullableFlightFilters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFlightFilters(val *FlightFilters) *NullableFlightFilters {
	return &NullableFlightFilters{value: val, isSet: true}
}

func (v NullableFlightFilters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFlightFilters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
