package models

import (
	"encoding/json"
	"github.com/yenvanio/amadeus-go/amadeus/utils"
)

// checks if the OriginDestinationLight type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &OriginDestinationLight{}

// OriginDestinationLight struct for OriginDestinationLight
type OriginDestinationLight struct {
	Id *string `json:"id,omitempty"`
	// Origin location, such as a city or an airport. Currently, only the locations defined in [IATA](http://www.iata.org/publications/Pages/code-search.aspx) are supported.
	OriginLocationCode *string `json:"originLocationCode,omitempty"`
	// Destination location, such as a city or an airport. Currently, only the locations defined in [IATA](http://www.iata.org/publications/Pages/code-search.aspx) are supported.
	DestinationLocationCode *string `json:"destinationLocationCode,omitempty"`
	// List of included connections points. When an includedViaPoints option is specified, all FlightOffer returned must at least go via this Connecting Point. Currently, only the locations defined in IATA are supported. Used only by the AMADEUS provider
	IncludedConnectionPoints []string `json:"includedConnectionPoints,omitempty"`
	// List of excluded connections points. Any FlightOffer with these connections points will be present in response. Currently, only the locations defined in IATA are supported. Used only by the AMADEUS provider
	ExcludedConnectionPoints []string `json:"excludedConnectionPoints,omitempty"`
}

// NewOriginDestinationLight instantiates a new OriginDestinationLight object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOriginDestinationLight() *OriginDestinationLight {
	this := OriginDestinationLight{}
	return &this
}

// NewOriginDestinationLightWithDefaults instantiates a new OriginDestinationLight object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOriginDestinationLightWithDefaults() *OriginDestinationLight {
	this := OriginDestinationLight{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OriginDestinationLight) GetId() string {
	if o == nil || utils.IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginDestinationLight) GetIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OriginDestinationLight) HasId() bool {
	if o != nil && !utils.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *OriginDestinationLight) SetId(v string) {
	o.Id = &v
}

// GetOriginLocationCode returns the OriginLocationCode field value if set, zero value otherwise.
func (o *OriginDestinationLight) GetOriginLocationCode() string {
	if o == nil || utils.IsNil(o.OriginLocationCode) {
		var ret string
		return ret
	}
	return *o.OriginLocationCode
}

// GetOriginLocationCodeOk returns a tuple with the OriginLocationCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginDestinationLight) GetOriginLocationCodeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.OriginLocationCode) {
		return nil, false
	}
	return o.OriginLocationCode, true
}

// HasOriginLocationCode returns a boolean if a field has been set.
func (o *OriginDestinationLight) HasOriginLocationCode() bool {
	if o != nil && !utils.IsNil(o.OriginLocationCode) {
		return true
	}

	return false
}

// SetOriginLocationCode gets a reference to the given string and assigns it to the OriginLocationCode field.
func (o *OriginDestinationLight) SetOriginLocationCode(v string) {
	o.OriginLocationCode = &v
}

// GetDestinationLocationCode returns the DestinationLocationCode field value if set, zero value otherwise.
func (o *OriginDestinationLight) GetDestinationLocationCode() string {
	if o == nil || utils.IsNil(o.DestinationLocationCode) {
		var ret string
		return ret
	}
	return *o.DestinationLocationCode
}

// GetDestinationLocationCodeOk returns a tuple with the DestinationLocationCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginDestinationLight) GetDestinationLocationCodeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.DestinationLocationCode) {
		return nil, false
	}
	return o.DestinationLocationCode, true
}

// HasDestinationLocationCode returns a boolean if a field has been set.
func (o *OriginDestinationLight) HasDestinationLocationCode() bool {
	if o != nil && !utils.IsNil(o.DestinationLocationCode) {
		return true
	}

	return false
}

// SetDestinationLocationCode gets a reference to the given string and assigns it to the DestinationLocationCode field.
func (o *OriginDestinationLight) SetDestinationLocationCode(v string) {
	o.DestinationLocationCode = &v
}

// GetIncludedConnectionPoints returns the IncludedConnectionPoints field value if set, zero value otherwise.
func (o *OriginDestinationLight) GetIncludedConnectionPoints() []string {
	if o == nil || utils.IsNil(o.IncludedConnectionPoints) {
		var ret []string
		return ret
	}
	return o.IncludedConnectionPoints
}

// GetIncludedConnectionPointsOk returns a tuple with the IncludedConnectionPoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginDestinationLight) GetIncludedConnectionPointsOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.IncludedConnectionPoints) {
		return nil, false
	}
	return o.IncludedConnectionPoints, true
}

// HasIncludedConnectionPoints returns a boolean if a field has been set.
func (o *OriginDestinationLight) HasIncludedConnectionPoints() bool {
	if o != nil && !utils.IsNil(o.IncludedConnectionPoints) {
		return true
	}

	return false
}

// SetIncludedConnectionPoints gets a reference to the given []string and assigns it to the IncludedConnectionPoints field.
func (o *OriginDestinationLight) SetIncludedConnectionPoints(v []string) {
	o.IncludedConnectionPoints = v
}

// GetExcludedConnectionPoints returns the ExcludedConnectionPoints field value if set, zero value otherwise.
func (o *OriginDestinationLight) GetExcludedConnectionPoints() []string {
	if o == nil || utils.IsNil(o.ExcludedConnectionPoints) {
		var ret []string
		return ret
	}
	return o.ExcludedConnectionPoints
}

// GetExcludedConnectionPointsOk returns a tuple with the ExcludedConnectionPoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginDestinationLight) GetExcludedConnectionPointsOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.ExcludedConnectionPoints) {
		return nil, false
	}
	return o.ExcludedConnectionPoints, true
}

// HasExcludedConnectionPoints returns a boolean if a field has been set.
func (o *OriginDestinationLight) HasExcludedConnectionPoints() bool {
	if o != nil && !utils.IsNil(o.ExcludedConnectionPoints) {
		return true
	}

	return false
}

// SetExcludedConnectionPoints gets a reference to the given []string and assigns it to the ExcludedConnectionPoints field.
func (o *OriginDestinationLight) SetExcludedConnectionPoints(v []string) {
	o.ExcludedConnectionPoints = v
}

func (o OriginDestinationLight) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OriginDestinationLight) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !utils.IsNil(o.OriginLocationCode) {
		toSerialize["originLocationCode"] = o.OriginLocationCode
	}
	if !utils.IsNil(o.DestinationLocationCode) {
		toSerialize["destinationLocationCode"] = o.DestinationLocationCode
	}
	if !utils.IsNil(o.IncludedConnectionPoints) {
		toSerialize["includedConnectionPoints"] = o.IncludedConnectionPoints
	}
	if !utils.IsNil(o.ExcludedConnectionPoints) {
		toSerialize["excludedConnectionPoints"] = o.ExcludedConnectionPoints
	}
	return toSerialize, nil
}

type NullableOriginDestinationLight struct {
	value *OriginDestinationLight
	isSet bool
}

func (v NullableOriginDestinationLight) Get() *OriginDestinationLight {
	return v.value
}

func (v *NullableOriginDestinationLight) Set(val *OriginDestinationLight) {
	v.value = val
	v.isSet = true
}

func (v NullableOriginDestinationLight) IsSet() bool {
	return v.isSet
}

func (v *NullableOriginDestinationLight) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOriginDestinationLight(val *OriginDestinationLight) *NullableOriginDestinationLight {
	return &NullableOriginDestinationLight{value: val, isSet: true}
}

func (v NullableOriginDestinationLight) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOriginDestinationLight) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
