package models

import (
	"encoding/json"
	"github.com/yenvanio/amadeus-go/amadeus/utils"
)

// checks if the Segment type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Segment{}

// Segment struct for Segment
type Segment struct {
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
	// Id of the segment
	Id *string `json:"id,omitempty"`
	// Number of stops planned on the segment for technical or operation purpose i.e. refueling
	NumberOfStops *int32 `json:"numberOfStops,omitempty"`
	// When the flight has a marketing or/and operating airline that is identified as blacklisted by the European Commission.   To improve travel safety, the European Commission regularly updates the list of the banned carriers from operating in Europe. It allows any Travel Agency located in the European Union to easily identify and hide any travel recommendation based on some unsafe airlines.  The [list of the banned airlines](https://ec.europa.eu/transport/sites/transport/files/air-safety-list_en.pdf) is published in the Official Journal of the European Union, where they are included as annexes A and B to the Commission Regulation. The blacklist of an airline can concern all its flights or some specific aircraft types pertaining to the airline
	BlacklistedInEU *bool `json:"blacklistedInEU,omitempty"`
	// Co2 informations
	Co2Emissions []Co2Emission `json:"co2Emissions,omitempty"`
}

// NewSegment instantiates a new Segment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSegment() *Segment {
	this := Segment{}
	return &this
}

// NewSegmentWithDefaults instantiates a new Segment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSegmentWithDefaults() *Segment {
	this := Segment{}
	return &this
}

// GetDeparture returns the Departure field value if set, zero value otherwise.
func (o *Segment) GetDeparture() FlightEndPoint {
	if o == nil || utils.IsNil(o.Departure) {
		var ret FlightEndPoint
		return ret
	}
	return *o.Departure
}

// GetDepartureOk returns a tuple with the Departure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Segment) GetDepartureOk() (*FlightEndPoint, bool) {
	if o == nil || utils.IsNil(o.Departure) {
		return nil, false
	}
	return o.Departure, true
}

// HasDeparture returns a boolean if a field has been set.
func (o *Segment) HasDeparture() bool {
	if o != nil && !utils.IsNil(o.Departure) {
		return true
	}

	return false
}

// SetDeparture gets a reference to the given FlightEndPoint and assigns it to the Departure field.
func (o *Segment) SetDeparture(v FlightEndPoint) {
	o.Departure = &v
}

// GetArrival returns the Arrival field value if set, zero value otherwise.
func (o *Segment) GetArrival() FlightEndPoint {
	if o == nil || utils.IsNil(o.Arrival) {
		var ret FlightEndPoint
		return ret
	}
	return *o.Arrival
}

// GetArrivalOk returns a tuple with the Arrival field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Segment) GetArrivalOk() (*FlightEndPoint, bool) {
	if o == nil || utils.IsNil(o.Arrival) {
		return nil, false
	}
	return o.Arrival, true
}

// HasArrival returns a boolean if a field has been set.
func (o *Segment) HasArrival() bool {
	if o != nil && !utils.IsNil(o.Arrival) {
		return true
	}

	return false
}

// SetArrival gets a reference to the given FlightEndPoint and assigns it to the Arrival field.
func (o *Segment) SetArrival(v FlightEndPoint) {
	o.Arrival = &v
}

// GetCarrierCode returns the CarrierCode field value if set, zero value otherwise.
func (o *Segment) GetCarrierCode() string {
	if o == nil || utils.IsNil(o.CarrierCode) {
		var ret string
		return ret
	}
	return *o.CarrierCode
}

// GetCarrierCodeOk returns a tuple with the CarrierCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Segment) GetCarrierCodeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.CarrierCode) {
		return nil, false
	}
	return o.CarrierCode, true
}

// HasCarrierCode returns a boolean if a field has been set.
func (o *Segment) HasCarrierCode() bool {
	if o != nil && !utils.IsNil(o.CarrierCode) {
		return true
	}

	return false
}

// SetCarrierCode gets a reference to the given string and assigns it to the CarrierCode field.
func (o *Segment) SetCarrierCode(v string) {
	o.CarrierCode = &v
}

// GetNumber returns the Number field value if set, zero value otherwise.
func (o *Segment) GetNumber() string {
	if o == nil || utils.IsNil(o.Number) {
		var ret string
		return ret
	}
	return *o.Number
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Segment) GetNumberOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Number) {
		return nil, false
	}
	return o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *Segment) HasNumber() bool {
	if o != nil && !utils.IsNil(o.Number) {
		return true
	}

	return false
}

// SetNumber gets a reference to the given string and assigns it to the Number field.
func (o *Segment) SetNumber(v string) {
	o.Number = &v
}

// GetAircraft returns the Aircraft field value if set, zero value otherwise.
func (o *Segment) GetAircraft() AircraftEquipment {
	if o == nil || utils.IsNil(o.Aircraft) {
		var ret AircraftEquipment
		return ret
	}
	return *o.Aircraft
}

// GetAircraftOk returns a tuple with the Aircraft field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Segment) GetAircraftOk() (*AircraftEquipment, bool) {
	if o == nil || utils.IsNil(o.Aircraft) {
		return nil, false
	}
	return o.Aircraft, true
}

// HasAircraft returns a boolean if a field has been set.
func (o *Segment) HasAircraft() bool {
	if o != nil && !utils.IsNil(o.Aircraft) {
		return true
	}

	return false
}

// SetAircraft gets a reference to the given AircraftEquipment and assigns it to the Aircraft field.
func (o *Segment) SetAircraft(v AircraftEquipment) {
	o.Aircraft = &v
}

// GetOperating returns the Operating field value if set, zero value otherwise.
func (o *Segment) GetOperating() OperatingFlight {
	if o == nil || utils.IsNil(o.Operating) {
		var ret OperatingFlight
		return ret
	}
	return *o.Operating
}

// GetOperatingOk returns a tuple with the Operating field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Segment) GetOperatingOk() (*OperatingFlight, bool) {
	if o == nil || utils.IsNil(o.Operating) {
		return nil, false
	}
	return o.Operating, true
}

// HasOperating returns a boolean if a field has been set.
func (o *Segment) HasOperating() bool {
	if o != nil && !utils.IsNil(o.Operating) {
		return true
	}

	return false
}

// SetOperating gets a reference to the given OperatingFlight and assigns it to the Operating field.
func (o *Segment) SetOperating(v OperatingFlight) {
	o.Operating = &v
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *Segment) GetDuration() string {
	if o == nil || utils.IsNil(o.Duration) {
		var ret string
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Segment) GetDurationOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Duration) {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *Segment) HasDuration() bool {
	if o != nil && !utils.IsNil(o.Duration) {
		return true
	}

	return false
}

// SetDuration gets a reference to the given string and assigns it to the Duration field.
func (o *Segment) SetDuration(v string) {
	o.Duration = &v
}

// GetStops returns the Stops field value if set, zero value otherwise.
func (o *Segment) GetStops() []FlightStop {
	if o == nil || utils.IsNil(o.Stops) {
		var ret []FlightStop
		return ret
	}
	return o.Stops
}

// GetStopsOk returns a tuple with the Stops field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Segment) GetStopsOk() ([]FlightStop, bool) {
	if o == nil || utils.IsNil(o.Stops) {
		return nil, false
	}
	return o.Stops, true
}

// HasStops returns a boolean if a field has been set.
func (o *Segment) HasStops() bool {
	if o != nil && !utils.IsNil(o.Stops) {
		return true
	}

	return false
}

// SetStops gets a reference to the given []FlightStop and assigns it to the Stops field.
func (o *Segment) SetStops(v []FlightStop) {
	o.Stops = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Segment) GetId() string {
	if o == nil || utils.IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Segment) GetIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Segment) HasId() bool {
	if o != nil && !utils.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Segment) SetId(v string) {
	o.Id = &v
}

// GetNumberOfStops returns the NumberOfStops field value if set, zero value otherwise.
func (o *Segment) GetNumberOfStops() int32 {
	if o == nil || utils.IsNil(o.NumberOfStops) {
		var ret int32
		return ret
	}
	return *o.NumberOfStops
}

// GetNumberOfStopsOk returns a tuple with the NumberOfStops field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Segment) GetNumberOfStopsOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.NumberOfStops) {
		return nil, false
	}
	return o.NumberOfStops, true
}

// HasNumberOfStops returns a boolean if a field has been set.
func (o *Segment) HasNumberOfStops() bool {
	if o != nil && !utils.IsNil(o.NumberOfStops) {
		return true
	}

	return false
}

// SetNumberOfStops gets a reference to the given int32 and assigns it to the NumberOfStops field.
func (o *Segment) SetNumberOfStops(v int32) {
	o.NumberOfStops = &v
}

// GetBlacklistedInEU returns the BlacklistedInEU field value if set, zero value otherwise.
func (o *Segment) GetBlacklistedInEU() bool {
	if o == nil || utils.IsNil(o.BlacklistedInEU) {
		var ret bool
		return ret
	}
	return *o.BlacklistedInEU
}

// GetBlacklistedInEUOk returns a tuple with the BlacklistedInEU field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Segment) GetBlacklistedInEUOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.BlacklistedInEU) {
		return nil, false
	}
	return o.BlacklistedInEU, true
}

// HasBlacklistedInEU returns a boolean if a field has been set.
func (o *Segment) HasBlacklistedInEU() bool {
	if o != nil && !utils.IsNil(o.BlacklistedInEU) {
		return true
	}

	return false
}

// SetBlacklistedInEU gets a reference to the given bool and assigns it to the BlacklistedInEU field.
func (o *Segment) SetBlacklistedInEU(v bool) {
	o.BlacklistedInEU = &v
}

// GetCo2Emissions returns the Co2Emissions field value if set, zero value otherwise.
func (o *Segment) GetCo2Emissions() []Co2Emission {
	if o == nil || utils.IsNil(o.Co2Emissions) {
		var ret []Co2Emission
		return ret
	}
	return o.Co2Emissions
}

// GetCo2EmissionsOk returns a tuple with the Co2Emissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Segment) GetCo2EmissionsOk() ([]Co2Emission, bool) {
	if o == nil || utils.IsNil(o.Co2Emissions) {
		return nil, false
	}
	return o.Co2Emissions, true
}

// HasCo2Emissions returns a boolean if a field has been set.
func (o *Segment) HasCo2Emissions() bool {
	if o != nil && !utils.IsNil(o.Co2Emissions) {
		return true
	}

	return false
}

// SetCo2Emissions gets a reference to the given []Co2Emission and assigns it to the Co2Emissions field.
func (o *Segment) SetCo2Emissions(v []Co2Emission) {
	o.Co2Emissions = v
}

func (o Segment) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Segment) ToMap() (map[string]interface{}, error) {
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
	if !utils.IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !utils.IsNil(o.NumberOfStops) {
		toSerialize["numberOfStops"] = o.NumberOfStops
	}
	if !utils.IsNil(o.BlacklistedInEU) {
		toSerialize["blacklistedInEU"] = o.BlacklistedInEU
	}
	if !utils.IsNil(o.Co2Emissions) {
		toSerialize["co2Emissions"] = o.Co2Emissions
	}
	return toSerialize, nil
}

type NullableSegment struct {
	value *Segment
	isSet bool
}

func (v NullableSegment) Get() *Segment {
	return v.value
}

func (v *NullableSegment) Set(val *Segment) {
	v.value = val
	v.isSet = true
}

func (v NullableSegment) IsSet() bool {
	return v.isSet
}

func (v *NullableSegment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSegment(val *Segment) *NullableSegment {
	return &NullableSegment{value: val, isSet: true}
}

func (v NullableSegment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSegment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
