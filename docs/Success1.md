# Success1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Warnings** | Pointer to [**[]Issue**](Issue.md) |  | [optional] 
**Meta** | Pointer to [**CollectionMeta**](CollectionMeta.md) |  | [optional] 
**Data** | [**[]FlightOffer**](FlightOffer.md) |  | 
**Dictionaries** | Pointer to [**Dictionaries**](Dictionaries.md) |  | [optional] 

## Methods

### NewSuccess1

`func NewSuccess1(data []FlightOffer, ) *Success1`

NewSuccess1 instantiates a new Success1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSuccess1WithDefaults

`func NewSuccess1WithDefaults() *Success1`

NewSuccess1WithDefaults instantiates a new Success1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWarnings

`func (o *Success1) GetWarnings() []Issue`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *Success1) GetWarningsOk() (*[]Issue, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *Success1) SetWarnings(v []Issue)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *Success1) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.

### GetMeta

`func (o *Success1) GetMeta() CollectionMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *Success1) GetMetaOk() (*CollectionMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *Success1) SetMeta(v CollectionMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *Success1) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetData

`func (o *Success1) GetData() []FlightOffer`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *Success1) GetDataOk() (*[]FlightOffer, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *Success1) SetData(v []FlightOffer)`

SetData sets Data field to given value.


### GetDictionaries

`func (o *Success1) GetDictionaries() Dictionaries`

GetDictionaries returns the Dictionaries field if non-nil, zero value otherwise.

### GetDictionariesOk

`func (o *Success1) GetDictionariesOk() (*Dictionaries, bool)`

GetDictionariesOk returns a tuple with the Dictionaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDictionaries

`func (o *Success1) SetDictionaries(v Dictionaries)`

SetDictionaries sets Dictionaries field to given value.

### HasDictionaries

`func (o *Success1) HasDictionaries() bool`

HasDictionaries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


