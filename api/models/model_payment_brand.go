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

// PaymentBrand credit card brand
type PaymentBrand string

// List of PaymentBrand
const (
	VISA               PaymentBrand = "VISA"
	AMERICAN_EXPRESS   PaymentBrand = "AMERICAN_EXPRESS"
	MASTERCARD         PaymentBrand = "MASTERCARD"
	VISA_ELECTRON      PaymentBrand = "VISA_ELECTRON"
	VISA_DEBIT         PaymentBrand = "VISA_DEBIT"
	MASTERCARD_DEBIT   PaymentBrand = "MASTERCARD_DEBIT"
	MAESTRO            PaymentBrand = "MAESTRO"
	DINERS             PaymentBrand = "DINERS"
	MASTERCARD_IXARIS  PaymentBrand = "MASTERCARD_IXARIS"
	VISA_IXARIS        PaymentBrand = "VISA_IXARIS"
	MASTERCARD_AIRPLUS PaymentBrand = "MASTERCARD_AIRPLUS"
	UATP_AIRPLUS       PaymentBrand = "UATP_AIRPLUS"
)

// All allowed values of PaymentBrand enum
var AllowedPaymentBrandEnumValues = []PaymentBrand{
	"VISA",
	"AMERICAN_EXPRESS",
	"MASTERCARD",
	"VISA_ELECTRON",
	"VISA_DEBIT",
	"MASTERCARD_DEBIT",
	"MAESTRO",
	"DINERS",
	"MASTERCARD_IXARIS",
	"VISA_IXARIS",
	"MASTERCARD_AIRPLUS",
	"UATP_AIRPLUS",
}

func (v *PaymentBrand) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PaymentBrand(value)
	for _, existing := range AllowedPaymentBrandEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PaymentBrand", value)
}

// NewPaymentBrandFromValue returns a pointer to a valid PaymentBrand
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPaymentBrandFromValue(v string) (*PaymentBrand, error) {
	ev := PaymentBrand(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PaymentBrand: valid values are %v", v, AllowedPaymentBrandEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PaymentBrand) IsValid() bool {
	for _, existing := range AllowedPaymentBrandEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PaymentBrand value
func (v PaymentBrand) Ptr() *PaymentBrand {
	return &v
}

type NullablePaymentBrand struct {
	value *PaymentBrand
	isSet bool
}

func (v NullablePaymentBrand) Get() *PaymentBrand {
	return v.value
}

func (v *NullablePaymentBrand) Set(val *PaymentBrand) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentBrand) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentBrand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentBrand(val *PaymentBrand) *NullablePaymentBrand {
	return &NullablePaymentBrand{value: val, isSet: true}
}

func (v NullablePaymentBrand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentBrand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
