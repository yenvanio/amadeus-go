package models

import (
	"encoding/json"
	"github.com/yenvanio/amadeus-go/amadeus/utils"
)

// checks if the Error500 type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Error500{}

// Error500 struct for Error500
type Error500 struct {
	Errors []Issue `json:"errors"`
}

// NewError500 instantiates a new Error500 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewError500(errors []Issue) *Error500 {
	this := Error500{}
	this.Errors = errors
	return &this
}

// NewError500WithDefaults instantiates a new Error500 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewError500WithDefaults() *Error500 {
	this := Error500{}
	return &this
}

// GetErrors returns the Errors field value
func (o *Error500) GetErrors() []Issue {
	if o == nil {
		var ret []Issue
		return ret
	}

	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value
// and a boolean to check if the value has been set.
func (o *Error500) GetErrorsOk() ([]Issue, bool) {
	if o == nil {
		return nil, false
	}
	return o.Errors, true
}

// SetErrors sets field value
func (o *Error500) SetErrors(v []Issue) {
	o.Errors = v
}

func (o Error500) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Error500) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["errors"] = o.Errors
	return toSerialize, nil
}

type NullableError500 struct {
	value *Error500
	isSet bool
}

func (v NullableError500) Get() *Error500 {
	return v.value
}

func (v *NullableError500) Set(val *Error500) {
	v.value = val
	v.isSet = true
}

func (v NullableError500) IsSet() bool {
	return v.isSet
}

func (v *NullableError500) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableError500(val *Error500) *NullableError500 {
	return &NullableError500{value: val, isSet: true}
}

func (v NullableError500) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableError500) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
