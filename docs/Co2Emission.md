# Co2Emission

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Weight** | Pointer to **int32** | Weight of Co2 emitted for the concerned segment | [optional] 
**WeightUnit** | Pointer to **string** | Code to qualify unit as pounds or kilos | [optional] 
**Cabin** | Pointer to [**TravelClass**](TravelClass.md) |  | [optional] 

## Methods

### NewCo2Emission

`func NewCo2Emission() *Co2Emission`

NewCo2Emission instantiates a new Co2Emission object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCo2EmissionWithDefaults

`func NewCo2EmissionWithDefaults() *Co2Emission`

NewCo2EmissionWithDefaults instantiates a new Co2Emission object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWeight

`func (o *Co2Emission) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *Co2Emission) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *Co2Emission) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *Co2Emission) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetWeightUnit

`func (o *Co2Emission) GetWeightUnit() string`

GetWeightUnit returns the WeightUnit field if non-nil, zero value otherwise.

### GetWeightUnitOk

`func (o *Co2Emission) GetWeightUnitOk() (*string, bool)`

GetWeightUnitOk returns a tuple with the WeightUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeightUnit

`func (o *Co2Emission) SetWeightUnit(v string)`

SetWeightUnit sets WeightUnit field to given value.

### HasWeightUnit

`func (o *Co2Emission) HasWeightUnit() bool`

HasWeightUnit returns a boolean if a field has been set.

### GetCabin

`func (o *Co2Emission) GetCabin() TravelClass`

GetCabin returns the Cabin field if non-nil, zero value otherwise.

### GetCabinOk

`func (o *Co2Emission) GetCabinOk() (*TravelClass, bool)`

GetCabinOk returns a tuple with the Cabin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCabin

`func (o *Co2Emission) SetCabin(v TravelClass)`

SetCabin sets Cabin field to given value.

### HasCabin

`func (o *Co2Emission) HasCabin() bool`

HasCabin returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


