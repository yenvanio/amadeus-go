# FlightOfferPricingOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | the resource name | 
**FlightOffers** | [**[]FlightOffer**](FlightOffer.md) | list of flight offer to price | 
**BookingRequirements** | Pointer to [**BookingRequirements**](BookingRequirements.md) |  | [optional] 

## Methods

### NewFlightOfferPricingOut

`func NewFlightOfferPricingOut(type_ string, flightOffers []FlightOffer, ) *FlightOfferPricingOut`

NewFlightOfferPricingOut instantiates a new FlightOfferPricingOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlightOfferPricingOutWithDefaults

`func NewFlightOfferPricingOutWithDefaults() *FlightOfferPricingOut`

NewFlightOfferPricingOutWithDefaults instantiates a new FlightOfferPricingOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *FlightOfferPricingOut) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FlightOfferPricingOut) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FlightOfferPricingOut) SetType(v string)`

SetType sets Type field to given value.


### GetFlightOffers

`func (o *FlightOfferPricingOut) GetFlightOffers() []FlightOffer`

GetFlightOffers returns the FlightOffers field if non-nil, zero value otherwise.

### GetFlightOffersOk

`func (o *FlightOfferPricingOut) GetFlightOffersOk() (*[]FlightOffer, bool)`

GetFlightOffersOk returns a tuple with the FlightOffers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlightOffers

`func (o *FlightOfferPricingOut) SetFlightOffers(v []FlightOffer)`

SetFlightOffers sets FlightOffers field to given value.


### GetBookingRequirements

`func (o *FlightOfferPricingOut) GetBookingRequirements() BookingRequirements`

GetBookingRequirements returns the BookingRequirements field if non-nil, zero value otherwise.

### GetBookingRequirementsOk

`func (o *FlightOfferPricingOut) GetBookingRequirementsOk() (*BookingRequirements, bool)`

GetBookingRequirementsOk returns a tuple with the BookingRequirements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBookingRequirements

`func (o *FlightOfferPricingOut) SetBookingRequirements(v BookingRequirements)`

SetBookingRequirements sets BookingRequirements field to given value.

### HasBookingRequirements

`func (o *FlightOfferPricingOut) HasBookingRequirements() bool`

HasBookingRequirements returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


