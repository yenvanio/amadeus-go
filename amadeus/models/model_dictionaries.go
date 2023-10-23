/*
Flight Offers Search

 Before using this API, we recommend you read our **[Authorization Guide](https://developers.amadeus.com/self-service/apis-docs/guides/authorization-262)** for more information on how to generate an access token.   Please also be aware that our test environment is based on a subset of the production, if you are not returning any results try with big cities/airports like LON (London) or NYC (New-York).

API version: 2.9.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"
	"github.com/yenvanio/amadeus-go/amadeus"
)

// checks if the Dictionaries type satisfies the MappedNullable interface at compile time
var _ amadeus.MappedNullable = &Dictionaries{}

// Dictionaries struct for Dictionaries
type Dictionaries struct {
	Locations *map[string]LocationValue `json:"locations,omitempty"`
	Aircraft *map[string]string         `json:"aircraft,omitempty"`
	Currencies *map[string]string `json:"currencies,omitempty"`
	Carriers *map[string]string `json:"carriers,omitempty"`
}

// NewDictionaries instantiates a new Dictionaries object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDictionaries() *Dictionaries {
	this := Dictionaries{}
	return &this
}

// NewDictionariesWithDefaults instantiates a new Dictionaries object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDictionariesWithDefaults() *Dictionaries {
	this := Dictionaries{}
	return &this
}

// GetLocations returns the Locations field value if set, zero value otherwise.
func (o *Dictionaries) GetLocations() map[string]LocationValue {
	if o == nil || amadeus.IsNil(o.Locations) {
		var ret map[string]LocationValue
		return ret
	}
	return *o.Locations
}

// GetLocationsOk returns a tuple with the Locations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dictionaries) GetLocationsOk() (*map[string]LocationValue, bool) {
	if o == nil || amadeus.IsNil(o.Locations) {
		return nil, false
	}
	return o.Locations, true
}

// HasLocations returns a boolean if a field has been set.
func (o *Dictionaries) HasLocations() bool {
	if o != nil && !amadeus.IsNil(o.Locations) {
		return true
	}

	return false
}

// SetLocations gets a reference to the given map[string]LocationValue and assigns it to the Locations field.
func (o *Dictionaries) SetLocations(v map[string]LocationValue) {
	o.Locations = &v
}

// GetAircraft returns the Aircraft field value if set, zero value otherwise.
func (o *Dictionaries) GetAircraft() map[string]string {
	if o == nil || amadeus.IsNil(o.Aircraft) {
		var ret map[string]string
		return ret
	}
	return *o.Aircraft
}

// GetAircraftOk returns a tuple with the Aircraft field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dictionaries) GetAircraftOk() (*map[string]string, bool) {
	if o == nil || amadeus.IsNil(o.Aircraft) {
		return nil, false
	}
	return o.Aircraft, true
}

// HasAircraft returns a boolean if a field has been set.
func (o *Dictionaries) HasAircraft() bool {
	if o != nil && !amadeus.IsNil(o.Aircraft) {
		return true
	}

	return false
}

// SetAircraft gets a reference to the given map[string]string and assigns it to the Aircraft field.
func (o *Dictionaries) SetAircraft(v map[string]string) {
	o.Aircraft = &v
}

// GetCurrencies returns the Currencies field value if set, zero value otherwise.
func (o *Dictionaries) GetCurrencies() map[string]string {
	if o == nil || amadeus.IsNil(o.Currencies) {
		var ret map[string]string
		return ret
	}
	return *o.Currencies
}

// GetCurrenciesOk returns a tuple with the Currencies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dictionaries) GetCurrenciesOk() (*map[string]string, bool) {
	if o == nil || amadeus.IsNil(o.Currencies) {
		return nil, false
	}
	return o.Currencies, true
}

// HasCurrencies returns a boolean if a field has been set.
func (o *Dictionaries) HasCurrencies() bool {
	if o != nil && !amadeus.IsNil(o.Currencies) {
		return true
	}

	return false
}

// SetCurrencies gets a reference to the given map[string]string and assigns it to the Currencies field.
func (o *Dictionaries) SetCurrencies(v map[string]string) {
	o.Currencies = &v
}

// GetCarriers returns the Carriers field value if set, zero value otherwise.
func (o *Dictionaries) GetCarriers() map[string]string {
	if o == nil || amadeus.IsNil(o.Carriers) {
		var ret map[string]string
		return ret
	}
	return *o.Carriers
}

// GetCarriersOk returns a tuple with the Carriers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dictionaries) GetCarriersOk() (*map[string]string, bool) {
	if o == nil || amadeus.IsNil(o.Carriers) {
		return nil, false
	}
	return o.Carriers, true
}

// HasCarriers returns a boolean if a field has been set.
func (o *Dictionaries) HasCarriers() bool {
	if o != nil && !amadeus.IsNil(o.Carriers) {
		return true
	}

	return false
}

// SetCarriers gets a reference to the given map[string]string and assigns it to the Carriers field.
func (o *Dictionaries) SetCarriers(v map[string]string) {
	o.Carriers = &v
}

func (o Dictionaries) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Dictionaries) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !amadeus.IsNil(o.Locations) {
		toSerialize["locations"] = o.Locations
	}
	if !amadeus.IsNil(o.Aircraft) {
		toSerialize["aircraft"] = o.Aircraft
	}
	if !amadeus.IsNil(o.Currencies) {
		toSerialize["currencies"] = o.Currencies
	}
	if !amadeus.IsNil(o.Carriers) {
		toSerialize["carriers"] = o.Carriers
	}
	return toSerialize, nil
}

type NullableDictionaries struct {
	value *Dictionaries
	isSet bool
}

func (v NullableDictionaries) Get() *Dictionaries {
	return v.value
}

func (v *NullableDictionaries) Set(val *Dictionaries) {
	v.value = val
	v.isSet = true
}

func (v NullableDictionaries) IsSet() bool {
	return v.isSet
}

func (v *NullableDictionaries) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDictionaries(val *Dictionaries) *NullableDictionaries {
	return &NullableDictionaries{value: val, isSet: true}
}

func (v NullableDictionaries) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDictionaries) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

