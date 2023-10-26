package models

import (
	"encoding/json"
	"fmt"
)

// TravelerPricingFareOption option specifying a group of fares, which may be valid under certain conditons Can be used to specify special fare discount for a passenger
type TravelerPricingFareOption string

// List of TravelerPricingFareOption
const (
	STANDARD_TRAVELER_PRICING_FARE_OPTION TravelerPricingFareOption = "STANDARD"
	INCLUSIVE_TOUR                        TravelerPricingFareOption = "INCLUSIVE_TOUR"
	SPANISH_MELILLA_RESIDENT              TravelerPricingFareOption = "SPANISH_MELILLA_RESIDENT"
	SPANISH_CEUTA_RESIDENT                TravelerPricingFareOption = "SPANISH_CEUTA_RESIDENT"
	SPANISH_CANARY_RESIDENT               TravelerPricingFareOption = "SPANISH_CANARY_RESIDENT"
	SPANISH_BALEARIC_RESIDENT             TravelerPricingFareOption = "SPANISH_BALEARIC_RESIDENT"
	AIR_FRANCE_METROPOLITAN_DISCOUNT_PASS TravelerPricingFareOption = "AIR_FRANCE_METROPOLITAN_DISCOUNT_PASS"
	AIR_FRANCE_DOM_DISCOUNT_PASS          TravelerPricingFareOption = "AIR_FRANCE_DOM_DISCOUNT_PASS"
	AIR_FRANCE_COMBINED_DISCOUNT_PASS     TravelerPricingFareOption = "AIR_FRANCE_COMBINED_DISCOUNT_PASS"
	AIR_FRANCE_FAMILY                     TravelerPricingFareOption = "AIR_FRANCE_FAMILY"
	ADULT_WITH_COMPANION                  TravelerPricingFareOption = "ADULT_WITH_COMPANION"
	COMPANION                             TravelerPricingFareOption = "COMPANION"
)

// All allowed values of TravelerPricingFareOption enum
var AllowedTravelerPricingFareOptionEnumValues = []TravelerPricingFareOption{
	"STANDARD",
	"INCLUSIVE_TOUR",
	"SPANISH_MELILLA_RESIDENT",
	"SPANISH_CEUTA_RESIDENT",
	"SPANISH_CANARY_RESIDENT",
	"SPANISH_BALEARIC_RESIDENT",
	"AIR_FRANCE_METROPOLITAN_DISCOUNT_PASS",
	"AIR_FRANCE_DOM_DISCOUNT_PASS",
	"AIR_FRANCE_COMBINED_DISCOUNT_PASS",
	"AIR_FRANCE_FAMILY",
	"ADULT_WITH_COMPANION",
	"COMPANION",
}

func (v *TravelerPricingFareOption) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TravelerPricingFareOption(value)
	for _, existing := range AllowedTravelerPricingFareOptionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TravelerPricingFareOption", value)
}

// NewTravelerPricingFareOptionFromValue returns a pointer to a valid TravelerPricingFareOption
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTravelerPricingFareOptionFromValue(v string) (*TravelerPricingFareOption, error) {
	ev := TravelerPricingFareOption(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TravelerPricingFareOption: valid values are %v", v, AllowedTravelerPricingFareOptionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TravelerPricingFareOption) IsValid() bool {
	for _, existing := range AllowedTravelerPricingFareOptionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TravelerPricingFareOption value
func (v TravelerPricingFareOption) Ptr() *TravelerPricingFareOption {
	return &v
}

type NullableTravelerPricingFareOption struct {
	value *TravelerPricingFareOption
	isSet bool
}

func (v NullableTravelerPricingFareOption) Get() *TravelerPricingFareOption {
	return v.value
}

func (v *NullableTravelerPricingFareOption) Set(val *TravelerPricingFareOption) {
	v.value = val
	v.isSet = true
}

func (v NullableTravelerPricingFareOption) IsSet() bool {
	return v.isSet
}

func (v *NullableTravelerPricingFareOption) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTravelerPricingFareOption(val *TravelerPricingFareOption) *NullableTravelerPricingFareOption {
	return &NullableTravelerPricingFareOption{value: val, isSet: true}
}

func (v NullableTravelerPricingFareOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTravelerPricingFareOption) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
