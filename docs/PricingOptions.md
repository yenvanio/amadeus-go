# PricingOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FareType** | Pointer to **[]string** | type of fare of the flight-offer | [optional] 
**IncludedCheckedBagsOnly** | Pointer to **bool** | If true, returns the flight-offers with included checked bags only | [optional] 
**RefundableFare** | Pointer to **bool** | If true, returns the flight-offers with refundable fares only | [optional] 
**NoRestrictionFare** | Pointer to **bool** | If true, returns the flight-offers with no restriction fares only | [optional] 
**NoPenaltyFare** | Pointer to **bool** | If true, returns the flight-offers with no penalty fares only | [optional] 

## Methods

### NewPricingOptions

`func NewPricingOptions() *PricingOptions`

NewPricingOptions instantiates a new PricingOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPricingOptionsWithDefaults

`func NewPricingOptionsWithDefaults() *PricingOptions`

NewPricingOptionsWithDefaults instantiates a new PricingOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFareType

`func (o *PricingOptions) GetFareType() []string`

GetFareType returns the FareType field if non-nil, zero value otherwise.

### GetFareTypeOk

`func (o *PricingOptions) GetFareTypeOk() (*[]string, bool)`

GetFareTypeOk returns a tuple with the FareType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFareType

`func (o *PricingOptions) SetFareType(v []string)`

SetFareType sets FareType field to given value.

### HasFareType

`func (o *PricingOptions) HasFareType() bool`

HasFareType returns a boolean if a field has been set.

### GetIncludedCheckedBagsOnly

`func (o *PricingOptions) GetIncludedCheckedBagsOnly() bool`

GetIncludedCheckedBagsOnly returns the IncludedCheckedBagsOnly field if non-nil, zero value otherwise.

### GetIncludedCheckedBagsOnlyOk

`func (o *PricingOptions) GetIncludedCheckedBagsOnlyOk() (*bool, bool)`

GetIncludedCheckedBagsOnlyOk returns a tuple with the IncludedCheckedBagsOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedCheckedBagsOnly

`func (o *PricingOptions) SetIncludedCheckedBagsOnly(v bool)`

SetIncludedCheckedBagsOnly sets IncludedCheckedBagsOnly field to given value.

### HasIncludedCheckedBagsOnly

`func (o *PricingOptions) HasIncludedCheckedBagsOnly() bool`

HasIncludedCheckedBagsOnly returns a boolean if a field has been set.

### GetRefundableFare

`func (o *PricingOptions) GetRefundableFare() bool`

GetRefundableFare returns the RefundableFare field if non-nil, zero value otherwise.

### GetRefundableFareOk

`func (o *PricingOptions) GetRefundableFareOk() (*bool, bool)`

GetRefundableFareOk returns a tuple with the RefundableFare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundableFare

`func (o *PricingOptions) SetRefundableFare(v bool)`

SetRefundableFare sets RefundableFare field to given value.

### HasRefundableFare

`func (o *PricingOptions) HasRefundableFare() bool`

HasRefundableFare returns a boolean if a field has been set.

### GetNoRestrictionFare

`func (o *PricingOptions) GetNoRestrictionFare() bool`

GetNoRestrictionFare returns the NoRestrictionFare field if non-nil, zero value otherwise.

### GetNoRestrictionFareOk

`func (o *PricingOptions) GetNoRestrictionFareOk() (*bool, bool)`

GetNoRestrictionFareOk returns a tuple with the NoRestrictionFare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoRestrictionFare

`func (o *PricingOptions) SetNoRestrictionFare(v bool)`

SetNoRestrictionFare sets NoRestrictionFare field to given value.

### HasNoRestrictionFare

`func (o *PricingOptions) HasNoRestrictionFare() bool`

HasNoRestrictionFare returns a boolean if a field has been set.

### GetNoPenaltyFare

`func (o *PricingOptions) GetNoPenaltyFare() bool`

GetNoPenaltyFare returns the NoPenaltyFare field if non-nil, zero value otherwise.

### GetNoPenaltyFareOk

`func (o *PricingOptions) GetNoPenaltyFareOk() (*bool, bool)`

GetNoPenaltyFareOk returns a tuple with the NoPenaltyFare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoPenaltyFare

`func (o *PricingOptions) SetNoPenaltyFare(v bool)`

SetNoPenaltyFare sets NoPenaltyFare field to given value.

### HasNoPenaltyFare

`func (o *PricingOptions) HasNoPenaltyFare() bool`

HasNoPenaltyFare returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


