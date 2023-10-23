# CollectionMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** |  | [optional] 
**OneWayCombinations** | Pointer to [**[]OneWayCombinations**](OneWayCombinations.md) |  | [optional] 

## Methods

### NewCollectionMeta

`func NewCollectionMeta() *CollectionMeta`

NewCollectionMeta instantiates a new CollectionMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionMetaWithDefaults

`func NewCollectionMetaWithDefaults() *CollectionMeta`

NewCollectionMetaWithDefaults instantiates a new CollectionMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *CollectionMeta) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *CollectionMeta) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *CollectionMeta) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *CollectionMeta) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetOneWayCombinations

`func (o *CollectionMeta) GetOneWayCombinations() []OneWayCombinations`

GetOneWayCombinations returns the OneWayCombinations field if non-nil, zero value otherwise.

### GetOneWayCombinationsOk

`func (o *CollectionMeta) GetOneWayCombinationsOk() (*[]OneWayCombinations, bool)`

GetOneWayCombinationsOk returns a tuple with the OneWayCombinations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOneWayCombinations

`func (o *CollectionMeta) SetOneWayCombinations(v []OneWayCombinations)`

SetOneWayCombinations sets OneWayCombinations field to given value.

### HasOneWayCombinations

`func (o *CollectionMeta) HasOneWayCombinations() bool`

HasOneWayCombinations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


