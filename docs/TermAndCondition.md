# TermAndCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | Pointer to **string** | This defines what type of modification is concerned in this rule. | [optional] 
**Circumstances** | Pointer to **string** |  | [optional] 
**NotApplicable** | Pointer to **bool** |  | [optional] 
**MaxPenaltyAmount** | Pointer to **string** |  | [optional] 
**Descriptions** | Pointer to [**[]Description**](Description.md) |  | [optional] 

## Methods

### NewTermAndCondition

`func NewTermAndCondition() *TermAndCondition`

NewTermAndCondition instantiates a new TermAndCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTermAndConditionWithDefaults

`func NewTermAndConditionWithDefaults() *TermAndCondition`

NewTermAndConditionWithDefaults instantiates a new TermAndCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *TermAndCondition) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *TermAndCondition) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *TermAndCondition) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *TermAndCondition) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetCircumstances

`func (o *TermAndCondition) GetCircumstances() string`

GetCircumstances returns the Circumstances field if non-nil, zero value otherwise.

### GetCircumstancesOk

`func (o *TermAndCondition) GetCircumstancesOk() (*string, bool)`

GetCircumstancesOk returns a tuple with the Circumstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCircumstances

`func (o *TermAndCondition) SetCircumstances(v string)`

SetCircumstances sets Circumstances field to given value.

### HasCircumstances

`func (o *TermAndCondition) HasCircumstances() bool`

HasCircumstances returns a boolean if a field has been set.

### GetNotApplicable

`func (o *TermAndCondition) GetNotApplicable() bool`

GetNotApplicable returns the NotApplicable field if non-nil, zero value otherwise.

### GetNotApplicableOk

`func (o *TermAndCondition) GetNotApplicableOk() (*bool, bool)`

GetNotApplicableOk returns a tuple with the NotApplicable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotApplicable

`func (o *TermAndCondition) SetNotApplicable(v bool)`

SetNotApplicable sets NotApplicable field to given value.

### HasNotApplicable

`func (o *TermAndCondition) HasNotApplicable() bool`

HasNotApplicable returns a boolean if a field has been set.

### GetMaxPenaltyAmount

`func (o *TermAndCondition) GetMaxPenaltyAmount() string`

GetMaxPenaltyAmount returns the MaxPenaltyAmount field if non-nil, zero value otherwise.

### GetMaxPenaltyAmountOk

`func (o *TermAndCondition) GetMaxPenaltyAmountOk() (*string, bool)`

GetMaxPenaltyAmountOk returns a tuple with the MaxPenaltyAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPenaltyAmount

`func (o *TermAndCondition) SetMaxPenaltyAmount(v string)`

SetMaxPenaltyAmount sets MaxPenaltyAmount field to given value.

### HasMaxPenaltyAmount

`func (o *TermAndCondition) HasMaxPenaltyAmount() bool`

HasMaxPenaltyAmount returns a boolean if a field has been set.

### GetDescriptions

`func (o *TermAndCondition) GetDescriptions() []Description`

GetDescriptions returns the Descriptions field if non-nil, zero value otherwise.

### GetDescriptionsOk

`func (o *TermAndCondition) GetDescriptionsOk() (*[]Description, bool)`

GetDescriptionsOk returns a tuple with the Descriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptions

`func (o *TermAndCondition) SetDescriptions(v []Description)`

SetDescriptions sets Descriptions field to given value.

### HasDescriptions

`func (o *TermAndCondition) HasDescriptions() bool`

HasDescriptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


