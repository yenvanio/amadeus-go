# SearchCriteria

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExcludeAllotments** | Pointer to **bool** | This option allows to exclude the isAllotment flag associated to a booking class in the search response when it exist. | [optional] 
**AddOneWayOffers** | Pointer to **bool** | This option allows activate the one-way combinable feature | [optional] 
**MaxFlightOffers** | Pointer to **float32** | Maximum number of flight offers returned (Max 250) | [optional] 
**MaxPrice** | Pointer to **int32** | maximum price per traveler. By default, no limit is applied. If specified, the value should be a positive number with no decimals | [optional] 
**AllowAlternativeFareOptions** | Pointer to **bool** | This option allows to default to a standard fareOption if no offers are found for the selected fareOption. | [optional] 
**OneFlightOfferPerDay** | Pointer to **bool** | Requests the system to find at least one flight-offer per day, if possible, when a range of dates is specified. Default is false. | [optional] 
**AdditionalInformation** | Pointer to [**AdditionalInformation**](AdditionalInformation.md) |  | [optional] 
**PricingOptions** | Pointer to [**ExtendedPricingOptions**](ExtendedPricingOptions.md) |  | [optional] 
**FlightFilters** | Pointer to [**FlightFilters**](FlightFilters.md) |  | [optional] 

## Methods

### NewSearchCriteria

`func NewSearchCriteria() *SearchCriteria`

NewSearchCriteria instantiates a new SearchCriteria object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchCriteriaWithDefaults

`func NewSearchCriteriaWithDefaults() *SearchCriteria`

NewSearchCriteriaWithDefaults instantiates a new SearchCriteria object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExcludeAllotments

`func (o *SearchCriteria) GetExcludeAllotments() bool`

GetExcludeAllotments returns the ExcludeAllotments field if non-nil, zero value otherwise.

### GetExcludeAllotmentsOk

`func (o *SearchCriteria) GetExcludeAllotmentsOk() (*bool, bool)`

GetExcludeAllotmentsOk returns a tuple with the ExcludeAllotments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeAllotments

`func (o *SearchCriteria) SetExcludeAllotments(v bool)`

SetExcludeAllotments sets ExcludeAllotments field to given value.

### HasExcludeAllotments

`func (o *SearchCriteria) HasExcludeAllotments() bool`

HasExcludeAllotments returns a boolean if a field has been set.

### GetAddOneWayOffers

`func (o *SearchCriteria) GetAddOneWayOffers() bool`

GetAddOneWayOffers returns the AddOneWayOffers field if non-nil, zero value otherwise.

### GetAddOneWayOffersOk

`func (o *SearchCriteria) GetAddOneWayOffersOk() (*bool, bool)`

GetAddOneWayOffersOk returns a tuple with the AddOneWayOffers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddOneWayOffers

`func (o *SearchCriteria) SetAddOneWayOffers(v bool)`

SetAddOneWayOffers sets AddOneWayOffers field to given value.

### HasAddOneWayOffers

`func (o *SearchCriteria) HasAddOneWayOffers() bool`

HasAddOneWayOffers returns a boolean if a field has been set.

### GetMaxFlightOffers

`func (o *SearchCriteria) GetMaxFlightOffers() float32`

GetMaxFlightOffers returns the MaxFlightOffers field if non-nil, zero value otherwise.

### GetMaxFlightOffersOk

`func (o *SearchCriteria) GetMaxFlightOffersOk() (*float32, bool)`

GetMaxFlightOffersOk returns a tuple with the MaxFlightOffers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxFlightOffers

`func (o *SearchCriteria) SetMaxFlightOffers(v float32)`

SetMaxFlightOffers sets MaxFlightOffers field to given value.

### HasMaxFlightOffers

`func (o *SearchCriteria) HasMaxFlightOffers() bool`

HasMaxFlightOffers returns a boolean if a field has been set.

### GetMaxPrice

`func (o *SearchCriteria) GetMaxPrice() int32`

GetMaxPrice returns the MaxPrice field if non-nil, zero value otherwise.

### GetMaxPriceOk

`func (o *SearchCriteria) GetMaxPriceOk() (*int32, bool)`

GetMaxPriceOk returns a tuple with the MaxPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPrice

`func (o *SearchCriteria) SetMaxPrice(v int32)`

SetMaxPrice sets MaxPrice field to given value.

### HasMaxPrice

`func (o *SearchCriteria) HasMaxPrice() bool`

HasMaxPrice returns a boolean if a field has been set.

### GetAllowAlternativeFareOptions

`func (o *SearchCriteria) GetAllowAlternativeFareOptions() bool`

GetAllowAlternativeFareOptions returns the AllowAlternativeFareOptions field if non-nil, zero value otherwise.

### GetAllowAlternativeFareOptionsOk

`func (o *SearchCriteria) GetAllowAlternativeFareOptionsOk() (*bool, bool)`

GetAllowAlternativeFareOptionsOk returns a tuple with the AllowAlternativeFareOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowAlternativeFareOptions

`func (o *SearchCriteria) SetAllowAlternativeFareOptions(v bool)`

SetAllowAlternativeFareOptions sets AllowAlternativeFareOptions field to given value.

### HasAllowAlternativeFareOptions

`func (o *SearchCriteria) HasAllowAlternativeFareOptions() bool`

HasAllowAlternativeFareOptions returns a boolean if a field has been set.

### GetOneFlightOfferPerDay

`func (o *SearchCriteria) GetOneFlightOfferPerDay() bool`

GetOneFlightOfferPerDay returns the OneFlightOfferPerDay field if non-nil, zero value otherwise.

### GetOneFlightOfferPerDayOk

`func (o *SearchCriteria) GetOneFlightOfferPerDayOk() (*bool, bool)`

GetOneFlightOfferPerDayOk returns a tuple with the OneFlightOfferPerDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOneFlightOfferPerDay

`func (o *SearchCriteria) SetOneFlightOfferPerDay(v bool)`

SetOneFlightOfferPerDay sets OneFlightOfferPerDay field to given value.

### HasOneFlightOfferPerDay

`func (o *SearchCriteria) HasOneFlightOfferPerDay() bool`

HasOneFlightOfferPerDay returns a boolean if a field has been set.

### GetAdditionalInformation

`func (o *SearchCriteria) GetAdditionalInformation() AdditionalInformation`

GetAdditionalInformation returns the AdditionalInformation field if non-nil, zero value otherwise.

### GetAdditionalInformationOk

`func (o *SearchCriteria) GetAdditionalInformationOk() (*AdditionalInformation, bool)`

GetAdditionalInformationOk returns a tuple with the AdditionalInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalInformation

`func (o *SearchCriteria) SetAdditionalInformation(v AdditionalInformation)`

SetAdditionalInformation sets AdditionalInformation field to given value.

### HasAdditionalInformation

`func (o *SearchCriteria) HasAdditionalInformation() bool`

HasAdditionalInformation returns a boolean if a field has been set.

### GetPricingOptions

`func (o *SearchCriteria) GetPricingOptions() ExtendedPricingOptions`

GetPricingOptions returns the PricingOptions field if non-nil, zero value otherwise.

### GetPricingOptionsOk

`func (o *SearchCriteria) GetPricingOptionsOk() (*ExtendedPricingOptions, bool)`

GetPricingOptionsOk returns a tuple with the PricingOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricingOptions

`func (o *SearchCriteria) SetPricingOptions(v ExtendedPricingOptions)`

SetPricingOptions sets PricingOptions field to given value.

### HasPricingOptions

`func (o *SearchCriteria) HasPricingOptions() bool`

HasPricingOptions returns a boolean if a field has been set.

### GetFlightFilters

`func (o *SearchCriteria) GetFlightFilters() FlightFilters`

GetFlightFilters returns the FlightFilters field if non-nil, zero value otherwise.

### GetFlightFiltersOk

`func (o *SearchCriteria) GetFlightFiltersOk() (*FlightFilters, bool)`

GetFlightFiltersOk returns a tuple with the FlightFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlightFilters

`func (o *SearchCriteria) SetFlightFilters(v FlightFilters)`

SetFlightFilters sets FlightFilters field to given value.

### HasFlightFilters

`func (o *SearchCriteria) HasFlightFilters() bool`

HasFlightFilters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


