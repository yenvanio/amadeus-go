# Dictionaries

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Locations** | Pointer to [**map[string]LocationValue**](LocationValue.md) |  | [optional] 
**Aircraft** | Pointer to **map[string]string** |  | [optional] 
**Currencies** | Pointer to **map[string]string** |  | [optional] 
**Carriers** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewDictionaries

`func NewDictionaries() *Dictionaries`

NewDictionaries instantiates a new Dictionaries object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDictionariesWithDefaults

`func NewDictionariesWithDefaults() *Dictionaries`

NewDictionariesWithDefaults instantiates a new Dictionaries object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocations

`func (o *Dictionaries) GetLocations() map[string]LocationValue`

GetLocations returns the Locations field if non-nil, zero value otherwise.

### GetLocationsOk

`func (o *Dictionaries) GetLocationsOk() (*map[string]LocationValue, bool)`

GetLocationsOk returns a tuple with the Locations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocations

`func (o *Dictionaries) SetLocations(v map[string]LocationValue)`

SetLocations sets Locations field to given value.

### HasLocations

`func (o *Dictionaries) HasLocations() bool`

HasLocations returns a boolean if a field has been set.

### GetAircraft

`func (o *Dictionaries) GetAircraft() map[string]string`

GetAircraft returns the Aircraft field if non-nil, zero value otherwise.

### GetAircraftOk

`func (o *Dictionaries) GetAircraftOk() (*map[string]string, bool)`

GetAircraftOk returns a tuple with the Aircraft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAircraft

`func (o *Dictionaries) SetAircraft(v map[string]string)`

SetAircraft sets Aircraft field to given value.

### HasAircraft

`func (o *Dictionaries) HasAircraft() bool`

HasAircraft returns a boolean if a field has been set.

### GetCurrencies

`func (o *Dictionaries) GetCurrencies() map[string]string`

GetCurrencies returns the Currencies field if non-nil, zero value otherwise.

### GetCurrenciesOk

`func (o *Dictionaries) GetCurrenciesOk() (*map[string]string, bool)`

GetCurrenciesOk returns a tuple with the Currencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencies

`func (o *Dictionaries) SetCurrencies(v map[string]string)`

SetCurrencies sets Currencies field to given value.

### HasCurrencies

`func (o *Dictionaries) HasCurrencies() bool`

HasCurrencies returns a boolean if a field has been set.

### GetCarriers

`func (o *Dictionaries) GetCarriers() map[string]string`

GetCarriers returns the Carriers field if non-nil, zero value otherwise.

### GetCarriersOk

`func (o *Dictionaries) GetCarriersOk() (*map[string]string, bool)`

GetCarriersOk returns a tuple with the Carriers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarriers

`func (o *Dictionaries) SetCarriers(v map[string]string)`

SetCarriers sets Carriers field to given value.

### HasCarriers

`func (o *Dictionaries) HasCarriers() bool`

HasCarriers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


