# Price

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | Pointer to **string** |  | [optional] 
**Total** | Pointer to **string** | Total amount paid by the user | [optional] 
**Base** | Pointer to **string** | Amount without taxes | [optional] 
**Fees** | Pointer to [**[]Fee**](Fee.md) | List of applicable fees | [optional] 
**Taxes** | Pointer to [**[]Tax**](Tax.md) |  | [optional] 
**RefundableTaxes** | Pointer to **string** | The amount of taxes which are refundable | [optional] 

## Methods

### NewPrice

`func NewPrice() *Price`

NewPrice instantiates a new Price object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPriceWithDefaults

`func NewPriceWithDefaults() *Price`

NewPriceWithDefaults instantiates a new Price object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrency

`func (o *Price) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Price) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Price) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *Price) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetTotal

`func (o *Price) GetTotal() string`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *Price) GetTotalOk() (*string, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *Price) SetTotal(v string)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *Price) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetBase

`func (o *Price) GetBase() string`

GetBase returns the Base field if non-nil, zero value otherwise.

### GetBaseOk

`func (o *Price) GetBaseOk() (*string, bool)`

GetBaseOk returns a tuple with the Base field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBase

`func (o *Price) SetBase(v string)`

SetBase sets Base field to given value.

### HasBase

`func (o *Price) HasBase() bool`

HasBase returns a boolean if a field has been set.

### GetFees

`func (o *Price) GetFees() []Fee`

GetFees returns the Fees field if non-nil, zero value otherwise.

### GetFeesOk

`func (o *Price) GetFeesOk() (*[]Fee, bool)`

GetFeesOk returns a tuple with the Fees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFees

`func (o *Price) SetFees(v []Fee)`

SetFees sets Fees field to given value.

### HasFees

`func (o *Price) HasFees() bool`

HasFees returns a boolean if a field has been set.

### GetTaxes

`func (o *Price) GetTaxes() []Tax`

GetTaxes returns the Taxes field if non-nil, zero value otherwise.

### GetTaxesOk

`func (o *Price) GetTaxesOk() (*[]Tax, bool)`

GetTaxesOk returns a tuple with the Taxes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxes

`func (o *Price) SetTaxes(v []Tax)`

SetTaxes sets Taxes field to given value.

### HasTaxes

`func (o *Price) HasTaxes() bool`

HasTaxes returns a boolean if a field has been set.

### GetRefundableTaxes

`func (o *Price) GetRefundableTaxes() string`

GetRefundableTaxes returns the RefundableTaxes field if non-nil, zero value otherwise.

### GetRefundableTaxesOk

`func (o *Price) GetRefundableTaxesOk() (*string, bool)`

GetRefundableTaxesOk returns a tuple with the RefundableTaxes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundableTaxes

`func (o *Price) SetRefundableTaxes(v string)`

SetRefundableTaxes sets RefundableTaxes field to given value.

### HasRefundableTaxes

`func (o *Price) HasRefundableTaxes() bool`

HasRefundableTaxes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


