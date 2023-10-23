# Issue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **int32** | the HTTP status code applicable to this error | [optional] 
**Code** | Pointer to **int64** | an application-specific error code | [optional] 
**Title** | Pointer to **string** | a short summary of the error | [optional] 
**Detail** | Pointer to **string** | explanation of the error | [optional] 
**Source** | Pointer to [**IssueSource**](IssueSource.md) |  | [optional] 

## Methods

### NewIssue

`func NewIssue() *Issue`

NewIssue instantiates a new Issue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssueWithDefaults

`func NewIssueWithDefaults() *Issue`

NewIssueWithDefaults instantiates a new Issue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *Issue) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Issue) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Issue) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Issue) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCode

`func (o *Issue) GetCode() int64`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *Issue) GetCodeOk() (*int64, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *Issue) SetCode(v int64)`

SetCode sets Code field to given value.

### HasCode

`func (o *Issue) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetTitle

`func (o *Issue) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Issue) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Issue) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Issue) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDetail

`func (o *Issue) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *Issue) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *Issue) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *Issue) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### GetSource

`func (o *Issue) GetSource() IssueSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *Issue) GetSourceOk() (*IssueSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *Issue) SetSource(v IssueSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *Issue) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


