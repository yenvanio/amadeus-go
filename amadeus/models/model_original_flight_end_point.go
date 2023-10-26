package models

import (
	"encoding/json"
	"github.com/yenvanio/amadeus-go/amadeus/utils"
)

// checks if the OriginalFlightEndPoint type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &OriginalFlightEndPoint{}

// OriginalFlightEndPoint departure or arrival information
type OriginalFlightEndPoint struct {
	// [IATA airline codes](http://www.iata.org/publications/Pages/code-search.aspx)
	IataCode *string `json:"iataCode,omitempty"`
	// terminal name / number
	Terminal *string `json:"terminal,omitempty"`
}

// NewOriginalFlightEndPoint instantiates a new OriginalFlightEndPoint object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOriginalFlightEndPoint() *OriginalFlightEndPoint {
	this := OriginalFlightEndPoint{}
	return &this
}

// NewOriginalFlightEndPointWithDefaults instantiates a new OriginalFlightEndPoint object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOriginalFlightEndPointWithDefaults() *OriginalFlightEndPoint {
	this := OriginalFlightEndPoint{}
	return &this
}

// GetIataCode returns the IataCode field value if set, zero value otherwise.
func (o *OriginalFlightEndPoint) GetIataCode() string {
	if o == nil || utils.IsNil(o.IataCode) {
		var ret string
		return ret
	}
	return *o.IataCode
}

// GetIataCodeOk returns a tuple with the IataCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginalFlightEndPoint) GetIataCodeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.IataCode) {
		return nil, false
	}
	return o.IataCode, true
}

// HasIataCode returns a boolean if a field has been set.
func (o *OriginalFlightEndPoint) HasIataCode() bool {
	if o != nil && !utils.IsNil(o.IataCode) {
		return true
	}

	return false
}

// SetIataCode gets a reference to the given string and assigns it to the IataCode field.
func (o *OriginalFlightEndPoint) SetIataCode(v string) {
	o.IataCode = &v
}

// GetTerminal returns the Terminal field value if set, zero value otherwise.
func (o *OriginalFlightEndPoint) GetTerminal() string {
	if o == nil || utils.IsNil(o.Terminal) {
		var ret string
		return ret
	}
	return *o.Terminal
}

// GetTerminalOk returns a tuple with the Terminal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginalFlightEndPoint) GetTerminalOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Terminal) {
		return nil, false
	}
	return o.Terminal, true
}

// HasTerminal returns a boolean if a field has been set.
func (o *OriginalFlightEndPoint) HasTerminal() bool {
	if o != nil && !utils.IsNil(o.Terminal) {
		return true
	}

	return false
}

// SetTerminal gets a reference to the given string and assigns it to the Terminal field.
func (o *OriginalFlightEndPoint) SetTerminal(v string) {
	o.Terminal = &v
}

func (o OriginalFlightEndPoint) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OriginalFlightEndPoint) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.IataCode) {
		toSerialize["iataCode"] = o.IataCode
	}
	if !utils.IsNil(o.Terminal) {
		toSerialize["terminal"] = o.Terminal
	}
	return toSerialize, nil
}

type NullableOriginalFlightEndPoint struct {
	value *OriginalFlightEndPoint
	isSet bool
}

func (v NullableOriginalFlightEndPoint) Get() *OriginalFlightEndPoint {
	return v.value
}

func (v *NullableOriginalFlightEndPoint) Set(val *OriginalFlightEndPoint) {
	v.value = val
	v.isSet = true
}

func (v NullableOriginalFlightEndPoint) IsSet() bool {
	return v.isSet
}

func (v *NullableOriginalFlightEndPoint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOriginalFlightEndPoint(val *OriginalFlightEndPoint) *NullableOriginalFlightEndPoint {
	return &NullableOriginalFlightEndPoint{value: val, isSet: true}
}

func (v NullableOriginalFlightEndPoint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOriginalFlightEndPoint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
