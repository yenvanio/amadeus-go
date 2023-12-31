package models

import (
	"encoding/json"
	"github.com/yenvanio/amadeus-go/amadeus/utils"
)

// checks if the FlightOffer type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &FlightOffer{}

// FlightOffer struct for FlightOffer
type FlightOffer struct {
	// the resource name
	Type string `json:"type"`
	// Id of the flight offer
	Id     string             `json:"id"`
	Source *FlightOfferSource `json:"source,omitempty"`
	// If true, inform that a ticketing will be required at booking step.
	InstantTicketingRequired *bool `json:"instantTicketingRequired,omitempty"`
	// BOOK step ONLY - If true, allows to book a PNR without pricing. Only for the source \"GDS\"
	DisablePricing *bool `json:"disablePricing,omitempty"`
	// If true, upon completion of the booking, this pricing solution is expected to yield multiple records (a record contains booking information confirmed and stored, typically a Passenger Name Record (PNR), in the provider GDS or system)
	NonHomogeneous *bool `json:"nonHomogeneous,omitempty"`
	// If true, the flight offer can be combined with other oneWays flight-offers to complete the whole journey (one-Way combinable feature).
	OneWay *bool `json:"oneWay,omitempty"`
	// If true, a payment card is mandatory to book this flight offer
	PaymentCardRequired *bool `json:"paymentCardRequired,omitempty"`
	// If booked on the same day as the search (with respect to timezone), this flight offer is guaranteed to be thereafter valid for ticketing until this date (included). Unspecified when it does not make sense for this flight offer (e.g. no control over ticketing once booked). YYYY-MM-DD format, e.g. 2019-06-07
	LastTicketingDate *string `json:"lastTicketingDate,omitempty"`
	// If booked on the same day as the search (with respect to timezone), this flight offer is guaranteed to be thereafter valid for ticketing until this date/time (included). Unspecified when it does not make sense for this flight offer (e.g. no control over ticketing once booked). Information only this attribute is not used in input of pricing request. Local date and time in YYYY-MM-ddThh:mm:ss format, e.g. 2017-02-10T20:40:00
	LastTicketingDateTime *string `json:"lastTicketingDateTime,omitempty"`
	// Number of seats bookable in a single request. Can not be higher than 9.
	NumberOfBookableSeats *float32        `json:"numberOfBookableSeats,omitempty"`
	Itineraries           []Itineraries   `json:"itineraries,omitempty"`
	Price                 *ExtendedPrice  `json:"price,omitempty"`
	PricingOptions        *PricingOptions `json:"pricingOptions,omitempty"`
	// This option ensures that the system will only consider offers with these airlines as validating carrier.
	ValidatingAirlineCodes []string `json:"validatingAirlineCodes,omitempty"`
	// Fare information for each traveler/segment
	TravelerPricings []TravelerPricing `json:"travelerPricings,omitempty"`
}

// NewFlightOffer instantiates a new FlightOffer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFlightOffer(type_ string, id string) *FlightOffer {
	this := FlightOffer{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewFlightOfferWithDefaults instantiates a new FlightOffer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFlightOfferWithDefaults() *FlightOffer {
	this := FlightOffer{}
	return &this
}

// GetType returns the Type field value
func (o *FlightOffer) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *FlightOffer) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *FlightOffer) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *FlightOffer) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *FlightOffer) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *FlightOffer) SetId(v string) {
	o.Id = v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *FlightOffer) GetSource() FlightOfferSource {
	if o == nil || utils.IsNil(o.Source) {
		var ret FlightOfferSource
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlightOffer) GetSourceOk() (*FlightOfferSource, bool) {
	if o == nil || utils.IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *FlightOffer) HasSource() bool {
	if o != nil && !utils.IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given FlightOfferSource and assigns it to the Source field.
func (o *FlightOffer) SetSource(v FlightOfferSource) {
	o.Source = &v
}

// GetInstantTicketingRequired returns the InstantTicketingRequired field value if set, zero value otherwise.
func (o *FlightOffer) GetInstantTicketingRequired() bool {
	if o == nil || utils.IsNil(o.InstantTicketingRequired) {
		var ret bool
		return ret
	}
	return *o.InstantTicketingRequired
}

// GetInstantTicketingRequiredOk returns a tuple with the InstantTicketingRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlightOffer) GetInstantTicketingRequiredOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.InstantTicketingRequired) {
		return nil, false
	}
	return o.InstantTicketingRequired, true
}

// HasInstantTicketingRequired returns a boolean if a field has been set.
func (o *FlightOffer) HasInstantTicketingRequired() bool {
	if o != nil && !utils.IsNil(o.InstantTicketingRequired) {
		return true
	}

	return false
}

// SetInstantTicketingRequired gets a reference to the given bool and assigns it to the InstantTicketingRequired field.
func (o *FlightOffer) SetInstantTicketingRequired(v bool) {
	o.InstantTicketingRequired = &v
}

// GetDisablePricing returns the DisablePricing field value if set, zero value otherwise.
func (o *FlightOffer) GetDisablePricing() bool {
	if o == nil || utils.IsNil(o.DisablePricing) {
		var ret bool
		return ret
	}
	return *o.DisablePricing
}

// GetDisablePricingOk returns a tuple with the DisablePricing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlightOffer) GetDisablePricingOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.DisablePricing) {
		return nil, false
	}
	return o.DisablePricing, true
}

// HasDisablePricing returns a boolean if a field has been set.
func (o *FlightOffer) HasDisablePricing() bool {
	if o != nil && !utils.IsNil(o.DisablePricing) {
		return true
	}

	return false
}

// SetDisablePricing gets a reference to the given bool and assigns it to the DisablePricing field.
func (o *FlightOffer) SetDisablePricing(v bool) {
	o.DisablePricing = &v
}

// GetNonHomogeneous returns the NonHomogeneous field value if set, zero value otherwise.
func (o *FlightOffer) GetNonHomogeneous() bool {
	if o == nil || utils.IsNil(o.NonHomogeneous) {
		var ret bool
		return ret
	}
	return *o.NonHomogeneous
}

// GetNonHomogeneousOk returns a tuple with the NonHomogeneous field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlightOffer) GetNonHomogeneousOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.NonHomogeneous) {
		return nil, false
	}
	return o.NonHomogeneous, true
}

// HasNonHomogeneous returns a boolean if a field has been set.
func (o *FlightOffer) HasNonHomogeneous() bool {
	if o != nil && !utils.IsNil(o.NonHomogeneous) {
		return true
	}

	return false
}

// SetNonHomogeneous gets a reference to the given bool and assigns it to the NonHomogeneous field.
func (o *FlightOffer) SetNonHomogeneous(v bool) {
	o.NonHomogeneous = &v
}

// GetOneWay returns the OneWay field value if set, zero value otherwise.
func (o *FlightOffer) GetOneWay() bool {
	if o == nil || utils.IsNil(o.OneWay) {
		var ret bool
		return ret
	}
	return *o.OneWay
}

// GetOneWayOk returns a tuple with the OneWay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlightOffer) GetOneWayOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.OneWay) {
		return nil, false
	}
	return o.OneWay, true
}

// HasOneWay returns a boolean if a field has been set.
func (o *FlightOffer) HasOneWay() bool {
	if o != nil && !utils.IsNil(o.OneWay) {
		return true
	}

	return false
}

// SetOneWay gets a reference to the given bool and assigns it to the OneWay field.
func (o *FlightOffer) SetOneWay(v bool) {
	o.OneWay = &v
}

// GetPaymentCardRequired returns the PaymentCardRequired field value if set, zero value otherwise.
func (o *FlightOffer) GetPaymentCardRequired() bool {
	if o == nil || utils.IsNil(o.PaymentCardRequired) {
		var ret bool
		return ret
	}
	return *o.PaymentCardRequired
}

// GetPaymentCardRequiredOk returns a tuple with the PaymentCardRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlightOffer) GetPaymentCardRequiredOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.PaymentCardRequired) {
		return nil, false
	}
	return o.PaymentCardRequired, true
}

// HasPaymentCardRequired returns a boolean if a field has been set.
func (o *FlightOffer) HasPaymentCardRequired() bool {
	if o != nil && !utils.IsNil(o.PaymentCardRequired) {
		return true
	}

	return false
}

// SetPaymentCardRequired gets a reference to the given bool and assigns it to the PaymentCardRequired field.
func (o *FlightOffer) SetPaymentCardRequired(v bool) {
	o.PaymentCardRequired = &v
}

// GetLastTicketingDate returns the LastTicketingDate field value if set, zero value otherwise.
func (o *FlightOffer) GetLastTicketingDate() string {
	if o == nil || utils.IsNil(o.LastTicketingDate) {
		var ret string
		return ret
	}
	return *o.LastTicketingDate
}

// GetLastTicketingDateOk returns a tuple with the LastTicketingDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlightOffer) GetLastTicketingDateOk() (*string, bool) {
	if o == nil || utils.IsNil(o.LastTicketingDate) {
		return nil, false
	}
	return o.LastTicketingDate, true
}

// HasLastTicketingDate returns a boolean if a field has been set.
func (o *FlightOffer) HasLastTicketingDate() bool {
	if o != nil && !utils.IsNil(o.LastTicketingDate) {
		return true
	}

	return false
}

// SetLastTicketingDate gets a reference to the given string and assigns it to the LastTicketingDate field.
func (o *FlightOffer) SetLastTicketingDate(v string) {
	o.LastTicketingDate = &v
}

// GetLastTicketingDateTime returns the LastTicketingDateTime field value if set, zero value otherwise.
func (o *FlightOffer) GetLastTicketingDateTime() string {
	if o == nil || utils.IsNil(o.LastTicketingDateTime) {
		var ret string
		return ret
	}
	return *o.LastTicketingDateTime
}

// GetLastTicketingDateTimeOk returns a tuple with the LastTicketingDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlightOffer) GetLastTicketingDateTimeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.LastTicketingDateTime) {
		return nil, false
	}
	return o.LastTicketingDateTime, true
}

// HasLastTicketingDateTime returns a boolean if a field has been set.
func (o *FlightOffer) HasLastTicketingDateTime() bool {
	if o != nil && !utils.IsNil(o.LastTicketingDateTime) {
		return true
	}

	return false
}

// SetLastTicketingDateTime gets a reference to the given string and assigns it to the LastTicketingDateTime field.
func (o *FlightOffer) SetLastTicketingDateTime(v string) {
	o.LastTicketingDateTime = &v
}

// GetNumberOfBookableSeats returns the NumberOfBookableSeats field value if set, zero value otherwise.
func (o *FlightOffer) GetNumberOfBookableSeats() float32 {
	if o == nil || utils.IsNil(o.NumberOfBookableSeats) {
		var ret float32
		return ret
	}
	return *o.NumberOfBookableSeats
}

// GetNumberOfBookableSeatsOk returns a tuple with the NumberOfBookableSeats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlightOffer) GetNumberOfBookableSeatsOk() (*float32, bool) {
	if o == nil || utils.IsNil(o.NumberOfBookableSeats) {
		return nil, false
	}
	return o.NumberOfBookableSeats, true
}

// HasNumberOfBookableSeats returns a boolean if a field has been set.
func (o *FlightOffer) HasNumberOfBookableSeats() bool {
	if o != nil && !utils.IsNil(o.NumberOfBookableSeats) {
		return true
	}

	return false
}

// SetNumberOfBookableSeats gets a reference to the given float32 and assigns it to the NumberOfBookableSeats field.
func (o *FlightOffer) SetNumberOfBookableSeats(v float32) {
	o.NumberOfBookableSeats = &v
}

// GetItineraries returns the Itineraries field value if set, zero value otherwise.
func (o *FlightOffer) GetItineraries() []Itineraries {
	if o == nil || utils.IsNil(o.Itineraries) {
		var ret []Itineraries
		return ret
	}
	return o.Itineraries
}

// GetItinerariesOk returns a tuple with the Itineraries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlightOffer) GetItinerariesOk() ([]Itineraries, bool) {
	if o == nil || utils.IsNil(o.Itineraries) {
		return nil, false
	}
	return o.Itineraries, true
}

// HasItineraries returns a boolean if a field has been set.
func (o *FlightOffer) HasItineraries() bool {
	if o != nil && !utils.IsNil(o.Itineraries) {
		return true
	}

	return false
}

// SetItineraries gets a reference to the given []Itineraries and assigns it to the Itineraries field.
func (o *FlightOffer) SetItineraries(v []Itineraries) {
	o.Itineraries = v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *FlightOffer) GetPrice() ExtendedPrice {
	if o == nil || utils.IsNil(o.Price) {
		var ret ExtendedPrice
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlightOffer) GetPriceOk() (*ExtendedPrice, bool) {
	if o == nil || utils.IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *FlightOffer) HasPrice() bool {
	if o != nil && !utils.IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given ExtendedPrice and assigns it to the Price field.
func (o *FlightOffer) SetPrice(v ExtendedPrice) {
	o.Price = &v
}

// GetPricingOptions returns the PricingOptions field value if set, zero value otherwise.
func (o *FlightOffer) GetPricingOptions() PricingOptions {
	if o == nil || utils.IsNil(o.PricingOptions) {
		var ret PricingOptions
		return ret
	}
	return *o.PricingOptions
}

// GetPricingOptionsOk returns a tuple with the PricingOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlightOffer) GetPricingOptionsOk() (*PricingOptions, bool) {
	if o == nil || utils.IsNil(o.PricingOptions) {
		return nil, false
	}
	return o.PricingOptions, true
}

// HasPricingOptions returns a boolean if a field has been set.
func (o *FlightOffer) HasPricingOptions() bool {
	if o != nil && !utils.IsNil(o.PricingOptions) {
		return true
	}

	return false
}

// SetPricingOptions gets a reference to the given PricingOptions and assigns it to the PricingOptions field.
func (o *FlightOffer) SetPricingOptions(v PricingOptions) {
	o.PricingOptions = &v
}

// GetValidatingAirlineCodes returns the ValidatingAirlineCodes field value if set, zero value otherwise.
func (o *FlightOffer) GetValidatingAirlineCodes() []string {
	if o == nil || utils.IsNil(o.ValidatingAirlineCodes) {
		var ret []string
		return ret
	}
	return o.ValidatingAirlineCodes
}

// GetValidatingAirlineCodesOk returns a tuple with the ValidatingAirlineCodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlightOffer) GetValidatingAirlineCodesOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.ValidatingAirlineCodes) {
		return nil, false
	}
	return o.ValidatingAirlineCodes, true
}

// HasValidatingAirlineCodes returns a boolean if a field has been set.
func (o *FlightOffer) HasValidatingAirlineCodes() bool {
	if o != nil && !utils.IsNil(o.ValidatingAirlineCodes) {
		return true
	}

	return false
}

// SetValidatingAirlineCodes gets a reference to the given []string and assigns it to the ValidatingAirlineCodes field.
func (o *FlightOffer) SetValidatingAirlineCodes(v []string) {
	o.ValidatingAirlineCodes = v
}

// GetTravelerPricings returns the TravelerPricings field value if set, zero value otherwise.
func (o *FlightOffer) GetTravelerPricings() []TravelerPricing {
	if o == nil || utils.IsNil(o.TravelerPricings) {
		var ret []TravelerPricing
		return ret
	}
	return o.TravelerPricings
}

// GetTravelerPricingsOk returns a tuple with the TravelerPricings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlightOffer) GetTravelerPricingsOk() ([]TravelerPricing, bool) {
	if o == nil || utils.IsNil(o.TravelerPricings) {
		return nil, false
	}
	return o.TravelerPricings, true
}

// HasTravelerPricings returns a boolean if a field has been set.
func (o *FlightOffer) HasTravelerPricings() bool {
	if o != nil && !utils.IsNil(o.TravelerPricings) {
		return true
	}

	return false
}

// SetTravelerPricings gets a reference to the given []TravelerPricing and assigns it to the TravelerPricings field.
func (o *FlightOffer) SetTravelerPricings(v []TravelerPricing) {
	o.TravelerPricings = v
}

func (o FlightOffer) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FlightOffer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	if !utils.IsNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	if !utils.IsNil(o.InstantTicketingRequired) {
		toSerialize["instantTicketingRequired"] = o.InstantTicketingRequired
	}
	if !utils.IsNil(o.DisablePricing) {
		toSerialize["disablePricing"] = o.DisablePricing
	}
	if !utils.IsNil(o.NonHomogeneous) {
		toSerialize["nonHomogeneous"] = o.NonHomogeneous
	}
	if !utils.IsNil(o.OneWay) {
		toSerialize["oneWay"] = o.OneWay
	}
	if !utils.IsNil(o.PaymentCardRequired) {
		toSerialize["paymentCardRequired"] = o.PaymentCardRequired
	}
	if !utils.IsNil(o.LastTicketingDate) {
		toSerialize["lastTicketingDate"] = o.LastTicketingDate
	}
	if !utils.IsNil(o.LastTicketingDateTime) {
		toSerialize["lastTicketingDateTime"] = o.LastTicketingDateTime
	}
	if !utils.IsNil(o.NumberOfBookableSeats) {
		toSerialize["numberOfBookableSeats"] = o.NumberOfBookableSeats
	}
	if !utils.IsNil(o.Itineraries) {
		toSerialize["itineraries"] = o.Itineraries
	}
	if !utils.IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if !utils.IsNil(o.PricingOptions) {
		toSerialize["pricingOptions"] = o.PricingOptions
	}
	if !utils.IsNil(o.ValidatingAirlineCodes) {
		toSerialize["validatingAirlineCodes"] = o.ValidatingAirlineCodes
	}
	if !utils.IsNil(o.TravelerPricings) {
		toSerialize["travelerPricings"] = o.TravelerPricings
	}
	return toSerialize, nil
}

type NullableFlightOffer struct {
	value *FlightOffer
	isSet bool
}

func (v NullableFlightOffer) Get() *FlightOffer {
	return v.value
}

func (v *NullableFlightOffer) Set(val *FlightOffer) {
	v.value = val
	v.isSet = true
}

func (v NullableFlightOffer) IsSet() bool {
	return v.isSet
}

func (v *NullableFlightOffer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFlightOffer(val *FlightOffer) *NullableFlightOffer {
	return &NullableFlightOffer{value: val, isSet: true}
}

func (v NullableFlightOffer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFlightOffer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
