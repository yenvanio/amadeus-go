package models

import (
	"encoding/json"
	"github.com/yenvanio/amadeus-go/amadeus/utils"
)

// checks if the CollectionLinks type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &CollectionLinks{}

// CollectionLinks struct for CollectionLinks
type CollectionLinks struct {
	Self     *string `json:"self,omitempty"`
	Next     *string `json:"next,omitempty"`
	Previous *string `json:"previous,omitempty"`
	Last     *string `json:"last,omitempty"`
	First    *string `json:"first,omitempty"`
	Up       *string `json:"up,omitempty"`
}

// NewCollectionLinks instantiates a new CollectionLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionLinks() *CollectionLinks {
	this := CollectionLinks{}
	return &this
}

// NewCollectionLinksWithDefaults instantiates a new CollectionLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionLinksWithDefaults() *CollectionLinks {
	this := CollectionLinks{}
	return &this
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *CollectionLinks) GetSelf() string {
	if o == nil || utils.IsNil(o.Self) {
		var ret string
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionLinks) GetSelfOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Self) {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *CollectionLinks) HasSelf() bool {
	if o != nil && !utils.IsNil(o.Self) {
		return true
	}

	return false
}

// SetSelf gets a reference to the given string and assigns it to the Self field.
func (o *CollectionLinks) SetSelf(v string) {
	o.Self = &v
}

// GetNext returns the Next field value if set, zero value otherwise.
func (o *CollectionLinks) GetNext() string {
	if o == nil || utils.IsNil(o.Next) {
		var ret string
		return ret
	}
	return *o.Next
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionLinks) GetNextOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Next) {
		return nil, false
	}
	return o.Next, true
}

// HasNext returns a boolean if a field has been set.
func (o *CollectionLinks) HasNext() bool {
	if o != nil && !utils.IsNil(o.Next) {
		return true
	}

	return false
}

// SetNext gets a reference to the given string and assigns it to the Next field.
func (o *CollectionLinks) SetNext(v string) {
	o.Next = &v
}

// GetPrevious returns the Previous field value if set, zero value otherwise.
func (o *CollectionLinks) GetPrevious() string {
	if o == nil || utils.IsNil(o.Previous) {
		var ret string
		return ret
	}
	return *o.Previous
}

// GetPreviousOk returns a tuple with the Previous field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionLinks) GetPreviousOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Previous) {
		return nil, false
	}
	return o.Previous, true
}

// HasPrevious returns a boolean if a field has been set.
func (o *CollectionLinks) HasPrevious() bool {
	if o != nil && !utils.IsNil(o.Previous) {
		return true
	}

	return false
}

// SetPrevious gets a reference to the given string and assigns it to the Previous field.
func (o *CollectionLinks) SetPrevious(v string) {
	o.Previous = &v
}

// GetLast returns the Last field value if set, zero value otherwise.
func (o *CollectionLinks) GetLast() string {
	if o == nil || utils.IsNil(o.Last) {
		var ret string
		return ret
	}
	return *o.Last
}

// GetLastOk returns a tuple with the Last field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionLinks) GetLastOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Last) {
		return nil, false
	}
	return o.Last, true
}

// HasLast returns a boolean if a field has been set.
func (o *CollectionLinks) HasLast() bool {
	if o != nil && !utils.IsNil(o.Last) {
		return true
	}

	return false
}

// SetLast gets a reference to the given string and assigns it to the Last field.
func (o *CollectionLinks) SetLast(v string) {
	o.Last = &v
}

// GetFirst returns the First field value if set, zero value otherwise.
func (o *CollectionLinks) GetFirst() string {
	if o == nil || utils.IsNil(o.First) {
		var ret string
		return ret
	}
	return *o.First
}

// GetFirstOk returns a tuple with the First field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionLinks) GetFirstOk() (*string, bool) {
	if o == nil || utils.IsNil(o.First) {
		return nil, false
	}
	return o.First, true
}

// HasFirst returns a boolean if a field has been set.
func (o *CollectionLinks) HasFirst() bool {
	if o != nil && !utils.IsNil(o.First) {
		return true
	}

	return false
}

// SetFirst gets a reference to the given string and assigns it to the First field.
func (o *CollectionLinks) SetFirst(v string) {
	o.First = &v
}

// GetUp returns the Up field value if set, zero value otherwise.
func (o *CollectionLinks) GetUp() string {
	if o == nil || utils.IsNil(o.Up) {
		var ret string
		return ret
	}
	return *o.Up
}

// GetUpOk returns a tuple with the Up field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionLinks) GetUpOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Up) {
		return nil, false
	}
	return o.Up, true
}

// HasUp returns a boolean if a field has been set.
func (o *CollectionLinks) HasUp() bool {
	if o != nil && !utils.IsNil(o.Up) {
		return true
	}

	return false
}

// SetUp gets a reference to the given string and assigns it to the Up field.
func (o *CollectionLinks) SetUp(v string) {
	o.Up = &v
}

func (o CollectionLinks) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CollectionLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Self) {
		toSerialize["self"] = o.Self
	}
	if !utils.IsNil(o.Next) {
		toSerialize["next"] = o.Next
	}
	if !utils.IsNil(o.Previous) {
		toSerialize["previous"] = o.Previous
	}
	if !utils.IsNil(o.Last) {
		toSerialize["last"] = o.Last
	}
	if !utils.IsNil(o.First) {
		toSerialize["first"] = o.First
	}
	if !utils.IsNil(o.Up) {
		toSerialize["up"] = o.Up
	}
	return toSerialize, nil
}

type NullableCollectionLinks struct {
	value *CollectionLinks
	isSet bool
}

func (v NullableCollectionLinks) Get() *CollectionLinks {
	return v.value
}

func (v *NullableCollectionLinks) Set(val *CollectionLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionLinks(val *CollectionLinks) *NullableCollectionLinks {
	return &NullableCollectionLinks{value: val, isSet: true}
}

func (v NullableCollectionLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
