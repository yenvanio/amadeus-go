package models

import (
	"encoding/json"
	"github.com/yenvanio/amadeus-go/amadeus/utils"
)

// checks if the BaggageAllowance type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &BaggageAllowance{}

// BaggageAllowance baggageAllowance
type BaggageAllowance struct {
	// Total number of units
	Quantity *int32 `json:"quantity,omitempty"`
	// Weight of the baggage allowance
	Weight *int32 `json:"weight,omitempty"`
	// Code to qualify unit as pounds or kilos
	WeightUnit *string `json:"weightUnit,omitempty"`
}

// NewBaggageAllowance instantiates a new BaggageAllowance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaggageAllowance() *BaggageAllowance {
	this := BaggageAllowance{}
	return &this
}

// NewBaggageAllowanceWithDefaults instantiates a new BaggageAllowance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaggageAllowanceWithDefaults() *BaggageAllowance {
	this := BaggageAllowance{}
	return &this
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *BaggageAllowance) GetQuantity() int32 {
	if o == nil || utils.IsNil(o.Quantity) {
		var ret int32
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaggageAllowance) GetQuantityOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *BaggageAllowance) HasQuantity() bool {
	if o != nil && !utils.IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given int32 and assigns it to the Quantity field.
func (o *BaggageAllowance) SetQuantity(v int32) {
	o.Quantity = &v
}

// GetWeight returns the Weight field value if set, zero value otherwise.
func (o *BaggageAllowance) GetWeight() int32 {
	if o == nil || utils.IsNil(o.Weight) {
		var ret int32
		return ret
	}
	return *o.Weight
}

// GetWeightOk returns a tuple with the Weight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaggageAllowance) GetWeightOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.Weight) {
		return nil, false
	}
	return o.Weight, true
}

// HasWeight returns a boolean if a field has been set.
func (o *BaggageAllowance) HasWeight() bool {
	if o != nil && !utils.IsNil(o.Weight) {
		return true
	}

	return false
}

// SetWeight gets a reference to the given int32 and assigns it to the Weight field.
func (o *BaggageAllowance) SetWeight(v int32) {
	o.Weight = &v
}

// GetWeightUnit returns the WeightUnit field value if set, zero value otherwise.
func (o *BaggageAllowance) GetWeightUnit() string {
	if o == nil || utils.IsNil(o.WeightUnit) {
		var ret string
		return ret
	}
	return *o.WeightUnit
}

// GetWeightUnitOk returns a tuple with the WeightUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaggageAllowance) GetWeightUnitOk() (*string, bool) {
	if o == nil || utils.IsNil(o.WeightUnit) {
		return nil, false
	}
	return o.WeightUnit, true
}

// HasWeightUnit returns a boolean if a field has been set.
func (o *BaggageAllowance) HasWeightUnit() bool {
	if o != nil && !utils.IsNil(o.WeightUnit) {
		return true
	}

	return false
}

// SetWeightUnit gets a reference to the given string and assigns it to the WeightUnit field.
func (o *BaggageAllowance) SetWeightUnit(v string) {
	o.WeightUnit = &v
}

func (o BaggageAllowance) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BaggageAllowance) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Quantity) {
		toSerialize["quantity"] = o.Quantity
	}
	if !utils.IsNil(o.Weight) {
		toSerialize["weight"] = o.Weight
	}
	if !utils.IsNil(o.WeightUnit) {
		toSerialize["weightUnit"] = o.WeightUnit
	}
	return toSerialize, nil
}

type NullableBaggageAllowance struct {
	value *BaggageAllowance
	isSet bool
}

func (v NullableBaggageAllowance) Get() *BaggageAllowance {
	return v.value
}

func (v *NullableBaggageAllowance) Set(val *BaggageAllowance) {
	v.value = val
	v.isSet = true
}

func (v NullableBaggageAllowance) IsSet() bool {
	return v.isSet
}

func (v *NullableBaggageAllowance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaggageAllowance(val *BaggageAllowance) *NullableBaggageAllowance {
	return &NullableBaggageAllowance{value: val, isSet: true}
}

func (v NullableBaggageAllowance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaggageAllowance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
