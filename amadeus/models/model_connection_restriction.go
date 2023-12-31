package models

import (
	"encoding/json"
	"github.com/yenvanio/amadeus-go/amadeus/utils"
)

// checks if the ConnectionRestriction type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ConnectionRestriction{}

// ConnectionRestriction Restriction towards number of connections.
type ConnectionRestriction struct {
	// The maximal number of connections for each itinerary. Value can be 0, 1 or 2.
	MaxNumberOfConnections *float32 `json:"maxNumberOfConnections,omitempty"`
	// When this option is requested, recommendations made of Non-Stop flights only are favoured by the search, on the whole itinerary, with a weight of 75%. It is possible to override this default 75% weight of Non-Stop recommendations by providing the requested weight in the attribute nonStopPreferredWeight
	NonStopPreferred *bool `json:"nonStopPreferred,omitempty"`
	// This weight is the percentage of total number of recommendations that could be returned for Non-Stop flight recommendations. It is only relevant when combined with nonSTopPreferred set to true. If not specified, the default weight of 75% is applied
	NonStopPreferredWeight *float32 `json:"nonStopPreferredWeight,omitempty"`
	// Allow to change airport during connection
	AirportChangeAllowed *bool `json:"airportChangeAllowed,omitempty"`
	// This option allows the single segment to have one or more intermediate stops (technical stops).
	TechnicalStopsAllowed *bool `json:"technicalStopsAllowed,omitempty"`
}

// NewConnectionRestriction instantiates a new ConnectionRestriction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectionRestriction() *ConnectionRestriction {
	this := ConnectionRestriction{}
	return &this
}

// NewConnectionRestrictionWithDefaults instantiates a new ConnectionRestriction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectionRestrictionWithDefaults() *ConnectionRestriction {
	this := ConnectionRestriction{}
	return &this
}

// GetMaxNumberOfConnections returns the MaxNumberOfConnections field value if set, zero value otherwise.
func (o *ConnectionRestriction) GetMaxNumberOfConnections() float32 {
	if o == nil || utils.IsNil(o.MaxNumberOfConnections) {
		var ret float32
		return ret
	}
	return *o.MaxNumberOfConnections
}

// GetMaxNumberOfConnectionsOk returns a tuple with the MaxNumberOfConnections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionRestriction) GetMaxNumberOfConnectionsOk() (*float32, bool) {
	if o == nil || utils.IsNil(o.MaxNumberOfConnections) {
		return nil, false
	}
	return o.MaxNumberOfConnections, true
}

// HasMaxNumberOfConnections returns a boolean if a field has been set.
func (o *ConnectionRestriction) HasMaxNumberOfConnections() bool {
	if o != nil && !utils.IsNil(o.MaxNumberOfConnections) {
		return true
	}

	return false
}

// SetMaxNumberOfConnections gets a reference to the given float32 and assigns it to the MaxNumberOfConnections field.
func (o *ConnectionRestriction) SetMaxNumberOfConnections(v float32) {
	o.MaxNumberOfConnections = &v
}

// GetNonStopPreferred returns the NonStopPreferred field value if set, zero value otherwise.
func (o *ConnectionRestriction) GetNonStopPreferred() bool {
	if o == nil || utils.IsNil(o.NonStopPreferred) {
		var ret bool
		return ret
	}
	return *o.NonStopPreferred
}

// GetNonStopPreferredOk returns a tuple with the NonStopPreferred field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionRestriction) GetNonStopPreferredOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.NonStopPreferred) {
		return nil, false
	}
	return o.NonStopPreferred, true
}

// HasNonStopPreferred returns a boolean if a field has been set.
func (o *ConnectionRestriction) HasNonStopPreferred() bool {
	if o != nil && !utils.IsNil(o.NonStopPreferred) {
		return true
	}

	return false
}

// SetNonStopPreferred gets a reference to the given bool and assigns it to the NonStopPreferred field.
func (o *ConnectionRestriction) SetNonStopPreferred(v bool) {
	o.NonStopPreferred = &v
}

// GetNonStopPreferredWeight returns the NonStopPreferredWeight field value if set, zero value otherwise.
func (o *ConnectionRestriction) GetNonStopPreferredWeight() float32 {
	if o == nil || utils.IsNil(o.NonStopPreferredWeight) {
		var ret float32
		return ret
	}
	return *o.NonStopPreferredWeight
}

// GetNonStopPreferredWeightOk returns a tuple with the NonStopPreferredWeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionRestriction) GetNonStopPreferredWeightOk() (*float32, bool) {
	if o == nil || utils.IsNil(o.NonStopPreferredWeight) {
		return nil, false
	}
	return o.NonStopPreferredWeight, true
}

// HasNonStopPreferredWeight returns a boolean if a field has been set.
func (o *ConnectionRestriction) HasNonStopPreferredWeight() bool {
	if o != nil && !utils.IsNil(o.NonStopPreferredWeight) {
		return true
	}

	return false
}

// SetNonStopPreferredWeight gets a reference to the given float32 and assigns it to the NonStopPreferredWeight field.
func (o *ConnectionRestriction) SetNonStopPreferredWeight(v float32) {
	o.NonStopPreferredWeight = &v
}

// GetAirportChangeAllowed returns the AirportChangeAllowed field value if set, zero value otherwise.
func (o *ConnectionRestriction) GetAirportChangeAllowed() bool {
	if o == nil || utils.IsNil(o.AirportChangeAllowed) {
		var ret bool
		return ret
	}
	return *o.AirportChangeAllowed
}

// GetAirportChangeAllowedOk returns a tuple with the AirportChangeAllowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionRestriction) GetAirportChangeAllowedOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.AirportChangeAllowed) {
		return nil, false
	}
	return o.AirportChangeAllowed, true
}

// HasAirportChangeAllowed returns a boolean if a field has been set.
func (o *ConnectionRestriction) HasAirportChangeAllowed() bool {
	if o != nil && !utils.IsNil(o.AirportChangeAllowed) {
		return true
	}

	return false
}

// SetAirportChangeAllowed gets a reference to the given bool and assigns it to the AirportChangeAllowed field.
func (o *ConnectionRestriction) SetAirportChangeAllowed(v bool) {
	o.AirportChangeAllowed = &v
}

// GetTechnicalStopsAllowed returns the TechnicalStopsAllowed field value if set, zero value otherwise.
func (o *ConnectionRestriction) GetTechnicalStopsAllowed() bool {
	if o == nil || utils.IsNil(o.TechnicalStopsAllowed) {
		var ret bool
		return ret
	}
	return *o.TechnicalStopsAllowed
}

// GetTechnicalStopsAllowedOk returns a tuple with the TechnicalStopsAllowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionRestriction) GetTechnicalStopsAllowedOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.TechnicalStopsAllowed) {
		return nil, false
	}
	return o.TechnicalStopsAllowed, true
}

// HasTechnicalStopsAllowed returns a boolean if a field has been set.
func (o *ConnectionRestriction) HasTechnicalStopsAllowed() bool {
	if o != nil && !utils.IsNil(o.TechnicalStopsAllowed) {
		return true
	}

	return false
}

// SetTechnicalStopsAllowed gets a reference to the given bool and assigns it to the TechnicalStopsAllowed field.
func (o *ConnectionRestriction) SetTechnicalStopsAllowed(v bool) {
	o.TechnicalStopsAllowed = &v
}

func (o ConnectionRestriction) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConnectionRestriction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.MaxNumberOfConnections) {
		toSerialize["maxNumberOfConnections"] = o.MaxNumberOfConnections
	}
	if !utils.IsNil(o.NonStopPreferred) {
		toSerialize["nonStopPreferred"] = o.NonStopPreferred
	}
	if !utils.IsNil(o.NonStopPreferredWeight) {
		toSerialize["nonStopPreferredWeight"] = o.NonStopPreferredWeight
	}
	if !utils.IsNil(o.AirportChangeAllowed) {
		toSerialize["airportChangeAllowed"] = o.AirportChangeAllowed
	}
	if !utils.IsNil(o.TechnicalStopsAllowed) {
		toSerialize["technicalStopsAllowed"] = o.TechnicalStopsAllowed
	}
	return toSerialize, nil
}

type NullableConnectionRestriction struct {
	value *ConnectionRestriction
	isSet bool
}

func (v NullableConnectionRestriction) Get() *ConnectionRestriction {
	return v.value
}

func (v *NullableConnectionRestriction) Set(val *ConnectionRestriction) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectionRestriction) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectionRestriction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectionRestriction(val *ConnectionRestriction) *NullableConnectionRestriction {
	return &NullableConnectionRestriction{value: val, isSet: true}
}

func (v NullableConnectionRestriction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectionRestriction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
