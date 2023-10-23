# AdditionalService

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**AdditionalServiceType**](AdditionalServiceType.md) |  | [optional] 

## Methods

### NewAdditionalService

`func NewAdditionalService() *AdditionalService`

NewAdditionalService instantiates a new AdditionalService object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdditionalServiceWithDefaults

`func NewAdditionalServiceWithDefaults() *AdditionalService`

NewAdditionalServiceWithDefaults instantiates a new AdditionalService object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *AdditionalService) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *AdditionalService) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *AdditionalService) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *AdditionalService) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetType

`func (o *AdditionalService) GetType() AdditionalServiceType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AdditionalService) GetTypeOk() (*AdditionalServiceType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AdditionalService) SetType(v AdditionalServiceType)`

SetType sets Type field to given value.

### HasType

`func (o *AdditionalService) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


