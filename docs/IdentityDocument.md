# IdentityDocument

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
**DocumentType** | Pointer to [**DocumentType**](DocumentType.md) |  | [optional] 
**ValidityCountry** | Pointer to **string** | [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) of the country where the document is valid | [optional] 
**BirthCountry** | Pointer to **string** | [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) of the country of birth | [optional] 
**Holder** | Pointer to **bool** | boolean to specify if the traveler is the holder of the document | [optional] 

## Methods

### NewIdentityDocument

`func NewIdentityDocument() *IdentityDocument`

NewIdentityDocument instantiates a new IdentityDocument object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityDocumentWithDefaults

`func NewIdentityDocumentWithDefaults() *IdentityDocument`

NewIdentityDocumentWithDefaults instantiates a new IdentityDocument object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumber

`func (o *IdentityDocument) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *IdentityDocument) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *IdentityDocument) SetNumber(v string)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *IdentityDocument) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetIssuanceDate

`func (o *IdentityDocument) GetIssuanceDate() string`

GetIssuanceDate returns the IssuanceDate field if non-nil, zero value otherwise.

### GetIssuanceDateOk

`func (o *IdentityDocument) GetIssuanceDateOk() (*string, bool)`

GetIssuanceDateOk returns a tuple with the IssuanceDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuanceDate

`func (o *IdentityDocument) SetIssuanceDate(v string)`

SetIssuanceDate sets IssuanceDate field to given value.

### HasIssuanceDate

`func (o *IdentityDocument) HasIssuanceDate() bool`

HasIssuanceDate returns a boolean if a field has been set.

### GetExpiryDate

`func (o *IdentityDocument) GetExpiryDate() string`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *IdentityDocument) GetExpiryDateOk() (*string, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDate

`func (o *IdentityDocument) SetExpiryDate(v string)`

SetExpiryDate sets ExpiryDate field to given value.

### HasExpiryDate

`func (o *IdentityDocument) HasExpiryDate() bool`

HasExpiryDate returns a boolean if a field has been set.

### GetIssuanceCountry

`func (o *IdentityDocument) GetIssuanceCountry() string`

GetIssuanceCountry returns the IssuanceCountry field if non-nil, zero value otherwise.

### GetIssuanceCountryOk

`func (o *IdentityDocument) GetIssuanceCountryOk() (*string, bool)`

GetIssuanceCountryOk returns a tuple with the IssuanceCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuanceCountry

`func (o *IdentityDocument) SetIssuanceCountry(v string)`

SetIssuanceCountry sets IssuanceCountry field to given value.

### HasIssuanceCountry

`func (o *IdentityDocument) HasIssuanceCountry() bool`

HasIssuanceCountry returns a boolean if a field has been set.

### GetIssuanceLocation

`func (o *IdentityDocument) GetIssuanceLocation() string`

GetIssuanceLocation returns the IssuanceLocation field if non-nil, zero value otherwise.

### GetIssuanceLocationOk

`func (o *IdentityDocument) GetIssuanceLocationOk() (*string, bool)`

GetIssuanceLocationOk returns a tuple with the IssuanceLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuanceLocation

`func (o *IdentityDocument) SetIssuanceLocation(v string)`

SetIssuanceLocation sets IssuanceLocation field to given value.

### HasIssuanceLocation

`func (o *IdentityDocument) HasIssuanceLocation() bool`

HasIssuanceLocation returns a boolean if a field has been set.

### GetNationality

`func (o *IdentityDocument) GetNationality() string`

GetNationality returns the Nationality field if non-nil, zero value otherwise.

### GetNationalityOk

`func (o *IdentityDocument) GetNationalityOk() (*string, bool)`

GetNationalityOk returns a tuple with the Nationality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationality

`func (o *IdentityDocument) SetNationality(v string)`

SetNationality sets Nationality field to given value.

### HasNationality

`func (o *IdentityDocument) HasNationality() bool`

HasNationality returns a boolean if a field has been set.

### GetBirthPlace

`func (o *IdentityDocument) GetBirthPlace() string`

GetBirthPlace returns the BirthPlace field if non-nil, zero value otherwise.

### GetBirthPlaceOk

`func (o *IdentityDocument) GetBirthPlaceOk() (*string, bool)`

GetBirthPlaceOk returns a tuple with the BirthPlace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthPlace

`func (o *IdentityDocument) SetBirthPlace(v string)`

SetBirthPlace sets BirthPlace field to given value.

### HasBirthPlace

`func (o *IdentityDocument) HasBirthPlace() bool`

HasBirthPlace returns a boolean if a field has been set.

### GetDocumentType

`func (o *IdentityDocument) GetDocumentType() DocumentType`

GetDocumentType returns the DocumentType field if non-nil, zero value otherwise.

### GetDocumentTypeOk

`func (o *IdentityDocument) GetDocumentTypeOk() (*DocumentType, bool)`

GetDocumentTypeOk returns a tuple with the DocumentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentType

`func (o *IdentityDocument) SetDocumentType(v DocumentType)`

SetDocumentType sets DocumentType field to given value.

### HasDocumentType

`func (o *IdentityDocument) HasDocumentType() bool`

HasDocumentType returns a boolean if a field has been set.

### GetValidityCountry

`func (o *IdentityDocument) GetValidityCountry() string`

GetValidityCountry returns the ValidityCountry field if non-nil, zero value otherwise.

### GetValidityCountryOk

`func (o *IdentityDocument) GetValidityCountryOk() (*string, bool)`

GetValidityCountryOk returns a tuple with the ValidityCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidityCountry

`func (o *IdentityDocument) SetValidityCountry(v string)`

SetValidityCountry sets ValidityCountry field to given value.

### HasValidityCountry

`func (o *IdentityDocument) HasValidityCountry() bool`

HasValidityCountry returns a boolean if a field has been set.

### GetBirthCountry

`func (o *IdentityDocument) GetBirthCountry() string`

GetBirthCountry returns the BirthCountry field if non-nil, zero value otherwise.

### GetBirthCountryOk

`func (o *IdentityDocument) GetBirthCountryOk() (*string, bool)`

GetBirthCountryOk returns a tuple with the BirthCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthCountry

`func (o *IdentityDocument) SetBirthCountry(v string)`

SetBirthCountry sets BirthCountry field to given value.

### HasBirthCountry

`func (o *IdentityDocument) HasBirthCountry() bool`

HasBirthCountry returns a boolean if a field has been set.

### GetHolder

`func (o *IdentityDocument) GetHolder() bool`

GetHolder returns the Holder field if non-nil, zero value otherwise.

### GetHolderOk

`func (o *IdentityDocument) GetHolderOk() (*bool, bool)`

GetHolderOk returns a tuple with the Holder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHolder

`func (o *IdentityDocument) SetHolder(v bool)`

SetHolder sets Holder field to given value.

### HasHolder

`func (o *IdentityDocument) HasHolder() bool`

HasHolder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


