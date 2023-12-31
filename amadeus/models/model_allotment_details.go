package models

import (
	"encoding/json"
	"github.com/yenvanio/amadeus-go/amadeus/utils"
)

// checks if the AllotmentDetails type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &AllotmentDetails{}

// AllotmentDetails struct for AllotmentDetails
type AllotmentDetails struct {
	// The tour name agreed for this specific allotment.
	TourName *string `json:"tourName,omitempty"`
	// The tour reference agreed for this specific allotment.
	TourReference *string `json:"tourReference,omitempty"`
}

// NewAllotmentDetails instantiates a new AllotmentDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAllotmentDetails() *AllotmentDetails {
	this := AllotmentDetails{}
	return &this
}

// NewAllotmentDetailsWithDefaults instantiates a new AllotmentDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAllotmentDetailsWithDefaults() *AllotmentDetails {
	this := AllotmentDetails{}
	return &this
}

// GetTourName returns the TourName field value if set, zero value otherwise.
func (o *AllotmentDetails) GetTourName() string {
	if o == nil || utils.IsNil(o.TourName) {
		var ret string
		return ret
	}
	return *o.TourName
}

// GetTourNameOk returns a tuple with the TourName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllotmentDetails) GetTourNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.TourName) {
		return nil, false
	}
	return o.TourName, true
}

// HasTourName returns a boolean if a field has been set.
func (o *AllotmentDetails) HasTourName() bool {
	if o != nil && !utils.IsNil(o.TourName) {
		return true
	}

	return false
}

// SetTourName gets a reference to the given string and assigns it to the TourName field.
func (o *AllotmentDetails) SetTourName(v string) {
	o.TourName = &v
}

// GetTourReference returns the TourReference field value if set, zero value otherwise.
func (o *AllotmentDetails) GetTourReference() string {
	if o == nil || utils.IsNil(o.TourReference) {
		var ret string
		return ret
	}
	return *o.TourReference
}

// GetTourReferenceOk returns a tuple with the TourReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllotmentDetails) GetTourReferenceOk() (*string, bool) {
	if o == nil || utils.IsNil(o.TourReference) {
		return nil, false
	}
	return o.TourReference, true
}

// HasTourReference returns a boolean if a field has been set.
func (o *AllotmentDetails) HasTourReference() bool {
	if o != nil && !utils.IsNil(o.TourReference) {
		return true
	}

	return false
}

// SetTourReference gets a reference to the given string and assigns it to the TourReference field.
func (o *AllotmentDetails) SetTourReference(v string) {
	o.TourReference = &v
}

func (o AllotmentDetails) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AllotmentDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.TourName) {
		toSerialize["tourName"] = o.TourName
	}
	if !utils.IsNil(o.TourReference) {
		toSerialize["tourReference"] = o.TourReference
	}
	return toSerialize, nil
}

type NullableAllotmentDetails struct {
	value *AllotmentDetails
	isSet bool
}

func (v NullableAllotmentDetails) Get() *AllotmentDetails {
	return v.value
}

func (v *NullableAllotmentDetails) Set(val *AllotmentDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableAllotmentDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableAllotmentDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAllotmentDetails(val *AllotmentDetails) *NullableAllotmentDetails {
	return &NullableAllotmentDetails{value: val, isSet: true}
}

func (v NullableAllotmentDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAllotmentDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
