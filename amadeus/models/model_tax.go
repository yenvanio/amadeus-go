package models

import (
	"encoding/json"
	"github.com/yenvanio/amadeus-go/amadeus/utils"
)

// checks if the Tax type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Tax{}

// Tax a tax
type Tax struct {
	Amount *string `json:"amount,omitempty"`
	Code   *string `json:"code,omitempty"`
}

// NewTax instantiates a new Tax object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTax() *Tax {
	this := Tax{}
	return &this
}

// NewTaxWithDefaults instantiates a new Tax object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaxWithDefaults() *Tax {
	this := Tax{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *Tax) GetAmount() string {
	if o == nil || utils.IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tax) GetAmountOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *Tax) HasAmount() bool {
	if o != nil && !utils.IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *Tax) SetAmount(v string) {
	o.Amount = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *Tax) GetCode() string {
	if o == nil || utils.IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tax) GetCodeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *Tax) HasCode() bool {
	if o != nil && !utils.IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *Tax) SetCode(v string) {
	o.Code = &v
}

func (o Tax) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Tax) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !utils.IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	return toSerialize, nil
}

type NullableTax struct {
	value *Tax
	isSet bool
}

func (v NullableTax) Get() *Tax {
	return v.value
}

func (v *NullableTax) Set(val *Tax) {
	v.value = val
	v.isSet = true
}

func (v NullableTax) IsSet() bool {
	return v.isSet
}

func (v *NullableTax) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTax(val *Tax) *NullableTax {
	return &NullableTax{value: val, isSet: true}
}

func (v NullableTax) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTax) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
