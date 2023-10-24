package authorization

import (
	"bytes"
	"context"
	api "github.com/yenvanio/amadeus-go/amadeus"
	"github.com/yenvanio/amadeus-go/amadeus/models"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type ApiGetOauth2TokenInfoRequest struct {
	ctx         context.Context
	ApiService  *OAuth2AccessTokenAPIService
	accessToken string
}

func (r ApiGetOauth2TokenInfoRequest) Execute() (*models.AmadeusOAuth2Token, *http.Response, error) {
	return r.ApiService.GetOauth2TokenInfoExecute(r)
}

/*
GetOauth2TokenInfo The OAuth 2.0 token info endpoint

Get token information

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accessToken
	@return ApiGetOauth2TokenInfoRequest
*/
func (a *OAuth2AccessTokenAPIService) GetOauth2TokenInfo(ctx context.Context, accessToken string) ApiGetOauth2TokenInfoRequest {
	return ApiGetOauth2TokenInfoRequest{
		ApiService:  a,
		ctx:         ctx,
		accessToken: accessToken,
	}
}

// GetOauth2TokenInfoExecute Execute executes the request
//
//	@return AmadeusOAuth2Token
func (a *OAuth2AccessTokenAPIService) GetOauth2TokenInfoExecute(r ApiGetOauth2TokenInfoRequest) (*models.AmadeusOAuth2Token, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []api.FormFile
		localVarReturnValue *models.AmadeusOAuth2Token
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "OAuth2AccessTokenAPIService.GetOauth2TokenInfo")
	if err != nil {
		return localVarReturnValue, nil, api.NewGenericOpenAPIErrorWithError(err.Error())
	}

	localVarPath := localBasePath + "/security/oauth2/token/{access_token}"
	localVarPath = strings.Replace(localVarPath, "{"+"access_token"+"}", url.PathEscape(api.ParameterValueToString(r.accessToken, "accessToken")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := api.SelectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := api.SelectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.Client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.Client.CallAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := api.NewGenericOpenAPIErrorWithErrorAndBody(localVarHTTPResponse.Status, localVarBody)
		if localVarHTTPResponse.StatusCode == 404 {
			var v models.ErrorResponse
			err = a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.SetError(err.Error())
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.SetError(api.FormatErrorMessage(localVarHTTPResponse.Status, &v))
			newErr.SetModel(v)
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.Client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := api.NewGenericOpenAPIErrorWithErrorAndBody(err.Error(), localVarBody)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiOauth2TokenRequest struct {
	ctx          context.Context
	ApiService   *OAuth2AccessTokenAPIService
	grantType    *string
	clientId     *string
	clientSecret *string
}

func (r ApiOauth2TokenRequest) GrantType(grantType string) ApiOauth2TokenRequest {
	r.grantType = &grantType
	return r
}

func (r ApiOauth2TokenRequest) ClientId(clientId string) ApiOauth2TokenRequest {
	r.clientId = &clientId
	return r
}

func (r ApiOauth2TokenRequest) ClientSecret(clientSecret string) ApiOauth2TokenRequest {
	r.clientSecret = &clientSecret
	return r
}

func (r ApiOauth2TokenRequest) Execute() (*models.AmadeusOAuth2Token, *http.Response, error) {
	return r.ApiService.Oauth2TokenExecute(r)
}

/*
Oauth2Token The OAuth 2.0 token endpoint

The token endpoint is used by the client to obtain an access token by presenting its authorization grant.
To learn more about this endpoint please refer to the specification at https://tools.ietf.org/html/rfc6749#section-3.2

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOauth2TokenRequest
*/
func (a *OAuth2AccessTokenAPIService) Oauth2Token(ctx context.Context) ApiOauth2TokenRequest {
	return ApiOauth2TokenRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Oauth2TokenExecute Execute executes the request
//
//	@return AmadeusOAuth2Token
func (a *OAuth2AccessTokenAPIService) Oauth2TokenExecute(r ApiOauth2TokenRequest) (*models.AmadeusOAuth2Token, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []api.FormFile
		localVarReturnValue *models.AmadeusOAuth2Token
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "OAuth2AccessTokenAPIService.Oauth2Token")
	if err != nil {
		return localVarReturnValue, nil, api.NewGenericOpenAPIErrorWithError(err.Error())
	}

	localVarPath := localBasePath + "/security/oauth2/token"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.grantType == nil {
		return localVarReturnValue, nil, api.ReportError("grantType is required and must be specified")
	}
	if r.clientId == nil {
		return localVarReturnValue, nil, api.ReportError("clientId is required and must be specified")
	}
	if r.clientSecret == nil {
		return localVarReturnValue, nil, api.ReportError("clientSecret is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/x-www-form-urlencoded"}

	// set Content-Type header
	localVarHTTPContentType := api.SelectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := api.SelectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	api.ParameterAddToHeaderOrQuery(localVarFormParams, "grant_type", r.grantType, "")
	api.ParameterAddToHeaderOrQuery(localVarFormParams, "client_id", r.clientId, "")
	api.ParameterAddToHeaderOrQuery(localVarFormParams, "client_secret", r.clientSecret, "")
	req, err := a.Client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.Client.CallAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := api.NewGenericOpenAPIErrorWithErrorAndBody(localVarHTTPResponse.Status, localVarBody)
		if localVarHTTPResponse.StatusCode == 401 {
			var v models.ErrorResponse
			err = a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.SetError(err.Error())
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.SetError(api.FormatErrorMessage(localVarHTTPResponse.Status, &v))
			newErr.SetModel(v)
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v models.ErrorResponse
			err = a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.SetError(err.Error())
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.SetError(api.FormatErrorMessage(localVarHTTPResponse.Status, &v))
			newErr.SetModel(v)
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.Client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := api.NewGenericOpenAPIErrorWithErrorAndBody(err.Error(), localVarBody)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
