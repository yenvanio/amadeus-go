package models

import (
	"encoding/json"
	"github.com/yenvanio/amadeus-go/amadeus/utils"
)

// checks if the CarrierRestrictions type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &CarrierRestrictions{}

// CarrierRestrictions Restriction towards carriers.
type CarrierRestrictions struct {
	// This flag enable/disable filtering of blacklisted airline by EU. The list of the banned airlines is published in the Official Journal of the European Union, where they are included as annexes A and B to the Commission Regulation. The blacklist of an airline can concern all its flights or some specific aircraft types pertaining to the airline
	BlacklistedInEUAllowed *bool `json:"blacklistedInEUAllowed,omitempty"`
	// This option ensures that the system will only consider these airlines.
	ExcludedCarrierCodes []string `json:"excludedCarrierCodes,omitempty"`
	// This option ensures that the system will only consider these airlines.
	IncludedCarrierCodes []string `json:"includedCarrierCodes,omitempty"`
}

// NewCarrierRestrictions instantiates a new CarrierRestrictions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCarrierRestrictions() *CarrierRestrictions {
	this := CarrierRestrictions{}
	var blacklistedInEUAllowed bool = false
	this.BlacklistedInEUAllowed = &blacklistedInEUAllowed
	return &this
}

// NewCarrierRestrictionsWithDefaults instantiates a new CarrierRestrictions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCarrierRestrictionsWithDefaults() *CarrierRestrictions {
	this := CarrierRestrictions{}
	var blacklistedInEUAllowed bool = false
	this.BlacklistedInEUAllowed = &blacklistedInEUAllowed
	return &this
}

// GetBlacklistedInEUAllowed returns the BlacklistedInEUAllowed field value if set, zero value otherwise.
func (o *CarrierRestrictions) GetBlacklistedInEUAllowed() bool {
	if o == nil || utils.IsNil(o.BlacklistedInEUAllowed) {
		var ret bool
		return ret
	}
	return *o.BlacklistedInEUAllowed
}

// GetBlacklistedInEUAllowedOk returns a tuple with the BlacklistedInEUAllowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CarrierRestrictions) GetBlacklistedInEUAllowedOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.BlacklistedInEUAllowed) {
		return nil, false
	}
	return o.BlacklistedInEUAllowed, true
}

// HasBlacklistedInEUAllowed returns a boolean if a field has been set.
func (o *CarrierRestrictions) HasBlacklistedInEUAllowed() bool {
	if o != nil && !utils.IsNil(o.BlacklistedInEUAllowed) {
		return true
	}

	return false
}

// SetBlacklistedInEUAllowed gets a reference to the given bool and assigns it to the BlacklistedInEUAllowed field.
func (o *CarrierRestrictions) SetBlacklistedInEUAllowed(v bool) {
	o.BlacklistedInEUAllowed = &v
}

// GetExcludedCarrierCodes returns the ExcludedCarrierCodes field value if set, zero value otherwise.
func (o *CarrierRestrictions) GetExcludedCarrierCodes() []string {
	if o == nil || utils.IsNil(o.ExcludedCarrierCodes) {
		var ret []string
		return ret
	}
	return o.ExcludedCarrierCodes
}

// GetExcludedCarrierCodesOk returns a tuple with the ExcludedCarrierCodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CarrierRestrictions) GetExcludedCarrierCodesOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.ExcludedCarrierCodes) {
		return nil, false
	}
	return o.ExcludedCarrierCodes, true
}

// HasExcludedCarrierCodes returns a boolean if a field has been set.
func (o *CarrierRestrictions) HasExcludedCarrierCodes() bool {
	if o != nil && !utils.IsNil(o.ExcludedCarrierCodes) {
		return true
	}

	return false
}

// SetExcludedCarrierCodes gets a reference to the given []string and assigns it to the ExcludedCarrierCodes field.
func (o *CarrierRestrictions) SetExcludedCarrierCodes(v []string) {
	o.ExcludedCarrierCodes = v
}

// GetIncludedCarrierCodes returns the IncludedCarrierCodes field value if set, zero value otherwise.
func (o *CarrierRestrictions) GetIncludedCarrierCodes() []string {
	if o == nil || utils.IsNil(o.IncludedCarrierCodes) {
		var ret []string
		return ret
	}
	return o.IncludedCarrierCodes
}

// GetIncludedCarrierCodesOk returns a tuple with the IncludedCarrierCodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CarrierRestrictions) GetIncludedCarrierCodesOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.IncludedCarrierCodes) {
		return nil, false
	}
	return o.IncludedCarrierCodes, true
}

// HasIncludedCarrierCodes returns a boolean if a field has been set.
func (o *CarrierRestrictions) HasIncludedCarrierCodes() bool {
	if o != nil && !utils.IsNil(o.IncludedCarrierCodes) {
		return true
	}

	return false
}

// SetIncludedCarrierCodes gets a reference to the given []string and assigns it to the IncludedCarrierCodes field.
func (o *CarrierRestrictions) SetIncludedCarrierCodes(v []string) {
	o.IncludedCarrierCodes = v
}

func (o CarrierRestrictions) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CarrierRestrictions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.BlacklistedInEUAllowed) {
		toSerialize["blacklistedInEUAllowed"] = o.BlacklistedInEUAllowed
	}
	if !utils.IsNil(o.ExcludedCarrierCodes) {
		toSerialize["excludedCarrierCodes"] = o.ExcludedCarrierCodes
	}
	if !utils.IsNil(o.IncludedCarrierCodes) {
		toSerialize["includedCarrierCodes"] = o.IncludedCarrierCodes
	}
	return toSerialize, nil
}

type NullableCarrierRestrictions struct {
	value *CarrierRestrictions
	isSet bool
}

func (v NullableCarrierRestrictions) Get() *CarrierRestrictions {
	return v.value
}

func (v *NullableCarrierRestrictions) Set(val *CarrierRestrictions) {
	v.value = val
	v.isSet = true
}

func (v NullableCarrierRestrictions) IsSet() bool {
	return v.isSet
}

func (v *NullableCarrierRestrictions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCarrierRestrictions(val *CarrierRestrictions) *NullableCarrierRestrictions {
	return &NullableCarrierRestrictions{value: val, isSet: true}
}

func (v NullableCarrierRestrictions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCarrierRestrictions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
