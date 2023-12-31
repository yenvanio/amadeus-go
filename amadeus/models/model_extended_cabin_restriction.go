package models

import (
	"encoding/json"
	"github.com/yenvanio/amadeus-go/amadeus/utils"
)

// checks if the ExtendedCabinRestriction type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ExtendedCabinRestriction{}

// ExtendedCabinRestriction struct for ExtendedCabinRestriction
type ExtendedCabinRestriction struct {
	Cabin *TravelClass `json:"cabin,omitempty"`
	// The list of originDestination identifiers for which the cabinRestriction applies
	OriginDestinationIds []string  `json:"originDestinationIds,omitempty"`
	Coverage             *Coverage `json:"coverage,omitempty"`
}

// NewExtendedCabinRestriction instantiates a new ExtendedCabinRestriction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExtendedCabinRestriction() *ExtendedCabinRestriction {
	this := ExtendedCabinRestriction{}
	return &this
}

// NewExtendedCabinRestrictionWithDefaults instantiates a new ExtendedCabinRestriction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExtendedCabinRestrictionWithDefaults() *ExtendedCabinRestriction {
	this := ExtendedCabinRestriction{}
	return &this
}

// GetCabin returns the Cabin field value if set, zero value otherwise.
func (o *ExtendedCabinRestriction) GetCabin() TravelClass {
	if o == nil || utils.IsNil(o.Cabin) {
		var ret TravelClass
		return ret
	}
	return *o.Cabin
}

// GetCabinOk returns a tuple with the Cabin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtendedCabinRestriction) GetCabinOk() (*TravelClass, bool) {
	if o == nil || utils.IsNil(o.Cabin) {
		return nil, false
	}
	return o.Cabin, true
}

// HasCabin returns a boolean if a field has been set.
func (o *ExtendedCabinRestriction) HasCabin() bool {
	if o != nil && !utils.IsNil(o.Cabin) {
		return true
	}

	return false
}

// SetCabin gets a reference to the given TravelClass and assigns it to the Cabin field.
func (o *ExtendedCabinRestriction) SetCabin(v TravelClass) {
	o.Cabin = &v
}

// GetOriginDestinationIds returns the OriginDestinationIds field value if set, zero value otherwise.
func (o *ExtendedCabinRestriction) GetOriginDestinationIds() []string {
	if o == nil || utils.IsNil(o.OriginDestinationIds) {
		var ret []string
		return ret
	}
	return o.OriginDestinationIds
}

// GetOriginDestinationIdsOk returns a tuple with the OriginDestinationIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtendedCabinRestriction) GetOriginDestinationIdsOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.OriginDestinationIds) {
		return nil, false
	}
	return o.OriginDestinationIds, true
}

// HasOriginDestinationIds returns a boolean if a field has been set.
func (o *ExtendedCabinRestriction) HasOriginDestinationIds() bool {
	if o != nil && !utils.IsNil(o.OriginDestinationIds) {
		return true
	}

	return false
}

// SetOriginDestinationIds gets a reference to the given []string and assigns it to the OriginDestinationIds field.
func (o *ExtendedCabinRestriction) SetOriginDestinationIds(v []string) {
	o.OriginDestinationIds = v
}

// GetCoverage returns the Coverage field value if set, zero value otherwise.
func (o *ExtendedCabinRestriction) GetCoverage() Coverage {
	if o == nil || utils.IsNil(o.Coverage) {
		var ret Coverage
		return ret
	}
	return *o.Coverage
}

// GetCoverageOk returns a tuple with the Coverage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtendedCabinRestriction) GetCoverageOk() (*Coverage, bool) {
	if o == nil || utils.IsNil(o.Coverage) {
		return nil, false
	}
	return o.Coverage, true
}

// HasCoverage returns a boolean if a field has been set.
func (o *ExtendedCabinRestriction) HasCoverage() bool {
	if o != nil && !utils.IsNil(o.Coverage) {
		return true
	}

	return false
}

// SetCoverage gets a reference to the given Coverage and assigns it to the Coverage field.
func (o *ExtendedCabinRestriction) SetCoverage(v Coverage) {
	o.Coverage = &v
}

func (o ExtendedCabinRestriction) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExtendedCabinRestriction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Cabin) {
		toSerialize["cabin"] = o.Cabin
	}
	if !utils.IsNil(o.OriginDestinationIds) {
		toSerialize["originDestinationIds"] = o.OriginDestinationIds
	}
	if !utils.IsNil(o.Coverage) {
		toSerialize["coverage"] = o.Coverage
	}
	return toSerialize, nil
}

type NullableExtendedCabinRestriction struct {
	value *ExtendedCabinRestriction
	isSet bool
}

func (v NullableExtendedCabinRestriction) Get() *ExtendedCabinRestriction {
	return v.value
}

func (v *NullableExtendedCabinRestriction) Set(val *ExtendedCabinRestriction) {
	v.value = val
	v.isSet = true
}

func (v NullableExtendedCabinRestriction) IsSet() bool {
	return v.isSet
}

func (v *NullableExtendedCabinRestriction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExtendedCabinRestriction(val *ExtendedCabinRestriction) *NullableExtendedCabinRestriction {
	return &NullableExtendedCabinRestriction{value: val, isSet: true}
}

func (v NullableExtendedCabinRestriction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExtendedCabinRestriction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
