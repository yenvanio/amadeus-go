package models

import (
	"encoding/json"
	"github.com/yenvanio/amadeus-go/amadeus/utils"
)

// checks if the AdditionalService type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &AdditionalService{}

// AdditionalService struct for AdditionalService
type AdditionalService struct {
	Amount *string                `json:"amount,omitempty"`
	Type   *AdditionalServiceType `json:"type,omitempty"`
}

// NewAdditionalService instantiates a new AdditionalService object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdditionalService() *AdditionalService {
	this := AdditionalService{}
	return &this
}

// NewAdditionalServiceWithDefaults instantiates a new AdditionalService object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdditionalServiceWithDefaults() *AdditionalService {
	this := AdditionalService{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *AdditionalService) GetAmount() string {
	if o == nil || utils.IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalService) GetAmountOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *AdditionalService) HasAmount() bool {
	if o != nil && !utils.IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *AdditionalService) SetAmount(v string) {
	o.Amount = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AdditionalService) GetType() AdditionalServiceType {
	if o == nil || utils.IsNil(o.Type) {
		var ret AdditionalServiceType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalService) GetTypeOk() (*AdditionalServiceType, bool) {
	if o == nil || utils.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AdditionalService) HasType() bool {
	if o != nil && !utils.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given AdditionalServiceType and assigns it to the Type field.
func (o *AdditionalService) SetType(v AdditionalServiceType) {
	o.Type = &v
}

func (o AdditionalService) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdditionalService) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !utils.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableAdditionalService struct {
	value *AdditionalService
	isSet bool
}

func (v NullableAdditionalService) Get() *AdditionalService {
	return v.value
}

func (v *NullableAdditionalService) Set(val *AdditionalService) {
	v.value = val
	v.isSet = true
}

func (v NullableAdditionalService) IsSet() bool {
	return v.isSet
}

func (v *NullableAdditionalService) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdditionalService(val *AdditionalService) *NullableAdditionalService {
	return &NullableAdditionalService{value: val, isSet: true}
}

func (v NullableAdditionalService) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdditionalService) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
