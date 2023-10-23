# BaggageAllowance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Quantity** | Pointer to **int32** | Total number of units | [optional] 
**Weight** | Pointer to **int32** | Weight of the baggage allowance | [optional] 
**WeightUnit** | Pointer to **string** | Code to qualify unit as pounds or kilos | [optional] 

## Methods

### NewBaggageAllowance

`func NewBaggageAllowance() *BaggageAllowance`

NewBaggageAllowance instantiates a new BaggageAllowance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaggageAllowanceWithDefaults

`func NewBaggageAllowanceWithDefaults() *BaggageAllowance`

NewBaggageAllowanceWithDefaults instantiates a new BaggageAllowance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuantity

`func (o *BaggageAllowance) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *BaggageAllowance) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *BaggageAllowance) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *BaggageAllowance) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetWeight

`func (o *BaggageAllowance) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *BaggageAllowance) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *BaggageAllowance) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *BaggageAllowance) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetWeightUnit

`func (o *BaggageAllowance) GetWeightUnit() string`

GetWeightUnit returns the WeightUnit field if non-nil, zero value otherwise.

### GetWeightUnitOk

`func (o *BaggageAllowance) GetWeightUnitOk() (*string, bool)`

GetWeightUnitOk returns a tuple with the WeightUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeightUnit

`func (o *BaggageAllowance) SetWeightUnit(v string)`

SetWeightUnit sets WeightUnit field to given value.

### HasWeightUnit

`func (o *BaggageAllowance) HasWeightUnit() bool`

HasWeightUnit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


