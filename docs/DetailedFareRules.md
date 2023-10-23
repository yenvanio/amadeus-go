# DetailedFareRules

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FareBasis** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**FareNotes** | Pointer to [**TermAndCondition**](TermAndCondition.md) |  | [optional] 
**SegmentId** | Pointer to **string** | Id of the segment concerned by the fare rule | [optional] 

## Methods

### NewDetailedFareRules

`func NewDetailedFareRules() *DetailedFareRules`

NewDetailedFareRules instantiates a new DetailedFareRules object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDetailedFareRulesWithDefaults

`func NewDetailedFareRulesWithDefaults() *DetailedFareRules`

NewDetailedFareRulesWithDefaults instantiates a new DetailedFareRules object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFareBasis

`func (o *DetailedFareRules) GetFareBasis() string`

GetFareBasis returns the FareBasis field if non-nil, zero value otherwise.

### GetFareBasisOk

`func (o *DetailedFareRules) GetFareBasisOk() (*string, bool)`

GetFareBasisOk returns a tuple with the FareBasis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFareBasis

`func (o *DetailedFareRules) SetFareBasis(v string)`

SetFareBasis sets FareBasis field to given value.

### HasFareBasis

`func (o *DetailedFareRules) HasFareBasis() bool`

HasFareBasis returns a boolean if a field has been set.

### GetName

`func (o *DetailedFareRules) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DetailedFareRules) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DetailedFareRules) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DetailedFareRules) HasName() bool`

HasName returns a boolean if a field has been set.

### GetFareNotes

`func (o *DetailedFareRules) GetFareNotes() TermAndCondition`

GetFareNotes returns the FareNotes field if non-nil, zero value otherwise.

### GetFareNotesOk

`func (o *DetailedFareRules) GetFareNotesOk() (*TermAndCondition, bool)`

GetFareNotesOk returns a tuple with the FareNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFareNotes

`func (o *DetailedFareRules) SetFareNotes(v TermAndCondition)`

SetFareNotes sets FareNotes field to given value.

### HasFareNotes

`func (o *DetailedFareRules) HasFareNotes() bool`

HasFareNotes returns a boolean if a field has been set.

### GetSegmentId

`func (o *DetailedFareRules) GetSegmentId() string`

GetSegmentId returns the SegmentId field if non-nil, zero value otherwise.

### GetSegmentIdOk

`func (o *DetailedFareRules) GetSegmentIdOk() (*string, bool)`

GetSegmentIdOk returns a tuple with the SegmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegmentId

`func (o *DetailedFareRules) SetSegmentId(v string)`

SetSegmentId sets SegmentId field to given value.

### HasSegmentId

`func (o *DetailedFareRules) HasSegmentId() bool`

HasSegmentId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


