package models

import (
	"encoding/json"
	"github.com/yenvanio/amadeus-go/amadeus/utils"
)

// checks if the FlightOfferPricingIn type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &FlightOfferPricingIn{}

// FlightOfferPricingIn input parameter to price flight offers element
type FlightOfferPricingIn struct {
	// the resource name
	Type string `json:"type"`
	// list of flight offer to price
	FlightOffers []FlightOffer `json:"flightOffers"`
	// payment information for retrieve eventual credit card fees
	Payments []Payment `json:"payments,omitempty"`
	// list of travelers
	Travelers []Traveler `json:"travelers,omitempty"`
}

// NewFlightOfferPricingIn instantiates a new FlightOfferPricingIn object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFlightOfferPricingIn(type_ string, flightOffers []FlightOffer) *FlightOfferPricingIn {
	this := FlightOfferPricingIn{}
	this.Type = type_
	this.FlightOffers = flightOffers
	return &this
}

// NewFlightOfferPricingInWithDefaults instantiates a new FlightOfferPricingIn object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFlightOfferPricingInWithDefaults() *FlightOfferPricingIn {
	this := FlightOfferPricingIn{}
	return &this
}

// GetType returns the Type field value
func (o *FlightOfferPricingIn) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *FlightOfferPricingIn) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *FlightOfferPricingIn) SetType(v string) {
	o.Type = v
}

// GetFlightOffers returns the FlightOffers field value
func (o *FlightOfferPricingIn) GetFlightOffers() []FlightOffer {
	if o == nil {
		var ret []FlightOffer
		return ret
	}

	return o.FlightOffers
}

// GetFlightOffersOk returns a tuple with the FlightOffers field value
// and a boolean to check if the value has been set.
func (o *FlightOfferPricingIn) GetFlightOffersOk() ([]FlightOffer, bool) {
	if o == nil {
		return nil, false
	}
	return o.FlightOffers, true
}

// SetFlightOffers sets field value
func (o *FlightOfferPricingIn) SetFlightOffers(v []FlightOffer) {
	o.FlightOffers = v
}

// GetPayments returns the Payments field value if set, zero value otherwise.
func (o *FlightOfferPricingIn) GetPayments() []Payment {
	if o == nil || utils.IsNil(o.Payments) {
		var ret []Payment
		return ret
	}
	return o.Payments
}

// GetPaymentsOk returns a tuple with the Payments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlightOfferPricingIn) GetPaymentsOk() ([]Payment, bool) {
	if o == nil || utils.IsNil(o.Payments) {
		return nil, false
	}
	return o.Payments, true
}

// HasPayments returns a boolean if a field has been set.
func (o *FlightOfferPricingIn) HasPayments() bool {
	if o != nil && !utils.IsNil(o.Payments) {
		return true
	}

	return false
}

// SetPayments gets a reference to the given []Payment and assigns it to the Payments field.
func (o *FlightOfferPricingIn) SetPayments(v []Payment) {
	o.Payments = v
}

// GetTravelers returns the Travelers field value if set, zero value otherwise.
func (o *FlightOfferPricingIn) GetTravelers() []Traveler {
	if o == nil || utils.IsNil(o.Travelers) {
		var ret []Traveler
		return ret
	}
	return o.Travelers
}

// GetTravelersOk returns a tuple with the Travelers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlightOfferPricingIn) GetTravelersOk() ([]Traveler, bool) {
	if o == nil || utils.IsNil(o.Travelers) {
		return nil, false
	}
	return o.Travelers, true
}

// HasTravelers returns a boolean if a field has been set.
func (o *FlightOfferPricingIn) HasTravelers() bool {
	if o != nil && !utils.IsNil(o.Travelers) {
		return true
	}

	return false
}

// SetTravelers gets a reference to the given []Traveler and assigns it to the Travelers field.
func (o *FlightOfferPricingIn) SetTravelers(v []Traveler) {
	o.Travelers = v
}

func (o FlightOfferPricingIn) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FlightOfferPricingIn) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["flightOffers"] = o.FlightOffers
	if !utils.IsNil(o.Payments) {
		toSerialize["payments"] = o.Payments
	}
	if !utils.IsNil(o.Travelers) {
		toSerialize["travelers"] = o.Travelers
	}
	return toSerialize, nil
}

type NullableFlightOfferPricingIn struct {
	value *FlightOfferPricingIn
	isSet bool
}

func (v NullableFlightOfferPricingIn) Get() *FlightOfferPricingIn {
	return v.value
}

func (v *NullableFlightOfferPricingIn) Set(val *FlightOfferPricingIn) {
	v.value = val
	v.isSet = true
}

func (v NullableFlightOfferPricingIn) IsSet() bool {
	return v.isSet
}

func (v *NullableFlightOfferPricingIn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFlightOfferPricingIn(val *FlightOfferPricingIn) *NullableFlightOfferPricingIn {
	return &NullableFlightOfferPricingIn{value: val, isSet: true}
}

func (v NullableFlightOfferPricingIn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFlightOfferPricingIn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
