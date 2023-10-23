# BookingRequirements

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InvoiceAddressRequired** | Pointer to **bool** | If true, an invoice address is required for the creation of the flight-order | [optional] 
**MailingAddressRequired** | Pointer to **bool** | If true, a postal address is required for the creation of the flight-order | [optional] 
**EmailAddressRequired** | Pointer to **bool** | If true, an email address is required for the creation of the flight-order | [optional] 
**PhoneCountryCodeRequired** | Pointer to **bool** | If true, the phone country code (e.g. &#39;33&#39;) associated for each phone number is required for the creation of the flight-order | [optional] 
**MobilePhoneNumberRequired** | Pointer to **bool** | If true, a mobile phone number is required for the creation of the flight-order | [optional] 
**PhoneNumberRequired** | Pointer to **bool** | If true, a phone number is required for the creation of the flight-order | [optional] 
**PostalCodeRequired** | Pointer to **bool** | If true, a postal code is required for the creation of the flight-order | [optional] 
**TravelerRequirements** | Pointer to [**[]PassengerConditions**](PassengerConditions.md) | traveler pricing condition | [optional] 

## Methods

### NewBookingRequirements

`func NewBookingRequirements() *BookingRequirements`

NewBookingRequirements instantiates a new BookingRequirements object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBookingRequirementsWithDefaults

`func NewBookingRequirementsWithDefaults() *BookingRequirements`

NewBookingRequirementsWithDefaults instantiates a new BookingRequirements object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvoiceAddressRequired

`func (o *BookingRequirements) GetInvoiceAddressRequired() bool`

GetInvoiceAddressRequired returns the InvoiceAddressRequired field if non-nil, zero value otherwise.

### GetInvoiceAddressRequiredOk

`func (o *BookingRequirements) GetInvoiceAddressRequiredOk() (*bool, bool)`

GetInvoiceAddressRequiredOk returns a tuple with the InvoiceAddressRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceAddressRequired

`func (o *BookingRequirements) SetInvoiceAddressRequired(v bool)`

SetInvoiceAddressRequired sets InvoiceAddressRequired field to given value.

### HasInvoiceAddressRequired

`func (o *BookingRequirements) HasInvoiceAddressRequired() bool`

HasInvoiceAddressRequired returns a boolean if a field has been set.

### GetMailingAddressRequired

`func (o *BookingRequirements) GetMailingAddressRequired() bool`

GetMailingAddressRequired returns the MailingAddressRequired field if non-nil, zero value otherwise.

### GetMailingAddressRequiredOk

`func (o *BookingRequirements) GetMailingAddressRequiredOk() (*bool, bool)`

GetMailingAddressRequiredOk returns a tuple with the MailingAddressRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailingAddressRequired

`func (o *BookingRequirements) SetMailingAddressRequired(v bool)`

SetMailingAddressRequired sets MailingAddressRequired field to given value.

### HasMailingAddressRequired

`func (o *BookingRequirements) HasMailingAddressRequired() bool`

HasMailingAddressRequired returns a boolean if a field has been set.

### GetEmailAddressRequired

`func (o *BookingRequirements) GetEmailAddressRequired() bool`

GetEmailAddressRequired returns the EmailAddressRequired field if non-nil, zero value otherwise.

### GetEmailAddressRequiredOk

`func (o *BookingRequirements) GetEmailAddressRequiredOk() (*bool, bool)`

GetEmailAddressRequiredOk returns a tuple with the EmailAddressRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddressRequired

`func (o *BookingRequirements) SetEmailAddressRequired(v bool)`

SetEmailAddressRequired sets EmailAddressRequired field to given value.

### HasEmailAddressRequired

`func (o *BookingRequirements) HasEmailAddressRequired() bool`

HasEmailAddressRequired returns a boolean if a field has been set.

### GetPhoneCountryCodeRequired

`func (o *BookingRequirements) GetPhoneCountryCodeRequired() bool`

GetPhoneCountryCodeRequired returns the PhoneCountryCodeRequired field if non-nil, zero value otherwise.

### GetPhoneCountryCodeRequiredOk

`func (o *BookingRequirements) GetPhoneCountryCodeRequiredOk() (*bool, bool)`

GetPhoneCountryCodeRequiredOk returns a tuple with the PhoneCountryCodeRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneCountryCodeRequired

`func (o *BookingRequirements) SetPhoneCountryCodeRequired(v bool)`

SetPhoneCountryCodeRequired sets PhoneCountryCodeRequired field to given value.

### HasPhoneCountryCodeRequired

`func (o *BookingRequirements) HasPhoneCountryCodeRequired() bool`

HasPhoneCountryCodeRequired returns a boolean if a field has been set.

### GetMobilePhoneNumberRequired

`func (o *BookingRequirements) GetMobilePhoneNumberRequired() bool`

GetMobilePhoneNumberRequired returns the MobilePhoneNumberRequired field if non-nil, zero value otherwise.

### GetMobilePhoneNumberRequiredOk

`func (o *BookingRequirements) GetMobilePhoneNumberRequiredOk() (*bool, bool)`

GetMobilePhoneNumberRequiredOk returns a tuple with the MobilePhoneNumberRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobilePhoneNumberRequired

`func (o *BookingRequirements) SetMobilePhoneNumberRequired(v bool)`

SetMobilePhoneNumberRequired sets MobilePhoneNumberRequired field to given value.

### HasMobilePhoneNumberRequired

`func (o *BookingRequirements) HasMobilePhoneNumberRequired() bool`

HasMobilePhoneNumberRequired returns a boolean if a field has been set.

### GetPhoneNumberRequired

`func (o *BookingRequirements) GetPhoneNumberRequired() bool`

GetPhoneNumberRequired returns the PhoneNumberRequired field if non-nil, zero value otherwise.

### GetPhoneNumberRequiredOk

`func (o *BookingRequirements) GetPhoneNumberRequiredOk() (*bool, bool)`

GetPhoneNumberRequiredOk returns a tuple with the PhoneNumberRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumberRequired

`func (o *BookingRequirements) SetPhoneNumberRequired(v bool)`

SetPhoneNumberRequired sets PhoneNumberRequired field to given value.

### HasPhoneNumberRequired

`func (o *BookingRequirements) HasPhoneNumberRequired() bool`

HasPhoneNumberRequired returns a boolean if a field has been set.

### GetPostalCodeRequired

`func (o *BookingRequirements) GetPostalCodeRequired() bool`

GetPostalCodeRequired returns the PostalCodeRequired field if non-nil, zero value otherwise.

### GetPostalCodeRequiredOk

`func (o *BookingRequirements) GetPostalCodeRequiredOk() (*bool, bool)`

GetPostalCodeRequiredOk returns a tuple with the PostalCodeRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCodeRequired

`func (o *BookingRequirements) SetPostalCodeRequired(v bool)`

SetPostalCodeRequired sets PostalCodeRequired field to given value.

### HasPostalCodeRequired

`func (o *BookingRequirements) HasPostalCodeRequired() bool`

HasPostalCodeRequired returns a boolean if a field has been set.

### GetTravelerRequirements

`func (o *BookingRequirements) GetTravelerRequirements() []PassengerConditions`

GetTravelerRequirements returns the TravelerRequirements field if non-nil, zero value otherwise.

### GetTravelerRequirementsOk

`func (o *BookingRequirements) GetTravelerRequirementsOk() (*[]PassengerConditions, bool)`

GetTravelerRequirementsOk returns a tuple with the TravelerRequirements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTravelerRequirements

`func (o *BookingRequirements) SetTravelerRequirements(v []PassengerConditions)`

SetTravelerRequirements sets TravelerRequirements field to given value.

### HasTravelerRequirements

`func (o *BookingRequirements) HasTravelerRequirements() bool`

HasTravelerRequirements returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


