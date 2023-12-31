package models

import (
	"encoding/json"
	"github.com/yenvanio/amadeus-go/amadeus/utils"
)

// checks if the DateTimeRange type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DateTimeRange{}

// DateTimeRange struct for DateTimeRange
type DateTimeRange struct {
	// Dates are specified in the [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) YYYY-MM-DD format, e.g. 2018-12-25
	Date string `json:"date"`
	// Local time. hh:mm:ss format, e.g 10:30:00
	Time *string `json:"time,omitempty"`
	// Either 1, 2 or 3 extra days around the local date (IxD for +/- x days - Ex: I3D), Either 1 to 3 days after the local date (PxD for +x days - Ex: P3D), or 1 to 3 days before the local date (MxD for -x days - Ex: M3D)  Can not be combined with \"originRadius\" or \"destinationRadius\".
	DateWindow *string `json:"dateWindow,omitempty"`
	// 1 to 12 hours around (both +and -) the local time. Possibly limited by the number of extra days when specified, i.e.  in some situations, it may not be used to exceed the maximum date range. [1-12]H format, e.g. 6H  Can not be combined with \"originRadius\" or \"destinationRadius\".
	TimeWindow *string `json:"timeWindow,omitempty"`
}

// NewDateTimeRange instantiates a new DateTimeRange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDateTimeRange(date string) *DateTimeRange {
	this := DateTimeRange{}
	this.Date = date
	return &this
}

// NewDateTimeRangeWithDefaults instantiates a new DateTimeRange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDateTimeRangeWithDefaults() *DateTimeRange {
	this := DateTimeRange{}
	return &this
}

// GetDate returns the Date field value
func (o *DateTimeRange) GetDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Date
}

// GetDateOk returns a tuple with the Date field value
// and a boolean to check if the value has been set.
func (o *DateTimeRange) GetDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Date, true
}

// SetDate sets field value
func (o *DateTimeRange) SetDate(v string) {
	o.Date = v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *DateTimeRange) GetTime() string {
	if o == nil || utils.IsNil(o.Time) {
		var ret string
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DateTimeRange) GetTimeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *DateTimeRange) HasTime() bool {
	if o != nil && !utils.IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given string and assigns it to the Time field.
func (o *DateTimeRange) SetTime(v string) {
	o.Time = &v
}

// GetDateWindow returns the DateWindow field value if set, zero value otherwise.
func (o *DateTimeRange) GetDateWindow() string {
	if o == nil || utils.IsNil(o.DateWindow) {
		var ret string
		return ret
	}
	return *o.DateWindow
}

// GetDateWindowOk returns a tuple with the DateWindow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DateTimeRange) GetDateWindowOk() (*string, bool) {
	if o == nil || utils.IsNil(o.DateWindow) {
		return nil, false
	}
	return o.DateWindow, true
}

// HasDateWindow returns a boolean if a field has been set.
func (o *DateTimeRange) HasDateWindow() bool {
	if o != nil && !utils.IsNil(o.DateWindow) {
		return true
	}

	return false
}

// SetDateWindow gets a reference to the given string and assigns it to the DateWindow field.
func (o *DateTimeRange) SetDateWindow(v string) {
	o.DateWindow = &v
}

// GetTimeWindow returns the TimeWindow field value if set, zero value otherwise.
func (o *DateTimeRange) GetTimeWindow() string {
	if o == nil || utils.IsNil(o.TimeWindow) {
		var ret string
		return ret
	}
	return *o.TimeWindow
}

// GetTimeWindowOk returns a tuple with the TimeWindow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DateTimeRange) GetTimeWindowOk() (*string, bool) {
	if o == nil || utils.IsNil(o.TimeWindow) {
		return nil, false
	}
	return o.TimeWindow, true
}

// HasTimeWindow returns a boolean if a field has been set.
func (o *DateTimeRange) HasTimeWindow() bool {
	if o != nil && !utils.IsNil(o.TimeWindow) {
		return true
	}

	return false
}

// SetTimeWindow gets a reference to the given string and assigns it to the TimeWindow field.
func (o *DateTimeRange) SetTimeWindow(v string) {
	o.TimeWindow = &v
}

func (o DateTimeRange) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DateTimeRange) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["date"] = o.Date
	if !utils.IsNil(o.Time) {
		toSerialize["time"] = o.Time
	}
	if !utils.IsNil(o.DateWindow) {
		toSerialize["dateWindow"] = o.DateWindow
	}
	if !utils.IsNil(o.TimeWindow) {
		toSerialize["timeWindow"] = o.TimeWindow
	}
	return toSerialize, nil
}

type NullableDateTimeRange struct {
	value *DateTimeRange
	isSet bool
}

func (v NullableDateTimeRange) Get() *DateTimeRange {
	return v.value
}

func (v *NullableDateTimeRange) Set(val *DateTimeRange) {
	v.value = val
	v.isSet = true
}

func (v NullableDateTimeRange) IsSet() bool {
	return v.isSet
}

func (v *NullableDateTimeRange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDateTimeRange(val *DateTimeRange) *NullableDateTimeRange {
	return &NullableDateTimeRange{value: val, isSet: true}
}

func (v NullableDateTimeRange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDateTimeRange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
