# LoyaltyProgram

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProgramOwner** | Pointer to **string** | loyalty program airline code | [optional] 
**Id** | Pointer to **string** | loyalty program number | [optional] 

## Methods

### NewLoyaltyProgram

`func NewLoyaltyProgram() *LoyaltyProgram`

NewLoyaltyProgram instantiates a new LoyaltyProgram object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoyaltyProgramWithDefaults

`func NewLoyaltyProgramWithDefaults() *LoyaltyProgram`

NewLoyaltyProgramWithDefaults instantiates a new LoyaltyProgram object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProgramOwner

`func (o *LoyaltyProgram) GetProgramOwner() string`

GetProgramOwner returns the ProgramOwner field if non-nil, zero value otherwise.

### GetProgramOwnerOk

`func (o *LoyaltyProgram) GetProgramOwnerOk() (*string, bool)`

GetProgramOwnerOk returns a tuple with the ProgramOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgramOwner

`func (o *LoyaltyProgram) SetProgramOwner(v string)`

SetProgramOwner sets ProgramOwner field to given value.

### HasProgramOwner

`func (o *LoyaltyProgram) HasProgramOwner() bool`

HasProgramOwner returns a boolean if a field has been set.

### GetId

`func (o *LoyaltyProgram) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LoyaltyProgram) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LoyaltyProgram) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LoyaltyProgram) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


