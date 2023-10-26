package models

import (
	"encoding/json"
	"github.com/yenvanio/amadeus-go/amadeus/utils"
)

// checks if the Address type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Address{}

// Address address information
type Address struct {
	// Line 1 = Street address, Line 2 = Apartment, suite, unit, building, floor, etc
	Lines []string `json:"lines,omitempty"`
	// Example: 74130
	PostalCode *string `json:"postalCode,omitempty"`
	// country code [ISO 3166-1 country code](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)
	CountryCode *string `json:"countryCode,omitempty"`
	// Full city name. Example: Dublin
	CityName *string `json:"cityName,omitempty"`
	// Full state name
	StateName *string `json:"stateName,omitempty"`
	// E.g. BP 220
	PostalBox *string `json:"postalBox,omitempty"`
}

// NewAddress instantiates a new Address object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddress() *Address {
	this := Address{}
	return &this
}

// NewAddressWithDefaults instantiates a new Address object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressWithDefaults() *Address {
	this := Address{}
	return &this
}

// GetLines returns the Lines field value if set, zero value otherwise.
func (o *Address) GetLines() []string {
	if o == nil || utils.IsNil(o.Lines) {
		var ret []string
		return ret
	}
	return o.Lines
}

// GetLinesOk returns a tuple with the Lines field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Address) GetLinesOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.Lines) {
		return nil, false
	}
	return o.Lines, true
}

// HasLines returns a boolean if a field has been set.
func (o *Address) HasLines() bool {
	if o != nil && !utils.IsNil(o.Lines) {
		return true
	}

	return false
}

// SetLines gets a reference to the given []string and assigns it to the Lines field.
func (o *Address) SetLines(v []string) {
	o.Lines = v
}

// GetPostalCode returns the PostalCode field value if set, zero value otherwise.
func (o *Address) GetPostalCode() string {
	if o == nil || utils.IsNil(o.PostalCode) {
		var ret string
		return ret
	}
	return *o.PostalCode
}

// GetPostalCodeOk returns a tuple with the PostalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Address) GetPostalCodeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.PostalCode) {
		return nil, false
	}
	return o.PostalCode, true
}

// HasPostalCode returns a boolean if a field has been set.
func (o *Address) HasPostalCode() bool {
	if o != nil && !utils.IsNil(o.PostalCode) {
		return true
	}

	return false
}

// SetPostalCode gets a reference to the given string and assigns it to the PostalCode field.
func (o *Address) SetPostalCode(v string) {
	o.PostalCode = &v
}

// GetCountryCode returns the CountryCode field value if set, zero value otherwise.
func (o *Address) GetCountryCode() string {
	if o == nil || utils.IsNil(o.CountryCode) {
		var ret string
		return ret
	}
	return *o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Address) GetCountryCodeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.CountryCode) {
		return nil, false
	}
	return o.CountryCode, true
}

// HasCountryCode returns a boolean if a field has been set.
func (o *Address) HasCountryCode() bool {
	if o != nil && !utils.IsNil(o.CountryCode) {
		return true
	}

	return false
}

// SetCountryCode gets a reference to the given string and assigns it to the CountryCode field.
func (o *Address) SetCountryCode(v string) {
	o.CountryCode = &v
}

// GetCityName returns the CityName field value if set, zero value otherwise.
func (o *Address) GetCityName() string {
	if o == nil || utils.IsNil(o.CityName) {
		var ret string
		return ret
	}
	return *o.CityName
}

// GetCityNameOk returns a tuple with the CityName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Address) GetCityNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.CityName) {
		return nil, false
	}
	return o.CityName, true
}

// HasCityName returns a boolean if a field has been set.
func (o *Address) HasCityName() bool {
	if o != nil && !utils.IsNil(o.CityName) {
		return true
	}

	return false
}

// SetCityName gets a reference to the given string and assigns it to the CityName field.
func (o *Address) SetCityName(v string) {
	o.CityName = &v
}

// GetStateName returns the StateName field value if set, zero value otherwise.
func (o *Address) GetStateName() string {
	if o == nil || utils.IsNil(o.StateName) {
		var ret string
		return ret
	}
	return *o.StateName
}

// GetStateNameOk returns a tuple with the StateName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Address) GetStateNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.StateName) {
		return nil, false
	}
	return o.StateName, true
}

// HasStateName returns a boolean if a field has been set.
func (o *Address) HasStateName() bool {
	if o != nil && !utils.IsNil(o.StateName) {
		return true
	}

	return false
}

// SetStateName gets a reference to the given string and assigns it to the StateName field.
func (o *Address) SetStateName(v string) {
	o.StateName = &v
}

// GetPostalBox returns the PostalBox field value if set, zero value otherwise.
func (o *Address) GetPostalBox() string {
	if o == nil || utils.IsNil(o.PostalBox) {
		var ret string
		return ret
	}
	return *o.PostalBox
}

// GetPostalBoxOk returns a tuple with the PostalBox field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Address) GetPostalBoxOk() (*string, bool) {
	if o == nil || utils.IsNil(o.PostalBox) {
		return nil, false
	}
	return o.PostalBox, true
}

// HasPostalBox returns a boolean if a field has been set.
func (o *Address) HasPostalBox() bool {
	if o != nil && !utils.IsNil(o.PostalBox) {
		return true
	}

	return false
}

// SetPostalBox gets a reference to the given string and assigns it to the PostalBox field.
func (o *Address) SetPostalBox(v string) {
	o.PostalBox = &v
}

func (o Address) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Address) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Lines) {
		toSerialize["lines"] = o.Lines
	}
	if !utils.IsNil(o.PostalCode) {
		toSerialize["postalCode"] = o.PostalCode
	}
	if !utils.IsNil(o.CountryCode) {
		toSerialize["countryCode"] = o.CountryCode
	}
	if !utils.IsNil(o.CityName) {
		toSerialize["cityName"] = o.CityName
	}
	if !utils.IsNil(o.StateName) {
		toSerialize["stateName"] = o.StateName
	}
	if !utils.IsNil(o.PostalBox) {
		toSerialize["postalBox"] = o.PostalBox
	}
	return toSerialize, nil
}

type NullableAddress struct {
	value *Address
	isSet bool
}

func (v NullableAddress) Get() *Address {
	return v.value
}

func (v *NullableAddress) Set(val *Address) {
	v.value = val
	v.isSet = true
}

func (v NullableAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddress(val *Address) *NullableAddress {
	return &NullableAddress{value: val, isSet: true}
}

func (v NullableAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
