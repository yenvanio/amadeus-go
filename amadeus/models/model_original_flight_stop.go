package models

import (
	"encoding/json"
	"github.com/yenvanio/amadeus-go/amadeus/utils"
)

// checks if the OriginalFlightStop type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &OriginalFlightStop{}

// OriginalFlightStop details of stops for direct or change of gauge flights
type OriginalFlightStop struct {
	// [IATA airline codes](http://www.iata.org/publications/Pages/code-search.aspx)
	IataCode *string `json:"iataCode,omitempty"`
	// stop duration in [ISO8601](https://en.wikipedia.org/wiki/ISO_8601) PnYnMnDTnHnMnS format, e.g. PT2H10M
	Duration *string `json:"duration,omitempty"`
}

// NewOriginalFlightStop instantiates a new OriginalFlightStop object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOriginalFlightStop() *OriginalFlightStop {
	this := OriginalFlightStop{}
	return &this
}

// NewOriginalFlightStopWithDefaults instantiates a new OriginalFlightStop object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOriginalFlightStopWithDefaults() *OriginalFlightStop {
	this := OriginalFlightStop{}
	return &this
}

// GetIataCode returns the IataCode field value if set, zero value otherwise.
func (o *OriginalFlightStop) GetIataCode() string {
	if o == nil || utils.IsNil(o.IataCode) {
		var ret string
		return ret
	}
	return *o.IataCode
}

// GetIataCodeOk returns a tuple with the IataCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginalFlightStop) GetIataCodeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.IataCode) {
		return nil, false
	}
	return o.IataCode, true
}

// HasIataCode returns a boolean if a field has been set.
func (o *OriginalFlightStop) HasIataCode() bool {
	if o != nil && !utils.IsNil(o.IataCode) {
		return true
	}

	return false
}

// SetIataCode gets a reference to the given string and assigns it to the IataCode field.
func (o *OriginalFlightStop) SetIataCode(v string) {
	o.IataCode = &v
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *OriginalFlightStop) GetDuration() string {
	if o == nil || utils.IsNil(o.Duration) {
		var ret string
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginalFlightStop) GetDurationOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Duration) {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *OriginalFlightStop) HasDuration() bool {
	if o != nil && !utils.IsNil(o.Duration) {
		return true
	}

	return false
}

// SetDuration gets a reference to the given string and assigns it to the Duration field.
func (o *OriginalFlightStop) SetDuration(v string) {
	o.Duration = &v
}

func (o OriginalFlightStop) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OriginalFlightStop) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.IataCode) {
		toSerialize["iataCode"] = o.IataCode
	}
	if !utils.IsNil(o.Duration) {
		toSerialize["duration"] = o.Duration
	}
	return toSerialize, nil
}

type NullableOriginalFlightStop struct {
	value *OriginalFlightStop
	isSet bool
}

func (v NullableOriginalFlightStop) Get() *OriginalFlightStop {
	return v.value
}

func (v *NullableOriginalFlightStop) Set(val *OriginalFlightStop) {
	v.value = val
	v.isSet = true
}

func (v NullableOriginalFlightStop) IsSet() bool {
	return v.isSet
}

func (v *NullableOriginalFlightStop) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOriginalFlightStop(val *OriginalFlightStop) *NullableOriginalFlightStop {
	return &NullableOriginalFlightStop{value: val, isSet: true}
}

func (v NullableOriginalFlightStop) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOriginalFlightStop) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
