# ContactDictionary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddresseeName** | Pointer to [**Name**](Name.md) |  | [optional] 
**Address** | Pointer to [**Address**](Address.md) |  | [optional] 
**Language** | Pointer to **string** | the preferred language of communication with this Contact | [optional] 
**Purpose** | Pointer to [**ContactPurpose**](ContactPurpose.md) |  | [optional] 

## Methods

### NewContactDictionary

`func NewContactDictionary() *ContactDictionary`

NewContactDictionary instantiates a new ContactDictionary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContactDictionaryWithDefaults

`func NewContactDictionaryWithDefaults() *ContactDictionary`

NewContactDictionaryWithDefaults instantiates a new ContactDictionary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddresseeName

`func (o *ContactDictionary) GetAddresseeName() Name`

GetAddresseeName returns the AddresseeName field if non-nil, zero value otherwise.

### GetAddresseeNameOk

`func (o *ContactDictionary) GetAddresseeNameOk() (*Name, bool)`

GetAddresseeNameOk returns a tuple with the AddresseeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresseeName

`func (o *ContactDictionary) SetAddresseeName(v Name)`

SetAddresseeName sets AddresseeName field to given value.

### HasAddresseeName

`func (o *ContactDictionary) HasAddresseeName() bool`

HasAddresseeName returns a boolean if a field has been set.

### GetAddress

`func (o *ContactDictionary) GetAddress() Address`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *ContactDictionary) GetAddressOk() (*Address, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *ContactDictionary) SetAddress(v Address)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *ContactDictionary) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetLanguage

`func (o *ContactDictionary) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *ContactDictionary) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *ContactDictionary) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *ContactDictionary) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetPurpose

`func (o *ContactDictionary) GetPurpose() ContactPurpose`

GetPurpose returns the Purpose field if non-nil, zero value otherwise.

### GetPurposeOk

`func (o *ContactDictionary) GetPurposeOk() (*ContactPurpose, bool)`

GetPurposeOk returns a tuple with the Purpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurpose

`func (o *ContactDictionary) SetPurpose(v ContactPurpose)`

SetPurpose sets Purpose field to given value.

### HasPurpose

`func (o *ContactDictionary) HasPurpose() bool`

HasPurpose returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


