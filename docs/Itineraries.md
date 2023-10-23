# Itineraries

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Duration** | Pointer to **string** | duration in [ISO8601](https://en.wikipedia.org/wiki/ISO_8601) PnYnMnDTnHnMnS format, e.g. PT2H10M for a duration of 2h10m | [optional] 
**Segments** | [**[]Segment**](Segment.md) |  | 

## Methods

### NewItineraries

`func NewItineraries(segments []Segment, ) *Itineraries`

NewItineraries instantiates a new Itineraries object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItinerariesWithDefaults

`func NewItinerariesWithDefaults() *Itineraries`

NewItinerariesWithDefaults instantiates a new Itineraries object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDuration

`func (o *Itineraries) GetDuration() string`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *Itineraries) GetDurationOk() (*string, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *Itineraries) SetDuration(v string)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *Itineraries) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetSegments

`func (o *Itineraries) GetSegments() []Segment`

GetSegments returns the Segments field if non-nil, zero value otherwise.

### GetSegmentsOk

`func (o *Itineraries) GetSegmentsOk() (*[]Segment, bool)`

GetSegmentsOk returns a tuple with the Segments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegments

`func (o *Itineraries) SetSegments(v []Segment)`

SetSegments sets Segments field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


