# Payment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Brand** | Pointer to [**PaymentBrand**](PaymentBrand.md) |  | [optional] 
**BinNumber** | Pointer to **int32** | The first 6 digits of the credit card | [optional] 
**FlightOfferIds** | Pointer to **[]string** | Id of the flightOffers to pay | [optional] 

## Methods

### NewPayment

`func NewPayment() *Payment`

NewPayment instantiates a new Payment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentWithDefaults

`func NewPaymentWithDefaults() *Payment`

NewPaymentWithDefaults instantiates a new Payment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBrand

`func (o *Payment) GetBrand() PaymentBrand`

GetBrand returns the Brand field if non-nil, zero value otherwise.

### GetBrandOk

`func (o *Payment) GetBrandOk() (*PaymentBrand, bool)`

GetBrandOk returns a tuple with the Brand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrand

`func (o *Payment) SetBrand(v PaymentBrand)`

SetBrand sets Brand field to given value.

### HasBrand

`func (o *Payment) HasBrand() bool`

HasBrand returns a boolean if a field has been set.

### GetBinNumber

`func (o *Payment) GetBinNumber() int32`

GetBinNumber returns the BinNumber field if non-nil, zero value otherwise.

### GetBinNumberOk

`func (o *Payment) GetBinNumberOk() (*int32, bool)`

GetBinNumberOk returns a tuple with the BinNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinNumber

`func (o *Payment) SetBinNumber(v int32)`

SetBinNumber sets BinNumber field to given value.

### HasBinNumber

`func (o *Payment) HasBinNumber() bool`

HasBinNumber returns a boolean if a field has been set.

### GetFlightOfferIds

`func (o *Payment) GetFlightOfferIds() []string`

GetFlightOfferIds returns the FlightOfferIds field if non-nil, zero value otherwise.

### GetFlightOfferIdsOk

`func (o *Payment) GetFlightOfferIdsOk() (*[]string, bool)`

GetFlightOfferIdsOk returns a tuple with the FlightOfferIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlightOfferIds

`func (o *Payment) SetFlightOfferIds(v []string)`

SetFlightOfferIds sets FlightOfferIds field to given value.

### HasFlightOfferIds

`func (o *Payment) HasFlightOfferIds() bool`

HasFlightOfferIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


