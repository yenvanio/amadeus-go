# EmergencyContact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddresseeName** | Pointer to **string** | Adressee name (e.g. in case of emergency purpose it corresponds to name of the person to be contacted). | [optional] 
**CountryCode** | Pointer to **string** | Country code of the country (ISO3166-1). E.g. \&quot;US\&quot; for the United States | [optional] 
**Number** | Pointer to **string** | Phone number. Composed of digits only. The number of digits depends on the country. | [optional] 
**Text** | Pointer to **string** | additional details | [optional] 

## Methods

### NewEmergencyContact

`func NewEmergencyContact() *EmergencyContact`

NewEmergencyContact instantiates a new EmergencyContact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmergencyContactWithDefaults

`func NewEmergencyContactWithDefaults() *EmergencyContact`

NewEmergencyContactWithDefaults instantiates a new EmergencyContact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddresseeName

`func (o *EmergencyContact) GetAddresseeName() string`

GetAddresseeName returns the AddresseeName field if non-nil, zero value otherwise.

### GetAddresseeNameOk

`func (o *EmergencyContact) GetAddresseeNameOk() (*string, bool)`

GetAddresseeNameOk returns a tuple with the AddresseeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresseeName

`func (o *EmergencyContact) SetAddresseeName(v string)`

SetAddresseeName sets AddresseeName field to given value.

### HasAddresseeName

`func (o *EmergencyContact) HasAddresseeName() bool`

HasAddresseeName returns a boolean if a field has been set.

### GetCountryCode

`func (o *EmergencyContact) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *EmergencyContact) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *EmergencyContact) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *EmergencyContact) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetNumber

`func (o *EmergencyContact) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *EmergencyContact) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *EmergencyContact) SetNumber(v string)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *EmergencyContact) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetText

`func (o *EmergencyContact) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *EmergencyContact) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *EmergencyContact) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *EmergencyContact) HasText() bool`

HasText returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


