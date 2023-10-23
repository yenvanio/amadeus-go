# CarrierRestrictions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlacklistedInEUAllowed** | Pointer to **bool** | This flag enable/disable filtering of blacklisted airline by EU. The list of the banned airlines is published in the Official Journal of the European Union, where they are included as annexes A and B to the Commission Regulation. The blacklist of an airline can concern all its flights or some specific aircraft types pertaining to the airline | [optional] [default to false]
**ExcludedCarrierCodes** | Pointer to **[]string** | This option ensures that the system will only consider these airlines. | [optional] 
**IncludedCarrierCodes** | Pointer to **[]string** | This option ensures that the system will only consider these airlines. | [optional] 

## Methods

### NewCarrierRestrictions

`func NewCarrierRestrictions() *CarrierRestrictions`

NewCarrierRestrictions instantiates a new CarrierRestrictions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCarrierRestrictionsWithDefaults

`func NewCarrierRestrictionsWithDefaults() *CarrierRestrictions`

NewCarrierRestrictionsWithDefaults instantiates a new CarrierRestrictions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlacklistedInEUAllowed

`func (o *CarrierRestrictions) GetBlacklistedInEUAllowed() bool`

GetBlacklistedInEUAllowed returns the BlacklistedInEUAllowed field if non-nil, zero value otherwise.

### GetBlacklistedInEUAllowedOk

`func (o *CarrierRestrictions) GetBlacklistedInEUAllowedOk() (*bool, bool)`

GetBlacklistedInEUAllowedOk returns a tuple with the BlacklistedInEUAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlacklistedInEUAllowed

`func (o *CarrierRestrictions) SetBlacklistedInEUAllowed(v bool)`

SetBlacklistedInEUAllowed sets BlacklistedInEUAllowed field to given value.

### HasBlacklistedInEUAllowed

`func (o *CarrierRestrictions) HasBlacklistedInEUAllowed() bool`

HasBlacklistedInEUAllowed returns a boolean if a field has been set.

### GetExcludedCarrierCodes

`func (o *CarrierRestrictions) GetExcludedCarrierCodes() []string`

GetExcludedCarrierCodes returns the ExcludedCarrierCodes field if non-nil, zero value otherwise.

### GetExcludedCarrierCodesOk

`func (o *CarrierRestrictions) GetExcludedCarrierCodesOk() (*[]string, bool)`

GetExcludedCarrierCodesOk returns a tuple with the ExcludedCarrierCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedCarrierCodes

`func (o *CarrierRestrictions) SetExcludedCarrierCodes(v []string)`

SetExcludedCarrierCodes sets ExcludedCarrierCodes field to given value.

### HasExcludedCarrierCodes

`func (o *CarrierRestrictions) HasExcludedCarrierCodes() bool`

HasExcludedCarrierCodes returns a boolean if a field has been set.

### GetIncludedCarrierCodes

`func (o *CarrierRestrictions) GetIncludedCarrierCodes() []string`

GetIncludedCarrierCodes returns the IncludedCarrierCodes field if non-nil, zero value otherwise.

### GetIncludedCarrierCodesOk

`func (o *CarrierRestrictions) GetIncludedCarrierCodesOk() (*[]string, bool)`

GetIncludedCarrierCodesOk returns a tuple with the IncludedCarrierCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedCarrierCodes

`func (o *CarrierRestrictions) SetIncludedCarrierCodes(v []string)`

SetIncludedCarrierCodes sets IncludedCarrierCodes field to given value.

### HasIncludedCarrierCodes

`func (o *CarrierRestrictions) HasIncludedCarrierCodes() bool`

HasIncludedCarrierCodes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


