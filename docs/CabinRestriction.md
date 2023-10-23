# CabinRestriction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cabin** | Pointer to [**TravelClass**](TravelClass.md) |  | [optional] 
**OriginDestinationIds** | Pointer to **[]string** | The list of originDestination identifiers for which the cabinRestriction applies | [optional] 

## Methods

### NewCabinRestriction

`func NewCabinRestriction() *CabinRestriction`

NewCabinRestriction instantiates a new CabinRestriction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCabinRestrictionWithDefaults

`func NewCabinRestrictionWithDefaults() *CabinRestriction`

NewCabinRestrictionWithDefaults instantiates a new CabinRestriction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCabin

`func (o *CabinRestriction) GetCabin() TravelClass`

GetCabin returns the Cabin field if non-nil, zero value otherwise.

### GetCabinOk

`func (o *CabinRestriction) GetCabinOk() (*TravelClass, bool)`

GetCabinOk returns a tuple with the Cabin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCabin

`func (o *CabinRestriction) SetCabin(v TravelClass)`

SetCabin sets Cabin field to given value.

### HasCabin

`func (o *CabinRestriction) HasCabin() bool`

HasCabin returns a boolean if a field has been set.

### GetOriginDestinationIds

`func (o *CabinRestriction) GetOriginDestinationIds() []string`

GetOriginDestinationIds returns the OriginDestinationIds field if non-nil, zero value otherwise.

### GetOriginDestinationIdsOk

`func (o *CabinRestriction) GetOriginDestinationIdsOk() (*[]string, bool)`

GetOriginDestinationIdsOk returns a tuple with the OriginDestinationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginDestinationIds

`func (o *CabinRestriction) SetOriginDestinationIds(v []string)`

SetOriginDestinationIds sets OriginDestinationIds field to given value.

### HasOriginDestinationIds

`func (o *CabinRestriction) HasOriginDestinationIds() bool`

HasOriginDestinationIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


