# TravelerInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**TravelerType** | [**TravelerType**](TravelerType.md) |  | 
**AssociatedAdultId** | Pointer to **string** | if type&#x3D;\&quot;HELD_INFANT\&quot;, corresponds to the adult travelers&#39;s id who will share the seat | [optional] 

## Methods

### NewTravelerInfo

`func NewTravelerInfo(id string, travelerType TravelerType, ) *TravelerInfo`

NewTravelerInfo instantiates a new TravelerInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTravelerInfoWithDefaults

`func NewTravelerInfoWithDefaults() *TravelerInfo`

NewTravelerInfoWithDefaults instantiates a new TravelerInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TravelerInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TravelerInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TravelerInfo) SetId(v string)`

SetId sets Id field to given value.


### GetTravelerType

`func (o *TravelerInfo) GetTravelerType() TravelerType`

GetTravelerType returns the TravelerType field if non-nil, zero value otherwise.

### GetTravelerTypeOk

`func (o *TravelerInfo) GetTravelerTypeOk() (*TravelerType, bool)`

GetTravelerTypeOk returns a tuple with the TravelerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTravelerType

`func (o *TravelerInfo) SetTravelerType(v TravelerType)`

SetTravelerType sets TravelerType field to given value.


### GetAssociatedAdultId

`func (o *TravelerInfo) GetAssociatedAdultId() string`

GetAssociatedAdultId returns the AssociatedAdultId field if non-nil, zero value otherwise.

### GetAssociatedAdultIdOk

`func (o *TravelerInfo) GetAssociatedAdultIdOk() (*string, bool)`

GetAssociatedAdultIdOk returns a tuple with the AssociatedAdultId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedAdultId

`func (o *TravelerInfo) SetAssociatedAdultId(v string)`

SetAssociatedAdultId sets AssociatedAdultId field to given value.

### HasAssociatedAdultId

`func (o *TravelerInfo) HasAssociatedAdultId() bool`

HasAssociatedAdultId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


