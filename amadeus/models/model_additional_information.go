package models

import (
	"encoding/json"
	"github.com/yenvanio/amadeus-go/amadeus/utils"
)

// checks if the AdditionalInformation type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &AdditionalInformation{}

// AdditionalInformation struct for AdditionalInformation
type AdditionalInformation struct {
	// If true, returns the price of the first additional bag when the airline is an \"Amadeus Ancillary Services\" member.
	ChargeableCheckedBags *bool `json:"chargeableCheckedBags,omitempty"`
	// If true, returns the fare family name for each flight-offer which supports fare family
	BrandedFares *bool `json:"brandedFares,omitempty"`
}

// NewAdditionalInformation instantiates a new AdditionalInformation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdditionalInformation() *AdditionalInformation {
	this := AdditionalInformation{}
	return &this
}

// NewAdditionalInformationWithDefaults instantiates a new AdditionalInformation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdditionalInformationWithDefaults() *AdditionalInformation {
	this := AdditionalInformation{}
	return &this
}

// GetChargeableCheckedBags returns the ChargeableCheckedBags field value if set, zero value otherwise.
func (o *AdditionalInformation) GetChargeableCheckedBags() bool {
	if o == nil || utils.IsNil(o.ChargeableCheckedBags) {
		var ret bool
		return ret
	}
	return *o.ChargeableCheckedBags
}

// GetChargeableCheckedBagsOk returns a tuple with the ChargeableCheckedBags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalInformation) GetChargeableCheckedBagsOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.ChargeableCheckedBags) {
		return nil, false
	}
	return o.ChargeableCheckedBags, true
}

// HasChargeableCheckedBags returns a boolean if a field has been set.
func (o *AdditionalInformation) HasChargeableCheckedBags() bool {
	if o != nil && !utils.IsNil(o.ChargeableCheckedBags) {
		return true
	}

	return false
}

// SetChargeableCheckedBags gets a reference to the given bool and assigns it to the ChargeableCheckedBags field.
func (o *AdditionalInformation) SetChargeableCheckedBags(v bool) {
	o.ChargeableCheckedBags = &v
}

// GetBrandedFares returns the BrandedFares field value if set, zero value otherwise.
func (o *AdditionalInformation) GetBrandedFares() bool {
	if o == nil || utils.IsNil(o.BrandedFares) {
		var ret bool
		return ret
	}
	return *o.BrandedFares
}

// GetBrandedFaresOk returns a tuple with the BrandedFares field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalInformation) GetBrandedFaresOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.BrandedFares) {
		return nil, false
	}
	return o.BrandedFares, true
}

// HasBrandedFares returns a boolean if a field has been set.
func (o *AdditionalInformation) HasBrandedFares() bool {
	if o != nil && !utils.IsNil(o.BrandedFares) {
		return true
	}

	return false
}

// SetBrandedFares gets a reference to the given bool and assigns it to the BrandedFares field.
func (o *AdditionalInformation) SetBrandedFares(v bool) {
	o.BrandedFares = &v
}

func (o AdditionalInformation) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdditionalInformation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.ChargeableCheckedBags) {
		toSerialize["chargeableCheckedBags"] = o.ChargeableCheckedBags
	}
	if !utils.IsNil(o.BrandedFares) {
		toSerialize["brandedFares"] = o.BrandedFares
	}
	return toSerialize, nil
}

type NullableAdditionalInformation struct {
	value *AdditionalInformation
	isSet bool
}

func (v NullableAdditionalInformation) Get() *AdditionalInformation {
	return v.value
}

func (v *NullableAdditionalInformation) Set(val *AdditionalInformation) {
	v.value = val
	v.isSet = true
}

func (v NullableAdditionalInformation) IsSet() bool {
	return v.isSet
}

func (v *NullableAdditionalInformation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdditionalInformation(val *AdditionalInformation) *NullableAdditionalInformation {
	return &NullableAdditionalInformation{value: val, isSet: true}
}

func (v NullableAdditionalInformation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdditionalInformation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
