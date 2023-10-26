package models

import (
	"encoding/json"
	"github.com/yenvanio/amadeus-go/amadeus/utils"
)

// checks if the OneWayCombinations type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &OneWayCombinations{}

// OneWayCombinations struct for OneWayCombinations
type OneWayCombinations struct {
	OriginDestinationId *string  `json:"originDestinationId,omitempty"`
	FlightOfferIds      []string `json:"flightOfferIds,omitempty"`
}

// NewOneWayCombinations instantiates a new OneWayCombinations object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOneWayCombinations() *OneWayCombinations {
	this := OneWayCombinations{}
	return &this
}

// NewOneWayCombinationsWithDefaults instantiates a new OneWayCombinations object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOneWayCombinationsWithDefaults() *OneWayCombinations {
	this := OneWayCombinations{}
	return &this
}

// GetOriginDestinationId returns the OriginDestinationId field value if set, zero value otherwise.
func (o *OneWayCombinations) GetOriginDestinationId() string {
	if o == nil || utils.IsNil(o.OriginDestinationId) {
		var ret string
		return ret
	}
	return *o.OriginDestinationId
}

// GetOriginDestinationIdOk returns a tuple with the OriginDestinationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OneWayCombinations) GetOriginDestinationIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.OriginDestinationId) {
		return nil, false
	}
	return o.OriginDestinationId, true
}

// HasOriginDestinationId returns a boolean if a field has been set.
func (o *OneWayCombinations) HasOriginDestinationId() bool {
	if o != nil && !utils.IsNil(o.OriginDestinationId) {
		return true
	}

	return false
}

// SetOriginDestinationId gets a reference to the given string and assigns it to the OriginDestinationId field.
func (o *OneWayCombinations) SetOriginDestinationId(v string) {
	o.OriginDestinationId = &v
}

// GetFlightOfferIds returns the FlightOfferIds field value if set, zero value otherwise.
func (o *OneWayCombinations) GetFlightOfferIds() []string {
	if o == nil || utils.IsNil(o.FlightOfferIds) {
		var ret []string
		return ret
	}
	return o.FlightOfferIds
}

// GetFlightOfferIdsOk returns a tuple with the FlightOfferIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OneWayCombinations) GetFlightOfferIdsOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.FlightOfferIds) {
		return nil, false
	}
	return o.FlightOfferIds, true
}

// HasFlightOfferIds returns a boolean if a field has been set.
func (o *OneWayCombinations) HasFlightOfferIds() bool {
	if o != nil && !utils.IsNil(o.FlightOfferIds) {
		return true
	}

	return false
}

// SetFlightOfferIds gets a reference to the given []string and assigns it to the FlightOfferIds field.
func (o *OneWayCombinations) SetFlightOfferIds(v []string) {
	o.FlightOfferIds = v
}

func (o OneWayCombinations) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OneWayCombinations) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.OriginDestinationId) {
		toSerialize["originDestinationId"] = o.OriginDestinationId
	}
	if !utils.IsNil(o.FlightOfferIds) {
		toSerialize["flightOfferIds"] = o.FlightOfferIds
	}
	return toSerialize, nil
}

type NullableOneWayCombinations struct {
	value *OneWayCombinations
	isSet bool
}

func (v NullableOneWayCombinations) Get() *OneWayCombinations {
	return v.value
}

func (v *NullableOneWayCombinations) Set(val *OneWayCombinations) {
	v.value = val
	v.isSet = true
}

func (v NullableOneWayCombinations) IsSet() bool {
	return v.isSet
}

func (v *NullableOneWayCombinations) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOneWayCombinations(val *OneWayCombinations) *NullableOneWayCombinations {
	return &NullableOneWayCombinations{value: val, isSet: true}
}

func (v NullableOneWayCombinations) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOneWayCombinations) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
