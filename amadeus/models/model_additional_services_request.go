package models

import (
	"encoding/json"
	"github.com/yenvanio/amadeus-go/amadeus/utils"
)

// checks if the AdditionalServicesRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &AdditionalServicesRequest{}

// AdditionalServicesRequest struct for AdditionalServicesRequest
type AdditionalServicesRequest struct {
	ChargeableCheckedBags *ChargeableCheckdBags `json:"chargeableCheckedBags,omitempty"`
	ChargeableSeat        *ChargeableSeat       `json:"chargeableSeat,omitempty"`
	// DEPRECATED - use the chargeableSeat attribute -  seat number
	ChargeableSeatNumber *string `json:"chargeableSeatNumber,omitempty"`
	// Other services to add
	OtherServices []ServiceName `json:"otherServices,omitempty"`
}

// NewAdditionalServicesRequest instantiates a new AdditionalServicesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdditionalServicesRequest() *AdditionalServicesRequest {
	this := AdditionalServicesRequest{}
	return &this
}

// NewAdditionalServicesRequestWithDefaults instantiates a new AdditionalServicesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdditionalServicesRequestWithDefaults() *AdditionalServicesRequest {
	this := AdditionalServicesRequest{}
	return &this
}

// GetChargeableCheckedBags returns the ChargeableCheckedBags field value if set, zero value otherwise.
func (o *AdditionalServicesRequest) GetChargeableCheckedBags() ChargeableCheckdBags {
	if o == nil || utils.IsNil(o.ChargeableCheckedBags) {
		var ret ChargeableCheckdBags
		return ret
	}
	return *o.ChargeableCheckedBags
}

// GetChargeableCheckedBagsOk returns a tuple with the ChargeableCheckedBags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalServicesRequest) GetChargeableCheckedBagsOk() (*ChargeableCheckdBags, bool) {
	if o == nil || utils.IsNil(o.ChargeableCheckedBags) {
		return nil, false
	}
	return o.ChargeableCheckedBags, true
}

// HasChargeableCheckedBags returns a boolean if a field has been set.
func (o *AdditionalServicesRequest) HasChargeableCheckedBags() bool {
	if o != nil && !utils.IsNil(o.ChargeableCheckedBags) {
		return true
	}

	return false
}

// SetChargeableCheckedBags gets a reference to the given ChargeableCheckdBags and assigns it to the ChargeableCheckedBags field.
func (o *AdditionalServicesRequest) SetChargeableCheckedBags(v ChargeableCheckdBags) {
	o.ChargeableCheckedBags = &v
}

// GetChargeableSeat returns the ChargeableSeat field value if set, zero value otherwise.
func (o *AdditionalServicesRequest) GetChargeableSeat() ChargeableSeat {
	if o == nil || utils.IsNil(o.ChargeableSeat) {
		var ret ChargeableSeat
		return ret
	}
	return *o.ChargeableSeat
}

// GetChargeableSeatOk returns a tuple with the ChargeableSeat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalServicesRequest) GetChargeableSeatOk() (*ChargeableSeat, bool) {
	if o == nil || utils.IsNil(o.ChargeableSeat) {
		return nil, false
	}
	return o.ChargeableSeat, true
}

// HasChargeableSeat returns a boolean if a field has been set.
func (o *AdditionalServicesRequest) HasChargeableSeat() bool {
	if o != nil && !utils.IsNil(o.ChargeableSeat) {
		return true
	}

	return false
}

// SetChargeableSeat gets a reference to the given ChargeableSeat and assigns it to the ChargeableSeat field.
func (o *AdditionalServicesRequest) SetChargeableSeat(v ChargeableSeat) {
	o.ChargeableSeat = &v
}

// GetChargeableSeatNumber returns the ChargeableSeatNumber field value if set, zero value otherwise.
func (o *AdditionalServicesRequest) GetChargeableSeatNumber() string {
	if o == nil || utils.IsNil(o.ChargeableSeatNumber) {
		var ret string
		return ret
	}
	return *o.ChargeableSeatNumber
}

// GetChargeableSeatNumberOk returns a tuple with the ChargeableSeatNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalServicesRequest) GetChargeableSeatNumberOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ChargeableSeatNumber) {
		return nil, false
	}
	return o.ChargeableSeatNumber, true
}

// HasChargeableSeatNumber returns a boolean if a field has been set.
func (o *AdditionalServicesRequest) HasChargeableSeatNumber() bool {
	if o != nil && !utils.IsNil(o.ChargeableSeatNumber) {
		return true
	}

	return false
}

// SetChargeableSeatNumber gets a reference to the given string and assigns it to the ChargeableSeatNumber field.
func (o *AdditionalServicesRequest) SetChargeableSeatNumber(v string) {
	o.ChargeableSeatNumber = &v
}

// GetOtherServices returns the OtherServices field value if set, zero value otherwise.
func (o *AdditionalServicesRequest) GetOtherServices() []ServiceName {
	if o == nil || utils.IsNil(o.OtherServices) {
		var ret []ServiceName
		return ret
	}
	return o.OtherServices
}

// GetOtherServicesOk returns a tuple with the OtherServices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalServicesRequest) GetOtherServicesOk() ([]ServiceName, bool) {
	if o == nil || utils.IsNil(o.OtherServices) {
		return nil, false
	}
	return o.OtherServices, true
}

// HasOtherServices returns a boolean if a field has been set.
func (o *AdditionalServicesRequest) HasOtherServices() bool {
	if o != nil && !utils.IsNil(o.OtherServices) {
		return true
	}

	return false
}

// SetOtherServices gets a reference to the given []ServiceName and assigns it to the OtherServices field.
func (o *AdditionalServicesRequest) SetOtherServices(v []ServiceName) {
	o.OtherServices = v
}

func (o AdditionalServicesRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdditionalServicesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.ChargeableCheckedBags) {
		toSerialize["chargeableCheckedBags"] = o.ChargeableCheckedBags
	}
	if !utils.IsNil(o.ChargeableSeat) {
		toSerialize["chargeableSeat"] = o.ChargeableSeat
	}
	if !utils.IsNil(o.ChargeableSeatNumber) {
		toSerialize["chargeableSeatNumber"] = o.ChargeableSeatNumber
	}
	if !utils.IsNil(o.OtherServices) {
		toSerialize["otherServices"] = o.OtherServices
	}
	return toSerialize, nil
}

type NullableAdditionalServicesRequest struct {
	value *AdditionalServicesRequest
	isSet bool
}

func (v NullableAdditionalServicesRequest) Get() *AdditionalServicesRequest {
	return v.value
}

func (v *NullableAdditionalServicesRequest) Set(val *AdditionalServicesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAdditionalServicesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAdditionalServicesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdditionalServicesRequest(val *AdditionalServicesRequest) *NullableAdditionalServicesRequest {
	return &NullableAdditionalServicesRequest{value: val, isSet: true}
}

func (v NullableAdditionalServicesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdditionalServicesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
