# TravelerPricing

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TravelerId** | **string** | Id of the traveler | 
**FareOption** | [**TravelerPricingFareOption**](TravelerPricingFareOption.md) |  | 
**TravelerType** | [**TravelerType**](TravelerType.md) |  | 
**AssociatedAdultId** | Pointer to **string** | if type&#x3D;\&quot;HELD_INFANT\&quot;, corresponds to the adult traveler&#39;s id who will share the seat | [optional] 
**Price** | Pointer to [**Price**](Price.md) |  | [optional] 
**FareDetailsBySegment** | [**[]FareDetailsBySegment**](FareDetailsBySegment.md) |  | 

## Methods

### NewTravelerPricing

`func NewTravelerPricing(travelerId string, fareOption TravelerPricingFareOption, travelerType TravelerType, fareDetailsBySegment []FareDetailsBySegment, ) *TravelerPricing`

NewTravelerPricing instantiates a new TravelerPricing object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTravelerPricingWithDefaults

`func NewTravelerPricingWithDefaults() *TravelerPricing`

NewTravelerPricingWithDefaults instantiates a new TravelerPricing object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTravelerId

`func (o *TravelerPricing) GetTravelerId() string`

GetTravelerId returns the TravelerId field if non-nil, zero value otherwise.

### GetTravelerIdOk

`func (o *TravelerPricing) GetTravelerIdOk() (*string, bool)`

GetTravelerIdOk returns a tuple with the TravelerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTravelerId

`func (o *TravelerPricing) SetTravelerId(v string)`

SetTravelerId sets TravelerId field to given value.


### GetFareOption

`func (o *TravelerPricing) GetFareOption() TravelerPricingFareOption`

GetFareOption returns the FareOption field if non-nil, zero value otherwise.

### GetFareOptionOk

`func (o *TravelerPricing) GetFareOptionOk() (*TravelerPricingFareOption, bool)`

GetFareOptionOk returns a tuple with the FareOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFareOption

`func (o *TravelerPricing) SetFareOption(v TravelerPricingFareOption)`

SetFareOption sets FareOption field to given value.


### GetTravelerType

`func (o *TravelerPricing) GetTravelerType() TravelerType`

GetTravelerType returns the TravelerType field if non-nil, zero value otherwise.

### GetTravelerTypeOk

`func (o *TravelerPricing) GetTravelerTypeOk() (*TravelerType, bool)`

GetTravelerTypeOk returns a tuple with the TravelerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTravelerType

`func (o *TravelerPricing) SetTravelerType(v TravelerType)`

SetTravelerType sets TravelerType field to given value.


### GetAssociatedAdultId

`func (o *TravelerPricing) GetAssociatedAdultId() string`

GetAssociatedAdultId returns the AssociatedAdultId field if non-nil, zero value otherwise.

### GetAssociatedAdultIdOk

`func (o *TravelerPricing) GetAssociatedAdultIdOk() (*string, bool)`

GetAssociatedAdultIdOk returns a tuple with the AssociatedAdultId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedAdultId

`func (o *TravelerPricing) SetAssociatedAdultId(v string)`

SetAssociatedAdultId sets AssociatedAdultId field to given value.

### HasAssociatedAdultId

`func (o *TravelerPricing) HasAssociatedAdultId() bool`

HasAssociatedAdultId returns a boolean if a field has been set.

### GetPrice

`func (o *TravelerPricing) GetPrice() Price`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *TravelerPricing) GetPriceOk() (*Price, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *TravelerPricing) SetPrice(v Price)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *TravelerPricing) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetFareDetailsBySegment

`func (o *TravelerPricing) GetFareDetailsBySegment() []FareDetailsBySegment`

GetFareDetailsBySegment returns the FareDetailsBySegment field if non-nil, zero value otherwise.

### GetFareDetailsBySegmentOk

`func (o *TravelerPricing) GetFareDetailsBySegmentOk() (*[]FareDetailsBySegment, bool)`

GetFareDetailsBySegmentOk returns a tuple with the FareDetailsBySegment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFareDetailsBySegment

`func (o *TravelerPricing) SetFareDetailsBySegment(v []FareDetailsBySegment)`

SetFareDetailsBySegment sets FareDetailsBySegment field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


