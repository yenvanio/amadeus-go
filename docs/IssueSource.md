# IssueSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pointer** | Pointer to **string** | a JSON Pointer [RFC6901] to the associated entity in the request document | [optional] 
**Parameter** | Pointer to **string** | a string indicating which URI query parameter caused the issue | [optional] 
**Example** | Pointer to **string** | a string indicating an example of the right value | [optional] 

## Methods

### NewIssueSource

`func NewIssueSource() *IssueSource`

NewIssueSource instantiates a new IssueSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssueSourceWithDefaults

`func NewIssueSourceWithDefaults() *IssueSource`

NewIssueSourceWithDefaults instantiates a new IssueSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPointer

`func (o *IssueSource) GetPointer() string`

GetPointer returns the Pointer field if non-nil, zero value otherwise.

### GetPointerOk

`func (o *IssueSource) GetPointerOk() (*string, bool)`

GetPointerOk returns a tuple with the Pointer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPointer

`func (o *IssueSource) SetPointer(v string)`

SetPointer sets Pointer field to given value.

### HasPointer

`func (o *IssueSource) HasPointer() bool`

HasPointer returns a boolean if a field has been set.

### GetParameter

`func (o *IssueSource) GetParameter() string`

GetParameter returns the Parameter field if non-nil, zero value otherwise.

### GetParameterOk

`func (o *IssueSource) GetParameterOk() (*string, bool)`

GetParameterOk returns a tuple with the Parameter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameter

`func (o *IssueSource) SetParameter(v string)`

SetParameter sets Parameter field to given value.

### HasParameter

`func (o *IssueSource) HasParameter() bool`

HasParameter returns a boolean if a field has been set.

### GetExample

`func (o *IssueSource) GetExample() string`

GetExample returns the Example field if non-nil, zero value otherwise.

### GetExampleOk

`func (o *IssueSource) GetExampleOk() (*string, bool)`

GetExampleOk returns a tuple with the Example field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExample

`func (o *IssueSource) SetExample(v string)`

SetExample sets Example field to given value.

### HasExample

`func (o *IssueSource) HasExample() bool`

HasExample returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


