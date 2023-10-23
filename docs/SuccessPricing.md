# SuccessPricing

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Warnings** | Pointer to [**[]Issue**](Issue.md) |  | [optional] 
**Data** | [**FlightOfferPricingOut**](FlightOfferPricingOut.md) |  | 
**Included** | Pointer to [**IncludedResourcesMap**](IncludedResourcesMap.md) |  | [optional] 
**Dictionaries** | Pointer to [**Dictionaries**](Dictionaries.md) |  | [optional] 

## Methods

### NewSuccessPricing

`func NewSuccessPricing(data FlightOfferPricingOut, ) *SuccessPricing`

NewSuccessPricing instantiates a new SuccessPricing object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSuccessPricingWithDefaults

`func NewSuccessPricingWithDefaults() *SuccessPricing`

NewSuccessPricingWithDefaults instantiates a new SuccessPricing object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWarnings

`func (o *SuccessPricing) GetWarnings() []Issue`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *SuccessPricing) GetWarningsOk() (*[]Issue, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *SuccessPricing) SetWarnings(v []Issue)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *SuccessPricing) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.

### GetData

`func (o *SuccessPricing) GetData() FlightOfferPricingOut`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SuccessPricing) GetDataOk() (*FlightOfferPricingOut, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SuccessPricing) SetData(v FlightOfferPricingOut)`

SetData sets Data field to given value.


### GetIncluded

`func (o *SuccessPricing) GetIncluded() IncludedResourcesMap`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *SuccessPricing) GetIncludedOk() (*IncludedResourcesMap, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *SuccessPricing) SetIncluded(v IncludedResourcesMap)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *SuccessPricing) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetDictionaries

`func (o *SuccessPricing) GetDictionaries() Dictionaries`

GetDictionaries returns the Dictionaries field if non-nil, zero value otherwise.

### GetDictionariesOk

`func (o *SuccessPricing) GetDictionariesOk() (*Dictionaries, bool)`

GetDictionariesOk returns a tuple with the Dictionaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDictionaries

`func (o *SuccessPricing) SetDictionaries(v Dictionaries)`

SetDictionaries sets Dictionaries field to given value.

### HasDictionaries

`func (o *SuccessPricing) HasDictionaries() bool`

HasDictionaries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


