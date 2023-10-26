package models

import (
	"encoding/json"
	"fmt"
)

// TravelerType traveler type age restrictions : CHILD < 12y, HELD_INFANT < 2y, SEATED_INFANT < 2y, SENIOR >=60y
type TravelerType string

// List of TravelerType
const (
	ADULT         TravelerType = "ADULT"
	CHILD         TravelerType = "CHILD"
	SENIOR        TravelerType = "SENIOR"
	YOUNG         TravelerType = "YOUNG"
	HELD_INFANT   TravelerType = "HELD_INFANT"
	SEATED_INFANT TravelerType = "SEATED_INFANT"
	STUDENT       TravelerType = "STUDENT"
)

// All allowed values of TravelerType enum
var AllowedTravelerTypeEnumValues = []TravelerType{
	"ADULT",
	"CHILD",
	"SENIOR",
	"YOUNG",
	"HELD_INFANT",
	"SEATED_INFANT",
	"STUDENT",
}

func (v *TravelerType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TravelerType(value)
	for _, existing := range AllowedTravelerTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TravelerType", value)
}

// NewTravelerTypeFromValue returns a pointer to a valid TravelerType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTravelerTypeFromValue(v string) (*TravelerType, error) {
	ev := TravelerType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TravelerType: valid values are %v", v, AllowedTravelerTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TravelerType) IsValid() bool {
	for _, existing := range AllowedTravelerTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TravelerType value
func (v TravelerType) Ptr() *TravelerType {
	return &v
}

type NullableTravelerType struct {
	value *TravelerType
	isSet bool
}

func (v NullableTravelerType) Get() *TravelerType {
	return v.value
}

func (v *NullableTravelerType) Set(val *TravelerType) {
	v.value = val
	v.isSet = true
}

func (v NullableTravelerType) IsSet() bool {
	return v.isSet
}

func (v *NullableTravelerType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTravelerType(val *TravelerType) *NullableTravelerType {
	return &NullableTravelerType{value: val, isSet: true}
}

func (v NullableTravelerType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTravelerType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
