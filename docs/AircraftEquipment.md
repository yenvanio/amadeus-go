# AircraftEquipment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | IATA aircraft code (http://www.flugzeuginfo.net/table_accodes_iata_en.php)  | [optional] 

## Methods

### NewAircraftEquipment

`func NewAircraftEquipment() *AircraftEquipment`

NewAircraftEquipment instantiates a new AircraftEquipment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAircraftEquipmentWithDefaults

`func NewAircraftEquipmentWithDefaults() *AircraftEquipment`

NewAircraftEquipmentWithDefaults instantiates a new AircraftEquipment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *AircraftEquipment) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *AircraftEquipment) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *AircraftEquipment) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *AircraftEquipment) HasCode() bool`

HasCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


