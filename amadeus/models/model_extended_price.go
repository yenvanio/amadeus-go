package models

import (
	"encoding/json"
	"github.com/yenvanio/amadeus-go/amadeus/utils"
)

// checks if the ExtendedPrice type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ExtendedPrice{}

// ExtendedPrice price information
type ExtendedPrice struct {
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
	// BOOK step ONLY - The price margin percentage (plus or minus) that the booking can tolerate. When set to 0, then no price magin is tolerated.
	Margin *string `json:"margin,omitempty"`
	// Total amount paid by the user (including fees and selected additional services).
	GrandTotal *string `json:"grandTotal,omitempty"`
	// Currency of the payment. It may be different than the requested currency
	BillingCurrency    *string             `json:"billingCurrency,omitempty"`
	AdditionalServices []AdditionalService `json:"additionalServices,omitempty"`
}

// NewExtendedPrice instantiates a new ExtendedPrice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExtendedPrice() *ExtendedPrice {
	this := ExtendedPrice{}
	return &this
}

// NewExtendedPriceWithDefaults instantiates a new ExtendedPrice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExtendedPriceWithDefaults() *ExtendedPrice {
	this := ExtendedPrice{}
	return &this
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *ExtendedPrice) GetCurrency() string {
	if o == nil || utils.IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtendedPrice) GetCurrencyOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *ExtendedPrice) HasCurrency() bool {
	if o != nil && !utils.IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *ExtendedPrice) SetCurrency(v string) {
	o.Currency = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *ExtendedPrice) GetTotal() string {
	if o == nil || utils.IsNil(o.Total) {
		var ret string
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtendedPrice) GetTotalOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *ExtendedPrice) HasTotal() bool {
	if o != nil && !utils.IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given string and assigns it to the Total field.
func (o *ExtendedPrice) SetTotal(v string) {
	o.Total = &v
}

// GetBase returns the Base field value if set, zero value otherwise.
func (o *ExtendedPrice) GetBase() string {
	if o == nil || utils.IsNil(o.Base) {
		var ret string
		return ret
	}
	return *o.Base
}

// GetBaseOk returns a tuple with the Base field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtendedPrice) GetBaseOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Base) {
		return nil, false
	}
	return o.Base, true
}

// HasBase returns a boolean if a field has been set.
func (o *ExtendedPrice) HasBase() bool {
	if o != nil && !utils.IsNil(o.Base) {
		return true
	}

	return false
}

// SetBase gets a reference to the given string and assigns it to the Base field.
func (o *ExtendedPrice) SetBase(v string) {
	o.Base = &v
}

// GetFees returns the Fees field value if set, zero value otherwise.
func (o *ExtendedPrice) GetFees() []Fee {
	if o == nil || utils.IsNil(o.Fees) {
		var ret []Fee
		return ret
	}
	return o.Fees
}

// GetFeesOk returns a tuple with the Fees field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtendedPrice) GetFeesOk() ([]Fee, bool) {
	if o == nil || utils.IsNil(o.Fees) {
		return nil, false
	}
	return o.Fees, true
}

// HasFees returns a boolean if a field has been set.
func (o *ExtendedPrice) HasFees() bool {
	if o != nil && !utils.IsNil(o.Fees) {
		return true
	}

	return false
}

// SetFees gets a reference to the given []Fee and assigns it to the Fees field.
func (o *ExtendedPrice) SetFees(v []Fee) {
	o.Fees = v
}

// GetTaxes returns the Taxes field value if set, zero value otherwise.
func (o *ExtendedPrice) GetTaxes() []Tax {
	if o == nil || utils.IsNil(o.Taxes) {
		var ret []Tax
		return ret
	}
	return o.Taxes
}

// GetTaxesOk returns a tuple with the Taxes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtendedPrice) GetTaxesOk() ([]Tax, bool) {
	if o == nil || utils.IsNil(o.Taxes) {
		return nil, false
	}
	return o.Taxes, true
}

// HasTaxes returns a boolean if a field has been set.
func (o *ExtendedPrice) HasTaxes() bool {
	if o != nil && !utils.IsNil(o.Taxes) {
		return true
	}

	return false
}

// SetTaxes gets a reference to the given []Tax and assigns it to the Taxes field.
func (o *ExtendedPrice) SetTaxes(v []Tax) {
	o.Taxes = v
}

// GetRefundableTaxes returns the RefundableTaxes field value if set, zero value otherwise.
func (o *ExtendedPrice) GetRefundableTaxes() string {
	if o == nil || utils.IsNil(o.RefundableTaxes) {
		var ret string
		return ret
	}
	return *o.RefundableTaxes
}

// GetRefundableTaxesOk returns a tuple with the RefundableTaxes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtendedPrice) GetRefundableTaxesOk() (*string, bool) {
	if o == nil || utils.IsNil(o.RefundableTaxes) {
		return nil, false
	}
	return o.RefundableTaxes, true
}

// HasRefundableTaxes returns a boolean if a field has been set.
func (o *ExtendedPrice) HasRefundableTaxes() bool {
	if o != nil && !utils.IsNil(o.RefundableTaxes) {
		return true
	}

	return false
}

// SetRefundableTaxes gets a reference to the given string and assigns it to the RefundableTaxes field.
func (o *ExtendedPrice) SetRefundableTaxes(v string) {
	o.RefundableTaxes = &v
}

// GetMargin returns the Margin field value if set, zero value otherwise.
func (o *ExtendedPrice) GetMargin() string {
	if o == nil || utils.IsNil(o.Margin) {
		var ret string
		return ret
	}
	return *o.Margin
}

// GetMarginOk returns a tuple with the Margin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtendedPrice) GetMarginOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Margin) {
		return nil, false
	}
	return o.Margin, true
}

// HasMargin returns a boolean if a field has been set.
func (o *ExtendedPrice) HasMargin() bool {
	if o != nil && !utils.IsNil(o.Margin) {
		return true
	}

	return false
}

// SetMargin gets a reference to the given string and assigns it to the Margin field.
func (o *ExtendedPrice) SetMargin(v string) {
	o.Margin = &v
}

// GetGrandTotal returns the GrandTotal field value if set, zero value otherwise.
func (o *ExtendedPrice) GetGrandTotal() string {
	if o == nil || utils.IsNil(o.GrandTotal) {
		var ret string
		return ret
	}
	return *o.GrandTotal
}

// GetGrandTotalOk returns a tuple with the GrandTotal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtendedPrice) GetGrandTotalOk() (*string, bool) {
	if o == nil || utils.IsNil(o.GrandTotal) {
		return nil, false
	}
	return o.GrandTotal, true
}

// HasGrandTotal returns a boolean if a field has been set.
func (o *ExtendedPrice) HasGrandTotal() bool {
	if o != nil && !utils.IsNil(o.GrandTotal) {
		return true
	}

	return false
}

// SetGrandTotal gets a reference to the given string and assigns it to the GrandTotal field.
func (o *ExtendedPrice) SetGrandTotal(v string) {
	o.GrandTotal = &v
}

// GetBillingCurrency returns the BillingCurrency field value if set, zero value otherwise.
func (o *ExtendedPrice) GetBillingCurrency() string {
	if o == nil || utils.IsNil(o.BillingCurrency) {
		var ret string
		return ret
	}
	return *o.BillingCurrency
}

// GetBillingCurrencyOk returns a tuple with the BillingCurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtendedPrice) GetBillingCurrencyOk() (*string, bool) {
	if o == nil || utils.IsNil(o.BillingCurrency) {
		return nil, false
	}
	return o.BillingCurrency, true
}

// HasBillingCurrency returns a boolean if a field has been set.
func (o *ExtendedPrice) HasBillingCurrency() bool {
	if o != nil && !utils.IsNil(o.BillingCurrency) {
		return true
	}

	return false
}

// SetBillingCurrency gets a reference to the given string and assigns it to the BillingCurrency field.
func (o *ExtendedPrice) SetBillingCurrency(v string) {
	o.BillingCurrency = &v
}

// GetAdditionalServices returns the AdditionalServices field value if set, zero value otherwise.
func (o *ExtendedPrice) GetAdditionalServices() []AdditionalService {
	if o == nil || utils.IsNil(o.AdditionalServices) {
		var ret []AdditionalService
		return ret
	}
	return o.AdditionalServices
}

// GetAdditionalServicesOk returns a tuple with the AdditionalServices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtendedPrice) GetAdditionalServicesOk() ([]AdditionalService, bool) {
	if o == nil || utils.IsNil(o.AdditionalServices) {
		return nil, false
	}
	return o.AdditionalServices, true
}

// HasAdditionalServices returns a boolean if a field has been set.
func (o *ExtendedPrice) HasAdditionalServices() bool {
	if o != nil && !utils.IsNil(o.AdditionalServices) {
		return true
	}

	return false
}

// SetAdditionalServices gets a reference to the given []AdditionalService and assigns it to the AdditionalServices field.
func (o *ExtendedPrice) SetAdditionalServices(v []AdditionalService) {
	o.AdditionalServices = v
}

func (o ExtendedPrice) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExtendedPrice) ToMap() (map[string]interface{}, error) {
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
	if !utils.IsNil(o.Margin) {
		toSerialize["margin"] = o.Margin
	}
	if !utils.IsNil(o.GrandTotal) {
		toSerialize["grandTotal"] = o.GrandTotal
	}
	if !utils.IsNil(o.BillingCurrency) {
		toSerialize["billingCurrency"] = o.BillingCurrency
	}
	if !utils.IsNil(o.AdditionalServices) {
		toSerialize["additionalServices"] = o.AdditionalServices
	}
	return toSerialize, nil
}

type NullableExtendedPrice struct {
	value *ExtendedPrice
	isSet bool
}

func (v NullableExtendedPrice) Get() *ExtendedPrice {
	return v.value
}

func (v *NullableExtendedPrice) Set(val *ExtendedPrice) {
	v.value = val
	v.isSet = true
}

func (v NullableExtendedPrice) IsSet() bool {
	return v.isSet
}

func (v *NullableExtendedPrice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExtendedPrice(val *ExtendedPrice) *NullableExtendedPrice {
	return &NullableExtendedPrice{value: val, isSet: true}
}

func (v NullableExtendedPrice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExtendedPrice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
