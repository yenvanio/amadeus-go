package models

import (
	"encoding/json"
	"github.com/yenvanio/amadeus-go/amadeus/utils"
)

// checks if the AircraftEquipment type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &AircraftEquipment{}

// AircraftEquipment information related to the aircraft
type AircraftEquipment struct {
	// IATA aircraft code (http://www.flugzeuginfo.net/table_accodes_iata_en.php)
	Code *string `json:"code,omitempty"`
}

// NewAircraftEquipment instantiates a new AircraftEquipment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAircraftEquipment() *AircraftEquipment {
	this := AircraftEquipment{}
	return &this
}

// NewAircraftEquipmentWithDefaults instantiates a new AircraftEquipment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAircraftEquipmentWithDefaults() *AircraftEquipment {
	this := AircraftEquipment{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *AircraftEquipment) GetCode() string {
	if o == nil || utils.IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AircraftEquipment) GetCodeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *AircraftEquipment) HasCode() bool {
	if o != nil && !utils.IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *AircraftEquipment) SetCode(v string) {
	o.Code = &v
}

func (o AircraftEquipment) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AircraftEquipment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	return toSerialize, nil
}

type NullableAircraftEquipment struct {
	value *AircraftEquipment
	isSet bool
}

func (v NullableAircraftEquipment) Get() *AircraftEquipment {
	return v.value
}

func (v *NullableAircraftEquipment) Set(val *AircraftEquipment) {
	v.value = val
	v.isSet = true
}

func (v NullableAircraftEquipment) IsSet() bool {
	return v.isSet
}

func (v *NullableAircraftEquipment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAircraftEquipment(val *AircraftEquipment) *NullableAircraftEquipment {
	return &NullableAircraftEquipment{value: val, isSet: true}
}

func (v NullableAircraftEquipment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAircraftEquipment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
