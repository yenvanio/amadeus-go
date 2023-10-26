package models

import (
	"encoding/json"
	"github.com/yenvanio/amadeus-go/amadeus/utils"
)

// checks if the FlightEndPoint type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &FlightEndPoint{}

// FlightEndPoint departure or arrival information
type FlightEndPoint struct {
	// [IATA airline codes](http://www.iata.org/publications/Pages/code-search.aspx)
	IataCode *string `json:"iataCode,omitempty"`
	// terminal name / number
	Terminal *string `json:"terminal,omitempty"`
	// local date and time in [ISO8601](https://en.wikipedia.org/wiki/ISO_8601) YYYY-MM-ddThh:mm:ss format, e.g. 2017-02-10T20:40:00
	At *AmadeusTime `json:"at,omitempty"`
}

// NewFlightEndPoint instantiates a new FlightEndPoint object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFlightEndPoint() *FlightEndPoint {
	this := FlightEndPoint{}
	return &this
}

// NewFlightEndPointWithDefaults instantiates a new FlightEndPoint object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFlightEndPointWithDefaults() *FlightEndPoint {
	this := FlightEndPoint{}
	return &this
}

// GetIataCode returns the IataCode field value if set, zero value otherwise.
func (o *FlightEndPoint) GetIataCode() string {
	if o == nil || utils.IsNil(o.IataCode) {
		var ret string
		return ret
	}
	return *o.IataCode
}

// GetIataCodeOk returns a tuple with the IataCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlightEndPoint) GetIataCodeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.IataCode) {
		return nil, false
	}
	return o.IataCode, true
}

// HasIataCode returns a boolean if a field has been set.
func (o *FlightEndPoint) HasIataCode() bool {
	if o != nil && !utils.IsNil(o.IataCode) {
		return true
	}

	return false
}

// SetIataCode gets a reference to the given string and assigns it to the IataCode field.
func (o *FlightEndPoint) SetIataCode(v string) {
	o.IataCode = &v
}

// GetTerminal returns the Terminal field value if set, zero value otherwise.
func (o *FlightEndPoint) GetTerminal() string {
	if o == nil || utils.IsNil(o.Terminal) {
		var ret string
		return ret
	}
	return *o.Terminal
}

// GetTerminalOk returns a tuple with the Terminal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlightEndPoint) GetTerminalOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Terminal) {
		return nil, false
	}
	return o.Terminal, true
}

// HasTerminal returns a boolean if a field has been set.
func (o *FlightEndPoint) HasTerminal() bool {
	if o != nil && !utils.IsNil(o.Terminal) {
		return true
	}

	return false
}

// SetTerminal gets a reference to the given string and assigns it to the Terminal field.
func (o *FlightEndPoint) SetTerminal(v string) {
	o.Terminal = &v
}

// GetAt returns the At field value if set, zero value otherwise.
func (o *FlightEndPoint) GetAt() AmadeusTime {
	if o == nil || utils.IsNil(o.At) {
		var ret AmadeusTime
		return ret
	}
	return *o.At
}

// GetAtOk returns a tuple with the At field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlightEndPoint) GetAtOk() (*AmadeusTime, bool) {
	if o == nil || utils.IsNil(o.At) {
		return nil, false
	}
	return o.At, true
}

// HasAt returns a boolean if a field has been set.
func (o *FlightEndPoint) HasAt() bool {
	if o != nil && !utils.IsNil(o.At) {
		return true
	}

	return false
}

// SetAt gets a reference to the given AmadeusTime and assigns it to the At field.
func (o *FlightEndPoint) SetAt(v AmadeusTime) {
	o.At = &v
}

func (o FlightEndPoint) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FlightEndPoint) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.IataCode) {
		toSerialize["iataCode"] = o.IataCode
	}
	if !utils.IsNil(o.Terminal) {
		toSerialize["terminal"] = o.Terminal
	}
	if !utils.IsNil(o.At) {
		toSerialize["at"] = o.At
	}
	return toSerialize, nil
}

type NullableFlightEndPoint struct {
	value *FlightEndPoint
	isSet bool
}

func (v NullableFlightEndPoint) Get() *FlightEndPoint {
	return v.value
}

func (v *NullableFlightEndPoint) Set(val *FlightEndPoint) {
	v.value = val
	v.isSet = true
}

func (v NullableFlightEndPoint) IsSet() bool {
	return v.isSet
}

func (v *NullableFlightEndPoint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFlightEndPoint(val *FlightEndPoint) *NullableFlightEndPoint {
	return &NullableFlightEndPoint{value: val, isSet: true}
}

func (v NullableFlightEndPoint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFlightEndPoint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
