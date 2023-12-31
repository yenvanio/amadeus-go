package models

import (
	"encoding/json"
	"fmt"
)

// TravelClass quality of service offered in the cabin where the seat is located in this flight. Economy, premium economy, business or first class
type TravelClass string

// List of TravelClass
const (
	ECONOMY         TravelClass = "ECONOMY"
	PREMIUM_ECONOMY TravelClass = "PREMIUM_ECONOMY"
	BUSINESS        TravelClass = "BUSINESS"
	FIRST           TravelClass = "FIRST"
)

// All allowed values of TravelClass enum
var AllowedTravelClassEnumValues = []TravelClass{
	"ECONOMY",
	"PREMIUM_ECONOMY",
	"BUSINESS",
	"FIRST",
}

func (v *TravelClass) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TravelClass(value)
	for _, existing := range AllowedTravelClassEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TravelClass", value)
}

// NewTravelClassFromValue returns a pointer to a valid TravelClass
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTravelClassFromValue(v string) (*TravelClass, error) {
	ev := TravelClass(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TravelClass: valid values are %v", v, AllowedTravelClassEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TravelClass) IsValid() bool {
	for _, existing := range AllowedTravelClassEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TravelClass value
func (v TravelClass) Ptr() *TravelClass {
	return &v
}

type NullableTravelClass struct {
	value *TravelClass
	isSet bool
}

func (v NullableTravelClass) Get() *TravelClass {
	return v.value
}

func (v *NullableTravelClass) Set(val *TravelClass) {
	v.value = val
	v.isSet = true
}

func (v NullableTravelClass) IsSet() bool {
	return v.isSet
}

func (v *NullableTravelClass) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTravelClass(val *TravelClass) *NullableTravelClass {
	return &NullableTravelClass{value: val, isSet: true}
}

func (v NullableTravelClass) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTravelClass) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
