# OtherServices

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to [**ServiceName**](ServiceName.md) |  | [optional] 
**Price** | Pointer to [**ElementaryPrice**](ElementaryPrice.md) |  | [optional] 
**BookableByTraveler** | Pointer to **bool** | Specify if the service is bookable by traveler or for all travelers | [optional] 
**BookableByItinerary** | Pointer to **bool** | Specify if the service is bookable by itinerary or for all itineraries | [optional] 
**SegmentIds** | Pointer to **[]string** | Id of the segment concerned by the service | [optional] 
**TravelerIds** | Pointer to **[]string** | Id of the traveler concerned by the service | [optional] 

## Methods

### NewOtherServices

`func NewOtherServices() *OtherServices`

NewOtherServices instantiates a new OtherServices object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOtherServicesWithDefaults

`func NewOtherServicesWithDefaults() *OtherServices`

NewOtherServicesWithDefaults instantiates a new OtherServices object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *OtherServices) GetName() ServiceName`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OtherServices) GetNameOk() (*ServiceName, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OtherServices) SetName(v ServiceName)`

SetName sets Name field to given value.

### HasName

`func (o *OtherServices) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPrice

`func (o *OtherServices) GetPrice() ElementaryPrice`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *OtherServices) GetPriceOk() (*ElementaryPrice, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *OtherServices) SetPrice(v ElementaryPrice)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *OtherServices) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetBookableByTraveler

`func (o *OtherServices) GetBookableByTraveler() bool`

GetBookableByTraveler returns the BookableByTraveler field if non-nil, zero value otherwise.

### GetBookableByTravelerOk

`func (o *OtherServices) GetBookableByTravelerOk() (*bool, bool)`

GetBookableByTravelerOk returns a tuple with the BookableByTraveler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBookableByTraveler

`func (o *OtherServices) SetBookableByTraveler(v bool)`

SetBookableByTraveler sets BookableByTraveler field to given value.

### HasBookableByTraveler

`func (o *OtherServices) HasBookableByTraveler() bool`

HasBookableByTraveler returns a boolean if a field has been set.

### GetBookableByItinerary

`func (o *OtherServices) GetBookableByItinerary() bool`

GetBookableByItinerary returns the BookableByItinerary field if non-nil, zero value otherwise.

### GetBookableByItineraryOk

`func (o *OtherServices) GetBookableByItineraryOk() (*bool, bool)`

GetBookableByItineraryOk returns a tuple with the BookableByItinerary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBookableByItinerary

`func (o *OtherServices) SetBookableByItinerary(v bool)`

SetBookableByItinerary sets BookableByItinerary field to given value.

### HasBookableByItinerary

`func (o *OtherServices) HasBookableByItinerary() bool`

HasBookableByItinerary returns a boolean if a field has been set.

### GetSegmentIds

`func (o *OtherServices) GetSegmentIds() []string`

GetSegmentIds returns the SegmentIds field if non-nil, zero value otherwise.

### GetSegmentIdsOk

`func (o *OtherServices) GetSegmentIdsOk() (*[]string, bool)`

GetSegmentIdsOk returns a tuple with the SegmentIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegmentIds

`func (o *OtherServices) SetSegmentIds(v []string)`

SetSegmentIds sets SegmentIds field to given value.

### HasSegmentIds

`func (o *OtherServices) HasSegmentIds() bool`

HasSegmentIds returns a boolean if a field has been set.

### GetTravelerIds

`func (o *OtherServices) GetTravelerIds() []string`

GetTravelerIds returns the TravelerIds field if non-nil, zero value otherwise.

### GetTravelerIdsOk

`func (o *OtherServices) GetTravelerIdsOk() (*[]string, bool)`

GetTravelerIdsOk returns a tuple with the TravelerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTravelerIds

`func (o *OtherServices) SetTravelerIds(v []string)`

SetTravelerIds sets TravelerIds field to given value.

### HasTravelerIds

`func (o *OtherServices) HasTravelerIds() bool`

HasTravelerIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


