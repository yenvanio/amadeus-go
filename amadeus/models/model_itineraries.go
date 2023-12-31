package models

import (
	"encoding/json"
	"github.com/yenvanio/amadeus-go/amadeus/utils"
)

// checks if the Itineraries type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Itineraries{}

// Itineraries struct for Itineraries
type Itineraries struct {
	// duration in [ISO8601](https://en.wikipedia.org/wiki/ISO_8601) PnYnMnDTnHnMnS format, e.g. PT2H10M for a duration of 2h10m
	Duration *string   `json:"duration,omitempty"`
	Segments []Segment `json:"segments"`
}

// NewItineraries instantiates a new Itineraries object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewItineraries(segments []Segment) *Itineraries {
	this := Itineraries{}
	this.Segments = segments
	return &this
}

// NewItinerariesWithDefaults instantiates a new Itineraries object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewItinerariesWithDefaults() *Itineraries {
	this := Itineraries{}
	return &this
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *Itineraries) GetDuration() string {
	if o == nil || utils.IsNil(o.Duration) {
		var ret string
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Itineraries) GetDurationOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Duration) {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *Itineraries) HasDuration() bool {
	if o != nil && !utils.IsNil(o.Duration) {
		return true
	}

	return false
}

// SetDuration gets a reference to the given string and assigns it to the Duration field.
func (o *Itineraries) SetDuration(v string) {
	o.Duration = &v
}

// GetSegments returns the Segments field value
func (o *Itineraries) GetSegments() []Segment {
	if o == nil {
		var ret []Segment
		return ret
	}

	return o.Segments
}

// GetSegmentsOk returns a tuple with the Segments field value
// and a boolean to check if the value has been set.
func (o *Itineraries) GetSegmentsOk() ([]Segment, bool) {
	if o == nil {
		return nil, false
	}
	return o.Segments, true
}

// SetSegments sets field value
func (o *Itineraries) SetSegments(v []Segment) {
	o.Segments = v
}

func (o Itineraries) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Itineraries) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Duration) {
		toSerialize["duration"] = o.Duration
	}
	toSerialize["segments"] = o.Segments
	return toSerialize, nil
}

type NullableItineraries struct {
	value *Itineraries
	isSet bool
}

func (v NullableItineraries) Get() *Itineraries {
	return v.value
}

func (v *NullableItineraries) Set(val *Itineraries) {
	v.value = val
	v.isSet = true
}

func (v NullableItineraries) IsSet() bool {
	return v.isSet
}

func (v *NullableItineraries) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableItineraries(val *Itineraries) *NullableItineraries {
	return &NullableItineraries{value: val, isSet: true}
}

func (v NullableItineraries) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableItineraries) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
