# FlightEndPoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IataCode** | Pointer to **string** | [IATA airline codes](http://www.iata.org/publications/Pages/code-search.aspx) | [optional] 
**Terminal** | Pointer to **string** | terminal name / number | [optional] 
**At** | Pointer to **time.Time** | local date and time in [ISO8601](https://en.wikipedia.org/wiki/ISO_8601) YYYY-MM-ddThh:mm:ss format, e.g. 2017-02-10T20:40:00 | [optional] 

## Methods

### NewFlightEndPoint

`func NewFlightEndPoint() *FlightEndPoint`

NewFlightEndPoint instantiates a new FlightEndPoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlightEndPointWithDefaults

`func NewFlightEndPointWithDefaults() *FlightEndPoint`

NewFlightEndPointWithDefaults instantiates a new FlightEndPoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIataCode

`func (o *FlightEndPoint) GetIataCode() string`

GetIataCode returns the IataCode field if non-nil, zero value otherwise.

### GetIataCodeOk

`func (o *FlightEndPoint) GetIataCodeOk() (*string, bool)`

GetIataCodeOk returns a tuple with the IataCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIataCode

`func (o *FlightEndPoint) SetIataCode(v string)`

SetIataCode sets IataCode field to given value.

### HasIataCode

`func (o *FlightEndPoint) HasIataCode() bool`

HasIataCode returns a boolean if a field has been set.

### GetTerminal

`func (o *FlightEndPoint) GetTerminal() string`

GetTerminal returns the Terminal field if non-nil, zero value otherwise.

### GetTerminalOk

`func (o *FlightEndPoint) GetTerminalOk() (*string, bool)`

GetTerminalOk returns a tuple with the Terminal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminal

`func (o *FlightEndPoint) SetTerminal(v string)`

SetTerminal sets Terminal field to given value.

### HasTerminal

`func (o *FlightEndPoint) HasTerminal() bool`

HasTerminal returns a boolean if a field has been set.

### GetAt

`func (o *FlightEndPoint) GetAt() time.Time`

GetAt returns the At field if non-nil, zero value otherwise.

### GetAtOk

`func (o *FlightEndPoint) GetAtOk() (*time.Time, bool)`

GetAtOk returns a tuple with the At field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAt

`func (o *FlightEndPoint) SetAt(v time.Time)`

SetAt sets At field to given value.

### HasAt

`func (o *FlightEndPoint) HasAt() bool`

HasAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


