# FareRules

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | Pointer to **string** | The currency of the penalties | [optional] 
**Rules** | Pointer to [**[]TermAndCondition**](TermAndCondition.md) |  | [optional] 

## Methods

### NewFareRules

`func NewFareRules() *FareRules`

NewFareRules instantiates a new FareRules object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFareRulesWithDefaults

`func NewFareRulesWithDefaults() *FareRules`

NewFareRulesWithDefaults instantiates a new FareRules object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrency

`func (o *FareRules) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *FareRules) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *FareRules) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *FareRules) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetRules

`func (o *FareRules) GetRules() []TermAndCondition`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *FareRules) GetRulesOk() (*[]TermAndCondition, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *FareRules) SetRules(v []TermAndCondition)`

SetRules sets Rules field to given value.

### HasRules

`func (o *FareRules) HasRules() bool`

HasRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


