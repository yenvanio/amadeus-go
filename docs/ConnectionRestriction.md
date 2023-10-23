# ConnectionRestriction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxNumberOfConnections** | Pointer to **float32** | The maximal number of connections for each itinerary. Value can be 0, 1 or 2. | [optional] 
**NonStopPreferred** | Pointer to **bool** | When this option is requested, recommendations made of Non-Stop flights only are favoured by the search, on the whole itinerary, with a weight of 75%. It is possible to override this default 75% weight of Non-Stop recommendations by providing the requested weight in the attribute nonStopPreferredWeight | [optional] 
**NonStopPreferredWeight** | Pointer to **float32** | This weight is the percentage of total number of recommendations that could be returned for Non-Stop flight recommendations. It is only relevant when combined with nonSTopPreferred set to true. If not specified, the default weight of 75% is applied | [optional] 
**AirportChangeAllowed** | Pointer to **bool** | Allow to change airport during connection | [optional] 
**TechnicalStopsAllowed** | Pointer to **bool** | This option allows the single segment to have one or more intermediate stops (technical stops). | [optional] 

## Methods

### NewConnectionRestriction

`func NewConnectionRestriction() *ConnectionRestriction`

NewConnectionRestriction instantiates a new ConnectionRestriction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionRestrictionWithDefaults

`func NewConnectionRestrictionWithDefaults() *ConnectionRestriction`

NewConnectionRestrictionWithDefaults instantiates a new ConnectionRestriction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxNumberOfConnections

`func (o *ConnectionRestriction) GetMaxNumberOfConnections() float32`

GetMaxNumberOfConnections returns the MaxNumberOfConnections field if non-nil, zero value otherwise.

### GetMaxNumberOfConnectionsOk

`func (o *ConnectionRestriction) GetMaxNumberOfConnectionsOk() (*float32, bool)`

GetMaxNumberOfConnectionsOk returns a tuple with the MaxNumberOfConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNumberOfConnections

`func (o *ConnectionRestriction) SetMaxNumberOfConnections(v float32)`

SetMaxNumberOfConnections sets MaxNumberOfConnections field to given value.

### HasMaxNumberOfConnections

`func (o *ConnectionRestriction) HasMaxNumberOfConnections() bool`

HasMaxNumberOfConnections returns a boolean if a field has been set.

### GetNonStopPreferred

`func (o *ConnectionRestriction) GetNonStopPreferred() bool`

GetNonStopPreferred returns the NonStopPreferred field if non-nil, zero value otherwise.

### GetNonStopPreferredOk

`func (o *ConnectionRestriction) GetNonStopPreferredOk() (*bool, bool)`

GetNonStopPreferredOk returns a tuple with the NonStopPreferred field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonStopPreferred

`func (o *ConnectionRestriction) SetNonStopPreferred(v bool)`

SetNonStopPreferred sets NonStopPreferred field to given value.

### HasNonStopPreferred

`func (o *ConnectionRestriction) HasNonStopPreferred() bool`

HasNonStopPreferred returns a boolean if a field has been set.

### GetNonStopPreferredWeight

`func (o *ConnectionRestriction) GetNonStopPreferredWeight() float32`

GetNonStopPreferredWeight returns the NonStopPreferredWeight field if non-nil, zero value otherwise.

### GetNonStopPreferredWeightOk

`func (o *ConnectionRestriction) GetNonStopPreferredWeightOk() (*float32, bool)`

GetNonStopPreferredWeightOk returns a tuple with the NonStopPreferredWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonStopPreferredWeight

`func (o *ConnectionRestriction) SetNonStopPreferredWeight(v float32)`

SetNonStopPreferredWeight sets NonStopPreferredWeight field to given value.

### HasNonStopPreferredWeight

`func (o *ConnectionRestriction) HasNonStopPreferredWeight() bool`

HasNonStopPreferredWeight returns a boolean if a field has been set.

### GetAirportChangeAllowed

`func (o *ConnectionRestriction) GetAirportChangeAllowed() bool`

GetAirportChangeAllowed returns the AirportChangeAllowed field if non-nil, zero value otherwise.

### GetAirportChangeAllowedOk

`func (o *ConnectionRestriction) GetAirportChangeAllowedOk() (*bool, bool)`

GetAirportChangeAllowedOk returns a tuple with the AirportChangeAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAirportChangeAllowed

`func (o *ConnectionRestriction) SetAirportChangeAllowed(v bool)`

SetAirportChangeAllowed sets AirportChangeAllowed field to given value.

### HasAirportChangeAllowed

`func (o *ConnectionRestriction) HasAirportChangeAllowed() bool`

HasAirportChangeAllowed returns a boolean if a field has been set.

### GetTechnicalStopsAllowed

`func (o *ConnectionRestriction) GetTechnicalStopsAllowed() bool`

GetTechnicalStopsAllowed returns the TechnicalStopsAllowed field if non-nil, zero value otherwise.

### GetTechnicalStopsAllowedOk

`func (o *ConnectionRestriction) GetTechnicalStopsAllowedOk() (*bool, bool)`

GetTechnicalStopsAllowedOk returns a tuple with the TechnicalStopsAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTechnicalStopsAllowed

`func (o *ConnectionRestriction) SetTechnicalStopsAllowed(v bool)`

SetTechnicalStopsAllowed sets TechnicalStopsAllowed field to given value.

### HasTechnicalStopsAllowed

`func (o *ConnectionRestriction) HasTechnicalStopsAllowed() bool`

HasTechnicalStopsAllowed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


