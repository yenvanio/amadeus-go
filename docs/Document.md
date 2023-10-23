# Document

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Number** | Pointer to **string** | The document number (shown on the document) . E.g. QFU514563221J | [optional] 
**IssuanceDate** | Pointer to **string** | Date at which the document has been issued. | [optional] 
**ExpiryDate** | Pointer to **string** | Date after which the document is not valid anymore. | [optional] 
**IssuanceCountry** | Pointer to **string** | [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) of the country that issued the document | [optional] 
**IssuanceLocation** | Pointer to **string** | A more precise information concerning the place where the document has been issued, when available. It may be a country, a state, a city or any other type of location. e.g. New-York | [optional] 
**Nationality** | Pointer to **string** | [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) of the nationality appearing on the document | [optional] 
**BirthPlace** | Pointer to **string** | Birth place as indicated on the document | [optional] 

## Methods

### NewDocument

`func NewDocument() *Document`

NewDocument instantiates a new Document object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumentWithDefaults

`func NewDocumentWithDefaults() *Document`

NewDocumentWithDefaults instantiates a new Document object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumber

`func (o *Document) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *Document) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *Document) SetNumber(v string)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *Document) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetIssuanceDate

`func (o *Document) GetIssuanceDate() string`

GetIssuanceDate returns the IssuanceDate field if non-nil, zero value otherwise.

### GetIssuanceDateOk

`func (o *Document) GetIssuanceDateOk() (*string, bool)`

GetIssuanceDateOk returns a tuple with the IssuanceDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuanceDate

`func (o *Document) SetIssuanceDate(v string)`

SetIssuanceDate sets IssuanceDate field to given value.

### HasIssuanceDate

`func (o *Document) HasIssuanceDate() bool`

HasIssuanceDate returns a boolean if a field has been set.

### GetExpiryDate

`func (o *Document) GetExpiryDate() string`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *Document) GetExpiryDateOk() (*string, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDate

`func (o *Document) SetExpiryDate(v string)`

SetExpiryDate sets ExpiryDate field to given value.

### HasExpiryDate

`func (o *Document) HasExpiryDate() bool`

HasExpiryDate returns a boolean if a field has been set.

### GetIssuanceCountry

`func (o *Document) GetIssuanceCountry() string`

GetIssuanceCountry returns the IssuanceCountry field if non-nil, zero value otherwise.

### GetIssuanceCountryOk

`func (o *Document) GetIssuanceCountryOk() (*string, bool)`

GetIssuanceCountryOk returns a tuple with the IssuanceCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuanceCountry

`func (o *Document) SetIssuanceCountry(v string)`

SetIssuanceCountry sets IssuanceCountry field to given value.

### HasIssuanceCountry

`func (o *Document) HasIssuanceCountry() bool`

HasIssuanceCountry returns a boolean if a field has been set.

### GetIssuanceLocation

`func (o *Document) GetIssuanceLocation() string`

GetIssuanceLocation returns the IssuanceLocation field if non-nil, zero value otherwise.

### GetIssuanceLocationOk

`func (o *Document) GetIssuanceLocationOk() (*string, bool)`

GetIssuanceLocationOk returns a tuple with the IssuanceLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuanceLocation

`func (o *Document) SetIssuanceLocation(v string)`

SetIssuanceLocation sets IssuanceLocation field to given value.

### HasIssuanceLocation

`func (o *Document) HasIssuanceLocation() bool`

HasIssuanceLocation returns a boolean if a field has been set.

### GetNationality

`func (o *Document) GetNationality() string`

GetNationality returns the Nationality field if non-nil, zero value otherwise.

### GetNationalityOk

`func (o *Document) GetNationalityOk() (*string, bool)`

GetNationalityOk returns a tuple with the Nationality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationality

`func (o *Document) SetNationality(v string)`

SetNationality sets Nationality field to given value.

### HasNationality

`func (o *Document) HasNationality() bool`

HasNationality returns a boolean if a field has been set.

### GetBirthPlace

`func (o *Document) GetBirthPlace() string`

GetBirthPlace returns the BirthPlace field if non-nil, zero value otherwise.

### GetBirthPlaceOk

`func (o *Document) GetBirthPlaceOk() (*string, bool)`

GetBirthPlaceOk returns a tuple with the BirthPlace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthPlace

`func (o *Document) SetBirthPlace(v string)`

SetBirthPlace sets BirthPlace field to given value.

### HasBirthPlace

`func (o *Document) HasBirthPlace() bool`

HasBirthPlace returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


