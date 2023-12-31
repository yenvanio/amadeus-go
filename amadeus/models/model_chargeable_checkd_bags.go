package models

import (
	"encoding/json"
	"github.com/yenvanio/amadeus-go/amadeus/utils"
)

// checks if the ChargeableCheckdBags type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ChargeableCheckdBags{}

// ChargeableCheckdBags Details of chargeable checked bags
type ChargeableCheckdBags struct {
	// Total number of units
	Quantity *int32 `json:"quantity,omitempty"`
	// Weight of the baggage allowance
	Weight *int32 `json:"weight,omitempty"`
	// Code to qualify unit as pounds or kilos
	WeightUnit *string `json:"weightUnit,omitempty"`
	// Id of the chargeable bag
	Id *string `json:"id,omitempty"`
}

// NewChargeableCheckdBags instantiates a new ChargeableCheckdBags object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChargeableCheckdBags() *ChargeableCheckdBags {
	this := ChargeableCheckdBags{}
	return &this
}

// NewChargeableCheckdBagsWithDefaults instantiates a new ChargeableCheckdBags object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChargeableCheckdBagsWithDefaults() *ChargeableCheckdBags {
	this := ChargeableCheckdBags{}
	return &this
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *ChargeableCheckdBags) GetQuantity() int32 {
	if o == nil || utils.IsNil(o.Quantity) {
		var ret int32
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargeableCheckdBags) GetQuantityOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *ChargeableCheckdBags) HasQuantity() bool {
	if o != nil && !utils.IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given int32 and assigns it to the Quantity field.
func (o *ChargeableCheckdBags) SetQuantity(v int32) {
	o.Quantity = &v
}

// GetWeight returns the Weight field value if set, zero value otherwise.
func (o *ChargeableCheckdBags) GetWeight() int32 {
	if o == nil || utils.IsNil(o.Weight) {
		var ret int32
		return ret
	}
	return *o.Weight
}

// GetWeightOk returns a tuple with the Weight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargeableCheckdBags) GetWeightOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.Weight) {
		return nil, false
	}
	return o.Weight, true
}

// HasWeight returns a boolean if a field has been set.
func (o *ChargeableCheckdBags) HasWeight() bool {
	if o != nil && !utils.IsNil(o.Weight) {
		return true
	}

	return false
}

// SetWeight gets a reference to the given int32 and assigns it to the Weight field.
func (o *ChargeableCheckdBags) SetWeight(v int32) {
	o.Weight = &v
}

// GetWeightUnit returns the WeightUnit field value if set, zero value otherwise.
func (o *ChargeableCheckdBags) GetWeightUnit() string {
	if o == nil || utils.IsNil(o.WeightUnit) {
		var ret string
		return ret
	}
	return *o.WeightUnit
}

// GetWeightUnitOk returns a tuple with the WeightUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargeableCheckdBags) GetWeightUnitOk() (*string, bool) {
	if o == nil || utils.IsNil(o.WeightUnit) {
		return nil, false
	}
	return o.WeightUnit, true
}

// HasWeightUnit returns a boolean if a field has been set.
func (o *ChargeableCheckdBags) HasWeightUnit() bool {
	if o != nil && !utils.IsNil(o.WeightUnit) {
		return true
	}

	return false
}

// SetWeightUnit gets a reference to the given string and assigns it to the WeightUnit field.
func (o *ChargeableCheckdBags) SetWeightUnit(v string) {
	o.WeightUnit = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ChargeableCheckdBags) GetId() string {
	if o == nil || utils.IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargeableCheckdBags) GetIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ChargeableCheckdBags) HasId() bool {
	if o != nil && !utils.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ChargeableCheckdBags) SetId(v string) {
	o.Id = &v
}

func (o ChargeableCheckdBags) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChargeableCheckdBags) ToMap() (map[string]interface{}, error) {
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
	if !utils.IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullableChargeableCheckdBags struct {
	value *ChargeableCheckdBags
	isSet bool
}

func (v NullableChargeableCheckdBags) Get() *ChargeableCheckdBags {
	return v.value
}

func (v *NullableChargeableCheckdBags) Set(val *ChargeableCheckdBags) {
	v.value = val
	v.isSet = true
}

func (v NullableChargeableCheckdBags) IsSet() bool {
	return v.isSet
}

func (v *NullableChargeableCheckdBags) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChargeableCheckdBags(val *ChargeableCheckdBags) *NullableChargeableCheckdBags {
	return &NullableChargeableCheckdBags{value: val, isSet: true}
}

func (v NullableChargeableCheckdBags) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChargeableCheckdBags) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
