/*
Flight Offers Price

Before using this API, we recommend you read our **[Authorization Guide](https://developers.amadeus.com/self-service/apis-docs/guides/authorization-262)** for more information on how to generate an access token.   Please also be aware that our test environment is based on a subset of the production, if you are not returning any results try with big cities/airports like LON (London) or NYC (New-York).

API version: 1.3.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"
	"github.com/yenvanio/amadeus-go/amadeus/utils"
)

// checks if the Contact type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Contact{}

// Contact contact information
type Contact struct {
	AddresseeName *Name `json:"addresseeName,omitempty"`
	Address *Address    `json:"address,omitempty"`
	// the preferred language of communication with this Contact
	Language *string        `json:"language,omitempty"`
	Purpose *ContactPurpose `json:"purpose,omitempty"`
	// Phone numbers
	Phones []Phone `json:"phones,omitempty"`
	// Name of the company
	CompanyName *string `json:"companyName,omitempty"`
	// Email address (e.g. john@smith.com)
	EmailAddress *string `json:"emailAddress,omitempty"`
}

// NewContact instantiates a new Contact object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContact() *Contact {
	this := Contact{}
	return &this
}

// NewContactWithDefaults instantiates a new Contact object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContactWithDefaults() *Contact {
	this := Contact{}
	return &this
}

// GetAddresseeName returns the AddresseeName field value if set, zero value otherwise.
func (o *Contact) GetAddresseeName() Name {
	if o == nil || utils.IsNil(o.AddresseeName) {
		var ret Name
		return ret
	}
	return *o.AddresseeName
}

// GetAddresseeNameOk returns a tuple with the AddresseeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Contact) GetAddresseeNameOk() (*Name, bool) {
	if o == nil || utils.IsNil(o.AddresseeName) {
		return nil, false
	}
	return o.AddresseeName, true
}

// HasAddresseeName returns a boolean if a field has been set.
func (o *Contact) HasAddresseeName() bool {
	if o != nil && !utils.IsNil(o.AddresseeName) {
		return true
	}

	return false
}

// SetAddresseeName gets a reference to the given Name and assigns it to the AddresseeName field.
func (o *Contact) SetAddresseeName(v Name) {
	o.AddresseeName = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *Contact) GetAddress() Address {
	if o == nil || utils.IsNil(o.Address) {
		var ret Address
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Contact) GetAddressOk() (*Address, bool) {
	if o == nil || utils.IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *Contact) HasAddress() bool {
	if o != nil && !utils.IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given Address and assigns it to the Address field.
func (o *Contact) SetAddress(v Address) {
	o.Address = &v
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *Contact) GetLanguage() string {
	if o == nil || utils.IsNil(o.Language) {
		var ret string
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Contact) GetLanguageOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Language) {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *Contact) HasLanguage() bool {
	if o != nil && !utils.IsNil(o.Language) {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given string and assigns it to the Language field.
func (o *Contact) SetLanguage(v string) {
	o.Language = &v
}

// GetPurpose returns the Purpose field value if set, zero value otherwise.
func (o *Contact) GetPurpose() ContactPurpose {
	if o == nil || utils.IsNil(o.Purpose) {
		var ret ContactPurpose
		return ret
	}
	return *o.Purpose
}

// GetPurposeOk returns a tuple with the Purpose field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Contact) GetPurposeOk() (*ContactPurpose, bool) {
	if o == nil || utils.IsNil(o.Purpose) {
		return nil, false
	}
	return o.Purpose, true
}

// HasPurpose returns a boolean if a field has been set.
func (o *Contact) HasPurpose() bool {
	if o != nil && !utils.IsNil(o.Purpose) {
		return true
	}

	return false
}

// SetPurpose gets a reference to the given ContactPurpose and assigns it to the Purpose field.
func (o *Contact) SetPurpose(v ContactPurpose) {
	o.Purpose = &v
}

// GetPhones returns the Phones field value if set, zero value otherwise.
func (o *Contact) GetPhones() []Phone {
	if o == nil || utils.IsNil(o.Phones) {
		var ret []Phone
		return ret
	}
	return o.Phones
}

// GetPhonesOk returns a tuple with the Phones field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Contact) GetPhonesOk() ([]Phone, bool) {
	if o == nil || utils.IsNil(o.Phones) {
		return nil, false
	}
	return o.Phones, true
}

// HasPhones returns a boolean if a field has been set.
func (o *Contact) HasPhones() bool {
	if o != nil && !utils.IsNil(o.Phones) {
		return true
	}

	return false
}

// SetPhones gets a reference to the given []Phone and assigns it to the Phones field.
func (o *Contact) SetPhones(v []Phone) {
	o.Phones = v
}

// GetCompanyName returns the CompanyName field value if set, zero value otherwise.
func (o *Contact) GetCompanyName() string {
	if o == nil || utils.IsNil(o.CompanyName) {
		var ret string
		return ret
	}
	return *o.CompanyName
}

// GetCompanyNameOk returns a tuple with the CompanyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Contact) GetCompanyNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.CompanyName) {
		return nil, false
	}
	return o.CompanyName, true
}

// HasCompanyName returns a boolean if a field has been set.
func (o *Contact) HasCompanyName() bool {
	if o != nil && !utils.IsNil(o.CompanyName) {
		return true
	}

	return false
}

// SetCompanyName gets a reference to the given string and assigns it to the CompanyName field.
func (o *Contact) SetCompanyName(v string) {
	o.CompanyName = &v
}

// GetEmailAddress returns the EmailAddress field value if set, zero value otherwise.
func (o *Contact) GetEmailAddress() string {
	if o == nil || utils.IsNil(o.EmailAddress) {
		var ret string
		return ret
	}
	return *o.EmailAddress
}

// GetEmailAddressOk returns a tuple with the EmailAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Contact) GetEmailAddressOk() (*string, bool) {
	if o == nil || utils.IsNil(o.EmailAddress) {
		return nil, false
	}
	return o.EmailAddress, true
}

// HasEmailAddress returns a boolean if a field has been set.
func (o *Contact) HasEmailAddress() bool {
	if o != nil && !utils.IsNil(o.EmailAddress) {
		return true
	}

	return false
}

// SetEmailAddress gets a reference to the given string and assigns it to the EmailAddress field.
func (o *Contact) SetEmailAddress(v string) {
	o.EmailAddress = &v
}

func (o Contact) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Contact) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.AddresseeName) {
		toSerialize["addresseeName"] = o.AddresseeName
	}
	if !utils.IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !utils.IsNil(o.Language) {
		toSerialize["language"] = o.Language
	}
	if !utils.IsNil(o.Purpose) {
		toSerialize["purpose"] = o.Purpose
	}
	if !utils.IsNil(o.Phones) {
		toSerialize["phones"] = o.Phones
	}
	if !utils.IsNil(o.CompanyName) {
		toSerialize["companyName"] = o.CompanyName
	}
	if !utils.IsNil(o.EmailAddress) {
		toSerialize["emailAddress"] = o.EmailAddress
	}
	return toSerialize, nil
}

type NullableContact struct {
	value *Contact
	isSet bool
}

func (v NullableContact) Get() *Contact {
	return v.value
}

func (v *NullableContact) Set(val *Contact) {
	v.value = val
	v.isSet = true
}

func (v NullableContact) IsSet() bool {
	return v.isSet
}

func (v *NullableContact) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContact(val *Contact) *NullableContact {
	return &NullableContact{value: val, isSet: true}
}

func (v NullableContact) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContact) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


