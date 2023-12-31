package models

import (
	"encoding/json"
	"fmt"
)

// Coverage part of the trip covered by the travel class restriction (ALL_SEGMENTS if ommited)
type Coverage string

// List of Coverage
const (
	MOST_SEGMENTS        Coverage = "MOST_SEGMENTS"
	AT_LEAST_ONE_SEGMENT Coverage = "AT_LEAST_ONE_SEGMENT"
	ALL_SEGMENTS         Coverage = "ALL_SEGMENTS"
)

// All allowed values of Coverage enum
var AllowedCoverageEnumValues = []Coverage{
	"MOST_SEGMENTS",
	"AT_LEAST_ONE_SEGMENT",
	"ALL_SEGMENTS",
}

func (v *Coverage) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := Coverage(value)
	for _, existing := range AllowedCoverageEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid Coverage", value)
}

// NewCoverageFromValue returns a pointer to a valid Coverage
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCoverageFromValue(v string) (*Coverage, error) {
	ev := Coverage(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for Coverage: valid values are %v", v, AllowedCoverageEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v Coverage) IsValid() bool {
	for _, existing := range AllowedCoverageEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Coverage value
func (v Coverage) Ptr() *Coverage {
	return &v
}

type NullableCoverage struct {
	value *Coverage
	isSet bool
}

func (v NullableCoverage) Get() *Coverage {
	return v.value
}

func (v *NullableCoverage) Set(val *Coverage) {
	v.value = val
	v.isSet = true
}

func (v NullableCoverage) IsSet() bool {
	return v.isSet
}

func (v *NullableCoverage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoverage(val *Coverage) *NullableCoverage {
	return &NullableCoverage{value: val, isSet: true}
}

func (v NullableCoverage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoverage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
