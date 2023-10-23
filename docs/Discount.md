# Discount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubType** | Pointer to [**DiscountType**](DiscountType.md) |  | [optional] 
**CityName** | Pointer to **string** | city of residence | [optional] 
**TravelerType** | Pointer to [**DiscountTravelerType**](DiscountTravelerType.md) |  | [optional] 
**CardNumber** | Pointer to **string** | resident card number | [optional] 
**CertificateNumber** | Pointer to **string** | resident certificate number | [optional] 

## Methods

### NewDiscount

`func NewDiscount() *Discount`

NewDiscount instantiates a new Discount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiscountWithDefaults

`func NewDiscountWithDefaults() *Discount`

NewDiscountWithDefaults instantiates a new Discount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubType

`func (o *Discount) GetSubType() DiscountType`

GetSubType returns the SubType field if non-nil, zero value otherwise.

### GetSubTypeOk

`func (o *Discount) GetSubTypeOk() (*DiscountType, bool)`

GetSubTypeOk returns a tuple with the SubType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubType

`func (o *Discount) SetSubType(v DiscountType)`

SetSubType sets SubType field to given value.

### HasSubType

`func (o *Discount) HasSubType() bool`

HasSubType returns a boolean if a field has been set.

### GetCityName

`func (o *Discount) GetCityName() string`

GetCityName returns the CityName field if non-nil, zero value otherwise.

### GetCityNameOk

`func (o *Discount) GetCityNameOk() (*string, bool)`

GetCityNameOk returns a tuple with the CityName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCityName

`func (o *Discount) SetCityName(v string)`

SetCityName sets CityName field to given value.

### HasCityName

`func (o *Discount) HasCityName() bool`

HasCityName returns a boolean if a field has been set.

### GetTravelerType

`func (o *Discount) GetTravelerType() DiscountTravelerType`

GetTravelerType returns the TravelerType field if non-nil, zero value otherwise.

### GetTravelerTypeOk

`func (o *Discount) GetTravelerTypeOk() (*DiscountTravelerType, bool)`

GetTravelerTypeOk returns a tuple with the TravelerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTravelerType

`func (o *Discount) SetTravelerType(v DiscountTravelerType)`

SetTravelerType sets TravelerType field to given value.

### HasTravelerType

`func (o *Discount) HasTravelerType() bool`

HasTravelerType returns a boolean if a field has been set.

### GetCardNumber

`func (o *Discount) GetCardNumber() string`

GetCardNumber returns the CardNumber field if non-nil, zero value otherwise.

### GetCardNumberOk

`func (o *Discount) GetCardNumberOk() (*string, bool)`

GetCardNumberOk returns a tuple with the CardNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardNumber

`func (o *Discount) SetCardNumber(v string)`

SetCardNumber sets CardNumber field to given value.

### HasCardNumber

`func (o *Discount) HasCardNumber() bool`

HasCardNumber returns a boolean if a field has been set.

### GetCertificateNumber

`func (o *Discount) GetCertificateNumber() string`

GetCertificateNumber returns the CertificateNumber field if non-nil, zero value otherwise.

### GetCertificateNumberOk

`func (o *Discount) GetCertificateNumberOk() (*string, bool)`

GetCertificateNumberOk returns a tuple with the CertificateNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateNumber

`func (o *Discount) SetCertificateNumber(v string)`

SetCertificateNumber sets CertificateNumber field to given value.

### HasCertificateNumber

`func (o *Discount) HasCertificateNumber() bool`

HasCertificateNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


