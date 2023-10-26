package models

import (
	"encoding/json"
	"github.com/yenvanio/amadeus-go/amadeus/utils"
)

// checks if the DateTimeType type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DateTimeType{}

// DateTimeType struct for DateTimeType
type DateTimeType struct {
	// Dates are specified in the [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) YYYY-MM-DD format, e.g. 2018-12-25
	Date string `json:"date"`
	// Local time. hh:mm:ss format, e.g 10:30:00
	Time *string `json:"time,omitempty"`
}

// NewDateTimeType instantiates a new DateTimeType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDateTimeType(date string) *DateTimeType {
	this := DateTimeType{}
	this.Date = date
	return &this
}

// NewDateTimeTypeWithDefaults instantiates a new DateTimeType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDateTimeTypeWithDefaults() *DateTimeType {
	this := DateTimeType{}
	return &this
}

// GetDate returns the Date field value
func (o *DateTimeType) GetDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Date
}

// GetDateOk returns a tuple with the Date field value
// and a boolean to check if the value has been set.
func (o *DateTimeType) GetDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Date, true
}

// SetDate sets field value
func (o *DateTimeType) SetDate(v string) {
	o.Date = v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *DateTimeType) GetTime() string {
	if o == nil || utils.IsNil(o.Time) {
		var ret string
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DateTimeType) GetTimeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *DateTimeType) HasTime() bool {
	if o != nil && !utils.IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given string and assigns it to the Time field.
func (o *DateTimeType) SetTime(v string) {
	o.Time = &v
}

func (o DateTimeType) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DateTimeType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["date"] = o.Date
	if !utils.IsNil(o.Time) {
		toSerialize["time"] = o.Time
	}
	return toSerialize, nil
}

type NullableDateTimeType struct {
	value *DateTimeType
	isSet bool
}

func (v NullableDateTimeType) Get() *DateTimeType {
	return v.value
}

func (v *NullableDateTimeType) Set(val *DateTimeType) {
	v.value = val
	v.isSet = true
}

func (v NullableDateTimeType) IsSet() bool {
	return v.isSet
}

func (v *NullableDateTimeType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDateTimeType(val *DateTimeType) *NullableDateTimeType {
	return &NullableDateTimeType{value: val, isSet: true}
}

func (v NullableDateTimeType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDateTimeType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
