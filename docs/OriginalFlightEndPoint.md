# OriginalFlightEndPoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IataCode** | Pointer to **string** | [IATA airline codes](http://www.iata.org/publications/Pages/code-search.aspx) | [optional] 
**Terminal** | Pointer to **string** | terminal name / number | [optional] 

## Methods

### NewOriginalFlightEndPoint

`func NewOriginalFlightEndPoint() *OriginalFlightEndPoint`

NewOriginalFlightEndPoint instantiates a new OriginalFlightEndPoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOriginalFlightEndPointWithDefaults

`func NewOriginalFlightEndPointWithDefaults() *OriginalFlightEndPoint`

NewOriginalFlightEndPointWithDefaults instantiates a new OriginalFlightEndPoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIataCode

`func (o *OriginalFlightEndPoint) GetIataCode() string`

GetIataCode returns the IataCode field if non-nil, zero value otherwise.

### GetIataCodeOk

`func (o *OriginalFlightEndPoint) GetIataCodeOk() (*string, bool)`

GetIataCodeOk returns a tuple with the IataCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIataCode

`func (o *OriginalFlightEndPoint) SetIataCode(v string)`

SetIataCode sets IataCode field to given value.

### HasIataCode

`func (o *OriginalFlightEndPoint) HasIataCode() bool`

HasIataCode returns a boolean if a field has been set.

### GetTerminal

`func (o *OriginalFlightEndPoint) GetTerminal() string`

GetTerminal returns the Terminal field if non-nil, zero value otherwise.

### GetTerminalOk

`func (o *OriginalFlightEndPoint) GetTerminalOk() (*string, bool)`

GetTerminalOk returns a tuple with the Terminal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminal

`func (o *OriginalFlightEndPoint) SetTerminal(v string)`

SetTerminal sets Terminal field to given value.

### HasTerminal

`func (o *OriginalFlightEndPoint) HasTerminal() bool`

HasTerminal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


