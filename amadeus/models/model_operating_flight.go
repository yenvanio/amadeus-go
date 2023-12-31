package models

import (
	"encoding/json"
	"github.com/yenvanio/amadeus-go/amadeus/utils"
)

// checks if the OperatingFlight type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &OperatingFlight{}

// OperatingFlight information about the operating flight
type OperatingFlight struct {
	// providing the airline / carrier code
	CarrierCode *string `json:"carrierCode,omitempty"`
}

// NewOperatingFlight instantiates a new OperatingFlight object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOperatingFlight() *OperatingFlight {
	this := OperatingFlight{}
	return &this
}

// NewOperatingFlightWithDefaults instantiates a new OperatingFlight object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOperatingFlightWithDefaults() *OperatingFlight {
	this := OperatingFlight{}
	return &this
}

// GetCarrierCode returns the CarrierCode field value if set, zero value otherwise.
func (o *OperatingFlight) GetCarrierCode() string {
	if o == nil || utils.IsNil(o.CarrierCode) {
		var ret string
		return ret
	}
	return *o.CarrierCode
}

// GetCarrierCodeOk returns a tuple with the CarrierCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperatingFlight) GetCarrierCodeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.CarrierCode) {
		return nil, false
	}
	return o.CarrierCode, true
}

// HasCarrierCode returns a boolean if a field has been set.
func (o *OperatingFlight) HasCarrierCode() bool {
	if o != nil && !utils.IsNil(o.CarrierCode) {
		return true
	}

	return false
}

// SetCarrierCode gets a reference to the given string and assigns it to the CarrierCode field.
func (o *OperatingFlight) SetCarrierCode(v string) {
	o.CarrierCode = &v
}

func (o OperatingFlight) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OperatingFlight) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.CarrierCode) {
		toSerialize["carrierCode"] = o.CarrierCode
	}
	return toSerialize, nil
}

type NullableOperatingFlight struct {
	value *OperatingFlight
	isSet bool
}

func (v NullableOperatingFlight) Get() *OperatingFlight {
	return v.value
}

func (v *NullableOperatingFlight) Set(val *OperatingFlight) {
	v.value = val
	v.isSet = true
}

func (v NullableOperatingFlight) IsSet() bool {
	return v.isSet
}

func (v *NullableOperatingFlight) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOperatingFlight(val *OperatingFlight) *NullableOperatingFlight {
	return &NullableOperatingFlight{value: val, isSet: true}
}

func (v NullableOperatingFlight) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOperatingFlight) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
