# PassengerConditions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TravelerId** | Pointer to **string** | Id of the traveler | [optional] 
**GenderRequired** | Pointer to **bool** | If true, the gender is required for the concerned traveler for the creation of the flight-order | [optional] 
**DocumentRequired** | Pointer to **bool** | If true, a document is required for the concerned traveler for the creation of the flight-order | [optional] 
**DocumentIssuanceCityRequired** | Pointer to **bool** | If true, the isuance city of the document is required for the concerned traveler for the creation of the flight-order | [optional] 
**DateOfBirthRequired** | Pointer to **bool** | If true, the date of bitrth is required for the concerned traveler for the creation of the flight-order | [optional] 
**RedressRequiredIfAny** | Pointer to **bool** | If true, the redress is required if any for the concerned traveler for the creation of the flight-order | [optional] 
**AirFranceDiscountRequired** | Pointer to **bool** | If true, the Air France discount is required for the concerned traveler for the creation of the flight-order | [optional] 
**SpanishResidentDiscountRequired** | Pointer to **bool** | If true, the Spanish resident discount is required for the concerned traveler for the creation of the flight-order | [optional] 
**ResidenceRequired** | Pointer to **bool** | If true, the address is required for the concerned traveler for the creation of the flight-order | [optional] 

## Methods

### NewPassengerConditions

`func NewPassengerConditions() *PassengerConditions`

NewPassengerConditions instantiates a new PassengerConditions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPassengerConditionsWithDefaults

`func NewPassengerConditionsWithDefaults() *PassengerConditions`

NewPassengerConditionsWithDefaults instantiates a new PassengerConditions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTravelerId

`func (o *PassengerConditions) GetTravelerId() string`

GetTravelerId returns the TravelerId field if non-nil, zero value otherwise.

### GetTravelerIdOk

`func (o *PassengerConditions) GetTravelerIdOk() (*string, bool)`

GetTravelerIdOk returns a tuple with the TravelerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTravelerId

`func (o *PassengerConditions) SetTravelerId(v string)`

SetTravelerId sets TravelerId field to given value.

### HasTravelerId

`func (o *PassengerConditions) HasTravelerId() bool`

HasTravelerId returns a boolean if a field has been set.

### GetGenderRequired

`func (o *PassengerConditions) GetGenderRequired() bool`

GetGenderRequired returns the GenderRequired field if non-nil, zero value otherwise.

### GetGenderRequiredOk

`func (o *PassengerConditions) GetGenderRequiredOk() (*bool, bool)`

GetGenderRequiredOk returns a tuple with the GenderRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenderRequired

`func (o *PassengerConditions) SetGenderRequired(v bool)`

SetGenderRequired sets GenderRequired field to given value.

### HasGenderRequired

`func (o *PassengerConditions) HasGenderRequired() bool`

HasGenderRequired returns a boolean if a field has been set.

### GetDocumentRequired

`func (o *PassengerConditions) GetDocumentRequired() bool`

GetDocumentRequired returns the DocumentRequired field if non-nil, zero value otherwise.

### GetDocumentRequiredOk

`func (o *PassengerConditions) GetDocumentRequiredOk() (*bool, bool)`

GetDocumentRequiredOk returns a tuple with the DocumentRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentRequired

`func (o *PassengerConditions) SetDocumentRequired(v bool)`

SetDocumentRequired sets DocumentRequired field to given value.

### HasDocumentRequired

`func (o *PassengerConditions) HasDocumentRequired() bool`

HasDocumentRequired returns a boolean if a field has been set.

### GetDocumentIssuanceCityRequired

`func (o *PassengerConditions) GetDocumentIssuanceCityRequired() bool`

GetDocumentIssuanceCityRequired returns the DocumentIssuanceCityRequired field if non-nil, zero value otherwise.

### GetDocumentIssuanceCityRequiredOk

`func (o *PassengerConditions) GetDocumentIssuanceCityRequiredOk() (*bool, bool)`

GetDocumentIssuanceCityRequiredOk returns a tuple with the DocumentIssuanceCityRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentIssuanceCityRequired

`func (o *PassengerConditions) SetDocumentIssuanceCityRequired(v bool)`

SetDocumentIssuanceCityRequired sets DocumentIssuanceCityRequired field to given value.

### HasDocumentIssuanceCityRequired

`func (o *PassengerConditions) HasDocumentIssuanceCityRequired() bool`

HasDocumentIssuanceCityRequired returns a boolean if a field has been set.

### GetDateOfBirthRequired

`func (o *PassengerConditions) GetDateOfBirthRequired() bool`

GetDateOfBirthRequired returns the DateOfBirthRequired field if non-nil, zero value otherwise.

### GetDateOfBirthRequiredOk

`func (o *PassengerConditions) GetDateOfBirthRequiredOk() (*bool, bool)`

GetDateOfBirthRequiredOk returns a tuple with the DateOfBirthRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirthRequired

`func (o *PassengerConditions) SetDateOfBirthRequired(v bool)`

SetDateOfBirthRequired sets DateOfBirthRequired field to given value.

### HasDateOfBirthRequired

`func (o *PassengerConditions) HasDateOfBirthRequired() bool`

HasDateOfBirthRequired returns a boolean if a field has been set.

### GetRedressRequiredIfAny

`func (o *PassengerConditions) GetRedressRequiredIfAny() bool`

GetRedressRequiredIfAny returns the RedressRequiredIfAny field if non-nil, zero value otherwise.

### GetRedressRequiredIfAnyOk

`func (o *PassengerConditions) GetRedressRequiredIfAnyOk() (*bool, bool)`

GetRedressRequiredIfAnyOk returns a tuple with the RedressRequiredIfAny field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedressRequiredIfAny

`func (o *PassengerConditions) SetRedressRequiredIfAny(v bool)`

SetRedressRequiredIfAny sets RedressRequiredIfAny field to given value.

### HasRedressRequiredIfAny

`func (o *PassengerConditions) HasRedressRequiredIfAny() bool`

HasRedressRequiredIfAny returns a boolean if a field has been set.

### GetAirFranceDiscountRequired

`func (o *PassengerConditions) GetAirFranceDiscountRequired() bool`

GetAirFranceDiscountRequired returns the AirFranceDiscountRequired field if non-nil, zero value otherwise.

### GetAirFranceDiscountRequiredOk

`func (o *PassengerConditions) GetAirFranceDiscountRequiredOk() (*bool, bool)`

GetAirFranceDiscountRequiredOk returns a tuple with the AirFranceDiscountRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAirFranceDiscountRequired

`func (o *PassengerConditions) SetAirFranceDiscountRequired(v bool)`

SetAirFranceDiscountRequired sets AirFranceDiscountRequired field to given value.

### HasAirFranceDiscountRequired

`func (o *PassengerConditions) HasAirFranceDiscountRequired() bool`

HasAirFranceDiscountRequired returns a boolean if a field has been set.

### GetSpanishResidentDiscountRequired

`func (o *PassengerConditions) GetSpanishResidentDiscountRequired() bool`

GetSpanishResidentDiscountRequired returns the SpanishResidentDiscountRequired field if non-nil, zero value otherwise.

### GetSpanishResidentDiscountRequiredOk

`func (o *PassengerConditions) GetSpanishResidentDiscountRequiredOk() (*bool, bool)`

GetSpanishResidentDiscountRequiredOk returns a tuple with the SpanishResidentDiscountRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpanishResidentDiscountRequired

`func (o *PassengerConditions) SetSpanishResidentDiscountRequired(v bool)`

SetSpanishResidentDiscountRequired sets SpanishResidentDiscountRequired field to given value.

### HasSpanishResidentDiscountRequired

`func (o *PassengerConditions) HasSpanishResidentDiscountRequired() bool`

HasSpanishResidentDiscountRequired returns a boolean if a field has been set.

### GetResidenceRequired

`func (o *PassengerConditions) GetResidenceRequired() bool`

GetResidenceRequired returns the ResidenceRequired field if non-nil, zero value otherwise.

### GetResidenceRequiredOk

`func (o *PassengerConditions) GetResidenceRequiredOk() (*bool, bool)`

GetResidenceRequiredOk returns a tuple with the ResidenceRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResidenceRequired

`func (o *PassengerConditions) SetResidenceRequired(v bool)`

SetResidenceRequired sets ResidenceRequired field to given value.

### HasResidenceRequired

`func (o *PassengerConditions) HasResidenceRequired() bool`

HasResidenceRequired returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


