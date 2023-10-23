# FareDetailsBySegment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SegmentId** | **string** | Id of the segment | 
**Cabin** | Pointer to [**TravelClass**](TravelClass.md) |  | [optional] 
**FareBasis** | Pointer to **string** | Fare basis specifying the rules of a fare. Usually, though not always, is composed of the booking class code followed by a set of letters and digits representing other characteristics of the ticket, such as refundability, minimum stay requirements, discounts or special promotional elements. | [optional] 
**BrandedFare** | Pointer to **string** | The name of the Fare Family corresponding to the fares. Only for the GDS provider and if the airline has fare families filled | [optional] 
**Class** | Pointer to **string** | The code of the booking class, a.k.a. class of service or Reservations/Booking Designator (RBD) | [optional] 
**IsAllotment** | Pointer to **bool** | True if the corresponding booking class is in an allotment | [optional] 
**AllotmentDetails** | Pointer to [**AllotmentDetails**](AllotmentDetails.md) |  | [optional] 
**SliceDiceIndicator** | Pointer to [**SliceDiceIndicator**](SliceDiceIndicator.md) |  | [optional] 
**IncludedCheckedBags** | Pointer to [**BaggageAllowance**](BaggageAllowance.md) |  | [optional] 
**AdditionalServices** | Pointer to [**AdditionalServicesRequest**](AdditionalServicesRequest.md) |  | [optional] 

## Methods

### NewFareDetailsBySegment

`func NewFareDetailsBySegment(segmentId string, ) *FareDetailsBySegment`

NewFareDetailsBySegment instantiates a new FareDetailsBySegment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFareDetailsBySegmentWithDefaults

`func NewFareDetailsBySegmentWithDefaults() *FareDetailsBySegment`

NewFareDetailsBySegmentWithDefaults instantiates a new FareDetailsBySegment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSegmentId

`func (o *FareDetailsBySegment) GetSegmentId() string`

GetSegmentId returns the SegmentId field if non-nil, zero value otherwise.

### GetSegmentIdOk

`func (o *FareDetailsBySegment) GetSegmentIdOk() (*string, bool)`

GetSegmentIdOk returns a tuple with the SegmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegmentId

`func (o *FareDetailsBySegment) SetSegmentId(v string)`

SetSegmentId sets SegmentId field to given value.


### GetCabin

`func (o *FareDetailsBySegment) GetCabin() TravelClass`

GetCabin returns the Cabin field if non-nil, zero value otherwise.

### GetCabinOk

`func (o *FareDetailsBySegment) GetCabinOk() (*TravelClass, bool)`

GetCabinOk returns a tuple with the Cabin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCabin

`func (o *FareDetailsBySegment) SetCabin(v TravelClass)`

SetCabin sets Cabin field to given value.

### HasCabin

`func (o *FareDetailsBySegment) HasCabin() bool`

HasCabin returns a boolean if a field has been set.

### GetFareBasis

`func (o *FareDetailsBySegment) GetFareBasis() string`

GetFareBasis returns the FareBasis field if non-nil, zero value otherwise.

### GetFareBasisOk

`func (o *FareDetailsBySegment) GetFareBasisOk() (*string, bool)`

GetFareBasisOk returns a tuple with the FareBasis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFareBasis

`func (o *FareDetailsBySegment) SetFareBasis(v string)`

SetFareBasis sets FareBasis field to given value.

### HasFareBasis

`func (o *FareDetailsBySegment) HasFareBasis() bool`

HasFareBasis returns a boolean if a field has been set.

### GetBrandedFare

`func (o *FareDetailsBySegment) GetBrandedFare() string`

GetBrandedFare returns the BrandedFare field if non-nil, zero value otherwise.

### GetBrandedFareOk

`func (o *FareDetailsBySegment) GetBrandedFareOk() (*string, bool)`

GetBrandedFareOk returns a tuple with the BrandedFare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandedFare

`func (o *FareDetailsBySegment) SetBrandedFare(v string)`

SetBrandedFare sets BrandedFare field to given value.

### HasBrandedFare

`func (o *FareDetailsBySegment) HasBrandedFare() bool`

HasBrandedFare returns a boolean if a field has been set.

### GetClass

`func (o *FareDetailsBySegment) GetClass() string`

GetClass returns the Class field if non-nil, zero value otherwise.

### GetClassOk

`func (o *FareDetailsBySegment) GetClassOk() (*string, bool)`

GetClassOk returns a tuple with the Class field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClass

`func (o *FareDetailsBySegment) SetClass(v string)`

SetClass sets Class field to given value.

### HasClass

`func (o *FareDetailsBySegment) HasClass() bool`

HasClass returns a boolean if a field has been set.

### GetIsAllotment

`func (o *FareDetailsBySegment) GetIsAllotment() bool`

GetIsAllotment returns the IsAllotment field if non-nil, zero value otherwise.

### GetIsAllotmentOk

`func (o *FareDetailsBySegment) GetIsAllotmentOk() (*bool, bool)`

GetIsAllotmentOk returns a tuple with the IsAllotment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAllotment

`func (o *FareDetailsBySegment) SetIsAllotment(v bool)`

SetIsAllotment sets IsAllotment field to given value.

### HasIsAllotment

`func (o *FareDetailsBySegment) HasIsAllotment() bool`

HasIsAllotment returns a boolean if a field has been set.

### GetAllotmentDetails

`func (o *FareDetailsBySegment) GetAllotmentDetails() AllotmentDetails`

GetAllotmentDetails returns the AllotmentDetails field if non-nil, zero value otherwise.

### GetAllotmentDetailsOk

`func (o *FareDetailsBySegment) GetAllotmentDetailsOk() (*AllotmentDetails, bool)`

GetAllotmentDetailsOk returns a tuple with the AllotmentDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllotmentDetails

`func (o *FareDetailsBySegment) SetAllotmentDetails(v AllotmentDetails)`

SetAllotmentDetails sets AllotmentDetails field to given value.

### HasAllotmentDetails

`func (o *FareDetailsBySegment) HasAllotmentDetails() bool`

HasAllotmentDetails returns a boolean if a field has been set.

### GetSliceDiceIndicator

`func (o *FareDetailsBySegment) GetSliceDiceIndicator() SliceDiceIndicator`

GetSliceDiceIndicator returns the SliceDiceIndicator field if non-nil, zero value otherwise.

### GetSliceDiceIndicatorOk

`func (o *FareDetailsBySegment) GetSliceDiceIndicatorOk() (*SliceDiceIndicator, bool)`

GetSliceDiceIndicatorOk returns a tuple with the SliceDiceIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSliceDiceIndicator

`func (o *FareDetailsBySegment) SetSliceDiceIndicator(v SliceDiceIndicator)`

SetSliceDiceIndicator sets SliceDiceIndicator field to given value.

### HasSliceDiceIndicator

`func (o *FareDetailsBySegment) HasSliceDiceIndicator() bool`

HasSliceDiceIndicator returns a boolean if a field has been set.

### GetIncludedCheckedBags

`func (o *FareDetailsBySegment) GetIncludedCheckedBags() BaggageAllowance`

GetIncludedCheckedBags returns the IncludedCheckedBags field if non-nil, zero value otherwise.

### GetIncludedCheckedBagsOk

`func (o *FareDetailsBySegment) GetIncludedCheckedBagsOk() (*BaggageAllowance, bool)`

GetIncludedCheckedBagsOk returns a tuple with the IncludedCheckedBags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedCheckedBags

`func (o *FareDetailsBySegment) SetIncludedCheckedBags(v BaggageAllowance)`

SetIncludedCheckedBags sets IncludedCheckedBags field to given value.

### HasIncludedCheckedBags

`func (o *FareDetailsBySegment) HasIncludedCheckedBags() bool`

HasIncludedCheckedBags returns a boolean if a field has been set.

### GetAdditionalServices

`func (o *FareDetailsBySegment) GetAdditionalServices() AdditionalServicesRequest`

GetAdditionalServices returns the AdditionalServices field if non-nil, zero value otherwise.

### GetAdditionalServicesOk

`func (o *FareDetailsBySegment) GetAdditionalServicesOk() (*AdditionalServicesRequest, bool)`

GetAdditionalServicesOk returns a tuple with the AdditionalServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalServices

`func (o *FareDetailsBySegment) SetAdditionalServices(v AdditionalServicesRequest)`

SetAdditionalServices sets AdditionalServices field to given value.

### HasAdditionalServices

`func (o *FareDetailsBySegment) HasAdditionalServices() bool`

HasAdditionalServices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


