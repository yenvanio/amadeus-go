package models

import (
	"encoding/json"
	"github.com/yenvanio/amadeus-go/amadeus/utils"
)

// checks if the FareDetailsBySegment type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &FareDetailsBySegment{}

// FareDetailsBySegment Fare details of the segment
type FareDetailsBySegment struct {
	// Id of the segment
	SegmentId string       `json:"segmentId"`
	Cabin     *TravelClass `json:"cabin,omitempty"`
	// Fare basis specifying the rules of a fare. Usually, though not always, is composed of the booking class code followed by a set of letters and digits representing other characteristics of the ticket, such as refundability, minimum stay requirements, discounts or special promotional elements.
	FareBasis *string `json:"fareBasis,omitempty"`
	// The name of the Fare Family corresponding to the fares. Only for the GDS provider and if the airline has fare families filled
	BrandedFare *string `json:"brandedFare,omitempty"`
	// The code of the booking class, a.k.a. class of service or Reservations/Booking Designator (RBD)
	Class *string `json:"class,omitempty"`
	// True if the corresponding booking class is in an allotment
	IsAllotment         *bool                      `json:"isAllotment,omitempty"`
	AllotmentDetails    *AllotmentDetails          `json:"allotmentDetails,omitempty"`
	SliceDiceIndicator  *SliceDiceIndicator        `json:"sliceDiceIndicator,omitempty"`
	IncludedCheckedBags *BaggageAllowance          `json:"includedCheckedBags,omitempty"`
	AdditionalServices  *AdditionalServicesRequest `json:"additionalServices,omitempty"`
}

// NewFareDetailsBySegment instantiates a new FareDetailsBySegment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFareDetailsBySegment(segmentId string) *FareDetailsBySegment {
	this := FareDetailsBySegment{}
	this.SegmentId = segmentId
	return &this
}

// NewFareDetailsBySegmentWithDefaults instantiates a new FareDetailsBySegment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFareDetailsBySegmentWithDefaults() *FareDetailsBySegment {
	this := FareDetailsBySegment{}
	return &this
}

// GetSegmentId returns the SegmentId field value
func (o *FareDetailsBySegment) GetSegmentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SegmentId
}

// GetSegmentIdOk returns a tuple with the SegmentId field value
// and a boolean to check if the value has been set.
func (o *FareDetailsBySegment) GetSegmentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SegmentId, true
}

// SetSegmentId sets field value
func (o *FareDetailsBySegment) SetSegmentId(v string) {
	o.SegmentId = v
}

// GetCabin returns the Cabin field value if set, zero value otherwise.
func (o *FareDetailsBySegment) GetCabin() TravelClass {
	if o == nil || utils.IsNil(o.Cabin) {
		var ret TravelClass
		return ret
	}
	return *o.Cabin
}

// GetCabinOk returns a tuple with the Cabin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FareDetailsBySegment) GetCabinOk() (*TravelClass, bool) {
	if o == nil || utils.IsNil(o.Cabin) {
		return nil, false
	}
	return o.Cabin, true
}

// HasCabin returns a boolean if a field has been set.
func (o *FareDetailsBySegment) HasCabin() bool {
	if o != nil && !utils.IsNil(o.Cabin) {
		return true
	}

	return false
}

// SetCabin gets a reference to the given TravelClass and assigns it to the Cabin field.
func (o *FareDetailsBySegment) SetCabin(v TravelClass) {
	o.Cabin = &v
}

// GetFareBasis returns the FareBasis field value if set, zero value otherwise.
func (o *FareDetailsBySegment) GetFareBasis() string {
	if o == nil || utils.IsNil(o.FareBasis) {
		var ret string
		return ret
	}
	return *o.FareBasis
}

// GetFareBasisOk returns a tuple with the FareBasis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FareDetailsBySegment) GetFareBasisOk() (*string, bool) {
	if o == nil || utils.IsNil(o.FareBasis) {
		return nil, false
	}
	return o.FareBasis, true
}

// HasFareBasis returns a boolean if a field has been set.
func (o *FareDetailsBySegment) HasFareBasis() bool {
	if o != nil && !utils.IsNil(o.FareBasis) {
		return true
	}

	return false
}

// SetFareBasis gets a reference to the given string and assigns it to the FareBasis field.
func (o *FareDetailsBySegment) SetFareBasis(v string) {
	o.FareBasis = &v
}

// GetBrandedFare returns the BrandedFare field value if set, zero value otherwise.
func (o *FareDetailsBySegment) GetBrandedFare() string {
	if o == nil || utils.IsNil(o.BrandedFare) {
		var ret string
		return ret
	}
	return *o.BrandedFare
}

// GetBrandedFareOk returns a tuple with the BrandedFare field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FareDetailsBySegment) GetBrandedFareOk() (*string, bool) {
	if o == nil || utils.IsNil(o.BrandedFare) {
		return nil, false
	}
	return o.BrandedFare, true
}

// HasBrandedFare returns a boolean if a field has been set.
func (o *FareDetailsBySegment) HasBrandedFare() bool {
	if o != nil && !utils.IsNil(o.BrandedFare) {
		return true
	}

	return false
}

// SetBrandedFare gets a reference to the given string and assigns it to the BrandedFare field.
func (o *FareDetailsBySegment) SetBrandedFare(v string) {
	o.BrandedFare = &v
}

// GetClass returns the Class field value if set, zero value otherwise.
func (o *FareDetailsBySegment) GetClass() string {
	if o == nil || utils.IsNil(o.Class) {
		var ret string
		return ret
	}
	return *o.Class
}

// GetClassOk returns a tuple with the Class field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FareDetailsBySegment) GetClassOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Class) {
		return nil, false
	}
	return o.Class, true
}

// HasClass returns a boolean if a field has been set.
func (o *FareDetailsBySegment) HasClass() bool {
	if o != nil && !utils.IsNil(o.Class) {
		return true
	}

	return false
}

// SetClass gets a reference to the given string and assigns it to the Class field.
func (o *FareDetailsBySegment) SetClass(v string) {
	o.Class = &v
}

// GetIsAllotment returns the IsAllotment field value if set, zero value otherwise.
func (o *FareDetailsBySegment) GetIsAllotment() bool {
	if o == nil || utils.IsNil(o.IsAllotment) {
		var ret bool
		return ret
	}
	return *o.IsAllotment
}

// GetIsAllotmentOk returns a tuple with the IsAllotment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FareDetailsBySegment) GetIsAllotmentOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.IsAllotment) {
		return nil, false
	}
	return o.IsAllotment, true
}

// HasIsAllotment returns a boolean if a field has been set.
func (o *FareDetailsBySegment) HasIsAllotment() bool {
	if o != nil && !utils.IsNil(o.IsAllotment) {
		return true
	}

	return false
}

// SetIsAllotment gets a reference to the given bool and assigns it to the IsAllotment field.
func (o *FareDetailsBySegment) SetIsAllotment(v bool) {
	o.IsAllotment = &v
}

// GetAllotmentDetails returns the AllotmentDetails field value if set, zero value otherwise.
func (o *FareDetailsBySegment) GetAllotmentDetails() AllotmentDetails {
	if o == nil || utils.IsNil(o.AllotmentDetails) {
		var ret AllotmentDetails
		return ret
	}
	return *o.AllotmentDetails
}

// GetAllotmentDetailsOk returns a tuple with the AllotmentDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FareDetailsBySegment) GetAllotmentDetailsOk() (*AllotmentDetails, bool) {
	if o == nil || utils.IsNil(o.AllotmentDetails) {
		return nil, false
	}
	return o.AllotmentDetails, true
}

// HasAllotmentDetails returns a boolean if a field has been set.
func (o *FareDetailsBySegment) HasAllotmentDetails() bool {
	if o != nil && !utils.IsNil(o.AllotmentDetails) {
		return true
	}

	return false
}

// SetAllotmentDetails gets a reference to the given AllotmentDetails and assigns it to the AllotmentDetails field.
func (o *FareDetailsBySegment) SetAllotmentDetails(v AllotmentDetails) {
	o.AllotmentDetails = &v
}

// GetSliceDiceIndicator returns the SliceDiceIndicator field value if set, zero value otherwise.
func (o *FareDetailsBySegment) GetSliceDiceIndicator() SliceDiceIndicator {
	if o == nil || utils.IsNil(o.SliceDiceIndicator) {
		var ret SliceDiceIndicator
		return ret
	}
	return *o.SliceDiceIndicator
}

// GetSliceDiceIndicatorOk returns a tuple with the SliceDiceIndicator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FareDetailsBySegment) GetSliceDiceIndicatorOk() (*SliceDiceIndicator, bool) {
	if o == nil || utils.IsNil(o.SliceDiceIndicator) {
		return nil, false
	}
	return o.SliceDiceIndicator, true
}

// HasSliceDiceIndicator returns a boolean if a field has been set.
func (o *FareDetailsBySegment) HasSliceDiceIndicator() bool {
	if o != nil && !utils.IsNil(o.SliceDiceIndicator) {
		return true
	}

	return false
}

// SetSliceDiceIndicator gets a reference to the given SliceDiceIndicator and assigns it to the SliceDiceIndicator field.
func (o *FareDetailsBySegment) SetSliceDiceIndicator(v SliceDiceIndicator) {
	o.SliceDiceIndicator = &v
}

// GetIncludedCheckedBags returns the IncludedCheckedBags field value if set, zero value otherwise.
func (o *FareDetailsBySegment) GetIncludedCheckedBags() BaggageAllowance {
	if o == nil || utils.IsNil(o.IncludedCheckedBags) {
		var ret BaggageAllowance
		return ret
	}
	return *o.IncludedCheckedBags
}

// GetIncludedCheckedBagsOk returns a tuple with the IncludedCheckedBags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FareDetailsBySegment) GetIncludedCheckedBagsOk() (*BaggageAllowance, bool) {
	if o == nil || utils.IsNil(o.IncludedCheckedBags) {
		return nil, false
	}
	return o.IncludedCheckedBags, true
}

// HasIncludedCheckedBags returns a boolean if a field has been set.
func (o *FareDetailsBySegment) HasIncludedCheckedBags() bool {
	if o != nil && !utils.IsNil(o.IncludedCheckedBags) {
		return true
	}

	return false
}

// SetIncludedCheckedBags gets a reference to the given BaggageAllowance and assigns it to the IncludedCheckedBags field.
func (o *FareDetailsBySegment) SetIncludedCheckedBags(v BaggageAllowance) {
	o.IncludedCheckedBags = &v
}

// GetAdditionalServices returns the AdditionalServices field value if set, zero value otherwise.
func (o *FareDetailsBySegment) GetAdditionalServices() AdditionalServicesRequest {
	if o == nil || utils.IsNil(o.AdditionalServices) {
		var ret AdditionalServicesRequest
		return ret
	}
	return *o.AdditionalServices
}

// GetAdditionalServicesOk returns a tuple with the AdditionalServices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FareDetailsBySegment) GetAdditionalServicesOk() (*AdditionalServicesRequest, bool) {
	if o == nil || utils.IsNil(o.AdditionalServices) {
		return nil, false
	}
	return o.AdditionalServices, true
}

// HasAdditionalServices returns a boolean if a field has been set.
func (o *FareDetailsBySegment) HasAdditionalServices() bool {
	if o != nil && !utils.IsNil(o.AdditionalServices) {
		return true
	}

	return false
}

// SetAdditionalServices gets a reference to the given AdditionalServicesRequest and assigns it to the AdditionalServices field.
func (o *FareDetailsBySegment) SetAdditionalServices(v AdditionalServicesRequest) {
	o.AdditionalServices = &v
}

func (o FareDetailsBySegment) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FareDetailsBySegment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["segmentId"] = o.SegmentId
	if !utils.IsNil(o.Cabin) {
		toSerialize["cabin"] = o.Cabin
	}
	if !utils.IsNil(o.FareBasis) {
		toSerialize["fareBasis"] = o.FareBasis
	}
	if !utils.IsNil(o.BrandedFare) {
		toSerialize["brandedFare"] = o.BrandedFare
	}
	if !utils.IsNil(o.Class) {
		toSerialize["class"] = o.Class
	}
	if !utils.IsNil(o.IsAllotment) {
		toSerialize["isAllotment"] = o.IsAllotment
	}
	if !utils.IsNil(o.AllotmentDetails) {
		toSerialize["allotmentDetails"] = o.AllotmentDetails
	}
	if !utils.IsNil(o.SliceDiceIndicator) {
		toSerialize["sliceDiceIndicator"] = o.SliceDiceIndicator
	}
	if !utils.IsNil(o.IncludedCheckedBags) {
		toSerialize["includedCheckedBags"] = o.IncludedCheckedBags
	}
	if !utils.IsNil(o.AdditionalServices) {
		toSerialize["additionalServices"] = o.AdditionalServices
	}
	return toSerialize, nil
}

type NullableFareDetailsBySegment struct {
	value *FareDetailsBySegment
	isSet bool
}

func (v NullableFareDetailsBySegment) Get() *FareDetailsBySegment {
	return v.value
}

func (v *NullableFareDetailsBySegment) Set(val *FareDetailsBySegment) {
	v.value = val
	v.isSet = true
}

func (v NullableFareDetailsBySegment) IsSet() bool {
	return v.isSet
}

func (v *NullableFareDetailsBySegment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFareDetailsBySegment(val *FareDetailsBySegment) *NullableFareDetailsBySegment {
	return &NullableFareDetailsBySegment{value: val, isSet: true}
}

func (v NullableFareDetailsBySegment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFareDetailsBySegment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
