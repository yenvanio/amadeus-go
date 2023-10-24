# ErrorResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | **string** | Name is the error name. | 
**ErrorDescription** | **string** | A small description about the error | 
**Code** | **int64** | Debug contains debug information. | 
**Title** | **string** | Debug contains debug information. This is usually not available and has to be enabled. | 

## Methods

### NewErrorResponse

`func NewErrorResponse(error_ string, errorDescription string, code int64, title string, ) *ErrorResponse`

NewErrorResponse instantiates a new ErrorResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorResponseWithDefaults

`func NewErrorResponseWithDefaults() *ErrorResponse`

NewErrorResponseWithDefaults instantiates a new ErrorResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *ErrorResponse) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ErrorResponse) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ErrorResponse) SetError(v string)`

SetError sets Error field to given value.


### GetErrorDescription

`func (o *ErrorResponse) GetErrorDescription() string`

GetErrorDescription returns the ErrorDescription field if non-nil, zero value otherwise.

### GetErrorDescriptionOk

`func (o *ErrorResponse) GetErrorDescriptionOk() (*string, bool)`

GetErrorDescriptionOk returns a tuple with the ErrorDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDescription

`func (o *ErrorResponse) SetErrorDescription(v string)`

SetErrorDescription sets ErrorDescription field to given value.


### GetCode

`func (o *ErrorResponse) GetCode() int64`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ErrorResponse) GetCodeOk() (*int64, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ErrorResponse) SetCode(v int64)`

SetCode sets Code field to given value.


### GetTitle

`func (o *ErrorResponse) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ErrorResponse) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ErrorResponse) SetTitle(v string)`

SetTitle sets Title field to given value.



[[Back to Model list]](../../test/README.md#documentation-for-models) [[Back to API list]](../../test/README.md#documentation-for-api-endpoints) [[Back to README]](../../test/README.md)


