# FlightOffer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | the resource name | 
**Id** | **string** | Id of the flight offer | 
**Source** | Pointer to [**FlightOfferSource**](FlightOfferSource.md) |  | [optional] 
**InstantTicketingRequired** | Pointer to **bool** | If true, inform that a ticketing will be required at booking step. | [optional] 
**DisablePricing** | Pointer to **bool** | BOOK step ONLY - If true, allows to book a PNR without pricing. Only for the source \&quot;GDS\&quot; | [optional] 
**NonHomogeneous** | Pointer to **bool** | If true, upon completion of the booking, this pricing solution is expected to yield multiple records (a record contains booking information confirmed and stored, typically a Passenger Name Record (PNR), in the provider GDS or system) | [optional] 
**OneWay** | Pointer to **bool** | If true, the flight offer can be combined with other oneWays flight-offers to complete the whole journey (one-Way combinable feature). | [optional] 
**PaymentCardRequired** | Pointer to **bool** | If true, a payment card is mandatory to book this flight offer | [optional] 
**LastTicketingDate** | Pointer to **string** | If booked on the same day as the search (with respect to timezone), this flight offer is guaranteed to be thereafter valid for ticketing until this date (included). Unspecified when it does not make sense for this flight offer (e.g. no control over ticketing once booked). YYYY-MM-DD format, e.g. 2019-06-07 | [optional] 
**LastTicketingDateTime** | Pointer to **time.Time** | If booked on the same day as the search (with respect to timezone), this flight offer is guaranteed to be thereafter valid for ticketing until this date/time (included). Unspecified when it does not make sense for this flight offer (e.g. no control over ticketing once booked). Information only this attribute is not used in input of pricing request. Local date and time in YYYY-MM-ddThh:mm:ss format, e.g. 2017-02-10T20:40:00 | [optional] 
**NumberOfBookableSeats** | Pointer to **float32** | Number of seats bookable in a single request. Can not be higher than 9. | [optional] 
**Itineraries** | Pointer to [**[]Itineraries**](Itineraries.md) |  | [optional] 
**Price** | Pointer to [**ExtendedPrice**](ExtendedPrice.md) |  | [optional] 
**PricingOptions** | Pointer to [**PricingOptions**](PricingOptions.md) |  | [optional] 
**ValidatingAirlineCodes** | Pointer to **[]string** | This option ensures that the system will only consider offers with these airlines as validating carrier. | [optional] 
**TravelerPricings** | Pointer to [**[]TravelerPricing**](TravelerPricing.md) | Fare information for each traveler/segment | [optional] 

## Methods

### NewFlightOffer

`func NewFlightOffer(type_ string, id string, ) *FlightOffer`

NewFlightOffer instantiates a new FlightOffer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlightOfferWithDefaults

`func NewFlightOfferWithDefaults() *FlightOffer`

NewFlightOfferWithDefaults instantiates a new FlightOffer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *FlightOffer) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FlightOffer) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FlightOffer) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *FlightOffer) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FlightOffer) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FlightOffer) SetId(v string)`

SetId sets Id field to given value.


### GetSource

`func (o *FlightOffer) GetSource() FlightOfferSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *FlightOffer) GetSourceOk() (*FlightOfferSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *FlightOffer) SetSource(v FlightOfferSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *FlightOffer) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetInstantTicketingRequired

`func (o *FlightOffer) GetInstantTicketingRequired() bool`

GetInstantTicketingRequired returns the InstantTicketingRequired field if non-nil, zero value otherwise.

### GetInstantTicketingRequiredOk

`func (o *FlightOffer) GetInstantTicketingRequiredOk() (*bool, bool)`

GetInstantTicketingRequiredOk returns a tuple with the InstantTicketingRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstantTicketingRequired

`func (o *FlightOffer) SetInstantTicketingRequired(v bool)`

SetInstantTicketingRequired sets InstantTicketingRequired field to given value.

### HasInstantTicketingRequired

`func (o *FlightOffer) HasInstantTicketingRequired() bool`

HasInstantTicketingRequired returns a boolean if a field has been set.

### GetDisablePricing

`func (o *FlightOffer) GetDisablePricing() bool`

GetDisablePricing returns the DisablePricing field if non-nil, zero value otherwise.

### GetDisablePricingOk

`func (o *FlightOffer) GetDisablePricingOk() (*bool, bool)`

GetDisablePricingOk returns a tuple with the DisablePricing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisablePricing

`func (o *FlightOffer) SetDisablePricing(v bool)`

SetDisablePricing sets DisablePricing field to given value.

### HasDisablePricing

`func (o *FlightOffer) HasDisablePricing() bool`

HasDisablePricing returns a boolean if a field has been set.

### GetNonHomogeneous

`func (o *FlightOffer) GetNonHomogeneous() bool`

GetNonHomogeneous returns the NonHomogeneous field if non-nil, zero value otherwise.

### GetNonHomogeneousOk

`func (o *FlightOffer) GetNonHomogeneousOk() (*bool, bool)`

GetNonHomogeneousOk returns a tuple with the NonHomogeneous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonHomogeneous

`func (o *FlightOffer) SetNonHomogeneous(v bool)`

SetNonHomogeneous sets NonHomogeneous field to given value.

### HasNonHomogeneous

`func (o *FlightOffer) HasNonHomogeneous() bool`

HasNonHomogeneous returns a boolean if a field has been set.

### GetOneWay

`func (o *FlightOffer) GetOneWay() bool`

GetOneWay returns the OneWay field if non-nil, zero value otherwise.

### GetOneWayOk

`func (o *FlightOffer) GetOneWayOk() (*bool, bool)`

GetOneWayOk returns a tuple with the OneWay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOneWay

`func (o *FlightOffer) SetOneWay(v bool)`

SetOneWay sets OneWay field to given value.

### HasOneWay

`func (o *FlightOffer) HasOneWay() bool`

HasOneWay returns a boolean if a field has been set.

### GetPaymentCardRequired

`func (o *FlightOffer) GetPaymentCardRequired() bool`

GetPaymentCardRequired returns the PaymentCardRequired field if non-nil, zero value otherwise.

### GetPaymentCardRequiredOk

`func (o *FlightOffer) GetPaymentCardRequiredOk() (*bool, bool)`

GetPaymentCardRequiredOk returns a tuple with the PaymentCardRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentCardRequired

`func (o *FlightOffer) SetPaymentCardRequired(v bool)`

SetPaymentCardRequired sets PaymentCardRequired field to given value.

### HasPaymentCardRequired

`func (o *FlightOffer) HasPaymentCardRequired() bool`

HasPaymentCardRequired returns a boolean if a field has been set.

### GetLastTicketingDate

`func (o *FlightOffer) GetLastTicketingDate() string`

GetLastTicketingDate returns the LastTicketingDate field if non-nil, zero value otherwise.

### GetLastTicketingDateOk

`func (o *FlightOffer) GetLastTicketingDateOk() (*string, bool)`

GetLastTicketingDateOk returns a tuple with the LastTicketingDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTicketingDate

`func (o *FlightOffer) SetLastTicketingDate(v string)`

SetLastTicketingDate sets LastTicketingDate field to given value.

### HasLastTicketingDate

`func (o *FlightOffer) HasLastTicketingDate() bool`

HasLastTicketingDate returns a boolean if a field has been set.

### GetLastTicketingDateTime

`func (o *FlightOffer) GetLastTicketingDateTime() time.Time`

GetLastTicketingDateTime returns the LastTicketingDateTime field if non-nil, zero value otherwise.

### GetLastTicketingDateTimeOk

`func (o *FlightOffer) GetLastTicketingDateTimeOk() (*time.Time, bool)`

GetLastTicketingDateTimeOk returns a tuple with the LastTicketingDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTicketingDateTime

`func (o *FlightOffer) SetLastTicketingDateTime(v time.Time)`

SetLastTicketingDateTime sets LastTicketingDateTime field to given value.

### HasLastTicketingDateTime

`func (o *FlightOffer) HasLastTicketingDateTime() bool`

HasLastTicketingDateTime returns a boolean if a field has been set.

### GetNumberOfBookableSeats

`func (o *FlightOffer) GetNumberOfBookableSeats() float32`

GetNumberOfBookableSeats returns the NumberOfBookableSeats field if non-nil, zero value otherwise.

### GetNumberOfBookableSeatsOk

`func (o *FlightOffer) GetNumberOfBookableSeatsOk() (*float32, bool)`

GetNumberOfBookableSeatsOk returns a tuple with the NumberOfBookableSeats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfBookableSeats

`func (o *FlightOffer) SetNumberOfBookableSeats(v float32)`

SetNumberOfBookableSeats sets NumberOfBookableSeats field to given value.

### HasNumberOfBookableSeats

`func (o *FlightOffer) HasNumberOfBookableSeats() bool`

HasNumberOfBookableSeats returns a boolean if a field has been set.

### GetItineraries

`func (o *FlightOffer) GetItineraries() []Itineraries`

GetItineraries returns the Itineraries field if non-nil, zero value otherwise.

### GetItinerariesOk

`func (o *FlightOffer) GetItinerariesOk() (*[]Itineraries, bool)`

GetItinerariesOk returns a tuple with the Itineraries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItineraries

`func (o *FlightOffer) SetItineraries(v []Itineraries)`

SetItineraries sets Itineraries field to given value.

### HasItineraries

`func (o *FlightOffer) HasItineraries() bool`

HasItineraries returns a boolean if a field has been set.

### GetPrice

`func (o *FlightOffer) GetPrice() ExtendedPrice`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *FlightOffer) GetPriceOk() (*ExtendedPrice, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *FlightOffer) SetPrice(v ExtendedPrice)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *FlightOffer) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetPricingOptions

`func (o *FlightOffer) GetPricingOptions() PricingOptions`

GetPricingOptions returns the PricingOptions field if non-nil, zero value otherwise.

### GetPricingOptionsOk

`func (o *FlightOffer) GetPricingOptionsOk() (*PricingOptions, bool)`

GetPricingOptionsOk returns a tuple with the PricingOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricingOptions

`func (o *FlightOffer) SetPricingOptions(v PricingOptions)`

SetPricingOptions sets PricingOptions field to given value.

### HasPricingOptions

`func (o *FlightOffer) HasPricingOptions() bool`

HasPricingOptions returns a boolean if a field has been set.

### GetValidatingAirlineCodes

`func (o *FlightOffer) GetValidatingAirlineCodes() []string`

GetValidatingAirlineCodes returns the ValidatingAirlineCodes field if non-nil, zero value otherwise.

### GetValidatingAirlineCodesOk

`func (o *FlightOffer) GetValidatingAirlineCodesOk() (*[]string, bool)`

GetValidatingAirlineCodesOk returns a tuple with the ValidatingAirlineCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidatingAirlineCodes

`func (o *FlightOffer) SetValidatingAirlineCodes(v []string)`

SetValidatingAirlineCodes sets ValidatingAirlineCodes field to given value.

### HasValidatingAirlineCodes

`func (o *FlightOffer) HasValidatingAirlineCodes() bool`

HasValidatingAirlineCodes returns a boolean if a field has been set.

### GetTravelerPricings

`func (o *FlightOffer) GetTravelerPricings() []TravelerPricing`

GetTravelerPricings returns the TravelerPricings field if non-nil, zero value otherwise.

### GetTravelerPricingsOk

`func (o *FlightOffer) GetTravelerPricingsOk() (*[]TravelerPricing, bool)`

GetTravelerPricingsOk returns a tuple with the TravelerPricings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTravelerPricings

`func (o *FlightOffer) SetTravelerPricings(v []TravelerPricing)`

SetTravelerPricings sets TravelerPricings field to given value.

### HasTravelerPricings

`func (o *FlightOffer) HasTravelerPricings() bool`

HasTravelerPricings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


