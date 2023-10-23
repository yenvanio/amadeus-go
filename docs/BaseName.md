# BaseName

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstName** | Pointer to **string** | First name. | [optional] 
**LastName** | Pointer to **string** | Last name. | [optional] 
**MiddleName** | Pointer to **string** | Middle name(s), for example \&quot;Lee\&quot; in \&quot;John Lee Smith\&quot;. | [optional] 

## Methods

### NewBaseName

`func NewBaseName() *BaseName`

NewBaseName instantiates a new BaseName object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseNameWithDefaults

`func NewBaseNameWithDefaults() *BaseName`

NewBaseNameWithDefaults instantiates a new BaseName object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstName

`func (o *BaseName) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *BaseName) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *BaseName) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *BaseName) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *BaseName) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *BaseName) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *BaseName) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *BaseName) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetMiddleName

`func (o *BaseName) GetMiddleName() string`

GetMiddleName returns the MiddleName field if non-nil, zero value otherwise.

### GetMiddleNameOk

`func (o *BaseName) GetMiddleNameOk() (*string, bool)`

GetMiddleNameOk returns a tuple with the MiddleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddleName

`func (o *BaseName) SetMiddleName(v string)`

SetMiddleName sets MiddleName field to given value.

### HasMiddleName

`func (o *BaseName) HasMiddleName() bool`

HasMiddleName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


