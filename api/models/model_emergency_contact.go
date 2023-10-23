/*
Flight Offers Price

Before using this API, we recommend you read our **[Authorization Guide](https://developers.amadeus.com/self-service/apis-docs/guides/authorization-262)** for more information on how to generate an access token.   Please also be aware that our test environment is based on a subset of the production, if you are not returning any results try with big cities/airports like LON (London) or NYC (New-York).

API version: 1.3.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"
	"github.com/GIT_USER_ID/GIT_REPO_ID/api"
)

// checks if the EmergencyContact type satisfies the MappedNullable interface at compile time
var _ api.MappedNullable = &EmergencyContact{}

// EmergencyContact emergency contact number
type EmergencyContact struct {
	// Adressee name (e.g. in case of emergency purpose it corresponds to name of the person to be contacted).
	AddresseeName *string `json:"addresseeName,omitempty"`
	// Country code of the country (ISO3166-1). E.g. \"US\" for the United States
	CountryCode *string `json:"countryCode,omitempty"`
	// Phone number. Composed of digits only. The number of digits depends on the country.
	Number *string `json:"number,omitempty"`
	// additional details
	Text *string `json:"text,omitempty"`
}

// NewEmergencyContact instantiates a new EmergencyContact object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmergencyContact() *EmergencyContact {
	this := EmergencyContact{}
	return &this
}

// NewEmergencyContactWithDefaults instantiates a new EmergencyContact object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmergencyContactWithDefaults() *EmergencyContact {
	this := EmergencyContact{}
	return &this
}

// GetAddresseeName returns the AddresseeName field value if set, zero value otherwise.
func (o *EmergencyContact) GetAddresseeName() string {
	if o == nil || api.IsNil(o.AddresseeName) {
		var ret string
		return ret
	}
	return *o.AddresseeName
}

// GetAddresseeNameOk returns a tuple with the AddresseeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmergencyContact) GetAddresseeNameOk() (*string, bool) {
	if o == nil || api.IsNil(o.AddresseeName) {
		return nil, false
	}
	return o.AddresseeName, true
}

// HasAddresseeName returns a boolean if a field has been set.
func (o *EmergencyContact) HasAddresseeName() bool {
	if o != nil && !api.IsNil(o.AddresseeName) {
		return true
	}

	return false
}

// SetAddresseeName gets a reference to the given string and assigns it to the AddresseeName field.
func (o *EmergencyContact) SetAddresseeName(v string) {
	o.AddresseeName = &v
}

// GetCountryCode returns the CountryCode field value if set, zero value otherwise.
func (o *EmergencyContact) GetCountryCode() string {
	if o == nil || api.IsNil(o.CountryCode) {
		var ret string
		return ret
	}
	return *o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmergencyContact) GetCountryCodeOk() (*string, bool) {
	if o == nil || api.IsNil(o.CountryCode) {
		return nil, false
	}
	return o.CountryCode, true
}

// HasCountryCode returns a boolean if a field has been set.
func (o *EmergencyContact) HasCountryCode() bool {
	if o != nil && !api.IsNil(o.CountryCode) {
		return true
	}

	return false
}

// SetCountryCode gets a reference to the given string and assigns it to the CountryCode field.
func (o *EmergencyContact) SetCountryCode(v string) {
	o.CountryCode = &v
}

// GetNumber returns the Number field value if set, zero value otherwise.
func (o *EmergencyContact) GetNumber() string {
	if o == nil || api.IsNil(o.Number) {
		var ret string
		return ret
	}
	return *o.Number
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmergencyContact) GetNumberOk() (*string, bool) {
	if o == nil || api.IsNil(o.Number) {
		return nil, false
	}
	return o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *EmergencyContact) HasNumber() bool {
	if o != nil && !api.IsNil(o.Number) {
		return true
	}

	return false
}

// SetNumber gets a reference to the given string and assigns it to the Number field.
func (o *EmergencyContact) SetNumber(v string) {
	o.Number = &v
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *EmergencyContact) GetText() string {
	if o == nil || api.IsNil(o.Text) {
		var ret string
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmergencyContact) GetTextOk() (*string, bool) {
	if o == nil || api.IsNil(o.Text) {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *EmergencyContact) HasText() bool {
	if o != nil && !api.IsNil(o.Text) {
		return true
	}

	return false
}

// SetText gets a reference to the given string and assigns it to the Text field.
func (o *EmergencyContact) SetText(v string) {
	o.Text = &v
}

func (o EmergencyContact) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmergencyContact) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !api.IsNil(o.AddresseeName) {
		toSerialize["addresseeName"] = o.AddresseeName
	}
	if !api.IsNil(o.CountryCode) {
		toSerialize["countryCode"] = o.CountryCode
	}
	if !api.IsNil(o.Number) {
		toSerialize["number"] = o.Number
	}
	if !api.IsNil(o.Text) {
		toSerialize["text"] = o.Text
	}
	return toSerialize, nil
}

type NullableEmergencyContact struct {
	value *EmergencyContact
	isSet bool
}

func (v NullableEmergencyContact) Get() *EmergencyContact {
	return v.value
}

func (v *NullableEmergencyContact) Set(val *EmergencyContact) {
	v.value = val
	v.isSet = true
}

func (v NullableEmergencyContact) IsSet() bool {
	return v.isSet
}

func (v *NullableEmergencyContact) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmergencyContact(val *EmergencyContact) *NullableEmergencyContact {
	return &NullableEmergencyContact{value: val, isSet: true}
}

func (v NullableEmergencyContact) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmergencyContact) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


