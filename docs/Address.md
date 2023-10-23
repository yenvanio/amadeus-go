# Address

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Lines** | Pointer to **[]string** | Line 1 &#x3D; Street address, Line 2 &#x3D; Apartment, suite, unit, building, floor, etc | [optional] 
**PostalCode** | Pointer to **string** | Example: 74130 | [optional] 
**CountryCode** | Pointer to **string** | country code [ISO 3166-1 country code](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) | [optional] 
**CityName** | Pointer to **string** | Full city name. Example: Dublin | [optional] 
**StateName** | Pointer to **string** | Full state name | [optional] 
**PostalBox** | Pointer to **string** | E.g. BP 220 | [optional] 

## Methods

### NewAddress

`func NewAddress() *Address`

NewAddress instantiates a new Address object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressWithDefaults

`func NewAddressWithDefaults() *Address`

NewAddressWithDefaults instantiates a new Address object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLines

`func (o *Address) GetLines() []string`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *Address) GetLinesOk() (*[]string, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *Address) SetLines(v []string)`

SetLines sets Lines field to given value.

### HasLines

`func (o *Address) HasLines() bool`

HasLines returns a boolean if a field has been set.

### GetPostalCode

`func (o *Address) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *Address) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *Address) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *Address) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### GetCountryCode

`func (o *Address) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *Address) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *Address) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *Address) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetCityName

`func (o *Address) GetCityName() string`

GetCityName returns the CityName field if non-nil, zero value otherwise.

### GetCityNameOk

`func (o *Address) GetCityNameOk() (*string, bool)`

GetCityNameOk returns a tuple with the CityName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCityName

`func (o *Address) SetCityName(v string)`

SetCityName sets CityName field to given value.

### HasCityName

`func (o *Address) HasCityName() bool`

HasCityName returns a boolean if a field has been set.

### GetStateName

`func (o *Address) GetStateName() string`

GetStateName returns the StateName field if non-nil, zero value otherwise.

### GetStateNameOk

`func (o *Address) GetStateNameOk() (*string, bool)`

GetStateNameOk returns a tuple with the StateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateName

`func (o *Address) SetStateName(v string)`

SetStateName sets StateName field to given value.

### HasStateName

`func (o *Address) HasStateName() bool`

HasStateName returns a boolean if a field has been set.

### GetPostalBox

`func (o *Address) GetPostalBox() string`

GetPostalBox returns the PostalBox field if non-nil, zero value otherwise.

### GetPostalBoxOk

`func (o *Address) GetPostalBoxOk() (*string, bool)`

GetPostalBoxOk returns a tuple with the PostalBox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalBox

`func (o *Address) SetPostalBox(v string)`

SetPostalBox sets PostalBox field to given value.

### HasPostalBox

`func (o *Address) HasPostalBox() bool`

HasPostalBox returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


