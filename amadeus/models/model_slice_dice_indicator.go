package models

import (
	"encoding/json"
	"fmt"
)

// SliceDiceIndicator slice and Dice indicator, such as Local Availability, Sub OnD(Origin and Destination) 1 Availability and Sub OnD 2 Availability
type SliceDiceIndicator string

// List of SliceDiceIndicator
const (
	LOCAL_AVAILABILITY    SliceDiceIndicator = "LOCAL_AVAILABILITY"
	SUB_OD_AVAILABILITY_1 SliceDiceIndicator = "SUB_OD_AVAILABILITY_1"
	SUB_OD_AVAILABILITY_2 SliceDiceIndicator = "SUB_OD_AVAILABILITY_2"
)

// All allowed values of SliceDiceIndicator enum
var AllowedSliceDiceIndicatorEnumValues = []SliceDiceIndicator{
	"LOCAL_AVAILABILITY",
	"SUB_OD_AVAILABILITY_1",
	"SUB_OD_AVAILABILITY_2",
}

func (v *SliceDiceIndicator) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SliceDiceIndicator(value)
	for _, existing := range AllowedSliceDiceIndicatorEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SliceDiceIndicator", value)
}

// NewSliceDiceIndicatorFromValue returns a pointer to a valid SliceDiceIndicator
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSliceDiceIndicatorFromValue(v string) (*SliceDiceIndicator, error) {
	ev := SliceDiceIndicator(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SliceDiceIndicator: valid values are %v", v, AllowedSliceDiceIndicatorEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SliceDiceIndicator) IsValid() bool {
	for _, existing := range AllowedSliceDiceIndicatorEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SliceDiceIndicator value
func (v SliceDiceIndicator) Ptr() *SliceDiceIndicator {
	return &v
}

type NullableSliceDiceIndicator struct {
	value *SliceDiceIndicator
	isSet bool
}

func (v NullableSliceDiceIndicator) Get() *SliceDiceIndicator {
	return v.value
}

func (v *NullableSliceDiceIndicator) Set(val *SliceDiceIndicator) {
	v.value = val
	v.isSet = true
}

func (v NullableSliceDiceIndicator) IsSet() bool {
	return v.isSet
}

func (v *NullableSliceDiceIndicator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSliceDiceIndicator(val *SliceDiceIndicator) *NullableSliceDiceIndicator {
	return &NullableSliceDiceIndicator{value: val, isSet: true}
}

func (v NullableSliceDiceIndicator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSliceDiceIndicator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
