package models

import (
	"encoding/json"
	"github.com/yenvanio/amadeus-go/amadeus/utils"
)

// checks if the FlightStop type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &FlightStop{}

// FlightStop details of stops for direct or change of gauge flights
type FlightStop struct {
	// [IATA airline codes](http://www.iata.org/publications/Pages/code-search.aspx)
	IataCode *string `json:"iataCode,omitempty"`
	// stop duration in [ISO8601](https://en.wikipedia.org/wiki/ISO_8601) PnYnMnDTnHnMnS format, e.g. PT2H10M
	Duration *string `json:"duration,omitempty"`
	// arrival at the stop in [ISO8601](https://en.wikipedia.org/wiki/ISO_8601) YYYY-MM-ddThh:mm:ss format, e.g. 2017-02-10T20:40:00
	ArrivalAt *AmadeusTime `json:"arrivalAt,omitempty"`
	// departure from the stop in [ISO8601](https://en.wikipedia.org/wiki/ISO_8601) YYYY-MM-ddThh:mm:ss format, e.g. 2017-02-10T20:40:00
	DepartureAt *AmadeusTime `json:"departureAt,omitempty"`
}

// NewFlightStop instantiates a new FlightStop object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFlightStop() *FlightStop {
	this := FlightStop{}
	return &this
}

// NewFlightStopWithDefaults instantiates a new FlightStop object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFlightStopWithDefaults() *FlightStop {
	this := FlightStop{}
	return &this
}

// GetIataCode returns the IataCode field value if set, zero value otherwise.
func (o *FlightStop) GetIataCode() string {
	if o == nil || utils.IsNil(o.IataCode) {
		var ret string
		return ret
	}
	return *o.IataCode
}

// GetIataCodeOk returns a tuple with the IataCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlightStop) GetIataCodeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.IataCode) {
		return nil, false
	}
	return o.IataCode, true
}

// HasIataCode returns a boolean if a field has been set.
func (o *FlightStop) HasIataCode() bool {
	if o != nil && !utils.IsNil(o.IataCode) {
		return true
	}

	return false
}

// SetIataCode gets a reference to the given string and assigns it to the IataCode field.
func (o *FlightStop) SetIataCode(v string) {
	o.IataCode = &v
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *FlightStop) GetDuration() string {
	if o == nil || utils.IsNil(o.Duration) {
		var ret string
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlightStop) GetDurationOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Duration) {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *FlightStop) HasDuration() bool {
	if o != nil && !utils.IsNil(o.Duration) {
		return true
	}

	return false
}

// SetDuration gets a reference to the given string and assigns it to the Duration field.
func (o *FlightStop) SetDuration(v string) {
	o.Duration = &v
}

// GetArrivalAt returns the ArrivalAt field value if set, zero value otherwise.
func (o *FlightStop) GetArrivalAt() AmadeusTime {
	if o == nil || utils.IsNil(o.ArrivalAt) {
		var ret AmadeusTime
		return ret
	}
	return *o.ArrivalAt
}

// GetArrivalAtOk returns a tuple with the ArrivalAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlightStop) GetArrivalAtOk() (*AmadeusTime, bool) {
	if o == nil || utils.IsNil(o.ArrivalAt) {
		return nil, false
	}
	return o.ArrivalAt, true
}

// HasArrivalAt returns a boolean if a field has been set.
func (o *FlightStop) HasArrivalAt() bool {
	if o != nil && !utils.IsNil(o.ArrivalAt) {
		return true
	}

	return false
}

// SetArrivalAt gets a reference to the given AmadeusTime and assigns it to the ArrivalAt field.
func (o *FlightStop) SetArrivalAt(v AmadeusTime) {
	o.ArrivalAt = &v
}

// GetDepartureAt returns the DepartureAt field value if set, zero value otherwise.
func (o *FlightStop) GetDepartureAt() AmadeusTime {
	if o == nil || utils.IsNil(o.DepartureAt) {
		var ret AmadeusTime
		return ret
	}
	return *o.DepartureAt
}

// GetDepartureAtOk returns a tuple with the DepartureAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlightStop) GetDepartureAtOk() (*AmadeusTime, bool) {
	if o == nil || utils.IsNil(o.DepartureAt) {
		return nil, false
	}
	return o.DepartureAt, true
}

// HasDepartureAt returns a boolean if a field has been set.
func (o *FlightStop) HasDepartureAt() bool {
	if o != nil && !utils.IsNil(o.DepartureAt) {
		return true
	}

	return false
}

// SetDepartureAt gets a reference to the given time.Time and assigns it to the DepartureAt field.
func (o *FlightStop) SetDepartureAt(v AmadeusTime) {
	o.DepartureAt = &v
}

func (o FlightStop) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FlightStop) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.IataCode) {
		toSerialize["iataCode"] = o.IataCode
	}
	if !utils.IsNil(o.Duration) {
		toSerialize["duration"] = o.Duration
	}
	if !utils.IsNil(o.ArrivalAt) {
		toSerialize["arrivalAt"] = o.ArrivalAt
	}
	if !utils.IsNil(o.DepartureAt) {
		toSerialize["departureAt"] = o.DepartureAt
	}
	return toSerialize, nil
}

type NullableFlightStop struct {
	value *FlightStop
	isSet bool
}

func (v NullableFlightStop) Get() *FlightStop {
	return v.value
}

func (v *NullableFlightStop) Set(val *FlightStop) {
	v.value = val
	v.isSet = true
}

func (v NullableFlightStop) IsSet() bool {
	return v.isSet
}

func (v *NullableFlightStop) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFlightStop(val *FlightStop) *NullableFlightStop {
	return &NullableFlightStop{value: val, isSet: true}
}

func (v NullableFlightStop) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFlightStop) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
