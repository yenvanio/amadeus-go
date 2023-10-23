# OriginDestinationLight

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**OriginLocationCode** | Pointer to **string** | Origin location, such as a city or an airport. Currently, only the locations defined in [IATA](http://www.iata.org/publications/Pages/code-search.aspx) are supported. | [optional] 
**DestinationLocationCode** | Pointer to **string** | Destination location, such as a city or an airport. Currently, only the locations defined in [IATA](http://www.iata.org/publications/Pages/code-search.aspx) are supported. | [optional] 
**IncludedConnectionPoints** | Pointer to **[]string** | List of included connections points. When an includedViaPoints option is specified, all FlightOffer returned must at least go via this Connecting Point. Currently, only the locations defined in IATA are supported. Used only by the AMADEUS provider | [optional] 
**ExcludedConnectionPoints** | Pointer to **[]string** | List of excluded connections points. Any FlightOffer with these connections points will be present in response. Currently, only the locations defined in IATA are supported. Used only by the AMADEUS provider | [optional] 

## Methods

### NewOriginDestinationLight

`func NewOriginDestinationLight() *OriginDestinationLight`

NewOriginDestinationLight instantiates a new OriginDestinationLight object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOriginDestinationLightWithDefaults

`func NewOriginDestinationLightWithDefaults() *OriginDestinationLight`

NewOriginDestinationLightWithDefaults instantiates a new OriginDestinationLight object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OriginDestinationLight) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OriginDestinationLight) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OriginDestinationLight) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OriginDestinationLight) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOriginLocationCode

`func (o *OriginDestinationLight) GetOriginLocationCode() string`

GetOriginLocationCode returns the OriginLocationCode field if non-nil, zero value otherwise.

### GetOriginLocationCodeOk

`func (o *OriginDestinationLight) GetOriginLocationCodeOk() (*string, bool)`

GetOriginLocationCodeOk returns a tuple with the OriginLocationCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginLocationCode

`func (o *OriginDestinationLight) SetOriginLocationCode(v string)`

SetOriginLocationCode sets OriginLocationCode field to given value.

### HasOriginLocationCode

`func (o *OriginDestinationLight) HasOriginLocationCode() bool`

HasOriginLocationCode returns a boolean if a field has been set.

### GetDestinationLocationCode

`func (o *OriginDestinationLight) GetDestinationLocationCode() string`

GetDestinationLocationCode returns the DestinationLocationCode field if non-nil, zero value otherwise.

### GetDestinationLocationCodeOk

`func (o *OriginDestinationLight) GetDestinationLocationCodeOk() (*string, bool)`

GetDestinationLocationCodeOk returns a tuple with the DestinationLocationCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationLocationCode

`func (o *OriginDestinationLight) SetDestinationLocationCode(v string)`

SetDestinationLocationCode sets DestinationLocationCode field to given value.

### HasDestinationLocationCode

`func (o *OriginDestinationLight) HasDestinationLocationCode() bool`

HasDestinationLocationCode returns a boolean if a field has been set.

### GetIncludedConnectionPoints

`func (o *OriginDestinationLight) GetIncludedConnectionPoints() []string`

GetIncludedConnectionPoints returns the IncludedConnectionPoints field if non-nil, zero value otherwise.

### GetIncludedConnectionPointsOk

`func (o *OriginDestinationLight) GetIncludedConnectionPointsOk() (*[]string, bool)`

GetIncludedConnectionPointsOk returns a tuple with the IncludedConnectionPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedConnectionPoints

`func (o *OriginDestinationLight) SetIncludedConnectionPoints(v []string)`

SetIncludedConnectionPoints sets IncludedConnectionPoints field to given value.

### HasIncludedConnectionPoints

`func (o *OriginDestinationLight) HasIncludedConnectionPoints() bool`

HasIncludedConnectionPoints returns a boolean if a field has been set.

### GetExcludedConnectionPoints

`func (o *OriginDestinationLight) GetExcludedConnectionPoints() []string`

GetExcludedConnectionPoints returns the ExcludedConnectionPoints field if non-nil, zero value otherwise.

### GetExcludedConnectionPointsOk

`func (o *OriginDestinationLight) GetExcludedConnectionPointsOk() (*[]string, bool)`

GetExcludedConnectionPointsOk returns a tuple with the ExcludedConnectionPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedConnectionPoints

`func (o *OriginDestinationLight) SetExcludedConnectionPoints(v []string)`

SetExcludedConnectionPoints sets ExcludedConnectionPoints field to given value.

### HasExcludedConnectionPoints

`func (o *OriginDestinationLight) HasExcludedConnectionPoints() bool`

HasExcludedConnectionPoints returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


