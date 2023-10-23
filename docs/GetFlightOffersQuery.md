# GetFlightOffersQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrencyCode** | Pointer to **string** | The currency code, as defined in [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217), to reflect the currency in which this amount is expressed. | [optional] 
**OriginDestinations** | [**[]OriginDestination**](OriginDestination.md) | Origins and Destinations must be properly ordered in time (chronological order in accordance with the timezone of each location) to describe the journey consistently. Dates and times must not be past nor more than 365 days in the future, according to provider settings.Number of Origins and Destinations must not exceed the limit defined in provider settings. | 
**Travelers** | [**[]ExtendedTravelerInfo**](ExtendedTravelerInfo.md) | travelers in the trip.    Maximum number of passengers older than 2 yo (CHILD, ADULT, YOUGHT): 9.  Each adult can travel with one INFANT so maximum total number of passengers: 18 | 
**Sources** | [**[]FlightOfferSource**](FlightOfferSource.md) | Allows enable one or more sources. If present in the list, these sources will be called by the system. | 
**SearchCriteria** | Pointer to [**SearchCriteria**](SearchCriteria.md) |  | [optional] 

## Methods

### NewGetFlightOffersQuery

`func NewGetFlightOffersQuery(originDestinations []OriginDestination, travelers []ExtendedTravelerInfo, sources []FlightOfferSource, ) *GetFlightOffersQuery`

NewGetFlightOffersQuery instantiates a new GetFlightOffersQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetFlightOffersQueryWithDefaults

`func NewGetFlightOffersQueryWithDefaults() *GetFlightOffersQuery`

NewGetFlightOffersQueryWithDefaults instantiates a new GetFlightOffersQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrencyCode

`func (o *GetFlightOffersQuery) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *GetFlightOffersQuery) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *GetFlightOffersQuery) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *GetFlightOffersQuery) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetOriginDestinations

`func (o *GetFlightOffersQuery) GetOriginDestinations() []OriginDestination`

GetOriginDestinations returns the OriginDestinations field if non-nil, zero value otherwise.

### GetOriginDestinationsOk

`func (o *GetFlightOffersQuery) GetOriginDestinationsOk() (*[]OriginDestination, bool)`

GetOriginDestinationsOk returns a tuple with the OriginDestinations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginDestinations

`func (o *GetFlightOffersQuery) SetOriginDestinations(v []OriginDestination)`

SetOriginDestinations sets OriginDestinations field to given value.


### GetTravelers

`func (o *GetFlightOffersQuery) GetTravelers() []ExtendedTravelerInfo`

GetTravelers returns the Travelers field if non-nil, zero value otherwise.

### GetTravelersOk

`func (o *GetFlightOffersQuery) GetTravelersOk() (*[]ExtendedTravelerInfo, bool)`

GetTravelersOk returns a tuple with the Travelers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTravelers

`func (o *GetFlightOffersQuery) SetTravelers(v []ExtendedTravelerInfo)`

SetTravelers sets Travelers field to given value.


### GetSources

`func (o *GetFlightOffersQuery) GetSources() []FlightOfferSource`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *GetFlightOffersQuery) GetSourcesOk() (*[]FlightOfferSource, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *GetFlightOffersQuery) SetSources(v []FlightOfferSource)`

SetSources sets Sources field to given value.


### GetSearchCriteria

`func (o *GetFlightOffersQuery) GetSearchCriteria() SearchCriteria`

GetSearchCriteria returns the SearchCriteria field if non-nil, zero value otherwise.

### GetSearchCriteriaOk

`func (o *GetFlightOffersQuery) GetSearchCriteriaOk() (*SearchCriteria, bool)`

GetSearchCriteriaOk returns a tuple with the SearchCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchCriteria

`func (o *GetFlightOffersQuery) SetSearchCriteria(v SearchCriteria)`

SetSearchCriteria sets SearchCriteria field to given value.

### HasSearchCriteria

`func (o *GetFlightOffersQuery) HasSearchCriteria() bool`

HasSearchCriteria returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


