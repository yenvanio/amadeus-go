# OriginDestination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**OriginLocationCode** | Pointer to **string** | Origin location, such as a city or an airport. Currently, only the locations defined in [IATA](http://www.iata.org/publications/Pages/code-search.aspx) are supported. | [optional] 
**DestinationLocationCode** | Pointer to **string** | Destination location, such as a city or an airport. Currently, only the locations defined in [IATA](http://www.iata.org/publications/Pages/code-search.aspx) are supported. | [optional] 
**IncludedConnectionPoints** | Pointer to **[]string** | List of included connections points. When an includedViaPoints option is specified, all FlightOffer returned must at least go via this Connecting Point. Currently, only the locations defined in IATA are supported. Used only by the AMADEUS provider | [optional] 
**ExcludedConnectionPoints** | Pointer to **[]string** | List of excluded connections points. Any FlightOffer with these connections points will be present in response. Currently, only the locations defined in IATA are supported. Used only by the AMADEUS provider | [optional] 
**OriginRadius** | Pointer to **float32** | Include other possible locations around the point, located less than this distance in kilometers away. Max:300  Can not be combined with \&quot;dateWindow\&quot; or \&quot;timeWindow\&quot;.  | [optional] 
**AlternativeOriginsCodes** | Pointer to **[]string** | Set of alternative origin location, such as a city or an airport. Currently, only the locations defined in [IATA](http://www.iata.org/publications/Pages/code-search.aspx) are supported. | [optional] 
**DestinationRadius** | Pointer to **float32** | Include other possible locations around the point, located less than this distance in kilometers away. Max:300  Can not be combined with \&quot;dateWindow\&quot; or \&quot;timeWindow\&quot;.  | [optional] 
**AlternativeDestinationsCodes** | Pointer to **[]string** | Set of alternative destination location, such as a city or an airport. Currently, only the locations defined in [IATA](http://www.iata.org/publications/Pages/code-search.aspx) are supported. | [optional] 
**DepartureDateTimeRange** | Pointer to [**DateTimeRange**](DateTimeRange.md) |  | [optional] 
**ArrivalDateTimeRange** | Pointer to [**DateTimeRange**](DateTimeRange.md) |  | [optional] 

## Methods

### NewOriginDestination

`func NewOriginDestination() *OriginDestination`

NewOriginDestination instantiates a new OriginDestination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOriginDestinationWithDefaults

`func NewOriginDestinationWithDefaults() *OriginDestination`

NewOriginDestinationWithDefaults instantiates a new OriginDestination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OriginDestination) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OriginDestination) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OriginDestination) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OriginDestination) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOriginLocationCode

`func (o *OriginDestination) GetOriginLocationCode() string`

GetOriginLocationCode returns the OriginLocationCode field if non-nil, zero value otherwise.

### GetOriginLocationCodeOk

`func (o *OriginDestination) GetOriginLocationCodeOk() (*string, bool)`

GetOriginLocationCodeOk returns a tuple with the OriginLocationCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginLocationCode

`func (o *OriginDestination) SetOriginLocationCode(v string)`

SetOriginLocationCode sets OriginLocationCode field to given value.

### HasOriginLocationCode

`func (o *OriginDestination) HasOriginLocationCode() bool`

HasOriginLocationCode returns a boolean if a field has been set.

### GetDestinationLocationCode

`func (o *OriginDestination) GetDestinationLocationCode() string`

GetDestinationLocationCode returns the DestinationLocationCode field if non-nil, zero value otherwise.

### GetDestinationLocationCodeOk

`func (o *OriginDestination) GetDestinationLocationCodeOk() (*string, bool)`

GetDestinationLocationCodeOk returns a tuple with the DestinationLocationCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationLocationCode

`func (o *OriginDestination) SetDestinationLocationCode(v string)`

SetDestinationLocationCode sets DestinationLocationCode field to given value.

### HasDestinationLocationCode

`func (o *OriginDestination) HasDestinationLocationCode() bool`

HasDestinationLocationCode returns a boolean if a field has been set.

### GetIncludedConnectionPoints

`func (o *OriginDestination) GetIncludedConnectionPoints() []string`

GetIncludedConnectionPoints returns the IncludedConnectionPoints field if non-nil, zero value otherwise.

### GetIncludedConnectionPointsOk

`func (o *OriginDestination) GetIncludedConnectionPointsOk() (*[]string, bool)`

GetIncludedConnectionPointsOk returns a tuple with the IncludedConnectionPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedConnectionPoints

`func (o *OriginDestination) SetIncludedConnectionPoints(v []string)`

SetIncludedConnectionPoints sets IncludedConnectionPoints field to given value.

### HasIncludedConnectionPoints

`func (o *OriginDestination) HasIncludedConnectionPoints() bool`

HasIncludedConnectionPoints returns a boolean if a field has been set.

### GetExcludedConnectionPoints

`func (o *OriginDestination) GetExcludedConnectionPoints() []string`

GetExcludedConnectionPoints returns the ExcludedConnectionPoints field if non-nil, zero value otherwise.

### GetExcludedConnectionPointsOk

`func (o *OriginDestination) GetExcludedConnectionPointsOk() (*[]string, bool)`

GetExcludedConnectionPointsOk returns a tuple with the ExcludedConnectionPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedConnectionPoints

`func (o *OriginDestination) SetExcludedConnectionPoints(v []string)`

SetExcludedConnectionPoints sets ExcludedConnectionPoints field to given value.

### HasExcludedConnectionPoints

`func (o *OriginDestination) HasExcludedConnectionPoints() bool`

HasExcludedConnectionPoints returns a boolean if a field has been set.

### GetOriginRadius

`func (o *OriginDestination) GetOriginRadius() float32`

GetOriginRadius returns the OriginRadius field if non-nil, zero value otherwise.

### GetOriginRadiusOk

`func (o *OriginDestination) GetOriginRadiusOk() (*float32, bool)`

GetOriginRadiusOk returns a tuple with the OriginRadius field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginRadius

`func (o *OriginDestination) SetOriginRadius(v float32)`

SetOriginRadius sets OriginRadius field to given value.

### HasOriginRadius

`func (o *OriginDestination) HasOriginRadius() bool`

HasOriginRadius returns a boolean if a field has been set.

### GetAlternativeOriginsCodes

`func (o *OriginDestination) GetAlternativeOriginsCodes() []string`

GetAlternativeOriginsCodes returns the AlternativeOriginsCodes field if non-nil, zero value otherwise.

### GetAlternativeOriginsCodesOk

`func (o *OriginDestination) GetAlternativeOriginsCodesOk() (*[]string, bool)`

GetAlternativeOriginsCodesOk returns a tuple with the AlternativeOriginsCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternativeOriginsCodes

`func (o *OriginDestination) SetAlternativeOriginsCodes(v []string)`

SetAlternativeOriginsCodes sets AlternativeOriginsCodes field to given value.

### HasAlternativeOriginsCodes

`func (o *OriginDestination) HasAlternativeOriginsCodes() bool`

HasAlternativeOriginsCodes returns a boolean if a field has been set.

### GetDestinationRadius

`func (o *OriginDestination) GetDestinationRadius() float32`

GetDestinationRadius returns the DestinationRadius field if non-nil, zero value otherwise.

### GetDestinationRadiusOk

`func (o *OriginDestination) GetDestinationRadiusOk() (*float32, bool)`

GetDestinationRadiusOk returns a tuple with the DestinationRadius field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationRadius

`func (o *OriginDestination) SetDestinationRadius(v float32)`

SetDestinationRadius sets DestinationRadius field to given value.

### HasDestinationRadius

`func (o *OriginDestination) HasDestinationRadius() bool`

HasDestinationRadius returns a boolean if a field has been set.

### GetAlternativeDestinationsCodes

`func (o *OriginDestination) GetAlternativeDestinationsCodes() []string`

GetAlternativeDestinationsCodes returns the AlternativeDestinationsCodes field if non-nil, zero value otherwise.

### GetAlternativeDestinationsCodesOk

`func (o *OriginDestination) GetAlternativeDestinationsCodesOk() (*[]string, bool)`

GetAlternativeDestinationsCodesOk returns a tuple with the AlternativeDestinationsCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternativeDestinationsCodes

`func (o *OriginDestination) SetAlternativeDestinationsCodes(v []string)`

SetAlternativeDestinationsCodes sets AlternativeDestinationsCodes field to given value.

### HasAlternativeDestinationsCodes

`func (o *OriginDestination) HasAlternativeDestinationsCodes() bool`

HasAlternativeDestinationsCodes returns a boolean if a field has been set.

### GetDepartureDateTimeRange

`func (o *OriginDestination) GetDepartureDateTimeRange() DateTimeRange`

GetDepartureDateTimeRange returns the DepartureDateTimeRange field if non-nil, zero value otherwise.

### GetDepartureDateTimeRangeOk

`func (o *OriginDestination) GetDepartureDateTimeRangeOk() (*DateTimeRange, bool)`

GetDepartureDateTimeRangeOk returns a tuple with the DepartureDateTimeRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartureDateTimeRange

`func (o *OriginDestination) SetDepartureDateTimeRange(v DateTimeRange)`

SetDepartureDateTimeRange sets DepartureDateTimeRange field to given value.

### HasDepartureDateTimeRange

`func (o *OriginDestination) HasDepartureDateTimeRange() bool`

HasDepartureDateTimeRange returns a boolean if a field has been set.

### GetArrivalDateTimeRange

`func (o *OriginDestination) GetArrivalDateTimeRange() DateTimeRange`

GetArrivalDateTimeRange returns the ArrivalDateTimeRange field if non-nil, zero value otherwise.

### GetArrivalDateTimeRangeOk

`func (o *OriginDestination) GetArrivalDateTimeRangeOk() (*DateTimeRange, bool)`

GetArrivalDateTimeRangeOk returns a tuple with the ArrivalDateTimeRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrivalDateTimeRange

`func (o *OriginDestination) SetArrivalDateTimeRange(v DateTimeRange)`

SetArrivalDateTimeRange sets ArrivalDateTimeRange field to given value.

### HasArrivalDateTimeRange

`func (o *OriginDestination) HasArrivalDateTimeRange() bool`

HasArrivalDateTimeRange returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


