package models

import (
	"encoding/json"
	"github.com/yenvanio/amadeus-go/amadeus/utils"
)

// checks if the Issue type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Issue{}

// Issue struct for Issue
type Issue struct {
	// the HTTP status code applicable to this error
	Status *int32 `json:"status,omitempty"`
	// an application-specific error code
	Code *int64 `json:"code,omitempty"`
	// a short summary of the error
	Title *string `json:"title,omitempty"`
	// explanation of the error
	Detail *string      `json:"detail,omitempty"`
	Source *IssueSource `json:"source,omitempty"`
}

// NewIssue instantiates a new Issue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIssue() *Issue {
	this := Issue{}
	return &this
}

// NewIssueWithDefaults instantiates a new Issue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIssueWithDefaults() *Issue {
	this := Issue{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Issue) GetStatus() int32 {
	if o == nil || utils.IsNil(o.Status) {
		var ret int32
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Issue) GetStatusOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Issue) HasStatus() bool {
	if o != nil && !utils.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int32 and assigns it to the Status field.
func (o *Issue) SetStatus(v int32) {
	o.Status = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *Issue) GetCode() int64 {
	if o == nil || utils.IsNil(o.Code) {
		var ret int64
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Issue) GetCodeOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *Issue) HasCode() bool {
	if o != nil && !utils.IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given int64 and assigns it to the Code field.
func (o *Issue) SetCode(v int64) {
	o.Code = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *Issue) GetTitle() string {
	if o == nil || utils.IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Issue) GetTitleOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *Issue) HasTitle() bool {
	if o != nil && !utils.IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *Issue) SetTitle(v string) {
	o.Title = &v
}

// GetDetail returns the Detail field value if set, zero value otherwise.
func (o *Issue) GetDetail() string {
	if o == nil || utils.IsNil(o.Detail) {
		var ret string
		return ret
	}
	return *o.Detail
}

// GetDetailOk returns a tuple with the Detail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Issue) GetDetailOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Detail) {
		return nil, false
	}
	return o.Detail, true
}

// HasDetail returns a boolean if a field has been set.
func (o *Issue) HasDetail() bool {
	if o != nil && !utils.IsNil(o.Detail) {
		return true
	}

	return false
}

// SetDetail gets a reference to the given string and assigns it to the Detail field.
func (o *Issue) SetDetail(v string) {
	o.Detail = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *Issue) GetSource() IssueSource {
	if o == nil || utils.IsNil(o.Source) {
		var ret IssueSource
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Issue) GetSourceOk() (*IssueSource, bool) {
	if o == nil || utils.IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *Issue) HasSource() bool {
	if o != nil && !utils.IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given IssueSource and assigns it to the Source field.
func (o *Issue) SetSource(v IssueSource) {
	o.Source = &v
}

func (o Issue) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Issue) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !utils.IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !utils.IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !utils.IsNil(o.Detail) {
		toSerialize["detail"] = o.Detail
	}
	if !utils.IsNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	return toSerialize, nil
}

type NullableIssue struct {
	value *Issue
	isSet bool
}

func (v NullableIssue) Get() *Issue {
	return v.value
}

func (v *NullableIssue) Set(val *Issue) {
	v.value = val
	v.isSet = true
}

func (v NullableIssue) IsSet() bool {
	return v.isSet
}

func (v *NullableIssue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIssue(val *Issue) *NullableIssue {
	return &NullableIssue{value: val, isSet: true}
}

func (v NullableIssue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIssue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
