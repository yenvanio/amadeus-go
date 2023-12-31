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

// DiscountType type of discount applied
type DiscountType string

// List of DiscountType
const (
	SPANISH_RESIDENT        DiscountType = "SPANISH_RESIDENT"
	AIR_FRANCE_DOMESTIC     DiscountType = "AIR_FRANCE_DOMESTIC"
	AIR_FRANCE_COMBINED     DiscountType = "AIR_FRANCE_COMBINED"
	AIR_FRANCE_METROPOLITAN DiscountType = "AIR_FRANCE_METROPOLITAN"
)

// All allowed values of DiscountType enum
var AllowedDiscountTypeEnumValues = []DiscountType{
	"SPANISH_RESIDENT",
	"AIR_FRANCE_DOMESTIC",
	"AIR_FRANCE_COMBINED",
	"AIR_FRANCE_METROPOLITAN",
}

func (v *DiscountType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DiscountType(value)
	for _, existing := range AllowedDiscountTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DiscountType", value)
}

// NewDiscountTypeFromValue returns a pointer to a valid DiscountType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDiscountTypeFromValue(v string) (*DiscountType, error) {
	ev := DiscountType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DiscountType: valid values are %v", v, AllowedDiscountTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DiscountType) IsValid() bool {
	for _, existing := range AllowedDiscountTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DiscountType value
func (v DiscountType) Ptr() *DiscountType {
	return &v
}

type NullableDiscountType struct {
	value *DiscountType
	isSet bool
}

func (v NullableDiscountType) Get() *DiscountType {
	return v.value
}

func (v *NullableDiscountType) Set(val *DiscountType) {
	v.value = val
	v.isSet = true
}

func (v NullableDiscountType) IsSet() bool {
	return v.isSet
}

func (v *NullableDiscountType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDiscountType(val *DiscountType) *NullableDiscountType {
	return &NullableDiscountType{value: val, isSet: true}
}

func (v NullableDiscountType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDiscountType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

