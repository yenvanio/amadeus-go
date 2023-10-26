package models

import (
	"encoding/json"
	"github.com/yenvanio/amadeus-go/amadeus/utils"
)

// checks if the Error400 type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Error400{}

// Error400 struct for Error400
type Error400 struct {
	Errors []Issue `json:"errors"`
}

// NewError400 instantiates a new Error400 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewError400(errors []Issue) *Error400 {
	this := Error400{}
	this.Errors = errors
	return &this
}

// NewError400WithDefaults instantiates a new Error400 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewError400WithDefaults() *Error400 {
	this := Error400{}
	return &this
}

// GetErrors returns the Errors field value
func (o *Error400) GetErrors() []Issue {
	if o == nil {
		var ret []Issue
		return ret
	}

	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value
// and a boolean to check if the value has been set.
func (o *Error400) GetErrorsOk() ([]Issue, bool) {
	if o == nil {
		return nil, false
	}
	return o.Errors, true
}

// SetErrors sets field value
func (o *Error400) SetErrors(v []Issue) {
	o.Errors = v
}

func (o Error400) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Error400) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["errors"] = o.Errors
	return toSerialize, nil
}

type NullableError400 struct {
	value *Error400
	isSet bool
}

func (v NullableError400) Get() *Error400 {
	return v.value
}

func (v *NullableError400) Set(val *Error400) {
	v.value = val
	v.isSet = true
}

func (v NullableError400) IsSet() bool {
	return v.isSet
}

func (v *NullableError400) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableError400(val *Error400) *NullableError400 {
	return &NullableError400{value: val, isSet: true}
}

func (v NullableError400) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableError400) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
