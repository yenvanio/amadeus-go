package models

import (
	"encoding/json"
	"github.com/yenvanio/amadeus-go/amadeus/utils"
)

// checks if the IssueSource type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &IssueSource{}

// IssueSource an object containing references to the source of the error
type IssueSource struct {
	// a JSON Pointer [RFC6901] to the associated entity in the request document
	Pointer *string `json:"pointer,omitempty"`
	// a string indicating which URI query parameter caused the issue
	Parameter *string `json:"parameter,omitempty"`
	// a string indicating an example of the right value
	Example *string `json:"example,omitempty"`
}

// NewIssueSource instantiates a new IssueSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIssueSource() *IssueSource {
	this := IssueSource{}
	return &this
}

// NewIssueSourceWithDefaults instantiates a new IssueSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIssueSourceWithDefaults() *IssueSource {
	this := IssueSource{}
	return &this
}

// GetPointer returns the Pointer field value if set, zero value otherwise.
func (o *IssueSource) GetPointer() string {
	if o == nil || utils.IsNil(o.Pointer) {
		var ret string
		return ret
	}
	return *o.Pointer
}

// GetPointerOk returns a tuple with the Pointer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueSource) GetPointerOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Pointer) {
		return nil, false
	}
	return o.Pointer, true
}

// HasPointer returns a boolean if a field has been set.
func (o *IssueSource) HasPointer() bool {
	if o != nil && !utils.IsNil(o.Pointer) {
		return true
	}

	return false
}

// SetPointer gets a reference to the given string and assigns it to the Pointer field.
func (o *IssueSource) SetPointer(v string) {
	o.Pointer = &v
}

// GetParameter returns the Parameter field value if set, zero value otherwise.
func (o *IssueSource) GetParameter() string {
	if o == nil || utils.IsNil(o.Parameter) {
		var ret string
		return ret
	}
	return *o.Parameter
}

// GetParameterOk returns a tuple with the Parameter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueSource) GetParameterOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Parameter) {
		return nil, false
	}
	return o.Parameter, true
}

// HasParameter returns a boolean if a field has been set.
func (o *IssueSource) HasParameter() bool {
	if o != nil && !utils.IsNil(o.Parameter) {
		return true
	}

	return false
}

// SetParameter gets a reference to the given string and assigns it to the Parameter field.
func (o *IssueSource) SetParameter(v string) {
	o.Parameter = &v
}

// GetExample returns the Example field value if set, zero value otherwise.
func (o *IssueSource) GetExample() string {
	if o == nil || utils.IsNil(o.Example) {
		var ret string
		return ret
	}
	return *o.Example
}

// GetExampleOk returns a tuple with the Example field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueSource) GetExampleOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Example) {
		return nil, false
	}
	return o.Example, true
}

// HasExample returns a boolean if a field has been set.
func (o *IssueSource) HasExample() bool {
	if o != nil && !utils.IsNil(o.Example) {
		return true
	}

	return false
}

// SetExample gets a reference to the given string and assigns it to the Example field.
func (o *IssueSource) SetExample(v string) {
	o.Example = &v
}

func (o IssueSource) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IssueSource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Pointer) {
		toSerialize["pointer"] = o.Pointer
	}
	if !utils.IsNil(o.Parameter) {
		toSerialize["parameter"] = o.Parameter
	}
	if !utils.IsNil(o.Example) {
		toSerialize["example"] = o.Example
	}
	return toSerialize, nil
}

type NullableIssueSource struct {
	value *IssueSource
	isSet bool
}

func (v NullableIssueSource) Get() *IssueSource {
	return v.value
}

func (v *NullableIssueSource) Set(val *IssueSource) {
	v.value = val
	v.isSet = true
}

func (v NullableIssueSource) IsSet() bool {
	return v.isSet
}

func (v *NullableIssueSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIssueSource(val *IssueSource) *NullableIssueSource {
	return &NullableIssueSource{value: val, isSet: true}
}

func (v NullableIssueSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIssueSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
