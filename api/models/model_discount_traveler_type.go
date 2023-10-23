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

// DiscountTravelerType type of discount applied
type DiscountTravelerType string

// List of DiscountTravelerType
const (
	SPANISH_CITIZEN   DiscountTravelerType = "SPANISH_CITIZEN"
	EUROPEAN_CITIZEN  DiscountTravelerType = "EUROPEAN_CITIZEN"
	GOVERNMENT_WORKER DiscountTravelerType = "GOVERNMENT_WORKER"
	MILITARY          DiscountTravelerType = "MILITARY"
	MINOR_WITHOUT_ID  DiscountTravelerType = "MINOR_WITHOUT_ID"
)

// All allowed values of DiscountTravelerType enum
var AllowedDiscountTravelerTypeEnumValues = []DiscountTravelerType{
	"SPANISH_CITIZEN",
	"EUROPEAN_CITIZEN",
	"GOVERNMENT_WORKER",
	"MILITARY",
	"MINOR_WITHOUT_ID",
}

func (v *DiscountTravelerType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DiscountTravelerType(value)
	for _, existing := range AllowedDiscountTravelerTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DiscountTravelerType", value)
}

// NewDiscountTravelerTypeFromValue returns a pointer to a valid DiscountTravelerType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDiscountTravelerTypeFromValue(v string) (*DiscountTravelerType, error) {
	ev := DiscountTravelerType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DiscountTravelerType: valid values are %v", v, AllowedDiscountTravelerTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DiscountTravelerType) IsValid() bool {
	for _, existing := range AllowedDiscountTravelerTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DiscountTravelerType value
func (v DiscountTravelerType) Ptr() *DiscountTravelerType {
	return &v
}

type NullableDiscountTravelerType struct {
	value *DiscountTravelerType
	isSet bool
}

func (v NullableDiscountTravelerType) Get() *DiscountTravelerType {
	return v.value
}

func (v *NullableDiscountTravelerType) Set(val *DiscountTravelerType) {
	v.value = val
	v.isSet = true
}

func (v NullableDiscountTravelerType) IsSet() bool {
	return v.isSet
}

func (v *NullableDiscountTravelerType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDiscountTravelerType(val *DiscountTravelerType) *NullableDiscountTravelerType {
	return &NullableDiscountTravelerType{value: val, isSet: true}
}

func (v NullableDiscountTravelerType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDiscountTravelerType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
