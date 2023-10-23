# AdditionalServicesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChargeableCheckedBags** | Pointer to [**ChargeableCheckdBags**](ChargeableCheckdBags.md) |  | [optional] 
**ChargeableSeat** | Pointer to [**ChargeableSeat**](ChargeableSeat.md) |  | [optional] 
**ChargeableSeatNumber** | Pointer to **string** | DEPRECATED - use the chargeableSeat attribute -  seat number | [optional] 
**OtherServices** | Pointer to [**[]ServiceName**](ServiceName.md) | Other services to add | [optional] 

## Methods

### NewAdditionalServicesRequest

`func NewAdditionalServicesRequest() *AdditionalServicesRequest`

NewAdditionalServicesRequest instantiates a new AdditionalServicesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdditionalServicesRequestWithDefaults

`func NewAdditionalServicesRequestWithDefaults() *AdditionalServicesRequest`

NewAdditionalServicesRequestWithDefaults instantiates a new AdditionalServicesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChargeableCheckedBags

`func (o *AdditionalServicesRequest) GetChargeableCheckedBags() ChargeableCheckdBags`

GetChargeableCheckedBags returns the ChargeableCheckedBags field if non-nil, zero value otherwise.

### GetChargeableCheckedBagsOk

`func (o *AdditionalServicesRequest) GetChargeableCheckedBagsOk() (*ChargeableCheckdBags, bool)`

GetChargeableCheckedBagsOk returns a tuple with the ChargeableCheckedBags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargeableCheckedBags

`func (o *AdditionalServicesRequest) SetChargeableCheckedBags(v ChargeableCheckdBags)`

SetChargeableCheckedBags sets ChargeableCheckedBags field to given value.

### HasChargeableCheckedBags

`func (o *AdditionalServicesRequest) HasChargeableCheckedBags() bool`

HasChargeableCheckedBags returns a boolean if a field has been set.

### GetChargeableSeat

`func (o *AdditionalServicesRequest) GetChargeableSeat() ChargeableSeat`

GetChargeableSeat returns the ChargeableSeat field if non-nil, zero value otherwise.

### GetChargeableSeatOk

`func (o *AdditionalServicesRequest) GetChargeableSeatOk() (*ChargeableSeat, bool)`

GetChargeableSeatOk returns a tuple with the ChargeableSeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargeableSeat

`func (o *AdditionalServicesRequest) SetChargeableSeat(v ChargeableSeat)`

SetChargeableSeat sets ChargeableSeat field to given value.

### HasChargeableSeat

`func (o *AdditionalServicesRequest) HasChargeableSeat() bool`

HasChargeableSeat returns a boolean if a field has been set.

### GetChargeableSeatNumber

`func (o *AdditionalServicesRequest) GetChargeableSeatNumber() string`

GetChargeableSeatNumber returns the ChargeableSeatNumber field if non-nil, zero value otherwise.

### GetChargeableSeatNumberOk

`func (o *AdditionalServicesRequest) GetChargeableSeatNumberOk() (*string, bool)`

GetChargeableSeatNumberOk returns a tuple with the ChargeableSeatNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargeableSeatNumber

`func (o *AdditionalServicesRequest) SetChargeableSeatNumber(v string)`

SetChargeableSeatNumber sets ChargeableSeatNumber field to given value.

### HasChargeableSeatNumber

`func (o *AdditionalServicesRequest) HasChargeableSeatNumber() bool`

HasChargeableSeatNumber returns a boolean if a field has been set.

### GetOtherServices

`func (o *AdditionalServicesRequest) GetOtherServices() []ServiceName`

GetOtherServices returns the OtherServices field if non-nil, zero value otherwise.

### GetOtherServicesOk

`func (o *AdditionalServicesRequest) GetOtherServicesOk() (*[]ServiceName, bool)`

GetOtherServicesOk returns a tuple with the OtherServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherServices

`func (o *AdditionalServicesRequest) SetOtherServices(v []ServiceName)`

SetOtherServices sets OtherServices field to given value.

### HasOtherServices

`func (o *AdditionalServicesRequest) HasOtherServices() bool`

HasOtherServices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


