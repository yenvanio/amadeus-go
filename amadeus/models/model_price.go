package models

import (
	"encoding/json"
	"github.com/yenvanio/amadeus-go/amadeus/utils"
)

// checks if the Price type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Price{}

// Price struct for Price
type Price struct {
	Currency *string `json:"currency,omitempty"`
	// Total amount paid by the user
	Total *string `json:"total,omitempty"`
	// Amount without taxes
	Base *string `json:"base,omitempty"`
	// List of applicable fees
	Fees  []Fee `json:"fees,omitempty"`
	Taxes []Tax `json:"taxes,omitempty"`
	// The amount of taxes which are refundable
	RefundableTaxes *string `json:"refundableTaxes,omitempty"`
}

// NewPrice instantiates a new Price object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrice() *Price {
	this := Price{}
	return &this
}

// NewPriceWithDefaults instantiates a new Price object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPriceWithDefaults() *Price {
	this := Price{}
	return &this
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *Price) GetCurrency() string {
	if o == nil || utils.IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Price) GetCurrencyOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *Price) HasCurrency() bool {
	if o != nil && !utils.IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *Price) SetCurrency(v string) {
	o.Currency = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *Price) GetTotal() string {
	if o == nil || utils.IsNil(o.Total) {
		var ret string
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Price) GetTotalOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *Price) HasTotal() bool {
	if o != nil && !utils.IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given string and assigns it to the Total field.
func (o *Price) SetTotal(v string) {
	o.Total = &v
}

// GetBase returns the Base field value if set, zero value otherwise.
func (o *Price) GetBase() string {
	if o == nil || utils.IsNil(o.Base) {
		var ret string
		return ret
	}
	return *o.Base
}

// GetBaseOk returns a tuple with the Base field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Price) GetBaseOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Base) {
		return nil, false
	}
	return o.Base, true
}

// HasBase returns a boolean if a field has been set.
func (o *Price) HasBase() bool {
	if o != nil && !utils.IsNil(o.Base) {
		return true
	}

	return false
}

// SetBase gets a reference to the given string and assigns it to the Base field.
func (o *Price) SetBase(v string) {
	o.Base = &v
}

// GetFees returns the Fees field value if set, zero value otherwise.
func (o *Price) GetFees() []Fee {
	if o == nil || utils.IsNil(o.Fees) {
		var ret []Fee
		return ret
	}
	return o.Fees
}

// GetFeesOk returns a tuple with the Fees field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Price) GetFeesOk() ([]Fee, bool) {
	if o == nil || utils.IsNil(o.Fees) {
		return nil, false
	}
	return o.Fees, true
}

// HasFees returns a boolean if a field has been set.
func (o *Price) HasFees() bool {
	if o != nil && !utils.IsNil(o.Fees) {
		return true
	}

	return false
}

// SetFees gets a reference to the given []Fee and assigns it to the Fees field.
func (o *Price) SetFees(v []Fee) {
	o.Fees = v
}

// GetTaxes returns the Taxes field value if set, zero value otherwise.
func (o *Price) GetTaxes() []Tax {
	if o == nil || utils.IsNil(o.Taxes) {
		var ret []Tax
		return ret
	}
	return o.Taxes
}

// GetTaxesOk returns a tuple with the Taxes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Price) GetTaxesOk() ([]Tax, bool) {
	if o == nil || utils.IsNil(o.Taxes) {
		return nil, false
	}
	return o.Taxes, true
}

// HasTaxes returns a boolean if a field has been set.
func (o *Price) HasTaxes() bool {
	if o != nil && !utils.IsNil(o.Taxes) {
		return true
	}

	return false
}

// SetTaxes gets a reference to the given []Tax and assigns it to the Taxes field.
func (o *Price) SetTaxes(v []Tax) {
	o.Taxes = v
}

// GetRefundableTaxes returns the RefundableTaxes field value if set, zero value otherwise.
func (o *Price) GetRefundableTaxes() string {
	if o == nil || utils.IsNil(o.RefundableTaxes) {
		var ret string
		return ret
	}
	return *o.RefundableTaxes
}

// GetRefundableTaxesOk returns a tuple with the RefundableTaxes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Price) GetRefundableTaxesOk() (*string, bool) {
	if o == nil || utils.IsNil(o.RefundableTaxes) {
		return nil, false
	}
	return o.RefundableTaxes, true
}

// HasRefundableTaxes returns a boolean if a field has been set.
func (o *Price) HasRefundableTaxes() bool {
	if o != nil && !utils.IsNil(o.RefundableTaxes) {
		return true
	}

	return false
}

// SetRefundableTaxes gets a reference to the given string and assigns it to the RefundableTaxes field.
func (o *Price) SetRefundableTaxes(v string) {
	o.RefundableTaxes = &v
}

func (o Price) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Price) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if !utils.IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	if !utils.IsNil(o.Base) {
		toSerialize["base"] = o.Base
	}
	if !utils.IsNil(o.Fees) {
		toSerialize["fees"] = o.Fees
	}
	if !utils.IsNil(o.Taxes) {
		toSerialize["taxes"] = o.Taxes
	}
	if !utils.IsNil(o.RefundableTaxes) {
		toSerialize["refundableTaxes"] = o.RefundableTaxes
	}
	return toSerialize, nil
}

type NullablePrice struct {
	value *Price
	isSet bool
}

func (v NullablePrice) Get() *Price {
	return v.value
}

func (v *NullablePrice) Set(val *Price) {
	v.value = val
	v.isSet = true
}

func (v NullablePrice) IsSet() bool {
	return v.isSet
}

func (v *NullablePrice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrice(val *Price) *NullablePrice {
	return &NullablePrice{value: val, isSet: true}
}

func (v NullablePrice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
