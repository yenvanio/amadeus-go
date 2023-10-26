package models

import (
	"encoding/json"
	"github.com/yenvanio/amadeus-go/amadeus/utils"
)

// checks if the Co2Emission type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Co2Emission{}

// Co2Emission struct for Co2Emission
type Co2Emission struct {
	// Weight of Co2 emitted for the concerned segment
	Weight *int32 `json:"weight,omitempty"`
	// Code to qualify unit as pounds or kilos
	WeightUnit *string      `json:"weightUnit,omitempty"`
	Cabin      *TravelClass `json:"cabin,omitempty"`
}

// NewCo2Emission instantiates a new Co2Emission object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCo2Emission() *Co2Emission {
	this := Co2Emission{}
	return &this
}

// NewCo2EmissionWithDefaults instantiates a new Co2Emission object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCo2EmissionWithDefaults() *Co2Emission {
	this := Co2Emission{}
	return &this
}

// GetWeight returns the Weight field value if set, zero value otherwise.
func (o *Co2Emission) GetWeight() int32 {
	if o == nil || utils.IsNil(o.Weight) {
		var ret int32
		return ret
	}
	return *o.Weight
}

// GetWeightOk returns a tuple with the Weight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Co2Emission) GetWeightOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.Weight) {
		return nil, false
	}
	return o.Weight, true
}

// HasWeight returns a boolean if a field has been set.
func (o *Co2Emission) HasWeight() bool {
	if o != nil && !utils.IsNil(o.Weight) {
		return true
	}

	return false
}

// SetWeight gets a reference to the given int32 and assigns it to the Weight field.
func (o *Co2Emission) SetWeight(v int32) {
	o.Weight = &v
}

// GetWeightUnit returns the WeightUnit field value if set, zero value otherwise.
func (o *Co2Emission) GetWeightUnit() string {
	if o == nil || utils.IsNil(o.WeightUnit) {
		var ret string
		return ret
	}
	return *o.WeightUnit
}

// GetWeightUnitOk returns a tuple with the WeightUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Co2Emission) GetWeightUnitOk() (*string, bool) {
	if o == nil || utils.IsNil(o.WeightUnit) {
		return nil, false
	}
	return o.WeightUnit, true
}

// HasWeightUnit returns a boolean if a field has been set.
func (o *Co2Emission) HasWeightUnit() bool {
	if o != nil && !utils.IsNil(o.WeightUnit) {
		return true
	}

	return false
}

// SetWeightUnit gets a reference to the given string and assigns it to the WeightUnit field.
func (o *Co2Emission) SetWeightUnit(v string) {
	o.WeightUnit = &v
}

// GetCabin returns the Cabin field value if set, zero value otherwise.
func (o *Co2Emission) GetCabin() TravelClass {
	if o == nil || utils.IsNil(o.Cabin) {
		var ret TravelClass
		return ret
	}
	return *o.Cabin
}

// GetCabinOk returns a tuple with the Cabin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Co2Emission) GetCabinOk() (*TravelClass, bool) {
	if o == nil || utils.IsNil(o.Cabin) {
		return nil, false
	}
	return o.Cabin, true
}

// HasCabin returns a boolean if a field has been set.
func (o *Co2Emission) HasCabin() bool {
	if o != nil && !utils.IsNil(o.Cabin) {
		return true
	}

	return false
}

// SetCabin gets a reference to the given TravelClass and assigns it to the Cabin field.
func (o *Co2Emission) SetCabin(v TravelClass) {
	o.Cabin = &v
}

func (o Co2Emission) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Co2Emission) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Weight) {
		toSerialize["weight"] = o.Weight
	}
	if !utils.IsNil(o.WeightUnit) {
		toSerialize["weightUnit"] = o.WeightUnit
	}
	if !utils.IsNil(o.Cabin) {
		toSerialize["cabin"] = o.Cabin
	}
	return toSerialize, nil
}

type NullableCo2Emission struct {
	value *Co2Emission
	isSet bool
}

func (v NullableCo2Emission) Get() *Co2Emission {
	return v.value
}

func (v *NullableCo2Emission) Set(val *Co2Emission) {
	v.value = val
	v.isSet = true
}

func (v NullableCo2Emission) IsSet() bool {
	return v.isSet
}

func (v *NullableCo2Emission) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCo2Emission(val *Co2Emission) *NullableCo2Emission {
	return &NullableCo2Emission{value: val, isSet: true}
}

func (v NullableCo2Emission) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCo2Emission) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
