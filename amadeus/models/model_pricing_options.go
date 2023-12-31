package models

import (
	"encoding/json"
	"github.com/yenvanio/amadeus-go/amadeus/utils"
)

// checks if the PricingOptions type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &PricingOptions{}

// PricingOptions struct for PricingOptions
type PricingOptions struct {
	// type of fare of the flight-offer
	FareType []string `json:"fareType,omitempty"`
	// If true, returns the flight-offers with included checked bags only
	IncludedCheckedBagsOnly *bool `json:"includedCheckedBagsOnly,omitempty"`
	// If true, returns the flight-offers with refundable fares only
	RefundableFare *bool `json:"refundableFare,omitempty"`
	// If true, returns the flight-offers with no restriction fares only
	NoRestrictionFare *bool `json:"noRestrictionFare,omitempty"`
	// If true, returns the flight-offers with no penalty fares only
	NoPenaltyFare *bool `json:"noPenaltyFare,omitempty"`
}

// NewPricingOptions instantiates a new PricingOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPricingOptions() *PricingOptions {
	this := PricingOptions{}
	return &this
}

// NewPricingOptionsWithDefaults instantiates a new PricingOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPricingOptionsWithDefaults() *PricingOptions {
	this := PricingOptions{}
	return &this
}

// GetFareType returns the FareType field value if set, zero value otherwise.
func (o *PricingOptions) GetFareType() []string {
	if o == nil || utils.IsNil(o.FareType) {
		var ret []string
		return ret
	}
	return o.FareType
}

// GetFareTypeOk returns a tuple with the FareType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PricingOptions) GetFareTypeOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.FareType) {
		return nil, false
	}
	return o.FareType, true
}

// HasFareType returns a boolean if a field has been set.
func (o *PricingOptions) HasFareType() bool {
	if o != nil && !utils.IsNil(o.FareType) {
		return true
	}

	return false
}

// SetFareType gets a reference to the given []string and assigns it to the FareType field.
func (o *PricingOptions) SetFareType(v []string) {
	o.FareType = v
}

// GetIncludedCheckedBagsOnly returns the IncludedCheckedBagsOnly field value if set, zero value otherwise.
func (o *PricingOptions) GetIncludedCheckedBagsOnly() bool {
	if o == nil || utils.IsNil(o.IncludedCheckedBagsOnly) {
		var ret bool
		return ret
	}
	return *o.IncludedCheckedBagsOnly
}

// GetIncludedCheckedBagsOnlyOk returns a tuple with the IncludedCheckedBagsOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PricingOptions) GetIncludedCheckedBagsOnlyOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.IncludedCheckedBagsOnly) {
		return nil, false
	}
	return o.IncludedCheckedBagsOnly, true
}

// HasIncludedCheckedBagsOnly returns a boolean if a field has been set.
func (o *PricingOptions) HasIncludedCheckedBagsOnly() bool {
	if o != nil && !utils.IsNil(o.IncludedCheckedBagsOnly) {
		return true
	}

	return false
}

// SetIncludedCheckedBagsOnly gets a reference to the given bool and assigns it to the IncludedCheckedBagsOnly field.
func (o *PricingOptions) SetIncludedCheckedBagsOnly(v bool) {
	o.IncludedCheckedBagsOnly = &v
}

// GetRefundableFare returns the RefundableFare field value if set, zero value otherwise.
func (o *PricingOptions) GetRefundableFare() bool {
	if o == nil || utils.IsNil(o.RefundableFare) {
		var ret bool
		return ret
	}
	return *o.RefundableFare
}

// GetRefundableFareOk returns a tuple with the RefundableFare field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PricingOptions) GetRefundableFareOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.RefundableFare) {
		return nil, false
	}
	return o.RefundableFare, true
}

// HasRefundableFare returns a boolean if a field has been set.
func (o *PricingOptions) HasRefundableFare() bool {
	if o != nil && !utils.IsNil(o.RefundableFare) {
		return true
	}

	return false
}

// SetRefundableFare gets a reference to the given bool and assigns it to the RefundableFare field.
func (o *PricingOptions) SetRefundableFare(v bool) {
	o.RefundableFare = &v
}

// GetNoRestrictionFare returns the NoRestrictionFare field value if set, zero value otherwise.
func (o *PricingOptions) GetNoRestrictionFare() bool {
	if o == nil || utils.IsNil(o.NoRestrictionFare) {
		var ret bool
		return ret
	}
	return *o.NoRestrictionFare
}

// GetNoRestrictionFareOk returns a tuple with the NoRestrictionFare field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PricingOptions) GetNoRestrictionFareOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.NoRestrictionFare) {
		return nil, false
	}
	return o.NoRestrictionFare, true
}

// HasNoRestrictionFare returns a boolean if a field has been set.
func (o *PricingOptions) HasNoRestrictionFare() bool {
	if o != nil && !utils.IsNil(o.NoRestrictionFare) {
		return true
	}

	return false
}

// SetNoRestrictionFare gets a reference to the given bool and assigns it to the NoRestrictionFare field.
func (o *PricingOptions) SetNoRestrictionFare(v bool) {
	o.NoRestrictionFare = &v
}

// GetNoPenaltyFare returns the NoPenaltyFare field value if set, zero value otherwise.
func (o *PricingOptions) GetNoPenaltyFare() bool {
	if o == nil || utils.IsNil(o.NoPenaltyFare) {
		var ret bool
		return ret
	}
	return *o.NoPenaltyFare
}

// GetNoPenaltyFareOk returns a tuple with the NoPenaltyFare field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PricingOptions) GetNoPenaltyFareOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.NoPenaltyFare) {
		return nil, false
	}
	return o.NoPenaltyFare, true
}

// HasNoPenaltyFare returns a boolean if a field has been set.
func (o *PricingOptions) HasNoPenaltyFare() bool {
	if o != nil && !utils.IsNil(o.NoPenaltyFare) {
		return true
	}

	return false
}

// SetNoPenaltyFare gets a reference to the given bool and assigns it to the NoPenaltyFare field.
func (o *PricingOptions) SetNoPenaltyFare(v bool) {
	o.NoPenaltyFare = &v
}

func (o PricingOptions) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PricingOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.FareType) {
		toSerialize["fareType"] = o.FareType
	}
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

type NullablePricingOptions struct {
	value *PricingOptions
	isSet bool
}

func (v NullablePricingOptions) Get() *PricingOptions {
	return v.value
}

func (v *NullablePricingOptions) Set(val *PricingOptions) {
	v.value = val
	v.isSet = true
}

func (v NullablePricingOptions) IsSet() bool {
	return v.isSet
}

func (v *NullablePricingOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePricingOptions(val *PricingOptions) *NullablePricingOptions {
	return &NullablePricingOptions{value: val, isSet: true}
}

func (v NullablePricingOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePricingOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
