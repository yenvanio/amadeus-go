# Traveler

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | item identifier | [optional] 
**DateOfBirth** | Pointer to **string** | The date of birth in ISO 8601 format (yyyy-mm-dd) | [optional] 
**Gender** | Pointer to [**StakeholderGender**](StakeholderGender.md) |  | [optional] 
**Name** | Pointer to [**Name**](Name.md) |  | [optional] 
**Documents** | Pointer to [**[]IdentityDocument**](IdentityDocument.md) | Advanced Passenger Information - regulatory identity documents - SSR DOCS &amp; DOCO elements | [optional] 
**EmergencyContact** | Pointer to [**EmergencyContact**](EmergencyContact.md) |  | [optional] 
**LoyaltyPrograms** | Pointer to [**[]LoyaltyProgram**](LoyaltyProgram.md) | list of loyalty program followed by the traveler | [optional] 
**DiscountEligibility** | Pointer to [**[]Discount**](Discount.md) | list of element that allow a discount. | [optional] 
**Contact** | Pointer to [**Contact**](Contact.md) |  | [optional] 

## Methods

### NewTraveler

`func NewTraveler() *Traveler`

NewTraveler instantiates a new Traveler object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTravelerWithDefaults

`func NewTravelerWithDefaults() *Traveler`

NewTravelerWithDefaults instantiates a new Traveler object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Traveler) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Traveler) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Traveler) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Traveler) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDateOfBirth

`func (o *Traveler) GetDateOfBirth() string`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *Traveler) GetDateOfBirthOk() (*string, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *Traveler) SetDateOfBirth(v string)`

SetDateOfBirth sets DateOfBirth field to given value.

### HasDateOfBirth

`func (o *Traveler) HasDateOfBirth() bool`

HasDateOfBirth returns a boolean if a field has been set.

### GetGender

`func (o *Traveler) GetGender() StakeholderGender`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *Traveler) GetGenderOk() (*StakeholderGender, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *Traveler) SetGender(v StakeholderGender)`

SetGender sets Gender field to given value.

### HasGender

`func (o *Traveler) HasGender() bool`

HasGender returns a boolean if a field has been set.

### GetName

`func (o *Traveler) GetName() Name`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Traveler) GetNameOk() (*Name, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Traveler) SetName(v Name)`

SetName sets Name field to given value.

### HasName

`func (o *Traveler) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDocuments

`func (o *Traveler) GetDocuments() []IdentityDocument`

GetDocuments returns the Documents field if non-nil, zero value otherwise.

### GetDocumentsOk

`func (o *Traveler) GetDocumentsOk() (*[]IdentityDocument, bool)`

GetDocumentsOk returns a tuple with the Documents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocuments

`func (o *Traveler) SetDocuments(v []IdentityDocument)`

SetDocuments sets Documents field to given value.

### HasDocuments

`func (o *Traveler) HasDocuments() bool`

HasDocuments returns a boolean if a field has been set.

### GetEmergencyContact

`func (o *Traveler) GetEmergencyContact() EmergencyContact`

GetEmergencyContact returns the EmergencyContact field if non-nil, zero value otherwise.

### GetEmergencyContactOk

`func (o *Traveler) GetEmergencyContactOk() (*EmergencyContact, bool)`

GetEmergencyContactOk returns a tuple with the EmergencyContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmergencyContact

`func (o *Traveler) SetEmergencyContact(v EmergencyContact)`

SetEmergencyContact sets EmergencyContact field to given value.

### HasEmergencyContact

`func (o *Traveler) HasEmergencyContact() bool`

HasEmergencyContact returns a boolean if a field has been set.

### GetLoyaltyPrograms

`func (o *Traveler) GetLoyaltyPrograms() []LoyaltyProgram`

GetLoyaltyPrograms returns the LoyaltyPrograms field if non-nil, zero value otherwise.

### GetLoyaltyProgramsOk

`func (o *Traveler) GetLoyaltyProgramsOk() (*[]LoyaltyProgram, bool)`

GetLoyaltyProgramsOk returns a tuple with the LoyaltyPrograms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoyaltyPrograms

`func (o *Traveler) SetLoyaltyPrograms(v []LoyaltyProgram)`

SetLoyaltyPrograms sets LoyaltyPrograms field to given value.

### HasLoyaltyPrograms

`func (o *Traveler) HasLoyaltyPrograms() bool`

HasLoyaltyPrograms returns a boolean if a field has been set.

### GetDiscountEligibility

`func (o *Traveler) GetDiscountEligibility() []Discount`

GetDiscountEligibility returns the DiscountEligibility field if non-nil, zero value otherwise.

### GetDiscountEligibilityOk

`func (o *Traveler) GetDiscountEligibilityOk() (*[]Discount, bool)`

GetDiscountEligibilityOk returns a tuple with the DiscountEligibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountEligibility

`func (o *Traveler) SetDiscountEligibility(v []Discount)`

SetDiscountEligibility sets DiscountEligibility field to given value.

### HasDiscountEligibility

`func (o *Traveler) HasDiscountEligibility() bool`

HasDiscountEligibility returns a boolean if a field has been set.

### GetContact

`func (o *Traveler) GetContact() Contact`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *Traveler) GetContactOk() (*Contact, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *Traveler) SetContact(v Contact)`

SetContact sets Contact field to given value.

### HasContact

`func (o *Traveler) HasContact() bool`

HasContact returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


