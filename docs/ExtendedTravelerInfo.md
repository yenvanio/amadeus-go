# ExtendedTravelerInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**TravelerType** | [**TravelerType**](TravelerType.md) |  | 
**AssociatedAdultId** | Pointer to **string** | if type&#x3D;\&quot;HELD_INFANT\&quot;, corresponds to the adult travelers&#39;s id who will share the seat | [optional] 

## Methods

### NewExtendedTravelerInfo

`func NewExtendedTravelerInfo(id string, travelerType TravelerType, ) *ExtendedTravelerInfo`

NewExtendedTravelerInfo instantiates a new ExtendedTravelerInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtendedTravelerInfoWithDefaults

`func NewExtendedTravelerInfoWithDefaults() *ExtendedTravelerInfo`

NewExtendedTravelerInfoWithDefaults instantiates a new ExtendedTravelerInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ExtendedTravelerInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExtendedTravelerInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExtendedTravelerInfo) SetId(v string)`

SetId sets Id field to given value.


### GetTravelerType

`func (o *ExtendedTravelerInfo) GetTravelerType() TravelerType`

GetTravelerType returns the TravelerType field if non-nil, zero value otherwise.

### GetTravelerTypeOk

`func (o *ExtendedTravelerInfo) GetTravelerTypeOk() (*TravelerType, bool)`

GetTravelerTypeOk returns a tuple with the TravelerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTravelerType

`func (o *ExtendedTravelerInfo) SetTravelerType(v TravelerType)`

SetTravelerType sets TravelerType field to given value.


### GetAssociatedAdultId

`func (o *ExtendedTravelerInfo) GetAssociatedAdultId() string`

GetAssociatedAdultId returns the AssociatedAdultId field if non-nil, zero value otherwise.

### GetAssociatedAdultIdOk

`func (o *ExtendedTravelerInfo) GetAssociatedAdultIdOk() (*string, bool)`

GetAssociatedAdultIdOk returns a tuple with the AssociatedAdultId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedAdultId

`func (o *ExtendedTravelerInfo) SetAssociatedAdultId(v string)`

SetAssociatedAdultId sets AssociatedAdultId field to given value.

### HasAssociatedAdultId

`func (o *ExtendedTravelerInfo) HasAssociatedAdultId() bool`

HasAssociatedAdultId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


