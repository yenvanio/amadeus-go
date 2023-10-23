# ChargeableCheckdBags

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Quantity** | Pointer to **int32** | Total number of units | [optional] 
**Weight** | Pointer to **int32** | Weight of the baggage allowance | [optional] 
**WeightUnit** | Pointer to **string** | Code to qualify unit as pounds or kilos | [optional] 
**Id** | Pointer to **string** | Id of the chargeable bag | [optional] 

## Methods

### NewChargeableCheckdBags

`func NewChargeableCheckdBags() *ChargeableCheckdBags`

NewChargeableCheckdBags instantiates a new ChargeableCheckdBags object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChargeableCheckdBagsWithDefaults

`func NewChargeableCheckdBagsWithDefaults() *ChargeableCheckdBags`

NewChargeableCheckdBagsWithDefaults instantiates a new ChargeableCheckdBags object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuantity

`func (o *ChargeableCheckdBags) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *ChargeableCheckdBags) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *ChargeableCheckdBags) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *ChargeableCheckdBags) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetWeight

`func (o *ChargeableCheckdBags) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *ChargeableCheckdBags) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *ChargeableCheckdBags) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *ChargeableCheckdBags) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetWeightUnit

`func (o *ChargeableCheckdBags) GetWeightUnit() string`

GetWeightUnit returns the WeightUnit field if non-nil, zero value otherwise.

### GetWeightUnitOk

`func (o *ChargeableCheckdBags) GetWeightUnitOk() (*string, bool)`

GetWeightUnitOk returns a tuple with the WeightUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeightUnit

`func (o *ChargeableCheckdBags) SetWeightUnit(v string)`

SetWeightUnit sets WeightUnit field to given value.

### HasWeightUnit

`func (o *ChargeableCheckdBags) HasWeightUnit() bool`

HasWeightUnit returns a boolean if a field has been set.

### GetId

`func (o *ChargeableCheckdBags) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ChargeableCheckdBags) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ChargeableCheckdBags) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ChargeableCheckdBags) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


