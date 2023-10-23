/*
Flight Offers Price

Before using this API, we recommend you read our **[Authorization Guide](https://developers.amadeus.com/self-service/apis-docs/guides/authorization-262)** for more information on how to generate an access token.   Please also be aware that our test environment is based on a subset of the production, if you are not returning any results try with big cities/airports like LON (London) or NYC (New-York).

API version: 1.3.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"
	"fmt"
)

// ContactPurpose the purpose for which this contact is to be used
type ContactPurpose string

// List of ContactPurpose
const (
	STANDARD                      ContactPurpose = "STANDARD"
	INVOICE                       ContactPurpose = "INVOICE"
	STANDARD_WITHOUT_TRANSMISSION ContactPurpose = "STANDARD_WITHOUT_TRANSMISSION"
)

// All allowed values of ContactPurpose enum
var AllowedContactPurposeEnumValues = []ContactPurpose{
	"STANDARD",
	"INVOICE",
	"STANDARD_WITHOUT_TRANSMISSION",
}

func (v *ContactPurpose) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ContactPurpose(value)
	for _, existing := range AllowedContactPurposeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ContactPurpose", value)
}

// NewContactPurposeFromValue returns a pointer to a valid ContactPurpose
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewContactPurposeFromValue(v string) (*ContactPurpose, error) {
	ev := ContactPurpose(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ContactPurpose: valid values are %v", v, AllowedContactPurposeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ContactPurpose) IsValid() bool {
	for _, existing := range AllowedContactPurposeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ContactPurpose value
func (v ContactPurpose) Ptr() *ContactPurpose {
	return &v
}

type NullableContactPurpose struct {
	value *ContactPurpose
	isSet bool
}

func (v NullableContactPurpose) Get() *ContactPurpose {
	return v.value
}

func (v *NullableContactPurpose) Set(val *ContactPurpose) {
	v.value = val
	v.isSet = true
}

func (v NullableContactPurpose) IsSet() bool {
	return v.isSet
}

func (v *NullableContactPurpose) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContactPurpose(val *ContactPurpose) *NullableContactPurpose {
	return &NullableContactPurpose{value: val, isSet: true}
}

func (v NullableContactPurpose) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContactPurpose) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

