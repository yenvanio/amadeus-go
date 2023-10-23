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

// checks if the PassengerConditions type satisfies the MappedNullable interface at compile time
var _ api.MappedNullable = &PassengerConditions{}

// PassengerConditions struct for PassengerConditions
type PassengerConditions struct {
	// Id of the traveler
	TravelerId *string `json:"travelerId,omitempty"`
	// If true, the gender is required for the concerned traveler for the creation of the flight-order
	GenderRequired *bool `json:"genderRequired,omitempty"`
	// If true, a document is required for the concerned traveler for the creation of the flight-order
	DocumentRequired *bool `json:"documentRequired,omitempty"`
	// If true, the isuance city of the document is required for the concerned traveler for the creation of the flight-order
	DocumentIssuanceCityRequired *bool `json:"documentIssuanceCityRequired,omitempty"`
	// If true, the date of bitrth is required for the concerned traveler for the creation of the flight-order
	DateOfBirthRequired *bool `json:"dateOfBirthRequired,omitempty"`
	// If true, the redress is required if any for the concerned traveler for the creation of the flight-order
	RedressRequiredIfAny *bool `json:"redressRequiredIfAny,omitempty"`
	// If true, the Air France discount is required for the concerned traveler for the creation of the flight-order
	AirFranceDiscountRequired *bool `json:"airFranceDiscountRequired,omitempty"`
	// If true, the Spanish resident discount is required for the concerned traveler for the creation of the flight-order
	SpanishResidentDiscountRequired *bool `json:"spanishResidentDiscountRequired,omitempty"`
	// If true, the address is required for the concerned traveler for the creation of the flight-order
	ResidenceRequired *bool `json:"residenceRequired,omitempty"`
}

// NewPassengerConditions instantiates a new PassengerConditions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPassengerConditions() *PassengerConditions {
	this := PassengerConditions{}
	return &this
}

// NewPassengerConditionsWithDefaults instantiates a new PassengerConditions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPassengerConditionsWithDefaults() *PassengerConditions {
	this := PassengerConditions{}
	return &this
}

// GetTravelerId returns the TravelerId field value if set, zero value otherwise.
func (o *PassengerConditions) GetTravelerId() string {
	if o == nil || api.IsNil(o.TravelerId) {
		var ret string
		return ret
	}
	return *o.TravelerId
}

// GetTravelerIdOk returns a tuple with the TravelerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PassengerConditions) GetTravelerIdOk() (*string, bool) {
	if o == nil || api.IsNil(o.TravelerId) {
		return nil, false
	}
	return o.TravelerId, true
}

// HasTravelerId returns a boolean if a field has been set.
func (o *PassengerConditions) HasTravelerId() bool {
	if o != nil && !api.IsNil(o.TravelerId) {
		return true
	}

	return false
}

// SetTravelerId gets a reference to the given string and assigns it to the TravelerId field.
func (o *PassengerConditions) SetTravelerId(v string) {
	o.TravelerId = &v
}

// GetGenderRequired returns the GenderRequired field value if set, zero value otherwise.
func (o *PassengerConditions) GetGenderRequired() bool {
	if o == nil || api.IsNil(o.GenderRequired) {
		var ret bool
		return ret
	}
	return *o.GenderRequired
}

// GetGenderRequiredOk returns a tuple with the GenderRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PassengerConditions) GetGenderRequiredOk() (*bool, bool) {
	if o == nil || api.IsNil(o.GenderRequired) {
		return nil, false
	}
	return o.GenderRequired, true
}

// HasGenderRequired returns a boolean if a field has been set.
func (o *PassengerConditions) HasGenderRequired() bool {
	if o != nil && !api.IsNil(o.GenderRequired) {
		return true
	}

	return false
}

// SetGenderRequired gets a reference to the given bool and assigns it to the GenderRequired field.
func (o *PassengerConditions) SetGenderRequired(v bool) {
	o.GenderRequired = &v
}

// GetDocumentRequired returns the DocumentRequired field value if set, zero value otherwise.
func (o *PassengerConditions) GetDocumentRequired() bool {
	if o == nil || api.IsNil(o.DocumentRequired) {
		var ret bool
		return ret
	}
	return *o.DocumentRequired
}

// GetDocumentRequiredOk returns a tuple with the DocumentRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PassengerConditions) GetDocumentRequiredOk() (*bool, bool) {
	if o == nil || api.IsNil(o.DocumentRequired) {
		return nil, false
	}
	return o.DocumentRequired, true
}

// HasDocumentRequired returns a boolean if a field has been set.
func (o *PassengerConditions) HasDocumentRequired() bool {
	if o != nil && !api.IsNil(o.DocumentRequired) {
		return true
	}

	return false
}

// SetDocumentRequired gets a reference to the given bool and assigns it to the DocumentRequired field.
func (o *PassengerConditions) SetDocumentRequired(v bool) {
	o.DocumentRequired = &v
}

// GetDocumentIssuanceCityRequired returns the DocumentIssuanceCityRequired field value if set, zero value otherwise.
func (o *PassengerConditions) GetDocumentIssuanceCityRequired() bool {
	if o == nil || api.IsNil(o.DocumentIssuanceCityRequired) {
		var ret bool
		return ret
	}
	return *o.DocumentIssuanceCityRequired
}

// GetDocumentIssuanceCityRequiredOk returns a tuple with the DocumentIssuanceCityRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PassengerConditions) GetDocumentIssuanceCityRequiredOk() (*bool, bool) {
	if o == nil || api.IsNil(o.DocumentIssuanceCityRequired) {
		return nil, false
	}
	return o.DocumentIssuanceCityRequired, true
}

// HasDocumentIssuanceCityRequired returns a boolean if a field has been set.
func (o *PassengerConditions) HasDocumentIssuanceCityRequired() bool {
	if o != nil && !api.IsNil(o.DocumentIssuanceCityRequired) {
		return true
	}

	return false
}

// SetDocumentIssuanceCityRequired gets a reference to the given bool and assigns it to the DocumentIssuanceCityRequired field.
func (o *PassengerConditions) SetDocumentIssuanceCityRequired(v bool) {
	o.DocumentIssuanceCityRequired = &v
}

// GetDateOfBirthRequired returns the DateOfBirthRequired field value if set, zero value otherwise.
func (o *PassengerConditions) GetDateOfBirthRequired() bool {
	if o == nil || api.IsNil(o.DateOfBirthRequired) {
		var ret bool
		return ret
	}
	return *o.DateOfBirthRequired
}

// GetDateOfBirthRequiredOk returns a tuple with the DateOfBirthRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PassengerConditions) GetDateOfBirthRequiredOk() (*bool, bool) {
	if o == nil || api.IsNil(o.DateOfBirthRequired) {
		return nil, false
	}
	return o.DateOfBirthRequired, true
}

// HasDateOfBirthRequired returns a boolean if a field has been set.
func (o *PassengerConditions) HasDateOfBirthRequired() bool {
	if o != nil && !api.IsNil(o.DateOfBirthRequired) {
		return true
	}

	return false
}

// SetDateOfBirthRequired gets a reference to the given bool and assigns it to the DateOfBirthRequired field.
func (o *PassengerConditions) SetDateOfBirthRequired(v bool) {
	o.DateOfBirthRequired = &v
}

// GetRedressRequiredIfAny returns the RedressRequiredIfAny field value if set, zero value otherwise.
func (o *PassengerConditions) GetRedressRequiredIfAny() bool {
	if o == nil || api.IsNil(o.RedressRequiredIfAny) {
		var ret bool
		return ret
	}
	return *o.RedressRequiredIfAny
}

// GetRedressRequiredIfAnyOk returns a tuple with the RedressRequiredIfAny field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PassengerConditions) GetRedressRequiredIfAnyOk() (*bool, bool) {
	if o == nil || api.IsNil(o.RedressRequiredIfAny) {
		return nil, false
	}
	return o.RedressRequiredIfAny, true
}

// HasRedressRequiredIfAny returns a boolean if a field has been set.
func (o *PassengerConditions) HasRedressRequiredIfAny() bool {
	if o != nil && !api.IsNil(o.RedressRequiredIfAny) {
		return true
	}

	return false
}

// SetRedressRequiredIfAny gets a reference to the given bool and assigns it to the RedressRequiredIfAny field.
func (o *PassengerConditions) SetRedressRequiredIfAny(v bool) {
	o.RedressRequiredIfAny = &v
}

// GetAirFranceDiscountRequired returns the AirFranceDiscountRequired field value if set, zero value otherwise.
func (o *PassengerConditions) GetAirFranceDiscountRequired() bool {
	if o == nil || api.IsNil(o.AirFranceDiscountRequired) {
		var ret bool
		return ret
	}
	return *o.AirFranceDiscountRequired
}

// GetAirFranceDiscountRequiredOk returns a tuple with the AirFranceDiscountRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PassengerConditions) GetAirFranceDiscountRequiredOk() (*bool, bool) {
	if o == nil || api.IsNil(o.AirFranceDiscountRequired) {
		return nil, false
	}
	return o.AirFranceDiscountRequired, true
}

// HasAirFranceDiscountRequired returns a boolean if a field has been set.
func (o *PassengerConditions) HasAirFranceDiscountRequired() bool {
	if o != nil && !api.IsNil(o.AirFranceDiscountRequired) {
		return true
	}

	return false
}

// SetAirFranceDiscountRequired gets a reference to the given bool and assigns it to the AirFranceDiscountRequired field.
func (o *PassengerConditions) SetAirFranceDiscountRequired(v bool) {
	o.AirFranceDiscountRequired = &v
}

// GetSpanishResidentDiscountRequired returns the SpanishResidentDiscountRequired field value if set, zero value otherwise.
func (o *PassengerConditions) GetSpanishResidentDiscountRequired() bool {
	if o == nil || api.IsNil(o.SpanishResidentDiscountRequired) {
		var ret bool
		return ret
	}
	return *o.SpanishResidentDiscountRequired
}

// GetSpanishResidentDiscountRequiredOk returns a tuple with the SpanishResidentDiscountRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PassengerConditions) GetSpanishResidentDiscountRequiredOk() (*bool, bool) {
	if o == nil || api.IsNil(o.SpanishResidentDiscountRequired) {
		return nil, false
	}
	return o.SpanishResidentDiscountRequired, true
}

// HasSpanishResidentDiscountRequired returns a boolean if a field has been set.
func (o *PassengerConditions) HasSpanishResidentDiscountRequired() bool {
	if o != nil && !api.IsNil(o.SpanishResidentDiscountRequired) {
		return true
	}

	return false
}

// SetSpanishResidentDiscountRequired gets a reference to the given bool and assigns it to the SpanishResidentDiscountRequired field.
func (o *PassengerConditions) SetSpanishResidentDiscountRequired(v bool) {
	o.SpanishResidentDiscountRequired = &v
}

// GetResidenceRequired returns the ResidenceRequired field value if set, zero value otherwise.
func (o *PassengerConditions) GetResidenceRequired() bool {
	if o == nil || api.IsNil(o.ResidenceRequired) {
		var ret bool
		return ret
	}
	return *o.ResidenceRequired
}

// GetResidenceRequiredOk returns a tuple with the ResidenceRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PassengerConditions) GetResidenceRequiredOk() (*bool, bool) {
	if o == nil || api.IsNil(o.ResidenceRequired) {
		return nil, false
	}
	return o.ResidenceRequired, true
}

// HasResidenceRequired returns a boolean if a field has been set.
func (o *PassengerConditions) HasResidenceRequired() bool {
	if o != nil && !api.IsNil(o.ResidenceRequired) {
		return true
	}

	return false
}

// SetResidenceRequired gets a reference to the given bool and assigns it to the ResidenceRequired field.
func (o *PassengerConditions) SetResidenceRequired(v bool) {
	o.ResidenceRequired = &v
}

func (o PassengerConditions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PassengerConditions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !api.IsNil(o.TravelerId) {
		toSerialize["travelerId"] = o.TravelerId
	}
	if !api.IsNil(o.GenderRequired) {
		toSerialize["genderRequired"] = o.GenderRequired
	}
	if !api.IsNil(o.DocumentRequired) {
		toSerialize["documentRequired"] = o.DocumentRequired
	}
	if !api.IsNil(o.DocumentIssuanceCityRequired) {
		toSerialize["documentIssuanceCityRequired"] = o.DocumentIssuanceCityRequired
	}
	if !api.IsNil(o.DateOfBirthRequired) {
		toSerialize["dateOfBirthRequired"] = o.DateOfBirthRequired
	}
	if !api.IsNil(o.RedressRequiredIfAny) {
		toSerialize["redressRequiredIfAny"] = o.RedressRequiredIfAny
	}
	if !api.IsNil(o.AirFranceDiscountRequired) {
		toSerialize["airFranceDiscountRequired"] = o.AirFranceDiscountRequired
	}
	if !api.IsNil(o.SpanishResidentDiscountRequired) {
		toSerialize["spanishResidentDiscountRequired"] = o.SpanishResidentDiscountRequired
	}
	if !api.IsNil(o.ResidenceRequired) {
		toSerialize["residenceRequired"] = o.ResidenceRequired
	}
	return toSerialize, nil
}

type NullablePassengerConditions struct {
	value *PassengerConditions
	isSet bool
}

func (v NullablePassengerConditions) Get() *PassengerConditions {
	return v.value
}

func (v *NullablePassengerConditions) Set(val *PassengerConditions) {
	v.value = val
	v.isSet = true
}

func (v NullablePassengerConditions) IsSet() bool {
	return v.isSet
}

func (v *NullablePassengerConditions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePassengerConditions(val *PassengerConditions) *NullablePassengerConditions {
	return &NullablePassengerConditions{value: val, isSet: true}
}

func (v NullablePassengerConditions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePassengerConditions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

