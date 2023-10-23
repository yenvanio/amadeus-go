# ElementaryPrice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **string** | Amount of the fare. could be alpha numeric. Ex- 500.20 or 514.13A, &#39;A&#39;signifies additional collection. | [optional] 
**CurrencyCode** | Pointer to **string** | Currency type of the fare. | [optional] 

## Methods

### NewElementaryPrice

`func NewElementaryPrice() *ElementaryPrice`

NewElementaryPrice instantiates a new ElementaryPrice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewElementaryPriceWithDefaults

`func NewElementaryPriceWithDefaults() *ElementaryPrice`

NewElementaryPriceWithDefaults instantiates a new ElementaryPrice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *ElementaryPrice) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ElementaryPrice) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ElementaryPrice) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *ElementaryPrice) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *ElementaryPrice) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *ElementaryPrice) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *ElementaryPrice) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *ElementaryPrice) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


