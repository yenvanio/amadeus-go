# Contact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddresseeName** | Pointer to [**Name**](Name.md) |  | [optional] 
**Address** | Pointer to [**Address**](Address.md) |  | [optional] 
**Language** | Pointer to **string** | the preferred language of communication with this Contact | [optional] 
**Purpose** | Pointer to [**ContactPurpose**](ContactPurpose.md) |  | [optional] 
**Phones** | Pointer to [**[]Phone**](Phone.md) | Phone numbers | [optional] 
**CompanyName** | Pointer to **string** | Name of the company | [optional] 
**EmailAddress** | Pointer to **string** | Email address (e.g. john@smith.com) | [optional] 

## Methods

### NewContact

`func NewContact() *Contact`

NewContact instantiates a new Contact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContactWithDefaults

`func NewContactWithDefaults() *Contact`

NewContactWithDefaults instantiates a new Contact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddresseeName

`func (o *Contact) GetAddresseeName() Name`

GetAddresseeName returns the AddresseeName field if non-nil, zero value otherwise.

### GetAddresseeNameOk

`func (o *Contact) GetAddresseeNameOk() (*Name, bool)`

GetAddresseeNameOk returns a tuple with the AddresseeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresseeName

`func (o *Contact) SetAddresseeName(v Name)`

SetAddresseeName sets AddresseeName field to given value.

### HasAddresseeName

`func (o *Contact) HasAddresseeName() bool`

HasAddresseeName returns a boolean if a field has been set.

### GetAddress

`func (o *Contact) GetAddress() Address`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *Contact) GetAddressOk() (*Address, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *Contact) SetAddress(v Address)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *Contact) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetLanguage

`func (o *Contact) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *Contact) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *Contact) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *Contact) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetPurpose

`func (o *Contact) GetPurpose() ContactPurpose`

GetPurpose returns the Purpose field if non-nil, zero value otherwise.

### GetPurposeOk

`func (o *Contact) GetPurposeOk() (*ContactPurpose, bool)`

GetPurposeOk returns a tuple with the Purpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurpose

`func (o *Contact) SetPurpose(v ContactPurpose)`

SetPurpose sets Purpose field to given value.

### HasPurpose

`func (o *Contact) HasPurpose() bool`

HasPurpose returns a boolean if a field has been set.

### GetPhones

`func (o *Contact) GetPhones() []Phone`

GetPhones returns the Phones field if non-nil, zero value otherwise.

### GetPhonesOk

`func (o *Contact) GetPhonesOk() (*[]Phone, bool)`

GetPhonesOk returns a tuple with the Phones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhones

`func (o *Contact) SetPhones(v []Phone)`

SetPhones sets Phones field to given value.

### HasPhones

`func (o *Contact) HasPhones() bool`

HasPhones returns a boolean if a field has been set.

### GetCompanyName

`func (o *Contact) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *Contact) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *Contact) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *Contact) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### GetEmailAddress

`func (o *Contact) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *Contact) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *Contact) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.

### HasEmailAddress

`func (o *Contact) HasEmailAddress() bool`

HasEmailAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


