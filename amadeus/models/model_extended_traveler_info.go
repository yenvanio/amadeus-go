package models

import (
	"encoding/json"
	"github.com/yenvanio/amadeus-go/amadeus/utils"
)

// checks if the ExtendedTravelerInfo type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ExtendedTravelerInfo{}

// ExtendedTravelerInfo struct for ExtendedTravelerInfo
type ExtendedTravelerInfo struct {
	Id           string       `json:"id"`
	TravelerType TravelerType `json:"travelerType"`
	// if type=\"HELD_INFANT\", corresponds to the adult travelers's id who will share the seat
	AssociatedAdultId *string `json:"associatedAdultId,omitempty"`
}

// NewExtendedTravelerInfo instantiates a new ExtendedTravelerInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExtendedTravelerInfo(id string, travelerType TravelerType) *ExtendedTravelerInfo {
	this := ExtendedTravelerInfo{}
	this.Id = id
	this.TravelerType = travelerType
	return &this
}

// NewExtendedTravelerInfoWithDefaults instantiates a new ExtendedTravelerInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExtendedTravelerInfoWithDefaults() *ExtendedTravelerInfo {
	this := ExtendedTravelerInfo{}
	return &this
}

// GetId returns the Id field value
func (o *ExtendedTravelerInfo) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ExtendedTravelerInfo) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ExtendedTravelerInfo) SetId(v string) {
	o.Id = v
}

// GetTravelerType returns the TravelerType field value
func (o *ExtendedTravelerInfo) GetTravelerType() TravelerType {
	if o == nil {
		var ret TravelerType
		return ret
	}

	return o.TravelerType
}

// GetTravelerTypeOk returns a tuple with the TravelerType field value
// and a boolean to check if the value has been set.
func (o *ExtendedTravelerInfo) GetTravelerTypeOk() (*TravelerType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TravelerType, true
}

// SetTravelerType sets field value
func (o *ExtendedTravelerInfo) SetTravelerType(v TravelerType) {
	o.TravelerType = v
}

// GetAssociatedAdultId returns the AssociatedAdultId field value if set, zero value otherwise.
func (o *ExtendedTravelerInfo) GetAssociatedAdultId() string {
	if o == nil || utils.IsNil(o.AssociatedAdultId) {
		var ret string
		return ret
	}
	return *o.AssociatedAdultId
}

// GetAssociatedAdultIdOk returns a tuple with the AssociatedAdultId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtendedTravelerInfo) GetAssociatedAdultIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.AssociatedAdultId) {
		return nil, false
	}
	return o.AssociatedAdultId, true
}

// HasAssociatedAdultId returns a boolean if a field has been set.
func (o *ExtendedTravelerInfo) HasAssociatedAdultId() bool {
	if o != nil && !utils.IsNil(o.AssociatedAdultId) {
		return true
	}

	return false
}

// SetAssociatedAdultId gets a reference to the given string and assigns it to the AssociatedAdultId field.
func (o *ExtendedTravelerInfo) SetAssociatedAdultId(v string) {
	o.AssociatedAdultId = &v
}

func (o ExtendedTravelerInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExtendedTravelerInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["travelerType"] = o.TravelerType
	if !utils.IsNil(o.AssociatedAdultId) {
		toSerialize["associatedAdultId"] = o.AssociatedAdultId
	}
	return toSerialize, nil
}

type NullableExtendedTravelerInfo struct {
	value *ExtendedTravelerInfo
	isSet bool
}

func (v NullableExtendedTravelerInfo) Get() *ExtendedTravelerInfo {
	return v.value
}

func (v *NullableExtendedTravelerInfo) Set(val *ExtendedTravelerInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableExtendedTravelerInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableExtendedTravelerInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExtendedTravelerInfo(val *ExtendedTravelerInfo) *NullableExtendedTravelerInfo {
	return &NullableExtendedTravelerInfo{value: val, isSet: true}
}

func (v NullableExtendedTravelerInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExtendedTravelerInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
