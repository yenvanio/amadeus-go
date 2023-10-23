# ExtendedPrice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | Pointer to **string** |  | [optional] 
**Total** | Pointer to **string** | Total amount paid by the user | [optional] 
**Base** | Pointer to **string** | Amount without taxes | [optional] 
**Fees** | Pointer to [**[]Fee**](Fee.md) | List of applicable fees | [optional] 
**Taxes** | Pointer to [**[]Tax**](Tax.md) |  | [optional] 
**RefundableTaxes** | Pointer to **string** | The amount of taxes which are refundable | [optional] 
**Margin** | Pointer to **string** | BOOK step ONLY - The price margin percentage (plus or minus) that the booking can tolerate. When set to 0, then no price magin is tolerated. | [optional] 
**GrandTotal** | Pointer to **string** | Total amount paid by the user (including fees and selected additional services). | [optional] 
**BillingCurrency** | Pointer to **string** | Currency of the payment. It may be different than the requested currency | [optional] 
**AdditionalServices** | Pointer to [**[]AdditionalService**](AdditionalService.md) |  | [optional] 

## Methods

### NewExtendedPrice

`func NewExtendedPrice() *ExtendedPrice`

NewExtendedPrice instantiates a new ExtendedPrice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtendedPriceWithDefaults

`func NewExtendedPriceWithDefaults() *ExtendedPrice`

NewExtendedPriceWithDefaults instantiates a new ExtendedPrice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrency

`func (o *ExtendedPrice) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *ExtendedPrice) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *ExtendedPrice) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *ExtendedPrice) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetTotal

`func (o *ExtendedPrice) GetTotal() string`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ExtendedPrice) GetTotalOk() (*string, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ExtendedPrice) SetTotal(v string)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *ExtendedPrice) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetBase

`func (o *ExtendedPrice) GetBase() string`

GetBase returns the Base field if non-nil, zero value otherwise.

### GetBaseOk

`func (o *ExtendedPrice) GetBaseOk() (*string, bool)`

GetBaseOk returns a tuple with the Base field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBase

`func (o *ExtendedPrice) SetBase(v string)`

SetBase sets Base field to given value.

### HasBase

`func (o *ExtendedPrice) HasBase() bool`

HasBase returns a boolean if a field has been set.

### GetFees

`func (o *ExtendedPrice) GetFees() []Fee`

GetFees returns the Fees field if non-nil, zero value otherwise.

### GetFeesOk

`func (o *ExtendedPrice) GetFeesOk() (*[]Fee, bool)`

GetFeesOk returns a tuple with the Fees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFees

`func (o *ExtendedPrice) SetFees(v []Fee)`

SetFees sets Fees field to given value.

### HasFees

`func (o *ExtendedPrice) HasFees() bool`

HasFees returns a boolean if a field has been set.

### GetTaxes

`func (o *ExtendedPrice) GetTaxes() []Tax`

GetTaxes returns the Taxes field if non-nil, zero value otherwise.

### GetTaxesOk

`func (o *ExtendedPrice) GetTaxesOk() (*[]Tax, bool)`

GetTaxesOk returns a tuple with the Taxes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxes

`func (o *ExtendedPrice) SetTaxes(v []Tax)`

SetTaxes sets Taxes field to given value.

### HasTaxes

`func (o *ExtendedPrice) HasTaxes() bool`

HasTaxes returns a boolean if a field has been set.

### GetRefundableTaxes

`func (o *ExtendedPrice) GetRefundableTaxes() string`

GetRefundableTaxes returns the RefundableTaxes field if non-nil, zero value otherwise.

### GetRefundableTaxesOk

`func (o *ExtendedPrice) GetRefundableTaxesOk() (*string, bool)`

GetRefundableTaxesOk returns a tuple with the RefundableTaxes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundableTaxes

`func (o *ExtendedPrice) SetRefundableTaxes(v string)`

SetRefundableTaxes sets RefundableTaxes field to given value.

### HasRefundableTaxes

`func (o *ExtendedPrice) HasRefundableTaxes() bool`

HasRefundableTaxes returns a boolean if a field has been set.

### GetMargin

`func (o *ExtendedPrice) GetMargin() string`

GetMargin returns the Margin field if non-nil, zero value otherwise.

### GetMarginOk

`func (o *ExtendedPrice) GetMarginOk() (*string, bool)`

GetMarginOk returns a tuple with the Margin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMargin

`func (o *ExtendedPrice) SetMargin(v string)`

SetMargin sets Margin field to given value.

### HasMargin

`func (o *ExtendedPrice) HasMargin() bool`

HasMargin returns a boolean if a field has been set.

### GetGrandTotal

`func (o *ExtendedPrice) GetGrandTotal() string`

GetGrandTotal returns the GrandTotal field if non-nil, zero value otherwise.

### GetGrandTotalOk

`func (o *ExtendedPrice) GetGrandTotalOk() (*string, bool)`

GetGrandTotalOk returns a tuple with the GrandTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrandTotal

`func (o *ExtendedPrice) SetGrandTotal(v string)`

SetGrandTotal sets GrandTotal field to given value.

### HasGrandTotal

`func (o *ExtendedPrice) HasGrandTotal() bool`

HasGrandTotal returns a boolean if a field has been set.

### GetBillingCurrency

`func (o *ExtendedPrice) GetBillingCurrency() string`

GetBillingCurrency returns the BillingCurrency field if non-nil, zero value otherwise.

### GetBillingCurrencyOk

`func (o *ExtendedPrice) GetBillingCurrencyOk() (*string, bool)`

GetBillingCurrencyOk returns a tuple with the BillingCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingCurrency

`func (o *ExtendedPrice) SetBillingCurrency(v string)`

SetBillingCurrency sets BillingCurrency field to given value.

### HasBillingCurrency

`func (o *ExtendedPrice) HasBillingCurrency() bool`

HasBillingCurrency returns a boolean if a field has been set.

### GetAdditionalServices

`func (o *ExtendedPrice) GetAdditionalServices() []AdditionalService`

GetAdditionalServices returns the AdditionalServices field if non-nil, zero value otherwise.

### GetAdditionalServicesOk

`func (o *ExtendedPrice) GetAdditionalServicesOk() (*[]AdditionalService, bool)`

GetAdditionalServicesOk returns a tuple with the AdditionalServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalServices

`func (o *ExtendedPrice) SetAdditionalServices(v []AdditionalService)`

SetAdditionalServices sets AdditionalServices field to given value.

### HasAdditionalServices

`func (o *ExtendedPrice) HasAdditionalServices() bool`

HasAdditionalServices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


