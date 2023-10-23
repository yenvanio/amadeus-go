# Bags

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Quantity** | Pointer to **int32** | Total number of units | [optional] 
**Weight** | Pointer to **int32** | Weight of the baggage allowance | [optional] 
**WeightUnit** | Pointer to **string** | Code to qualify unit as pounds or kilos | [optional] 
**Name** | Pointer to **string** | Type of service | [optional] 
**Price** | Pointer to [**ElementaryPrice**](ElementaryPrice.md) |  | [optional] 
**BookableByItinerary** | Pointer to **bool** | Specify if the service is bookable by itinerary  or for all itineraries | [optional] 
**SegmentIds** | Pointer to **[]string** | Id of the segment concerned by the service | [optional] 
**TravelerIds** | Pointer to **[]string** | Id of the traveler concerned by the service | [optional] 

## Methods

### NewBags

`func NewBags() *Bags`

NewBags instantiates a new Bags object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBagsWithDefaults

`func NewBagsWithDefaults() *Bags`

NewBagsWithDefaults instantiates a new Bags object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuantity

`func (o *Bags) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *Bags) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *Bags) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *Bags) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetWeight

`func (o *Bags) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *Bags) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *Bags) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *Bags) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetWeightUnit

`func (o *Bags) GetWeightUnit() string`

GetWeightUnit returns the WeightUnit field if non-nil, zero value otherwise.

### GetWeightUnitOk

`func (o *Bags) GetWeightUnitOk() (*string, bool)`

GetWeightUnitOk returns a tuple with the WeightUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeightUnit

`func (o *Bags) SetWeightUnit(v string)`

SetWeightUnit sets WeightUnit field to given value.

### HasWeightUnit

`func (o *Bags) HasWeightUnit() bool`

HasWeightUnit returns a boolean if a field has been set.

### GetName

`func (o *Bags) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Bags) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Bags) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Bags) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPrice

`func (o *Bags) GetPrice() ElementaryPrice`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *Bags) GetPriceOk() (*ElementaryPrice, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *Bags) SetPrice(v ElementaryPrice)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *Bags) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetBookableByItinerary

`func (o *Bags) GetBookableByItinerary() bool`

GetBookableByItinerary returns the BookableByItinerary field if non-nil, zero value otherwise.

### GetBookableByItineraryOk

`func (o *Bags) GetBookableByItineraryOk() (*bool, bool)`

GetBookableByItineraryOk returns a tuple with the BookableByItinerary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBookableByItinerary

`func (o *Bags) SetBookableByItinerary(v bool)`

SetBookableByItinerary sets BookableByItinerary field to given value.

### HasBookableByItinerary

`func (o *Bags) HasBookableByItinerary() bool`

HasBookableByItinerary returns a boolean if a field has been set.

### GetSegmentIds

`func (o *Bags) GetSegmentIds() []string`

GetSegmentIds returns the SegmentIds field if non-nil, zero value otherwise.

### GetSegmentIdsOk

`func (o *Bags) GetSegmentIdsOk() (*[]string, bool)`

GetSegmentIdsOk returns a tuple with the SegmentIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegmentIds

`func (o *Bags) SetSegmentIds(v []string)`

SetSegmentIds sets SegmentIds field to given value.

### HasSegmentIds

`func (o *Bags) HasSegmentIds() bool`

HasSegmentIds returns a boolean if a field has been set.

### GetTravelerIds

`func (o *Bags) GetTravelerIds() []string`

GetTravelerIds returns the TravelerIds field if non-nil, zero value otherwise.

### GetTravelerIdsOk

`func (o *Bags) GetTravelerIdsOk() (*[]string, bool)`

GetTravelerIdsOk returns a tuple with the TravelerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTravelerIds

`func (o *Bags) SetTravelerIds(v []string)`

SetTravelerIds sets TravelerIds field to given value.

### HasTravelerIds

`func (o *Bags) HasTravelerIds() bool`

HasTravelerIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


