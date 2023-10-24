# AmadeusOAuth2Token

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The access token issued by the authorization server. | [optional] 
**Username** | Pointer to **string** | The user who requested the access_token | [optional] 
**ApplicationName** | Pointer to **string** | The application which is requested the access_token | [optional] 
**ClientId** | Pointer to **string** | The client_id is a public identifier for apps | [optional] 
**TokenType** | Pointer to **string** | token_type is a parameter in Access Token generate call to Authorization server, which essentially represents how an access_token will be generated and presented for resource access calls | [optional] 
**AccessToken** | Pointer to **string** | Access tokens are a String which applications use to make API requests on behalf of a user. | [optional] 
**ExpiresIn** | Pointer to **int64** | The lifetime in seconds of the access token | [optional] 
**State** | Pointer to **string** | The state | [optional] 
**Scope** | Pointer to **string** | Scope is a mechanism in OAuth 2.0 to limit an application&#39;s access to a user&#39;s account | [optional] 

## Methods

### NewAmadeusOAuth2Token

`func NewAmadeusOAuth2Token() *AmadeusOAuth2Token`

NewAmadeusOAuth2Token instantiates a new AmadeusOAuth2Token object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAmadeusOAuth2TokenWithDefaults

`func NewAmadeusOAuth2TokenWithDefaults() *AmadeusOAuth2Token`

NewAmadeusOAuth2TokenWithDefaults instantiates a new AmadeusOAuth2Token object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AmadeusOAuth2Token) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AmadeusOAuth2Token) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AmadeusOAuth2Token) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AmadeusOAuth2Token) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUsername

`func (o *AmadeusOAuth2Token) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *AmadeusOAuth2Token) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *AmadeusOAuth2Token) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *AmadeusOAuth2Token) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetApplicationName

`func (o *AmadeusOAuth2Token) GetApplicationName() string`

GetApplicationName returns the ApplicationName field if non-nil, zero value otherwise.

### GetApplicationNameOk

`func (o *AmadeusOAuth2Token) GetApplicationNameOk() (*string, bool)`

GetApplicationNameOk returns a tuple with the ApplicationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationName

`func (o *AmadeusOAuth2Token) SetApplicationName(v string)`

SetApplicationName sets ApplicationName field to given value.

### HasApplicationName

`func (o *AmadeusOAuth2Token) HasApplicationName() bool`

HasApplicationName returns a boolean if a field has been set.

### GetClientId

`func (o *AmadeusOAuth2Token) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *AmadeusOAuth2Token) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *AmadeusOAuth2Token) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *AmadeusOAuth2Token) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetTokenType

`func (o *AmadeusOAuth2Token) GetTokenType() string`

GetTokenType returns the TokenType field if non-nil, zero value otherwise.

### GetTokenTypeOk

`func (o *AmadeusOAuth2Token) GetTokenTypeOk() (*string, bool)`

GetTokenTypeOk returns a tuple with the TokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenType

`func (o *AmadeusOAuth2Token) SetTokenType(v string)`

SetTokenType sets TokenType field to given value.

### HasTokenType

`func (o *AmadeusOAuth2Token) HasTokenType() bool`

HasTokenType returns a boolean if a field has been set.

### GetAccessToken

`func (o *AmadeusOAuth2Token) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *AmadeusOAuth2Token) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *AmadeusOAuth2Token) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.

### HasAccessToken

`func (o *AmadeusOAuth2Token) HasAccessToken() bool`

HasAccessToken returns a boolean if a field has been set.

### GetExpiresIn

`func (o *AmadeusOAuth2Token) GetExpiresIn() int64`

GetExpiresIn returns the ExpiresIn field if non-nil, zero value otherwise.

### GetExpiresInOk

`func (o *AmadeusOAuth2Token) GetExpiresInOk() (*int64, bool)`

GetExpiresInOk returns a tuple with the ExpiresIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresIn

`func (o *AmadeusOAuth2Token) SetExpiresIn(v int64)`

SetExpiresIn sets ExpiresIn field to given value.

### HasExpiresIn

`func (o *AmadeusOAuth2Token) HasExpiresIn() bool`

HasExpiresIn returns a boolean if a field has been set.

### GetState

`func (o *AmadeusOAuth2Token) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AmadeusOAuth2Token) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AmadeusOAuth2Token) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *AmadeusOAuth2Token) HasState() bool`

HasState returns a boolean if a field has been set.

### GetScope

`func (o *AmadeusOAuth2Token) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *AmadeusOAuth2Token) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *AmadeusOAuth2Token) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *AmadeusOAuth2Token) HasScope() bool`

HasScope returns a boolean if a field has been set.


[[Back to Model list]](../../test/README.md#documentation-for-models) [[Back to API list]](../../test/README.md#documentation-for-api-endpoints) [[Back to README]](../../test/README.md)


