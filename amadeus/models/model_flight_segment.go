package models

import (
	"encoding/json"
	"github.com/yenvanio/amadeus-go/amadeus/utils"
)

// checks if the FlightSegment type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &FlightSegment{}

// FlightSegment defining a flight segment; including both operating and marketing details when applicable
type FlightSegment struct {
	Departure *FlightEndPoint `json:"departure,omitempty"`
	Arrival   *FlightEndPoint `json:"arrival,omitempty"`
	// providing the airline / carrier code
	CarrierCode *string `json:"carrierCode,omitempty"`
	// the flight number as assigned by the carrier
	Number    *string            `json:"number,omitempty"`
	Aircraft  *AircraftEquipment `json:"aircraft,omitempty"`
	Operating *OperatingFlight   `json:"operating,omitempty"`
	// stop duration in [ISO8601](https://en.wikipedia.org/wiki/ISO_8601) PnYnMnDTnHnMnS format, e.g. PT2H10M
	Duration *string `json:"duration,omitempty"`
	// information regarding the different stops composing the flight segment. E.g. technical stop, change of gauge...
	Stops []FlightStop `json:"stops,omitempty"`
}

// NewFlightSegment instantiates a new FlightSegment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFlightSegment() *FlightSegment {
	this := FlightSegment{}
	return &this
}

// NewFlightSegmentWithDefaults instantiates a new FlightSegment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFlightSegmentWithDefaults() *FlightSegment {
	this := FlightSegment{}
	return &this
}

// GetDeparture returns the Departure field value if set, zero value otherwise.
func (o *FlightSegment) GetDeparture() FlightEndPoint {
	if o == nil || utils.IsNil(o.Departure) {
		var ret FlightEndPoint
		return ret
	}
	return *o.Departure
}

// GetDepartureOk returns a tuple with the Departure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlightSegment) GetDepartureOk() (*FlightEndPoint, bool) {
	if o == nil || utils.IsNil(o.Departure) {
		return nil, false
	}
	return o.Departure, true
}

// HasDeparture returns a boolean if a field has been set.
func (o *FlightSegment) HasDeparture() bool {
	if o != nil && !utils.IsNil(o.Departure) {
		return true
	}

	return false
}

// SetDeparture gets a reference to the given FlightEndPoint and assigns it to the Departure field.
func (o *FlightSegment) SetDeparture(v FlightEndPoint) {
	o.Departure = &v
}

// GetArrival returns the Arrival field value if set, zero value otherwise.
func (o *FlightSegment) GetArrival() FlightEndPoint {
	if o == nil || utils.IsNil(o.Arrival) {
		var ret FlightEndPoint
		return ret
	}
	return *o.Arrival
}

// GetArrivalOk returns a tuple with the Arrival field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlightSegment) GetArrivalOk() (*FlightEndPoint, bool) {
	if o == nil || utils.IsNil(o.Arrival) {
		return nil, false
	}
	return o.Arrival, true
}

// HasArrival returns a boolean if a field has been set.
func (o *FlightSegment) HasArrival() bool {
	if o != nil && !utils.IsNil(o.Arrival) {
		return true
	}

	return false
}

// SetArrival gets a reference to the given FlightEndPoint and assigns it to the Arrival field.
func (o *FlightSegment) SetArrival(v FlightEndPoint) {
	o.Arrival = &v
}

// GetCarrierCode returns the CarrierCode field value if set, zero value otherwise.
func (o *FlightSegment) GetCarrierCode() string {
	if o == nil || utils.IsNil(o.CarrierCode) {
		var ret string
		return ret
	}
	return *o.CarrierCode
}

// GetCarrierCodeOk returns a tuple with the CarrierCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlightSegment) GetCarrierCodeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.CarrierCode) {
		return nil, false
	}
	return o.CarrierCode, true
}

// HasCarrierCode returns a boolean if a field has been set.
func (o *FlightSegment) HasCarrierCode() bool {
	if o != nil && !utils.IsNil(o.CarrierCode) {
		return true
	}

	return false
}

// SetCarrierCode gets a reference to the given string and assigns it to the CarrierCode field.
func (o *FlightSegment) SetCarrierCode(v string) {
	o.CarrierCode = &v
}

// GetNumber returns the Number field value if set, zero value otherwise.
func (o *FlightSegment) GetNumber() string {
	if o == nil || utils.IsNil(o.Number) {
		var ret string
		return ret
	}
	return *o.Number
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlightSegment) GetNumberOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Number) {
		return nil, false
	}
	return o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *FlightSegment) HasNumber() bool {
	if o != nil && !utils.IsNil(o.Number) {
		return true
	}

	return false
}

// SetNumber gets a reference to the given string and assigns it to the Number field.
func (o *FlightSegment) SetNumber(v string) {
	o.Number = &v
}

// GetAircraft returns the Aircraft field value if set, zero value otherwise.
func (o *FlightSegment) GetAircraft() AircraftEquipment {
	if o == nil || utils.IsNil(o.Aircraft) {
		var ret AircraftEquipment
		return ret
	}
	return *o.Aircraft
}

// GetAircraftOk returns a tuple with the Aircraft field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlightSegment) GetAircraftOk() (*AircraftEquipment, bool) {
	if o == nil || utils.IsNil(o.Aircraft) {
		return nil, false
	}
	return o.Aircraft, true
}

// HasAircraft returns a boolean if a field has been set.
func (o *FlightSegment) HasAircraft() bool {
	if o != nil && !utils.IsNil(o.Aircraft) {
		return true
	}

	return false
}

// SetAircraft gets a reference to the given AircraftEquipment and assigns it to the Aircraft field.
func (o *FlightSegment) SetAircraft(v AircraftEquipment) {
	o.Aircraft = &v
}

// GetOperating returns the Operating field value if set, zero value otherwise.
func (o *FlightSegment) GetOperating() OperatingFlight {
	if o == nil || utils.IsNil(o.Operating) {
		var ret OperatingFlight
		return ret
	}
	return *o.Operating
}

// GetOperatingOk returns a tuple with the Operating field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlightSegment) GetOperatingOk() (*OperatingFlight, bool) {
	if o == nil || utils.IsNil(o.Operating) {
		return nil, false
	}
	return o.Operating, true
}

// HasOperating returns a boolean if a field has been set.
func (o *FlightSegment) HasOperating() bool {
	if o != nil && !utils.IsNil(o.Operating) {
		return true
	}

	return false
}

// SetOperating gets a reference to the given OperatingFlight and assigns it to the Operating field.
func (o *FlightSegment) SetOperating(v OperatingFlight) {
	o.Operating = &v
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *FlightSegment) GetDuration() string {
	if o == nil || utils.IsNil(o.Duration) {
		var ret string
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlightSegment) GetDurationOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Duration) {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *FlightSegment) HasDuration() bool {
	if o != nil && !utils.IsNil(o.Duration) {
		return true
	}

	return false
}

// SetDuration gets a reference to the given string and assigns it to the Duration field.
func (o *FlightSegment) SetDuration(v string) {
	o.Duration = &v
}

// GetStops returns the Stops field value if set, zero value otherwise.
func (o *FlightSegment) GetStops() []FlightStop {
	if o == nil || utils.IsNil(o.Stops) {
		var ret []FlightStop
		return ret
	}
	return o.Stops
}

// GetStopsOk returns a tuple with the Stops field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlightSegment) GetStopsOk() ([]FlightStop, bool) {
	if o == nil || utils.IsNil(o.Stops) {
		return nil, false
	}
	return o.Stops, true
}

// HasStops returns a boolean if a field has been set.
func (o *FlightSegment) HasStops() bool {
	if o != nil && !utils.IsNil(o.Stops) {
		return true
	}

	return false
}

// SetStops gets a reference to the given []FlightStop and assigns it to the Stops field.
func (o *FlightSegment) SetStops(v []FlightStop) {
	o.Stops = v
}

func (o FlightSegment) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FlightSegment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Departure) {
		toSerialize["departure"] = o.Departure
	}
	if !utils.IsNil(o.Arrival) {
		toSerialize["arrival"] = o.Arrival
	}
	if !utils.IsNil(o.CarrierCode) {
		toSerialize["carrierCode"] = o.CarrierCode
	}
	if !utils.IsNil(o.Number) {
		toSerialize["number"] = o.Number
	}
	if !utils.IsNil(o.Aircraft) {
		toSerialize["aircraft"] = o.Aircraft
	}
	if !utils.IsNil(o.Operating) {
		toSerialize["operating"] = o.Operating
	}
	if !utils.IsNil(o.Duration) {
		toSerialize["duration"] = o.Duration
	}
	if !utils.IsNil(o.Stops) {
		toSerialize["stops"] = o.Stops
	}
	return toSerialize, nil
}

type NullableFlightSegment struct {
	value *FlightSegment
	isSet bool
}

func (v NullableFlightSegment) Get() *FlightSegment {
	return v.value
}

func (v *NullableFlightSegment) Set(val *FlightSegment) {
	v.value = val
	v.isSet = true
}

func (v NullableFlightSegment) IsSet() bool {
	return v.isSet
}

func (v *NullableFlightSegment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFlightSegment(val *FlightSegment) *NullableFlightSegment {
	return &NullableFlightSegment{value: val, isSet: true}
}

func (v NullableFlightSegment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFlightSegment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
