package models

import (
	"encoding/json"
	"github.com/yenvanio/amadeus-go/amadeus/utils"
)

// checks if the OriginDestination type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &OriginDestination{}

// OriginDestination struct for OriginDestination
type OriginDestination struct {
	Id *string `json:"id,omitempty"`
	// Origin location, such as a city or an airport. Currently, only the locations defined in [IATA](http://www.iata.org/publications/Pages/code-search.aspx) are supported.
	OriginLocationCode *string `json:"originLocationCode,omitempty"`
	// Destination location, such as a city or an airport. Currently, only the locations defined in [IATA](http://www.iata.org/publications/Pages/code-search.aspx) are supported.
	DestinationLocationCode *string `json:"destinationLocationCode,omitempty"`
	// List of included connections points. When an includedViaPoints option is specified, all FlightOffer returned must at least go via this Connecting Point. Currently, only the locations defined in IATA are supported. Used only by the AMADEUS provider
	IncludedConnectionPoints []string `json:"includedConnectionPoints,omitempty"`
	// List of excluded connections points. Any FlightOffer with these connections points will be present in response. Currently, only the locations defined in IATA are supported. Used only by the AMADEUS provider
	ExcludedConnectionPoints []string `json:"excludedConnectionPoints,omitempty"`
	// Include other possible locations around the point, located less than this distance in kilometers away. Max:300  Can not be combined with \"dateWindow\" or \"timeWindow\".
	OriginRadius *float32 `json:"originRadius,omitempty"`
	// Set of alternative origin location, such as a city or an airport. Currently, only the locations defined in [IATA](http://www.iata.org/publications/Pages/code-search.aspx) are supported.
	AlternativeOriginsCodes []string `json:"alternativeOriginsCodes,omitempty"`
	// Include other possible locations around the point, located less than this distance in kilometers away. Max:300  Can not be combined with \"dateWindow\" or \"timeWindow\".
	DestinationRadius *float32 `json:"destinationRadius,omitempty"`
	// Set of alternative destination location, such as a city or an airport. Currently, only the locations defined in [IATA](http://www.iata.org/publications/Pages/code-search.aspx) are supported.
	AlternativeDestinationsCodes []string       `json:"alternativeDestinationsCodes,omitempty"`
	DepartureDateTimeRange       *DateTimeRange `json:"departureDateTimeRange,omitempty"`
	ArrivalDateTimeRange         *DateTimeRange `json:"arrivalDateTimeRange,omitempty"`
}

// NewOriginDestination instantiates a new OriginDestination object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOriginDestination() *OriginDestination {
	this := OriginDestination{}
	return &this
}

// NewOriginDestinationWithDefaults instantiates a new OriginDestination object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOriginDestinationWithDefaults() *OriginDestination {
	this := OriginDestination{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OriginDestination) GetId() string {
	if o == nil || utils.IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginDestination) GetIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OriginDestination) HasId() bool {
	if o != nil && !utils.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *OriginDestination) SetId(v string) {
	o.Id = &v
}

// GetOriginLocationCode returns the OriginLocationCode field value if set, zero value otherwise.
func (o *OriginDestination) GetOriginLocationCode() string {
	if o == nil || utils.IsNil(o.OriginLocationCode) {
		var ret string
		return ret
	}
	return *o.OriginLocationCode
}

// GetOriginLocationCodeOk returns a tuple with the OriginLocationCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginDestination) GetOriginLocationCodeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.OriginLocationCode) {
		return nil, false
	}
	return o.OriginLocationCode, true
}

// HasOriginLocationCode returns a boolean if a field has been set.
func (o *OriginDestination) HasOriginLocationCode() bool {
	if o != nil && !utils.IsNil(o.OriginLocationCode) {
		return true
	}

	return false
}

// SetOriginLocationCode gets a reference to the given string and assigns it to the OriginLocationCode field.
func (o *OriginDestination) SetOriginLocationCode(v string) {
	o.OriginLocationCode = &v
}

// GetDestinationLocationCode returns the DestinationLocationCode field value if set, zero value otherwise.
func (o *OriginDestination) GetDestinationLocationCode() string {
	if o == nil || utils.IsNil(o.DestinationLocationCode) {
		var ret string
		return ret
	}
	return *o.DestinationLocationCode
}

// GetDestinationLocationCodeOk returns a tuple with the DestinationLocationCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginDestination) GetDestinationLocationCodeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.DestinationLocationCode) {
		return nil, false
	}
	return o.DestinationLocationCode, true
}

// HasDestinationLocationCode returns a boolean if a field has been set.
func (o *OriginDestination) HasDestinationLocationCode() bool {
	if o != nil && !utils.IsNil(o.DestinationLocationCode) {
		return true
	}

	return false
}

// SetDestinationLocationCode gets a reference to the given string and assigns it to the DestinationLocationCode field.
func (o *OriginDestination) SetDestinationLocationCode(v string) {
	o.DestinationLocationCode = &v
}

// GetIncludedConnectionPoints returns the IncludedConnectionPoints field value if set, zero value otherwise.
func (o *OriginDestination) GetIncludedConnectionPoints() []string {
	if o == nil || utils.IsNil(o.IncludedConnectionPoints) {
		var ret []string
		return ret
	}
	return o.IncludedConnectionPoints
}

// GetIncludedConnectionPointsOk returns a tuple with the IncludedConnectionPoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginDestination) GetIncludedConnectionPointsOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.IncludedConnectionPoints) {
		return nil, false
	}
	return o.IncludedConnectionPoints, true
}

// HasIncludedConnectionPoints returns a boolean if a field has been set.
func (o *OriginDestination) HasIncludedConnectionPoints() bool {
	if o != nil && !utils.IsNil(o.IncludedConnectionPoints) {
		return true
	}

	return false
}

// SetIncludedConnectionPoints gets a reference to the given []string and assigns it to the IncludedConnectionPoints field.
func (o *OriginDestination) SetIncludedConnectionPoints(v []string) {
	o.IncludedConnectionPoints = v
}

// GetExcludedConnectionPoints returns the ExcludedConnectionPoints field value if set, zero value otherwise.
func (o *OriginDestination) GetExcludedConnectionPoints() []string {
	if o == nil || utils.IsNil(o.ExcludedConnectionPoints) {
		var ret []string
		return ret
	}
	return o.ExcludedConnectionPoints
}

// GetExcludedConnectionPointsOk returns a tuple with the ExcludedConnectionPoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginDestination) GetExcludedConnectionPointsOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.ExcludedConnectionPoints) {
		return nil, false
	}
	return o.ExcludedConnectionPoints, true
}

// HasExcludedConnectionPoints returns a boolean if a field has been set.
func (o *OriginDestination) HasExcludedConnectionPoints() bool {
	if o != nil && !utils.IsNil(o.ExcludedConnectionPoints) {
		return true
	}

	return false
}

// SetExcludedConnectionPoints gets a reference to the given []string and assigns it to the ExcludedConnectionPoints field.
func (o *OriginDestination) SetExcludedConnectionPoints(v []string) {
	o.ExcludedConnectionPoints = v
}

// GetOriginRadius returns the OriginRadius field value if set, zero value otherwise.
func (o *OriginDestination) GetOriginRadius() float32 {
	if o == nil || utils.IsNil(o.OriginRadius) {
		var ret float32
		return ret
	}
	return *o.OriginRadius
}

// GetOriginRadiusOk returns a tuple with the OriginRadius field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginDestination) GetOriginRadiusOk() (*float32, bool) {
	if o == nil || utils.IsNil(o.OriginRadius) {
		return nil, false
	}
	return o.OriginRadius, true
}

// HasOriginRadius returns a boolean if a field has been set.
func (o *OriginDestination) HasOriginRadius() bool {
	if o != nil && !utils.IsNil(o.OriginRadius) {
		return true
	}

	return false
}

// SetOriginRadius gets a reference to the given float32 and assigns it to the OriginRadius field.
func (o *OriginDestination) SetOriginRadius(v float32) {
	o.OriginRadius = &v
}

// GetAlternativeOriginsCodes returns the AlternativeOriginsCodes field value if set, zero value otherwise.
func (o *OriginDestination) GetAlternativeOriginsCodes() []string {
	if o == nil || utils.IsNil(o.AlternativeOriginsCodes) {
		var ret []string
		return ret
	}
	return o.AlternativeOriginsCodes
}

// GetAlternativeOriginsCodesOk returns a tuple with the AlternativeOriginsCodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginDestination) GetAlternativeOriginsCodesOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.AlternativeOriginsCodes) {
		return nil, false
	}
	return o.AlternativeOriginsCodes, true
}

// HasAlternativeOriginsCodes returns a boolean if a field has been set.
func (o *OriginDestination) HasAlternativeOriginsCodes() bool {
	if o != nil && !utils.IsNil(o.AlternativeOriginsCodes) {
		return true
	}

	return false
}

// SetAlternativeOriginsCodes gets a reference to the given []string and assigns it to the AlternativeOriginsCodes field.
func (o *OriginDestination) SetAlternativeOriginsCodes(v []string) {
	o.AlternativeOriginsCodes = v
}

// GetDestinationRadius returns the DestinationRadius field value if set, zero value otherwise.
func (o *OriginDestination) GetDestinationRadius() float32 {
	if o == nil || utils.IsNil(o.DestinationRadius) {
		var ret float32
		return ret
	}
	return *o.DestinationRadius
}

// GetDestinationRadiusOk returns a tuple with the DestinationRadius field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginDestination) GetDestinationRadiusOk() (*float32, bool) {
	if o == nil || utils.IsNil(o.DestinationRadius) {
		return nil, false
	}
	return o.DestinationRadius, true
}

// HasDestinationRadius returns a boolean if a field has been set.
func (o *OriginDestination) HasDestinationRadius() bool {
	if o != nil && !utils.IsNil(o.DestinationRadius) {
		return true
	}

	return false
}

// SetDestinationRadius gets a reference to the given float32 and assigns it to the DestinationRadius field.
func (o *OriginDestination) SetDestinationRadius(v float32) {
	o.DestinationRadius = &v
}

// GetAlternativeDestinationsCodes returns the AlternativeDestinationsCodes field value if set, zero value otherwise.
func (o *OriginDestination) GetAlternativeDestinationsCodes() []string {
	if o == nil || utils.IsNil(o.AlternativeDestinationsCodes) {
		var ret []string
		return ret
	}
	return o.AlternativeDestinationsCodes
}

// GetAlternativeDestinationsCodesOk returns a tuple with the AlternativeDestinationsCodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginDestination) GetAlternativeDestinationsCodesOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.AlternativeDestinationsCodes) {
		return nil, false
	}
	return o.AlternativeDestinationsCodes, true
}

// HasAlternativeDestinationsCodes returns a boolean if a field has been set.
func (o *OriginDestination) HasAlternativeDestinationsCodes() bool {
	if o != nil && !utils.IsNil(o.AlternativeDestinationsCodes) {
		return true
	}

	return false
}

// SetAlternativeDestinationsCodes gets a reference to the given []string and assigns it to the AlternativeDestinationsCodes field.
func (o *OriginDestination) SetAlternativeDestinationsCodes(v []string) {
	o.AlternativeDestinationsCodes = v
}

// GetDepartureDateTimeRange returns the DepartureDateTimeRange field value if set, zero value otherwise.
func (o *OriginDestination) GetDepartureDateTimeRange() DateTimeRange {
	if o == nil || utils.IsNil(o.DepartureDateTimeRange) {
		var ret DateTimeRange
		return ret
	}
	return *o.DepartureDateTimeRange
}

// GetDepartureDateTimeRangeOk returns a tuple with the DepartureDateTimeRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginDestination) GetDepartureDateTimeRangeOk() (*DateTimeRange, bool) {
	if o == nil || utils.IsNil(o.DepartureDateTimeRange) {
		return nil, false
	}
	return o.DepartureDateTimeRange, true
}

// HasDepartureDateTimeRange returns a boolean if a field has been set.
func (o *OriginDestination) HasDepartureDateTimeRange() bool {
	if o != nil && !utils.IsNil(o.DepartureDateTimeRange) {
		return true
	}

	return false
}

// SetDepartureDateTimeRange gets a reference to the given DateTimeRange and assigns it to the DepartureDateTimeRange field.
func (o *OriginDestination) SetDepartureDateTimeRange(v DateTimeRange) {
	o.DepartureDateTimeRange = &v
}

// GetArrivalDateTimeRange returns the ArrivalDateTimeRange field value if set, zero value otherwise.
func (o *OriginDestination) GetArrivalDateTimeRange() DateTimeRange {
	if o == nil || utils.IsNil(o.ArrivalDateTimeRange) {
		var ret DateTimeRange
		return ret
	}
	return *o.ArrivalDateTimeRange
}

// GetArrivalDateTimeRangeOk returns a tuple with the ArrivalDateTimeRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginDestination) GetArrivalDateTimeRangeOk() (*DateTimeRange, bool) {
	if o == nil || utils.IsNil(o.ArrivalDateTimeRange) {
		return nil, false
	}
	return o.ArrivalDateTimeRange, true
}

// HasArrivalDateTimeRange returns a boolean if a field has been set.
func (o *OriginDestination) HasArrivalDateTimeRange() bool {
	if o != nil && !utils.IsNil(o.ArrivalDateTimeRange) {
		return true
	}

	return false
}

// SetArrivalDateTimeRange gets a reference to the given DateTimeRange and assigns it to the ArrivalDateTimeRange field.
func (o *OriginDestination) SetArrivalDateTimeRange(v DateTimeRange) {
	o.ArrivalDateTimeRange = &v
}

func (o OriginDestination) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OriginDestination) ToMap() (map[string]interface{}, error) {
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
	if !utils.IsNil(o.OriginRadius) {
		toSerialize["originRadius"] = o.OriginRadius
	}
	if !utils.IsNil(o.AlternativeOriginsCodes) {
		toSerialize["alternativeOriginsCodes"] = o.AlternativeOriginsCodes
	}
	if !utils.IsNil(o.DestinationRadius) {
		toSerialize["destinationRadius"] = o.DestinationRadius
	}
	if !utils.IsNil(o.AlternativeDestinationsCodes) {
		toSerialize["alternativeDestinationsCodes"] = o.AlternativeDestinationsCodes
	}
	if !utils.IsNil(o.DepartureDateTimeRange) {
		toSerialize["departureDateTimeRange"] = o.DepartureDateTimeRange
	}
	if !utils.IsNil(o.ArrivalDateTimeRange) {
		toSerialize["arrivalDateTimeRange"] = o.ArrivalDateTimeRange
	}
	return toSerialize, nil
}

type NullableOriginDestination struct {
	value *OriginDestination
	isSet bool
}

func (v NullableOriginDestination) Get() *OriginDestination {
	return v.value
}

func (v *NullableOriginDestination) Set(val *OriginDestination) {
	v.value = val
	v.isSet = true
}

func (v NullableOriginDestination) IsSet() bool {
	return v.isSet
}

func (v *NullableOriginDestination) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOriginDestination(val *OriginDestination) *NullableOriginDestination {
	return &NullableOriginDestination{value: val, isSet: true}
}

func (v NullableOriginDestination) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOriginDestination) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
