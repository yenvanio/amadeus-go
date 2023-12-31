package models

import (
	"encoding/json"
	"github.com/yenvanio/amadeus-go/amadeus/utils"
)

// checks if the CollectionMeta type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &CollectionMeta{}

// CollectionMeta struct for CollectionMeta
type CollectionMeta struct {
	Count              *int32               `json:"count,omitempty"`
	OneWayCombinations []OneWayCombinations `json:"oneWayCombinations,omitempty"`
}

// NewCollectionMeta instantiates a new CollectionMeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionMeta() *CollectionMeta {
	this := CollectionMeta{}
	return &this
}

// NewCollectionMetaWithDefaults instantiates a new CollectionMeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionMetaWithDefaults() *CollectionMeta {
	this := CollectionMeta{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *CollectionMeta) GetCount() int32 {
	if o == nil || utils.IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionMeta) GetCountOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *CollectionMeta) HasCount() bool {
	if o != nil && !utils.IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *CollectionMeta) SetCount(v int32) {
	o.Count = &v
}

// GetOneWayCombinations returns the OneWayCombinations field value if set, zero value otherwise.
func (o *CollectionMeta) GetOneWayCombinations() []OneWayCombinations {
	if o == nil || utils.IsNil(o.OneWayCombinations) {
		var ret []OneWayCombinations
		return ret
	}
	return o.OneWayCombinations
}

// GetOneWayCombinationsOk returns a tuple with the OneWayCombinations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionMeta) GetOneWayCombinationsOk() ([]OneWayCombinations, bool) {
	if o == nil || utils.IsNil(o.OneWayCombinations) {
		return nil, false
	}
	return o.OneWayCombinations, true
}

// HasOneWayCombinations returns a boolean if a field has been set.
func (o *CollectionMeta) HasOneWayCombinations() bool {
	if o != nil && !utils.IsNil(o.OneWayCombinations) {
		return true
	}

	return false
}

// SetOneWayCombinations gets a reference to the given []OneWayCombinations and assigns it to the OneWayCombinations field.
func (o *CollectionMeta) SetOneWayCombinations(v []OneWayCombinations) {
	o.OneWayCombinations = v
}

func (o CollectionMeta) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CollectionMeta) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !utils.IsNil(o.OneWayCombinations) {
		toSerialize["oneWayCombinations"] = o.OneWayCombinations
	}
	return toSerialize, nil
}

type NullableCollectionMeta struct {
	value *CollectionMeta
	isSet bool
}

func (v NullableCollectionMeta) Get() *CollectionMeta {
	return v.value
}

func (v *NullableCollectionMeta) Set(val *CollectionMeta) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionMeta) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionMeta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionMeta(val *CollectionMeta) *NullableCollectionMeta {
	return &NullableCollectionMeta{value: val, isSet: true}
}

func (v NullableCollectionMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
