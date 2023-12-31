package models

import (
	"encoding/json"
	"github.com/yenvanio/amadeus-go/amadeus/utils"
)

// checks if the SearchCriteria type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &SearchCriteria{}

// SearchCriteria struct for SearchCriteria
type SearchCriteria struct {
	// This option allows to exclude the isAllotment flag associated to a booking class in the search response when it exist.
	ExcludeAllotments *bool `json:"excludeAllotments,omitempty"`
	// This option allows activate the one-way combinable feature
	AddOneWayOffers *bool `json:"addOneWayOffers,omitempty"`
	// Maximum number of flight offers returned (Max 250)
	MaxFlightOffers *float32 `json:"maxFlightOffers,omitempty"`
	// maximum price per traveler. By default, no limit is applied. If specified, the value should be a positive number with no decimals
	MaxPrice *int32 `json:"maxPrice,omitempty"`
	// This option allows to default to a standard fareOption if no offers are found for the selected fareOption.
	AllowAlternativeFareOptions *bool `json:"allowAlternativeFareOptions,omitempty"`
	// Requests the system to find at least one flight-offer per day, if possible, when a range of dates is specified. Default is false.
	OneFlightOfferPerDay  *bool                   `json:"oneFlightOfferPerDay,omitempty"`
	AdditionalInformation *AdditionalInformation  `json:"additionalInformation,omitempty"`
	PricingOptions        *ExtendedPricingOptions `json:"pricingOptions,omitempty"`
	FlightFilters         *FlightFilters          `json:"flightFilters,omitempty"`
}

// NewSearchCriteria instantiates a new SearchCriteria object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchCriteria() *SearchCriteria {
	this := SearchCriteria{}
	return &this
}

// NewSearchCriteriaWithDefaults instantiates a new SearchCriteria object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchCriteriaWithDefaults() *SearchCriteria {
	this := SearchCriteria{}
	return &this
}

// GetExcludeAllotments returns the ExcludeAllotments field value if set, zero value otherwise.
func (o *SearchCriteria) GetExcludeAllotments() bool {
	if o == nil || utils.IsNil(o.ExcludeAllotments) {
		var ret bool
		return ret
	}
	return *o.ExcludeAllotments
}

// GetExcludeAllotmentsOk returns a tuple with the ExcludeAllotments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchCriteria) GetExcludeAllotmentsOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.ExcludeAllotments) {
		return nil, false
	}
	return o.ExcludeAllotments, true
}

// HasExcludeAllotments returns a boolean if a field has been set.
func (o *SearchCriteria) HasExcludeAllotments() bool {
	if o != nil && !utils.IsNil(o.ExcludeAllotments) {
		return true
	}

	return false
}

// SetExcludeAllotments gets a reference to the given bool and assigns it to the ExcludeAllotments field.
func (o *SearchCriteria) SetExcludeAllotments(v bool) {
	o.ExcludeAllotments = &v
}

// GetAddOneWayOffers returns the AddOneWayOffers field value if set, zero value otherwise.
func (o *SearchCriteria) GetAddOneWayOffers() bool {
	if o == nil || utils.IsNil(o.AddOneWayOffers) {
		var ret bool
		return ret
	}
	return *o.AddOneWayOffers
}

// GetAddOneWayOffersOk returns a tuple with the AddOneWayOffers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchCriteria) GetAddOneWayOffersOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.AddOneWayOffers) {
		return nil, false
	}
	return o.AddOneWayOffers, true
}

// HasAddOneWayOffers returns a boolean if a field has been set.
func (o *SearchCriteria) HasAddOneWayOffers() bool {
	if o != nil && !utils.IsNil(o.AddOneWayOffers) {
		return true
	}

	return false
}

// SetAddOneWayOffers gets a reference to the given bool and assigns it to the AddOneWayOffers field.
func (o *SearchCriteria) SetAddOneWayOffers(v bool) {
	o.AddOneWayOffers = &v
}

// GetMaxFlightOffers returns the MaxFlightOffers field value if set, zero value otherwise.
func (o *SearchCriteria) GetMaxFlightOffers() float32 {
	if o == nil || utils.IsNil(o.MaxFlightOffers) {
		var ret float32
		return ret
	}
	return *o.MaxFlightOffers
}

// GetMaxFlightOffersOk returns a tuple with the MaxFlightOffers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchCriteria) GetMaxFlightOffersOk() (*float32, bool) {
	if o == nil || utils.IsNil(o.MaxFlightOffers) {
		return nil, false
	}
	return o.MaxFlightOffers, true
}

// HasMaxFlightOffers returns a boolean if a field has been set.
func (o *SearchCriteria) HasMaxFlightOffers() bool {
	if o != nil && !utils.IsNil(o.MaxFlightOffers) {
		return true
	}

	return false
}

// SetMaxFlightOffers gets a reference to the given float32 and assigns it to the MaxFlightOffers field.
func (o *SearchCriteria) SetMaxFlightOffers(v float32) {
	o.MaxFlightOffers = &v
}

// GetMaxPrice returns the MaxPrice field value if set, zero value otherwise.
func (o *SearchCriteria) GetMaxPrice() int32 {
	if o == nil || utils.IsNil(o.MaxPrice) {
		var ret int32
		return ret
	}
	return *o.MaxPrice
}

// GetMaxPriceOk returns a tuple with the MaxPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchCriteria) GetMaxPriceOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.MaxPrice) {
		return nil, false
	}
	return o.MaxPrice, true
}

// HasMaxPrice returns a boolean if a field has been set.
func (o *SearchCriteria) HasMaxPrice() bool {
	if o != nil && !utils.IsNil(o.MaxPrice) {
		return true
	}

	return false
}

// SetMaxPrice gets a reference to the given int32 and assigns it to the MaxPrice field.
func (o *SearchCriteria) SetMaxPrice(v int32) {
	o.MaxPrice = &v
}

// GetAllowAlternativeFareOptions returns the AllowAlternativeFareOptions field value if set, zero value otherwise.
func (o *SearchCriteria) GetAllowAlternativeFareOptions() bool {
	if o == nil || utils.IsNil(o.AllowAlternativeFareOptions) {
		var ret bool
		return ret
	}
	return *o.AllowAlternativeFareOptions
}

// GetAllowAlternativeFareOptionsOk returns a tuple with the AllowAlternativeFareOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchCriteria) GetAllowAlternativeFareOptionsOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.AllowAlternativeFareOptions) {
		return nil, false
	}
	return o.AllowAlternativeFareOptions, true
}

// HasAllowAlternativeFareOptions returns a boolean if a field has been set.
func (o *SearchCriteria) HasAllowAlternativeFareOptions() bool {
	if o != nil && !utils.IsNil(o.AllowAlternativeFareOptions) {
		return true
	}

	return false
}

// SetAllowAlternativeFareOptions gets a reference to the given bool and assigns it to the AllowAlternativeFareOptions field.
func (o *SearchCriteria) SetAllowAlternativeFareOptions(v bool) {
	o.AllowAlternativeFareOptions = &v
}

// GetOneFlightOfferPerDay returns the OneFlightOfferPerDay field value if set, zero value otherwise.
func (o *SearchCriteria) GetOneFlightOfferPerDay() bool {
	if o == nil || utils.IsNil(o.OneFlightOfferPerDay) {
		var ret bool
		return ret
	}
	return *o.OneFlightOfferPerDay
}

// GetOneFlightOfferPerDayOk returns a tuple with the OneFlightOfferPerDay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchCriteria) GetOneFlightOfferPerDayOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.OneFlightOfferPerDay) {
		return nil, false
	}
	return o.OneFlightOfferPerDay, true
}

// HasOneFlightOfferPerDay returns a boolean if a field has been set.
func (o *SearchCriteria) HasOneFlightOfferPerDay() bool {
	if o != nil && !utils.IsNil(o.OneFlightOfferPerDay) {
		return true
	}

	return false
}

// SetOneFlightOfferPerDay gets a reference to the given bool and assigns it to the OneFlightOfferPerDay field.
func (o *SearchCriteria) SetOneFlightOfferPerDay(v bool) {
	o.OneFlightOfferPerDay = &v
}

// GetAdditionalInformation returns the AdditionalInformation field value if set, zero value otherwise.
func (o *SearchCriteria) GetAdditionalInformation() AdditionalInformation {
	if o == nil || utils.IsNil(o.AdditionalInformation) {
		var ret AdditionalInformation
		return ret
	}
	return *o.AdditionalInformation
}

// GetAdditionalInformationOk returns a tuple with the AdditionalInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchCriteria) GetAdditionalInformationOk() (*AdditionalInformation, bool) {
	if o == nil || utils.IsNil(o.AdditionalInformation) {
		return nil, false
	}
	return o.AdditionalInformation, true
}

// HasAdditionalInformation returns a boolean if a field has been set.
func (o *SearchCriteria) HasAdditionalInformation() bool {
	if o != nil && !utils.IsNil(o.AdditionalInformation) {
		return true
	}

	return false
}

// SetAdditionalInformation gets a reference to the given AdditionalInformation and assigns it to the AdditionalInformation field.
func (o *SearchCriteria) SetAdditionalInformation(v AdditionalInformation) {
	o.AdditionalInformation = &v
}

// GetPricingOptions returns the PricingOptions field value if set, zero value otherwise.
func (o *SearchCriteria) GetPricingOptions() ExtendedPricingOptions {
	if o == nil || utils.IsNil(o.PricingOptions) {
		var ret ExtendedPricingOptions
		return ret
	}
	return *o.PricingOptions
}

// GetPricingOptionsOk returns a tuple with the PricingOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchCriteria) GetPricingOptionsOk() (*ExtendedPricingOptions, bool) {
	if o == nil || utils.IsNil(o.PricingOptions) {
		return nil, false
	}
	return o.PricingOptions, true
}

// HasPricingOptions returns a boolean if a field has been set.
func (o *SearchCriteria) HasPricingOptions() bool {
	if o != nil && !utils.IsNil(o.PricingOptions) {
		return true
	}

	return false
}

// SetPricingOptions gets a reference to the given ExtendedPricingOptions and assigns it to the PricingOptions field.
func (o *SearchCriteria) SetPricingOptions(v ExtendedPricingOptions) {
	o.PricingOptions = &v
}

// GetFlightFilters returns the FlightFilters field value if set, zero value otherwise.
func (o *SearchCriteria) GetFlightFilters() FlightFilters {
	if o == nil || utils.IsNil(o.FlightFilters) {
		var ret FlightFilters
		return ret
	}
	return *o.FlightFilters
}

// GetFlightFiltersOk returns a tuple with the FlightFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchCriteria) GetFlightFiltersOk() (*FlightFilters, bool) {
	if o == nil || utils.IsNil(o.FlightFilters) {
		return nil, false
	}
	return o.FlightFilters, true
}

// HasFlightFilters returns a boolean if a field has been set.
func (o *SearchCriteria) HasFlightFilters() bool {
	if o != nil && !utils.IsNil(o.FlightFilters) {
		return true
	}

	return false
}

// SetFlightFilters gets a reference to the given FlightFilters and assigns it to the FlightFilters field.
func (o *SearchCriteria) SetFlightFilters(v FlightFilters) {
	o.FlightFilters = &v
}

func (o SearchCriteria) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchCriteria) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.ExcludeAllotments) {
		toSerialize["excludeAllotments"] = o.ExcludeAllotments
	}
	if !utils.IsNil(o.AddOneWayOffers) {
		toSerialize["addOneWayOffers"] = o.AddOneWayOffers
	}
	if !utils.IsNil(o.MaxFlightOffers) {
		toSerialize["maxFlightOffers"] = o.MaxFlightOffers
	}
	if !utils.IsNil(o.MaxPrice) {
		toSerialize["maxPrice"] = o.MaxPrice
	}
	if !utils.IsNil(o.AllowAlternativeFareOptions) {
		toSerialize["allowAlternativeFareOptions"] = o.AllowAlternativeFareOptions
	}
	if !utils.IsNil(o.OneFlightOfferPerDay) {
		toSerialize["oneFlightOfferPerDay"] = o.OneFlightOfferPerDay
	}
	if !utils.IsNil(o.AdditionalInformation) {
		toSerialize["additionalInformation"] = o.AdditionalInformation
	}
	if !utils.IsNil(o.PricingOptions) {
		toSerialize["pricingOptions"] = o.PricingOptions
	}
	if !utils.IsNil(o.FlightFilters) {
		toSerialize["flightFilters"] = o.FlightFilters
	}
	return toSerialize, nil
}

type NullableSearchCriteria struct {
	value *SearchCriteria
	isSet bool
}

func (v NullableSearchCriteria) Get() *SearchCriteria {
	return v.value
}

func (v *NullableSearchCriteria) Set(val *SearchCriteria) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchCriteria) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchCriteria) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchCriteria(val *SearchCriteria) *NullableSearchCriteria {
	return &NullableSearchCriteria{value: val, isSet: true}
}

func (v NullableSearchCriteria) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchCriteria) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
