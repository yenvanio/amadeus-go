# FlightOfferPricingIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | the resource name | 
**FlightOffers** | [**[]FlightOffer**](FlightOffer.md) | list of flight offer to price | 
**Payments** | Pointer to [**[]Payment**](Payment.md) | payment information for retrieve eventual credit card fees | [optional] 
**Travelers** | Pointer to [**[]Traveler**](Traveler.md) | list of travelers | [optional] 

## Methods

### NewFlightOfferPricingIn

`func NewFlightOfferPricingIn(type_ string, flightOffers []FlightOffer, ) *FlightOfferPricingIn`

NewFlightOfferPricingIn instantiates a new FlightOfferPricingIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlightOfferPricingInWithDefaults

`func NewFlightOfferPricingInWithDefaults() *FlightOfferPricingIn`

NewFlightOfferPricingInWithDefaults instantiates a new FlightOfferPricingIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *FlightOfferPricingIn) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FlightOfferPricingIn) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FlightOfferPricingIn) SetType(v string)`

SetType sets Type field to given value.


### GetFlightOffers

`func (o *FlightOfferPricingIn) GetFlightOffers() []FlightOffer`

GetFlightOffers returns the FlightOffers field if non-nil, zero value otherwise.

### GetFlightOffersOk

`func (o *FlightOfferPricingIn) GetFlightOffersOk() (*[]FlightOffer, bool)`

GetFlightOffersOk returns a tuple with the FlightOffers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlightOffers

`func (o *FlightOfferPricingIn) SetFlightOffers(v []FlightOffer)`

SetFlightOffers sets FlightOffers field to given value.


### GetPayments

`func (o *FlightOfferPricingIn) GetPayments() []Payment`

GetPayments returns the Payments field if non-nil, zero value otherwise.

### GetPaymentsOk

`func (o *FlightOfferPricingIn) GetPaymentsOk() (*[]Payment, bool)`

GetPaymentsOk returns a tuple with the Payments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayments

`func (o *FlightOfferPricingIn) SetPayments(v []Payment)`

SetPayments sets Payments field to given value.

### HasPayments

`func (o *FlightOfferPricingIn) HasPayments() bool`

HasPayments returns a boolean if a field has been set.

### GetTravelers

`func (o *FlightOfferPricingIn) GetTravelers() []Traveler`

GetTravelers returns the Travelers field if non-nil, zero value otherwise.

### GetTravelersOk

`func (o *FlightOfferPricingIn) GetTravelersOk() (*[]Traveler, bool)`

GetTravelersOk returns a tuple with the Travelers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTravelers

`func (o *FlightOfferPricingIn) SetTravelers(v []Traveler)`

SetTravelers sets Travelers field to given value.

### HasTravelers

`func (o *FlightOfferPricingIn) HasTravelers() bool`

HasTravelers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


