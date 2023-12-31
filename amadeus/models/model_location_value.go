package models

import (
	"encoding/json"
	"github.com/yenvanio/amadeus-go/amadeus/utils"
)

// checks if the LocationValue type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &LocationValue{}

// LocationValue struct for LocationValue
type LocationValue struct {
	// City code associated to the airport
	CityCode *string `json:"cityCode,omitempty"`
	// Country code of the airport
	CountryCode *string `json:"countryCode,omitempty"`
}

// NewLocationValue instantiates a new LocationValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLocationValue() *LocationValue {
	this := LocationValue{}
	return &this
}

// NewLocationValueWithDefaults instantiates a new LocationValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLocationValueWithDefaults() *LocationValue {
	this := LocationValue{}
	return &this
}

// GetCityCode returns the CityCode field value if set, zero value otherwise.
func (o *LocationValue) GetCityCode() string {
	if o == nil || utils.IsNil(o.CityCode) {
		var ret string
		return ret
	}
	return *o.CityCode
}

// GetCityCodeOk returns a tuple with the CityCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationValue) GetCityCodeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.CityCode) {
		return nil, false
	}
	return o.CityCode, true
}

// HasCityCode returns a boolean if a field has been set.
func (o *LocationValue) HasCityCode() bool {
	if o != nil && !utils.IsNil(o.CityCode) {
		return true
	}

	return false
}

// SetCityCode gets a reference to the given string and assigns it to the CityCode field.
func (o *LocationValue) SetCityCode(v string) {
	o.CityCode = &v
}

// GetCountryCode returns the CountryCode field value if set, zero value otherwise.
func (o *LocationValue) GetCountryCode() string {
	if o == nil || utils.IsNil(o.CountryCode) {
		var ret string
		return ret
	}
	return *o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationValue) GetCountryCodeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.CountryCode) {
		return nil, false
	}
	return o.CountryCode, true
}

// HasCountryCode returns a boolean if a field has been set.
func (o *LocationValue) HasCountryCode() bool {
	if o != nil && !utils.IsNil(o.CountryCode) {
		return true
	}

	return false
}

// SetCountryCode gets a reference to the given string and assigns it to the CountryCode field.
func (o *LocationValue) SetCountryCode(v string) {
	o.CountryCode = &v
}

func (o LocationValue) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LocationValue) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.CityCode) {
		toSerialize["cityCode"] = o.CityCode
	}
	if !utils.IsNil(o.CountryCode) {
		toSerialize["countryCode"] = o.CountryCode
	}
	return toSerialize, nil
}

type NullableLocationValue struct {
	value *LocationValue
	isSet bool
}

func (v NullableLocationValue) Get() *LocationValue {
	return v.value
}

func (v *NullableLocationValue) Set(val *LocationValue) {
	v.value = val
	v.isSet = true
}

func (v NullableLocationValue) IsSet() bool {
	return v.isSet
}

func (v *NullableLocationValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocationValue(val *LocationValue) *NullableLocationValue {
	return &NullableLocationValue{value: val, isSet: true}
}

func (v NullableLocationValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocationValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
