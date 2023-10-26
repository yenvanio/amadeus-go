package models

import (
	"encoding/json"
	"fmt"
)

// FlightOfferSource source of the flight offer
type FlightOfferSource string

// List of FlightOfferSource
const (
	GDS FlightOfferSource = "GDS"
)

// All allowed values of FlightOfferSource enum
var AllowedFlightOfferSourceEnumValues = []FlightOfferSource{
	"GDS",
}

func (v *FlightOfferSource) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FlightOfferSource(value)
	for _, existing := range AllowedFlightOfferSourceEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FlightOfferSource", value)
}

// NewFlightOfferSourceFromValue returns a pointer to a valid FlightOfferSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFlightOfferSourceFromValue(v string) (*FlightOfferSource, error) {
	ev := FlightOfferSource(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FlightOfferSource: valid values are %v", v, AllowedFlightOfferSourceEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FlightOfferSource) IsValid() bool {
	for _, existing := range AllowedFlightOfferSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FlightOfferSource value
func (v FlightOfferSource) Ptr() *FlightOfferSource {
	return &v
}

type NullableFlightOfferSource struct {
	value *FlightOfferSource
	isSet bool
}

func (v NullableFlightOfferSource) Get() *FlightOfferSource {
	return v.value
}

func (v *NullableFlightOfferSource) Set(val *FlightOfferSource) {
	v.value = val
	v.isSet = true
}

func (v NullableFlightOfferSource) IsSet() bool {
	return v.isSet
}

func (v *NullableFlightOfferSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFlightOfferSource(val *FlightOfferSource) *NullableFlightOfferSource {
	return &NullableFlightOfferSource{value: val, isSet: true}
}

func (v NullableFlightOfferSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFlightOfferSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
