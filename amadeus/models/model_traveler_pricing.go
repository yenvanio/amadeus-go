package models

import (
	"encoding/json"
	"github.com/yenvanio/amadeus-go/amadeus/utils"
)

// checks if the TravelerPricing type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &TravelerPricing{}

// TravelerPricing struct for TravelerPricing
type TravelerPricing struct {
	// Id of the traveler
	TravelerId   string                    `json:"travelerId"`
	FareOption   TravelerPricingFareOption `json:"fareOption"`
	TravelerType TravelerType              `json:"travelerType"`
	// if type=\"HELD_INFANT\", corresponds to the adult traveler's id who will share the seat
	AssociatedAdultId    *string                `json:"associatedAdultId,omitempty"`
	Price                *Price                 `json:"price,omitempty"`
	FareDetailsBySegment []FareDetailsBySegment `json:"fareDetailsBySegment"`
}

// NewTravelerPricing instantiates a new TravelerPricing object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTravelerPricing(travelerId string, fareOption TravelerPricingFareOption, travelerType TravelerType, fareDetailsBySegment []FareDetailsBySegment) *TravelerPricing {
	this := TravelerPricing{}
	this.TravelerId = travelerId
	this.FareOption = fareOption
	this.TravelerType = travelerType
	this.FareDetailsBySegment = fareDetailsBySegment
	return &this
}

// NewTravelerPricingWithDefaults instantiates a new TravelerPricing object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTravelerPricingWithDefaults() *TravelerPricing {
	this := TravelerPricing{}
	return &this
}

// GetTravelerId returns the TravelerId field value
func (o *TravelerPricing) GetTravelerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TravelerId
}

// GetTravelerIdOk returns a tuple with the TravelerId field value
// and a boolean to check if the value has been set.
func (o *TravelerPricing) GetTravelerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TravelerId, true
}

// SetTravelerId sets field value
func (o *TravelerPricing) SetTravelerId(v string) {
	o.TravelerId = v
}

// GetFareOption returns the FareOption field value
func (o *TravelerPricing) GetFareOption() TravelerPricingFareOption {
	if o == nil {
		var ret TravelerPricingFareOption
		return ret
	}

	return o.FareOption
}

// GetFareOptionOk returns a tuple with the FareOption field value
// and a boolean to check if the value has been set.
func (o *TravelerPricing) GetFareOptionOk() (*TravelerPricingFareOption, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FareOption, true
}

// SetFareOption sets field value
func (o *TravelerPricing) SetFareOption(v TravelerPricingFareOption) {
	o.FareOption = v
}

// GetTravelerType returns the TravelerType field value
func (o *TravelerPricing) GetTravelerType() TravelerType {
	if o == nil {
		var ret TravelerType
		return ret
	}

	return o.TravelerType
}

// GetTravelerTypeOk returns a tuple with the TravelerType field value
// and a boolean to check if the value has been set.
func (o *TravelerPricing) GetTravelerTypeOk() (*TravelerType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TravelerType, true
}

// SetTravelerType sets field value
func (o *TravelerPricing) SetTravelerType(v TravelerType) {
	o.TravelerType = v
}

// GetAssociatedAdultId returns the AssociatedAdultId field value if set, zero value otherwise.
func (o *TravelerPricing) GetAssociatedAdultId() string {
	if o == nil || utils.IsNil(o.AssociatedAdultId) {
		var ret string
		return ret
	}
	return *o.AssociatedAdultId
}

// GetAssociatedAdultIdOk returns a tuple with the AssociatedAdultId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TravelerPricing) GetAssociatedAdultIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.AssociatedAdultId) {
		return nil, false
	}
	return o.AssociatedAdultId, true
}

// HasAssociatedAdultId returns a boolean if a field has been set.
func (o *TravelerPricing) HasAssociatedAdultId() bool {
	if o != nil && !utils.IsNil(o.AssociatedAdultId) {
		return true
	}

	return false
}

// SetAssociatedAdultId gets a reference to the given string and assigns it to the AssociatedAdultId field.
func (o *TravelerPricing) SetAssociatedAdultId(v string) {
	o.AssociatedAdultId = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *TravelerPricing) GetPrice() Price {
	if o == nil || utils.IsNil(o.Price) {
		var ret Price
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TravelerPricing) GetPriceOk() (*Price, bool) {
	if o == nil || utils.IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *TravelerPricing) HasPrice() bool {
	if o != nil && !utils.IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given Price and assigns it to the Price field.
func (o *TravelerPricing) SetPrice(v Price) {
	o.Price = &v
}

// GetFareDetailsBySegment returns the FareDetailsBySegment field value
func (o *TravelerPricing) GetFareDetailsBySegment() []FareDetailsBySegment {
	if o == nil {
		var ret []FareDetailsBySegment
		return ret
	}

	return o.FareDetailsBySegment
}

// GetFareDetailsBySegmentOk returns a tuple with the FareDetailsBySegment field value
// and a boolean to check if the value has been set.
func (o *TravelerPricing) GetFareDetailsBySegmentOk() ([]FareDetailsBySegment, bool) {
	if o == nil {
		return nil, false
	}
	return o.FareDetailsBySegment, true
}

// SetFareDetailsBySegment sets field value
func (o *TravelerPricing) SetFareDetailsBySegment(v []FareDetailsBySegment) {
	o.FareDetailsBySegment = v
}

func (o TravelerPricing) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TravelerPricing) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["travelerId"] = o.TravelerId
	toSerialize["fareOption"] = o.FareOption
	toSerialize["travelerType"] = o.TravelerType
	if !utils.IsNil(o.AssociatedAdultId) {
		toSerialize["associatedAdultId"] = o.AssociatedAdultId
	}
	if !utils.IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	toSerialize["fareDetailsBySegment"] = o.FareDetailsBySegment
	return toSerialize, nil
}

type NullableTravelerPricing struct {
	value *TravelerPricing
	isSet bool
}

func (v NullableTravelerPricing) Get() *TravelerPricing {
	return v.value
}

func (v *NullableTravelerPricing) Set(val *TravelerPricing) {
	v.value = val
	v.isSet = true
}

func (v NullableTravelerPricing) IsSet() bool {
	return v.isSet
}

func (v *NullableTravelerPricing) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTravelerPricing(val *TravelerPricing) *NullableTravelerPricing {
	return &NullableTravelerPricing{value: val, isSet: true}
}

func (v NullableTravelerPricing) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTravelerPricing) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
