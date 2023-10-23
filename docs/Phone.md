# Phone

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceType** | Pointer to [**PhoneDeviceType**](PhoneDeviceType.md) |  | [optional] 
**CountryCallingCode** | Pointer to **string** | Country calling code of the phone number, as defined by the International Communication Union. Examples - \&quot;1\&quot; for US, \&quot;371\&quot; for Latvia. | [optional] 
**Number** | Pointer to **string** | Phone number. Composed of digits only. The number of digits depends on the country. | [optional] 

## Methods

### NewPhone

`func NewPhone() *Phone`

NewPhone instantiates a new Phone object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPhoneWithDefaults

`func NewPhoneWithDefaults() *Phone`

NewPhoneWithDefaults instantiates a new Phone object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceType

`func (o *Phone) GetDeviceType() PhoneDeviceType`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *Phone) GetDeviceTypeOk() (*PhoneDeviceType, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *Phone) SetDeviceType(v PhoneDeviceType)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *Phone) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### GetCountryCallingCode

`func (o *Phone) GetCountryCallingCode() string`

GetCountryCallingCode returns the CountryCallingCode field if non-nil, zero value otherwise.

### GetCountryCallingCodeOk

`func (o *Phone) GetCountryCallingCodeOk() (*string, bool)`

GetCountryCallingCodeOk returns a tuple with the CountryCallingCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCallingCode

`func (o *Phone) SetCountryCallingCode(v string)`

SetCountryCallingCode sets CountryCallingCode field to given value.

### HasCountryCallingCode

`func (o *Phone) HasCountryCallingCode() bool`

HasCountryCallingCode returns a boolean if a field has been set.

### GetNumber

`func (o *Phone) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *Phone) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *Phone) SetNumber(v string)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *Phone) HasNumber() bool`

HasNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


