# LocationValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CityCode** | Pointer to **string** | City code associated to the airport | [optional] 
**CountryCode** | Pointer to **string** | Country code of the airport | [optional] 

## Methods

### NewLocationValue

`func NewLocationValue() *LocationValue`

NewLocationValue instantiates a new LocationValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocationValueWithDefaults

`func NewLocationValueWithDefaults() *LocationValue`

NewLocationValueWithDefaults instantiates a new LocationValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCityCode

`func (o *LocationValue) GetCityCode() string`

GetCityCode returns the CityCode field if non-nil, zero value otherwise.

### GetCityCodeOk

`func (o *LocationValue) GetCityCodeOk() (*string, bool)`

GetCityCodeOk returns a tuple with the CityCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCityCode

`func (o *LocationValue) SetCityCode(v string)`

SetCityCode sets CityCode field to given value.

### HasCityCode

`func (o *LocationValue) HasCityCode() bool`

HasCityCode returns a boolean if a field has been set.

### GetCountryCode

`func (o *LocationValue) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *LocationValue) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *LocationValue) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *LocationValue) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


