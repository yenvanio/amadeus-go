# CreditCardFee

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Brand** | Pointer to [**PaymentBrand**](PaymentBrand.md) |  | [optional] 
**Amount** | Pointer to **string** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**FlightOfferId** | Pointer to **string** | Id of the flightOffer concerned by the credit card fee | [optional] 

## Methods

### NewCreditCardFee

`func NewCreditCardFee() *CreditCardFee`

NewCreditCardFee instantiates a new CreditCardFee object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreditCardFeeWithDefaults

`func NewCreditCardFeeWithDefaults() *CreditCardFee`

NewCreditCardFeeWithDefaults instantiates a new CreditCardFee object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBrand

`func (o *CreditCardFee) GetBrand() PaymentBrand`

GetBrand returns the Brand field if non-nil, zero value otherwise.

### GetBrandOk

`func (o *CreditCardFee) GetBrandOk() (*PaymentBrand, bool)`

GetBrandOk returns a tuple with the Brand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrand

`func (o *CreditCardFee) SetBrand(v PaymentBrand)`

SetBrand sets Brand field to given value.

### HasBrand

`func (o *CreditCardFee) HasBrand() bool`

HasBrand returns a boolean if a field has been set.

### GetAmount

`func (o *CreditCardFee) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CreditCardFee) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CreditCardFee) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *CreditCardFee) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetCurrency

`func (o *CreditCardFee) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CreditCardFee) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CreditCardFee) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *CreditCardFee) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetFlightOfferId

`func (o *CreditCardFee) GetFlightOfferId() string`

GetFlightOfferId returns the FlightOfferId field if non-nil, zero value otherwise.

### GetFlightOfferIdOk

`func (o *CreditCardFee) GetFlightOfferIdOk() (*string, bool)`

GetFlightOfferIdOk returns a tuple with the FlightOfferId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlightOfferId

`func (o *CreditCardFee) SetFlightOfferId(v string)`

SetFlightOfferId sets FlightOfferId field to given value.

### HasFlightOfferId

`func (o *CreditCardFee) HasFlightOfferId() bool`

HasFlightOfferId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


