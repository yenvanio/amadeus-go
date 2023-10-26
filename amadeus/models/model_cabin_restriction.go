package models

import (
	"encoding/json"
	"github.com/yenvanio/amadeus-go/amadeus/utils"
)

// checks if the CabinRestriction type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &CabinRestriction{}

// CabinRestriction struct for CabinRestriction
type CabinRestriction struct {
	Cabin *TravelClass `json:"cabin,omitempty"`
	// The list of originDestination identifiers for which the cabinRestriction applies
	OriginDestinationIds []string `json:"originDestinationIds,omitempty"`
}

// NewCabinRestriction instantiates a new CabinRestriction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCabinRestriction() *CabinRestriction {
	this := CabinRestriction{}
	return &this
}

// NewCabinRestrictionWithDefaults instantiates a new CabinRestriction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCabinRestrictionWithDefaults() *CabinRestriction {
	this := CabinRestriction{}
	return &this
}

// GetCabin returns the Cabin field value if set, zero value otherwise.
func (o *CabinRestriction) GetCabin() TravelClass {
	if o == nil || utils.IsNil(o.Cabin) {
		var ret TravelClass
		return ret
	}
	return *o.Cabin
}

// GetCabinOk returns a tuple with the Cabin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CabinRestriction) GetCabinOk() (*TravelClass, bool) {
	if o == nil || utils.IsNil(o.Cabin) {
		return nil, false
	}
	return o.Cabin, true
}

// HasCabin returns a boolean if a field has been set.
func (o *CabinRestriction) HasCabin() bool {
	if o != nil && !utils.IsNil(o.Cabin) {
		return true
	}

	return false
}

// SetCabin gets a reference to the given TravelClass and assigns it to the Cabin field.
func (o *CabinRestriction) SetCabin(v TravelClass) {
	o.Cabin = &v
}

// GetOriginDestinationIds returns the OriginDestinationIds field value if set, zero value otherwise.
func (o *CabinRestriction) GetOriginDestinationIds() []string {
	if o == nil || utils.IsNil(o.OriginDestinationIds) {
		var ret []string
		return ret
	}
	return o.OriginDestinationIds
}

// GetOriginDestinationIdsOk returns a tuple with the OriginDestinationIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CabinRestriction) GetOriginDestinationIdsOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.OriginDestinationIds) {
		return nil, false
	}
	return o.OriginDestinationIds, true
}

// HasOriginDestinationIds returns a boolean if a field has been set.
func (o *CabinRestriction) HasOriginDestinationIds() bool {
	if o != nil && !utils.IsNil(o.OriginDestinationIds) {
		return true
	}

	return false
}

// SetOriginDestinationIds gets a reference to the given []string and assigns it to the OriginDestinationIds field.
func (o *CabinRestriction) SetOriginDestinationIds(v []string) {
	o.OriginDestinationIds = v
}

func (o CabinRestriction) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CabinRestriction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Cabin) {
		toSerialize["cabin"] = o.Cabin
	}
	if !utils.IsNil(o.OriginDestinationIds) {
		toSerialize["originDestinationIds"] = o.OriginDestinationIds
	}
	return toSerialize, nil
}

type NullableCabinRestriction struct {
	value *CabinRestriction
	isSet bool
}

func (v NullableCabinRestriction) Get() *CabinRestriction {
	return v.value
}

func (v *NullableCabinRestriction) Set(val *CabinRestriction) {
	v.value = val
	v.isSet = true
}

func (v NullableCabinRestriction) IsSet() bool {
	return v.isSet
}

func (v *NullableCabinRestriction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCabinRestriction(val *CabinRestriction) *NullableCabinRestriction {
	return &NullableCabinRestriction{value: val, isSet: true}
}

func (v NullableCabinRestriction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCabinRestriction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
