package models

import (
	"encoding/json"
	"github.com/yenvanio/amadeus-go/amadeus/utils"
)

// checks if the GetFlightOffersQuery type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &GetFlightOffersQuery{}

// GetFlightOffersQuery struct for GetFlightOffersQuery
type GetFlightOffersQuery struct {
	// The currency code, as defined in [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217), to reflect the currency in which this amount is expressed.
	CurrencyCode *string `json:"currencyCode,omitempty"`
	// Origins and Destinations must be properly ordered in time (chronological order in accordance with the timezone of each location) to describe the journey consistently. Dates and times must not be past nor more than 365 days in the future, according to provider settings.Number of Origins and Destinations must not exceed the limit defined in provider settings.
	OriginDestinations []OriginDestination `json:"originDestinations"`
	// travelers in the trip.    Maximum number of passengers older than 2 yo (CHILD, ADULT, YOUGHT): 9.  Each adult can travel with one INFANT so maximum total number of passengers: 18
	Travelers []ExtendedTravelerInfo `json:"travelers"`
	// Allows enable one or more sources. If present in the list, these sources will be called by the system.
	Sources        []FlightOfferSource `json:"sources"`
	SearchCriteria *SearchCriteria     `json:"searchCriteria,omitempty"`
}

// NewGetFlightOffersQuery instantiates a new GetFlightOffersQuery object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetFlightOffersQuery(originDestinations []OriginDestination, travelers []ExtendedTravelerInfo, sources []FlightOfferSource) *GetFlightOffersQuery {
	this := GetFlightOffersQuery{}
	this.OriginDestinations = originDestinations
	this.Travelers = travelers
	this.Sources = sources
	return &this
}

// NewGetFlightOffersQueryWithDefaults instantiates a new GetFlightOffersQuery object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetFlightOffersQueryWithDefaults() *GetFlightOffersQuery {
	this := GetFlightOffersQuery{}
	return &this
}

// GetCurrencyCode returns the CurrencyCode field value if set, zero value otherwise.
func (o *GetFlightOffersQuery) GetCurrencyCode() string {
	if o == nil || utils.IsNil(o.CurrencyCode) {
		var ret string
		return ret
	}
	return *o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFlightOffersQuery) GetCurrencyCodeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.CurrencyCode) {
		return nil, false
	}
	return o.CurrencyCode, true
}

// HasCurrencyCode returns a boolean if a field has been set.
func (o *GetFlightOffersQuery) HasCurrencyCode() bool {
	if o != nil && !utils.IsNil(o.CurrencyCode) {
		return true
	}

	return false
}

// SetCurrencyCode gets a reference to the given string and assigns it to the CurrencyCode field.
func (o *GetFlightOffersQuery) SetCurrencyCode(v string) {
	o.CurrencyCode = &v
}

// GetOriginDestinations returns the OriginDestinations field value
func (o *GetFlightOffersQuery) GetOriginDestinations() []OriginDestination {
	if o == nil {
		var ret []OriginDestination
		return ret
	}

	return o.OriginDestinations
}

// GetOriginDestinationsOk returns a tuple with the OriginDestinations field value
// and a boolean to check if the value has been set.
func (o *GetFlightOffersQuery) GetOriginDestinationsOk() ([]OriginDestination, bool) {
	if o == nil {
		return nil, false
	}
	return o.OriginDestinations, true
}

// SetOriginDestinations sets field value
func (o *GetFlightOffersQuery) SetOriginDestinations(v []OriginDestination) {
	o.OriginDestinations = v
}

// GetTravelers returns the Travelers field value
func (o *GetFlightOffersQuery) GetTravelers() []ExtendedTravelerInfo {
	if o == nil {
		var ret []ExtendedTravelerInfo
		return ret
	}

	return o.Travelers
}

// GetTravelersOk returns a tuple with the Travelers field value
// and a boolean to check if the value has been set.
func (o *GetFlightOffersQuery) GetTravelersOk() ([]ExtendedTravelerInfo, bool) {
	if o == nil {
		return nil, false
	}
	return o.Travelers, true
}

// SetTravelers sets field value
func (o *GetFlightOffersQuery) SetTravelers(v []ExtendedTravelerInfo) {
	o.Travelers = v
}

// GetSources returns the Sources field value
func (o *GetFlightOffersQuery) GetSources() []FlightOfferSource {
	if o == nil {
		var ret []FlightOfferSource
		return ret
	}

	return o.Sources
}

// GetSourcesOk returns a tuple with the Sources field value
// and a boolean to check if the value has been set.
func (o *GetFlightOffersQuery) GetSourcesOk() ([]FlightOfferSource, bool) {
	if o == nil {
		return nil, false
	}
	return o.Sources, true
}

// SetSources sets field value
func (o *GetFlightOffersQuery) SetSources(v []FlightOfferSource) {
	o.Sources = v
}

// GetSearchCriteria returns the SearchCriteria field value if set, zero value otherwise.
func (o *GetFlightOffersQuery) GetSearchCriteria() SearchCriteria {
	if o == nil || utils.IsNil(o.SearchCriteria) {
		var ret SearchCriteria
		return ret
	}
	return *o.SearchCriteria
}

// GetSearchCriteriaOk returns a tuple with the SearchCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFlightOffersQuery) GetSearchCriteriaOk() (*SearchCriteria, bool) {
	if o == nil || utils.IsNil(o.SearchCriteria) {
		return nil, false
	}
	return o.SearchCriteria, true
}

// HasSearchCriteria returns a boolean if a field has been set.
func (o *GetFlightOffersQuery) HasSearchCriteria() bool {
	if o != nil && !utils.IsNil(o.SearchCriteria) {
		return true
	}

	return false
}

// SetSearchCriteria gets a reference to the given SearchCriteria and assigns it to the SearchCriteria field.
func (o *GetFlightOffersQuery) SetSearchCriteria(v SearchCriteria) {
	o.SearchCriteria = &v
}

func (o GetFlightOffersQuery) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetFlightOffersQuery) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.CurrencyCode) {
		toSerialize["currencyCode"] = o.CurrencyCode
	}
	toSerialize["originDestinations"] = o.OriginDestinations
	toSerialize["travelers"] = o.Travelers
	toSerialize["sources"] = o.Sources
	if !utils.IsNil(o.SearchCriteria) {
		toSerialize["searchCriteria"] = o.SearchCriteria
	}
	return toSerialize, nil
}

type NullableGetFlightOffersQuery struct {
	value *GetFlightOffersQuery
	isSet bool
}

func (v NullableGetFlightOffersQuery) Get() *GetFlightOffersQuery {
	return v.value
}

func (v *NullableGetFlightOffersQuery) Set(val *GetFlightOffersQuery) {
	v.value = val
	v.isSet = true
}

func (v NullableGetFlightOffersQuery) IsSet() bool {
	return v.isSet
}

func (v *NullableGetFlightOffersQuery) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetFlightOffersQuery(val *GetFlightOffersQuery) *NullableGetFlightOffersQuery {
	return &NullableGetFlightOffersQuery{value: val, isSet: true}
}

func (v NullableGetFlightOffersQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetFlightOffersQuery) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
