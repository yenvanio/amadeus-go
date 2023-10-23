# Error500

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | [**[]Issue**](Issue.md) |  | 

## Methods

### NewError500

`func NewError500(errors []Issue, ) *Error500`

NewError500 instantiates a new Error500 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewError500WithDefaults

`func NewError500WithDefaults() *Error500`

NewError500WithDefaults instantiates a new Error500 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrors

`func (o *Error500) GetErrors() []Issue`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *Error500) GetErrorsOk() (*[]Issue, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *Error500) SetErrors(v []Issue)`

SetErrors sets Errors field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


