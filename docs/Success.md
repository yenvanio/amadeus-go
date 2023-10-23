# Success

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Warnings** | Pointer to [**[]Issue**](Issue.md) |  | [optional] 
**Meta** | Pointer to [**CollectionMetaLink**](CollectionMetaLink.md) |  | [optional] 
**Data** | [**[]FlightOffer**](FlightOffer.md) |  | 
**Dictionaries** | Pointer to [**Dictionaries**](Dictionaries.md) |  | [optional] 

## Methods

### NewSuccess

`func NewSuccess(data []FlightOffer, ) *Success`

NewSuccess instantiates a new Success object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSuccessWithDefaults

`func NewSuccessWithDefaults() *Success`

NewSuccessWithDefaults instantiates a new Success object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWarnings

`func (o *Success) GetWarnings() []Issue`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *Success) GetWarningsOk() (*[]Issue, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *Success) SetWarnings(v []Issue)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *Success) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.

### GetMeta

`func (o *Success) GetMeta() CollectionMetaLink`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *Success) GetMetaOk() (*CollectionMetaLink, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *Success) SetMeta(v CollectionMetaLink)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *Success) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetData

`func (o *Success) GetData() []FlightOffer`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *Success) GetDataOk() (*[]FlightOffer, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *Success) SetData(v []FlightOffer)`

SetData sets Data field to given value.


### GetDictionaries

`func (o *Success) GetDictionaries() Dictionaries`

GetDictionaries returns the Dictionaries field if non-nil, zero value otherwise.

### GetDictionariesOk

`func (o *Success) GetDictionariesOk() (*Dictionaries, bool)`

GetDictionariesOk returns a tuple with the Dictionaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDictionaries

`func (o *Success) SetDictionaries(v Dictionaries)`

SetDictionaries sets Dictionaries field to given value.

### HasDictionaries

`func (o *Success) HasDictionaries() bool`

HasDictionaries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


