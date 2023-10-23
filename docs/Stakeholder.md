# Stakeholder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | item identifier | [optional] 
**DateOfBirth** | Pointer to **string** | The date of birth in ISO 8601 format (yyyy-mm-dd) | [optional] 
**Gender** | Pointer to [**StakeholderGender**](StakeholderGender.md) |  | [optional] 
**Name** | Pointer to [**Name**](Name.md) |  | [optional] 
**Documents** | Pointer to [**[]IdentityDocument**](IdentityDocument.md) | Advanced Passenger Information - regulatory identity documents - SSR DOCS &amp; DOCO elements | [optional] 

## Methods

### NewStakeholder

`func NewStakeholder() *Stakeholder`

NewStakeholder instantiates a new Stakeholder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStakeholderWithDefaults

`func NewStakeholderWithDefaults() *Stakeholder`

NewStakeholderWithDefaults instantiates a new Stakeholder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Stakeholder) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Stakeholder) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Stakeholder) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Stakeholder) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDateOfBirth

`func (o *Stakeholder) GetDateOfBirth() string`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *Stakeholder) GetDateOfBirthOk() (*string, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *Stakeholder) SetDateOfBirth(v string)`

SetDateOfBirth sets DateOfBirth field to given value.

### HasDateOfBirth

`func (o *Stakeholder) HasDateOfBirth() bool`

HasDateOfBirth returns a boolean if a field has been set.

### GetGender

`func (o *Stakeholder) GetGender() StakeholderGender`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *Stakeholder) GetGenderOk() (*StakeholderGender, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *Stakeholder) SetGender(v StakeholderGender)`

SetGender sets Gender field to given value.

### HasGender

`func (o *Stakeholder) HasGender() bool`

HasGender returns a boolean if a field has been set.

### GetName

`func (o *Stakeholder) GetName() Name`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Stakeholder) GetNameOk() (*Name, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Stakeholder) SetName(v Name)`

SetName sets Name field to given value.

### HasName

`func (o *Stakeholder) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDocuments

`func (o *Stakeholder) GetDocuments() []IdentityDocument`

GetDocuments returns the Documents field if non-nil, zero value otherwise.

### GetDocumentsOk

`func (o *Stakeholder) GetDocumentsOk() (*[]IdentityDocument, bool)`

GetDocumentsOk returns a tuple with the Documents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocuments

`func (o *Stakeholder) SetDocuments(v []IdentityDocument)`

SetDocuments sets Documents field to given value.

### HasDocuments

`func (o *Stakeholder) HasDocuments() bool`

HasDocuments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


