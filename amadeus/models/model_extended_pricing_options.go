package models

import (
	"encoding/json"
	"github.com/yenvanio/amadeus-go/amadeus/utils"
)

// checks if the ExtendedPricingOptions type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ExtendedPricingOptions{}

// ExtendedPricingOptions fare filter options
type ExtendedPricingOptions struct {
	// If true, returns the flight-offers with included checked bags only
	IncludedCheckedBagsOnly *bool `json:"includedCheckedBagsOnly,omitempty"`
	// If true, returns the flight-offers with refundable fares only
	RefundableFare *bool `json:"refundableFare,omitempty"`
	// If true, returns the flight-offers with no restriction fares only
	NoRestrictionFare *bool `json:"noRestrictionFare,omitempty"`
	// If true, returns the flight-offers with no penalty fares only
	NoPenaltyFare *bool `json:"noPenaltyFare,omitempty"`
}

// NewExtendedPricingOptions instantiates a new ExtendedPricingOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExtendedPricingOptions() *ExtendedPricingOptions {
	this := ExtendedPricingOptions{}
	return &this
}

// NewExtendedPricingOptionsWithDefaults instantiates a new ExtendedPricingOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExtendedPricingOptionsWithDefaults() *ExtendedPricingOptions {
	this := ExtendedPricingOptions{}
	return &this
}

// GetIncludedCheckedBagsOnly returns the IncludedCheckedBagsOnly field value if set, zero value otherwise.
func (o *ExtendedPricingOptions) GetIncludedCheckedBagsOnly() bool {
	if o == nil || utils.IsNil(o.IncludedCheckedBagsOnly) {
		var ret bool
		return ret
	}
	return *o.IncludedCheckedBagsOnly
}

// GetIncludedCheckedBagsOnlyOk returns a tuple with the IncludedCheckedBagsOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtendedPricingOptions) GetIncludedCheckedBagsOnlyOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.IncludedCheckedBagsOnly) {
		return nil, false
	}
	return o.IncludedCheckedBagsOnly, true
}

// HasIncludedCheckedBagsOnly returns a boolean if a field has been set.
func (o *ExtendedPricingOptions) HasIncludedCheckedBagsOnly() bool {
	if o != nil && !utils.IsNil(o.IncludedCheckedBagsOnly) {
		return true
	}

	return false
}

// SetIncludedCheckedBagsOnly gets a reference to the given bool and assigns it to the IncludedCheckedBagsOnly field.
func (o *ExtendedPricingOptions) SetIncludedCheckedBagsOnly(v bool) {
	o.IncludedCheckedBagsOnly = &v
}

// GetRefundableFare returns the RefundableFare field value if set, zero value otherwise.
func (o *ExtendedPricingOptions) GetRefundableFare() bool {
	if o == nil || utils.IsNil(o.RefundableFare) {
		var ret bool
		return ret
	}
	return *o.RefundableFare
}

// GetRefundableFareOk returns a tuple with the RefundableFare field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtendedPricingOptions) GetRefundableFareOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.RefundableFare) {
		return nil, false
	}
	return o.RefundableFare, true
}

// HasRefundableFare returns a boolean if a field has been set.
func (o *ExtendedPricingOptions) HasRefundableFare() bool {
	if o != nil && !utils.IsNil(o.RefundableFare) {
		return true
	}

	return false
}

// SetRefundableFare gets a reference to the given bool and assigns it to the RefundableFare field.
func (o *ExtendedPricingOptions) SetRefundableFare(v bool) {
	o.RefundableFare = &v
}

// GetNoRestrictionFare returns the NoRestrictionFare field value if set, zero value otherwise.
func (o *ExtendedPricingOptions) GetNoRestrictionFare() bool {
	if o == nil || utils.IsNil(o.NoRestrictionFare) {
		var ret bool
		return ret
	}
	return *o.NoRestrictionFare
}

// GetNoRestrictionFareOk returns a tuple with the NoRestrictionFare field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtendedPricingOptions) GetNoRestrictionFareOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.NoRestrictionFare) {
		return nil, false
	}
	return o.NoRestrictionFare, true
}

// HasNoRestrictionFare returns a boolean if a field has been set.
func (o *ExtendedPricingOptions) HasNoRestrictionFare() bool {
	if o != nil && !utils.IsNil(o.NoRestrictionFare) {
		return true
	}

	return false
}

// SetNoRestrictionFare gets a reference to the given bool and assigns it to the NoRestrictionFare field.
func (o *ExtendedPricingOptions) SetNoRestrictionFare(v bool) {
	o.NoRestrictionFare = &v
}

// GetNoPenaltyFare returns the NoPenaltyFare field value if set, zero value otherwise.
func (o *ExtendedPricingOptions) GetNoPenaltyFare() bool {
	if o == nil || utils.IsNil(o.NoPenaltyFare) {
		var ret bool
		return ret
	}
	return *o.NoPenaltyFare
}

// GetNoPenaltyFareOk returns a tuple with the NoPenaltyFare field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtendedPricingOptions) GetNoPenaltyFareOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.NoPenaltyFare) {
		return nil, false
	}
	return o.NoPenaltyFare, true
}

// HasNoPenaltyFare returns a boolean if a field has been set.
func (o *ExtendedPricingOptions) HasNoPenaltyFare() bool {
	if o != nil && !utils.IsNil(o.NoPenaltyFare) {
		return true
	}

	return false
}

// SetNoPenaltyFare gets a reference to the given bool and assigns it to the NoPenaltyFare field.
func (o *ExtendedPricingOptions) SetNoPenaltyFare(v bool) {
	o.NoPenaltyFare = &v
}

func (o ExtendedPricingOptions) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExtendedPricingOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.IncludedCheckedBagsOnly) {
		toSerialize["includedCheckedBagsOnly"] = o.IncludedCheckedBagsOnly
	}
	if !utils.IsNil(o.RefundableFare) {
		toSerialize["refundableFare"] = o.RefundableFare
	}
	if !utils.IsNil(o.NoRestrictionFare) {
		toSerialize["noRestrictionFare"] = o.NoRestrictionFare
	}
	if !utils.IsNil(o.NoPenaltyFare) {
		toSerialize["noPenaltyFare"] = o.NoPenaltyFare
	}
	return toSerialize, nil
}

type NullableExtendedPricingOptions struct {
	value *ExtendedPricingOptions
	isSet bool
}

func (v NullableExtendedPricingOptions) Get() *ExtendedPricingOptions {
	return v.value
}

func (v *NullableExtendedPricingOptions) Set(val *ExtendedPricingOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableExtendedPricingOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableExtendedPricingOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExtendedPricingOptions(val *ExtendedPricingOptions) *NullableExtendedPricingOptions {
	return &NullableExtendedPricingOptions{value: val, isSet: true}
}

func (v NullableExtendedPricingOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExtendedPricingOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
