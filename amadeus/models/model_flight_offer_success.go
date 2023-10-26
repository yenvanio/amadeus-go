package models

import (
	"encoding/json"
	"github.com/yenvanio/amadeus-go/amadeus/utils"
)

// checks if the Success type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &FlightOfferSuccess{}

// Success struct for Success
type FlightOfferSuccess struct {
	Warnings     []Issue             `json:"warnings,omitempty"`
	Meta         *CollectionMetaLink `json:"meta,omitempty"`
	Data         []FlightOffer       `json:"data"`
	Dictionaries *Dictionaries       `json:"dictionaries,omitempty"`
}

// NewSuccess instantiates a new Success object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSuccess(data []FlightOffer) *FlightOfferSuccess {
	this := FlightOfferSuccess{}
	this.Data = data
	return &this
}

// NewSuccessWithDefaults instantiates a new Success object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSuccessWithDefaults() *FlightOfferSuccess {
	this := FlightOfferSuccess{}
	return &this
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *FlightOfferSuccess) GetWarnings() []Issue {
	if o == nil || utils.IsNil(o.Warnings) {
		var ret []Issue
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlightOfferSuccess) GetWarningsOk() ([]Issue, bool) {
	if o == nil || utils.IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *FlightOfferSuccess) HasWarnings() bool {
	if o != nil && !utils.IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []Issue and assigns it to the Warnings field.
func (o *FlightOfferSuccess) SetWarnings(v []Issue) {
	o.Warnings = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *FlightOfferSuccess) GetMeta() CollectionMetaLink {
	if o == nil || utils.IsNil(o.Meta) {
		var ret CollectionMetaLink
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlightOfferSuccess) GetMetaOk() (*CollectionMetaLink, bool) {
	if o == nil || utils.IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *FlightOfferSuccess) HasMeta() bool {
	if o != nil && !utils.IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given CollectionMetaLink and assigns it to the Meta field.
func (o *FlightOfferSuccess) SetMeta(v CollectionMetaLink) {
	o.Meta = &v
}

// GetData returns the Data field value
func (o *FlightOfferSuccess) GetData() []FlightOffer {
	if o == nil {
		var ret []FlightOffer
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *FlightOfferSuccess) GetDataOk() ([]FlightOffer, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *FlightOfferSuccess) SetData(v []FlightOffer) {
	o.Data = v
}

// GetDictionaries returns the Dictionaries field value if set, zero value otherwise.
func (o *FlightOfferSuccess) GetDictionaries() Dictionaries {
	if o == nil || utils.IsNil(o.Dictionaries) {
		var ret Dictionaries
		return ret
	}
	return *o.Dictionaries
}

// GetDictionariesOk returns a tuple with the Dictionaries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlightOfferSuccess) GetDictionariesOk() (*Dictionaries, bool) {
	if o == nil || utils.IsNil(o.Dictionaries) {
		return nil, false
	}
	return o.Dictionaries, true
}

// HasDictionaries returns a boolean if a field has been set.
func (o *FlightOfferSuccess) HasDictionaries() bool {
	if o != nil && !utils.IsNil(o.Dictionaries) {
		return true
	}

	return false
}

// SetDictionaries gets a reference to the given Dictionaries and assigns it to the Dictionaries field.
func (o *FlightOfferSuccess) SetDictionaries(v Dictionaries) {
	o.Dictionaries = &v
}

func (o FlightOfferSuccess) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FlightOfferSuccess) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	if !utils.IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	toSerialize["data"] = o.Data
	if !utils.IsNil(o.Dictionaries) {
		toSerialize["dictionaries"] = o.Dictionaries
	}
	return toSerialize, nil
}

type NullableSuccess struct {
	value *FlightOfferSuccess
	isSet bool
}

func (v NullableSuccess) Get() *FlightOfferSuccess {
	return v.value
}

func (v *NullableSuccess) Set(val *FlightOfferSuccess) {
	v.value = val
	v.isSet = true
}

func (v NullableSuccess) IsSet() bool {
	return v.isSet
}

func (v *NullableSuccess) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSuccess(val *FlightOfferSuccess) *NullableSuccess {
	return &NullableSuccess{value: val, isSet: true}
}

func (v NullableSuccess) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSuccess) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
