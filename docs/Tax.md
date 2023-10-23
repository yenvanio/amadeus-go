# Tax

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 

## Methods

### NewTax

`func NewTax() *Tax`

NewTax instantiates a new Tax object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaxWithDefaults

`func NewTaxWithDefaults() *Tax`

NewTaxWithDefaults instantiates a new Tax object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *Tax) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *Tax) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *Tax) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *Tax) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetCode

`func (o *Tax) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *Tax) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *Tax) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *Tax) HasCode() bool`

HasCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


