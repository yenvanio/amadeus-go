/*
Flight Offers Price

Before using this API, we recommend you read our **[Authorization Guide](https://developers.amadeus.com/self-service/apis-docs/guides/authorization-262)** for more information on how to generate an access token.   Please also be aware that our test environment is based on a subset of the production, if you are not returning any results try with big cities/airports like LON (London) or NYC (New-York).

API version: 1.3.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"
	"github.com/GIT_USER_ID/GIT_REPO_ID/api"
)

// checks if the Payment type satisfies the MappedNullable interface at compile time
var _ api.MappedNullable = &Payment{}

// Payment struct for Payment
type Payment struct {
	Brand *PaymentBrand `json:"brand,omitempty"`
	// The first 6 digits of the credit card
	BinNumber *int32 `json:"binNumber,omitempty"`
	// Id of the flightOffers to pay
	FlightOfferIds []string `json:"flightOfferIds,omitempty"`
}

// NewPayment instantiates a new Payment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPayment() *Payment {
	this := Payment{}
	return &this
}

// NewPaymentWithDefaults instantiates a new Payment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentWithDefaults() *Payment {
	this := Payment{}
	return &this
}

// GetBrand returns the Brand field value if set, zero value otherwise.
func (o *Payment) GetBrand() PaymentBrand {
	if o == nil || api.IsNil(o.Brand) {
		var ret PaymentBrand
		return ret
	}
	return *o.Brand
}

// GetBrandOk returns a tuple with the Brand field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Payment) GetBrandOk() (*PaymentBrand, bool) {
	if o == nil || api.IsNil(o.Brand) {
		return nil, false
	}
	return o.Brand, true
}

// HasBrand returns a boolean if a field has been set.
func (o *Payment) HasBrand() bool {
	if o != nil && !api.IsNil(o.Brand) {
		return true
	}

	return false
}

// SetBrand gets a reference to the given PaymentBrand and assigns it to the Brand field.
func (o *Payment) SetBrand(v PaymentBrand) {
	o.Brand = &v
}

// GetBinNumber returns the BinNumber field value if set, zero value otherwise.
func (o *Payment) GetBinNumber() int32 {
	if o == nil || api.IsNil(o.BinNumber) {
		var ret int32
		return ret
	}
	return *o.BinNumber
}

// GetBinNumberOk returns a tuple with the BinNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Payment) GetBinNumberOk() (*int32, bool) {
	if o == nil || api.IsNil(o.BinNumber) {
		return nil, false
	}
	return o.BinNumber, true
}

// HasBinNumber returns a boolean if a field has been set.
func (o *Payment) HasBinNumber() bool {
	if o != nil && !api.IsNil(o.BinNumber) {
		return true
	}

	return false
}

// SetBinNumber gets a reference to the given int32 and assigns it to the BinNumber field.
func (o *Payment) SetBinNumber(v int32) {
	o.BinNumber = &v
}

// GetFlightOfferIds returns the FlightOfferIds field value if set, zero value otherwise.
func (o *Payment) GetFlightOfferIds() []string {
	if o == nil || api.IsNil(o.FlightOfferIds) {
		var ret []string
		return ret
	}
	return o.FlightOfferIds
}

// GetFlightOfferIdsOk returns a tuple with the FlightOfferIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Payment) GetFlightOfferIdsOk() ([]string, bool) {
	if o == nil || api.IsNil(o.FlightOfferIds) {
		return nil, false
	}
	return o.FlightOfferIds, true
}

// HasFlightOfferIds returns a boolean if a field has been set.
func (o *Payment) HasFlightOfferIds() bool {
	if o != nil && !api.IsNil(o.FlightOfferIds) {
		return true
	}

	return false
}

// SetFlightOfferIds gets a reference to the given []string and assigns it to the FlightOfferIds field.
func (o *Payment) SetFlightOfferIds(v []string) {
	o.FlightOfferIds = v
}

func (o Payment) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Payment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !api.IsNil(o.Brand) {
		toSerialize["brand"] = o.Brand
	}
	if !api.IsNil(o.BinNumber) {
		toSerialize["binNumber"] = o.BinNumber
	}
	if !api.IsNil(o.FlightOfferIds) {
		toSerialize["flightOfferIds"] = o.FlightOfferIds
	}
	return toSerialize, nil
}

type NullablePayment struct {
	value *Payment
	isSet bool
}

func (v NullablePayment) Get() *Payment {
	return v.value
}

func (v *NullablePayment) Set(val *Payment) {
	v.value = val
	v.isSet = true
}

func (v NullablePayment) IsSet() bool {
	return v.isSet
}

func (v *NullablePayment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePayment(val *Payment) *NullablePayment {
	return &NullablePayment{value: val, isSet: true}
}

func (v NullablePayment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePayment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


