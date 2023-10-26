package models

import (
	"encoding/json"
	"fmt"
)

// AdditionalServiceType additional service type
type AdditionalServiceType string

// List of AdditionalServiceType
const (
	CHECKED_BAGS   AdditionalServiceType = "CHECKED_BAGS"
	MEALS          AdditionalServiceType = "MEALS"
	SEATS          AdditionalServiceType = "SEATS"
	OTHER_SERVICES AdditionalServiceType = "OTHER_SERVICES"
)

// All allowed values of AdditionalServiceType enum
var AllowedAdditionalServiceTypeEnumValues = []AdditionalServiceType{
	"CHECKED_BAGS",
	"MEALS",
	"SEATS",
	"OTHER_SERVICES",
}

func (v *AdditionalServiceType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AdditionalServiceType(value)
	for _, existing := range AllowedAdditionalServiceTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AdditionalServiceType", value)
}

// NewAdditionalServiceTypeFromValue returns a pointer to a valid AdditionalServiceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdditionalServiceTypeFromValue(v string) (*AdditionalServiceType, error) {
	ev := AdditionalServiceType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdditionalServiceType: valid values are %v", v, AllowedAdditionalServiceTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdditionalServiceType) IsValid() bool {
	for _, existing := range AllowedAdditionalServiceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AdditionalServiceType value
func (v AdditionalServiceType) Ptr() *AdditionalServiceType {
	return &v
}

type NullableAdditionalServiceType struct {
	value *AdditionalServiceType
	isSet bool
}

func (v NullableAdditionalServiceType) Get() *AdditionalServiceType {
	return v.value
}

func (v *NullableAdditionalServiceType) Set(val *AdditionalServiceType) {
	v.value = val
	v.isSet = true
}

func (v NullableAdditionalServiceType) IsSet() bool {
	return v.isSet
}

func (v *NullableAdditionalServiceType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdditionalServiceType(val *AdditionalServiceType) *NullableAdditionalServiceType {
	return &NullableAdditionalServiceType{value: val, isSet: true}
}

func (v NullableAdditionalServiceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdditionalServiceType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
