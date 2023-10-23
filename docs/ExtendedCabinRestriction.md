# ExtendedCabinRestriction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cabin** | Pointer to [**TravelClass**](TravelClass.md) |  | [optional] 
**OriginDestinationIds** | Pointer to **[]string** | The list of originDestination identifiers for which the cabinRestriction applies | [optional] 
**Coverage** | Pointer to [**Coverage**](Coverage.md) |  | [optional] 

## Methods

### NewExtendedCabinRestriction

`func NewExtendedCabinRestriction() *ExtendedCabinRestriction`

NewExtendedCabinRestriction instantiates a new ExtendedCabinRestriction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtendedCabinRestrictionWithDefaults

`func NewExtendedCabinRestrictionWithDefaults() *ExtendedCabinRestriction`

NewExtendedCabinRestrictionWithDefaults instantiates a new ExtendedCabinRestriction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCabin

`func (o *ExtendedCabinRestriction) GetCabin() TravelClass`

GetCabin returns the Cabin field if non-nil, zero value otherwise.

### GetCabinOk

`func (o *ExtendedCabinRestriction) GetCabinOk() (*TravelClass, bool)`

GetCabinOk returns a tuple with the Cabin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCabin

`func (o *ExtendedCabinRestriction) SetCabin(v TravelClass)`

SetCabin sets Cabin field to given value.

### HasCabin

`func (o *ExtendedCabinRestriction) HasCabin() bool`

HasCabin returns a boolean if a field has been set.

### GetOriginDestinationIds

`func (o *ExtendedCabinRestriction) GetOriginDestinationIds() []string`

GetOriginDestinationIds returns the OriginDestinationIds field if non-nil, zero value otherwise.

### GetOriginDestinationIdsOk

`func (o *ExtendedCabinRestriction) GetOriginDestinationIdsOk() (*[]string, bool)`

GetOriginDestinationIdsOk returns a tuple with the OriginDestinationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginDestinationIds

`func (o *ExtendedCabinRestriction) SetOriginDestinationIds(v []string)`

SetOriginDestinationIds sets OriginDestinationIds field to given value.

### HasOriginDestinationIds

`func (o *ExtendedCabinRestriction) HasOriginDestinationIds() bool`

HasOriginDestinationIds returns a boolean if a field has been set.

### GetCoverage

`func (o *ExtendedCabinRestriction) GetCoverage() Coverage`

GetCoverage returns the Coverage field if non-nil, zero value otherwise.

### GetCoverageOk

`func (o *ExtendedCabinRestriction) GetCoverageOk() (*Coverage, bool)`

GetCoverageOk returns a tuple with the Coverage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoverage

`func (o *ExtendedCabinRestriction) SetCoverage(v Coverage)`

SetCoverage sets Coverage field to given value.

### HasCoverage

`func (o *ExtendedCabinRestriction) HasCoverage() bool`

HasCoverage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


