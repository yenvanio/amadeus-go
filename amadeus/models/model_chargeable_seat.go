package models

import (
	"encoding/json"
	"github.com/yenvanio/amadeus-go/amadeus/utils"
)

// checks if the ChargeableSeat type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ChargeableSeat{}

// ChargeableSeat Details of chargeable seat
type ChargeableSeat struct {
	// Id of the chargeable seat
	Id *string `json:"id,omitempty"`
	// seat number
	Number *string `json:"number,omitempty"`
}

// NewChargeableSeat instantiates a new ChargeableSeat object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChargeableSeat() *ChargeableSeat {
	this := ChargeableSeat{}
	return &this
}

// NewChargeableSeatWithDefaults instantiates a new ChargeableSeat object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChargeableSeatWithDefaults() *ChargeableSeat {
	this := ChargeableSeat{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ChargeableSeat) GetId() string {
	if o == nil || utils.IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargeableSeat) GetIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ChargeableSeat) HasId() bool {
	if o != nil && !utils.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ChargeableSeat) SetId(v string) {
	o.Id = &v
}

// GetNumber returns the Number field value if set, zero value otherwise.
func (o *ChargeableSeat) GetNumber() string {
	if o == nil || utils.IsNil(o.Number) {
		var ret string
		return ret
	}
	return *o.Number
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargeableSeat) GetNumberOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Number) {
		return nil, false
	}
	return o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *ChargeableSeat) HasNumber() bool {
	if o != nil && !utils.IsNil(o.Number) {
		return true
	}

	return false
}

// SetNumber gets a reference to the given string and assigns it to the Number field.
func (o *ChargeableSeat) SetNumber(v string) {
	o.Number = &v
}

func (o ChargeableSeat) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChargeableSeat) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !utils.IsNil(o.Number) {
		toSerialize["number"] = o.Number
	}
	return toSerialize, nil
}

type NullableChargeableSeat struct {
	value *ChargeableSeat
	isSet bool
}

func (v NullableChargeableSeat) Get() *ChargeableSeat {
	return v.value
}

func (v *NullableChargeableSeat) Set(val *ChargeableSeat) {
	v.value = val
	v.isSet = true
}

func (v NullableChargeableSeat) IsSet() bool {
	return v.isSet
}

func (v *NullableChargeableSeat) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChargeableSeat(val *ChargeableSeat) *NullableChargeableSeat {
	return &NullableChargeableSeat{value: val, isSet: true}
}

func (v NullableChargeableSeat) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChargeableSeat) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
