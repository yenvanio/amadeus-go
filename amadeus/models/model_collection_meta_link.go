package models

import (
	"encoding/json"
	"github.com/yenvanio/amadeus-go/amadeus/utils"
)

// checks if the CollectionMetaLink type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &CollectionMetaLink{}

// CollectionMetaLink struct for CollectionMetaLink
type CollectionMetaLink struct {
	Count *int32           `json:"count,omitempty"`
	Links *CollectionLinks `json:"links,omitempty"`
}

// NewCollectionMetaLink instantiates a new CollectionMetaLink object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionMetaLink() *CollectionMetaLink {
	this := CollectionMetaLink{}
	return &this
}

// NewCollectionMetaLinkWithDefaults instantiates a new CollectionMetaLink object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionMetaLinkWithDefaults() *CollectionMetaLink {
	this := CollectionMetaLink{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *CollectionMetaLink) GetCount() int32 {
	if o == nil || utils.IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionMetaLink) GetCountOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *CollectionMetaLink) HasCount() bool {
	if o != nil && !utils.IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *CollectionMetaLink) SetCount(v int32) {
	o.Count = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *CollectionMetaLink) GetLinks() CollectionLinks {
	if o == nil || utils.IsNil(o.Links) {
		var ret CollectionLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionMetaLink) GetLinksOk() (*CollectionLinks, bool) {
	if o == nil || utils.IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *CollectionMetaLink) HasLinks() bool {
	if o != nil && !utils.IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given CollectionLinks and assigns it to the Links field.
func (o *CollectionMetaLink) SetLinks(v CollectionLinks) {
	o.Links = &v
}

func (o CollectionMetaLink) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CollectionMetaLink) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !utils.IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	return toSerialize, nil
}

type NullableCollectionMetaLink struct {
	value *CollectionMetaLink
	isSet bool
}

func (v NullableCollectionMetaLink) Get() *CollectionMetaLink {
	return v.value
}

func (v *NullableCollectionMetaLink) Set(val *CollectionMetaLink) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionMetaLink) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionMetaLink) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionMetaLink(val *CollectionMetaLink) *NullableCollectionMetaLink {
	return &NullableCollectionMetaLink{value: val, isSet: true}
}

func (v NullableCollectionMetaLink) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionMetaLink) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
