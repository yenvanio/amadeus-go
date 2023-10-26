package api

import (
	"bytes"
	"context"
	"github.com/yenvanio/amadeus-go/amadeus/models"
	"io"
	"net/http"
	"net/url"
)

type ApiQuoteAirOffersRequest struct {
	Context               context.Context
	ApiService            *APIService
	XHTTPMethodOverride   *string
	PriceFlightOffersBody *models.GetPriceQuery
	Include               *[]string
	ForceClass            *bool
}

func (r ApiQuoteAirOffersRequest) Execute() (*models.SuccessPricing, *http.Response, error) {
	return r.ApiService.QuoteAirOffersExecute(r)
}

/*
QuoteAirOffers Confirm pricing of given flightOffers.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiQuoteAirOffersRequest
*/
func (a *APIService) QuoteAirOffers(ctx context.Context) ApiQuoteAirOffersRequest {
	return ApiQuoteAirOffersRequest{
		ApiService: a,
		Context:    ctx,
	}
}

// QuoteAirOffersExecute Execute executes the request
//
//	@return SuccessPricing
func (a *APIService) QuoteAirOffersExecute(r ApiQuoteAirOffersRequest) (*models.SuccessPricing, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []FormFile
		localVarReturnValue *models.SuccessPricing
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.Context, "ShoppingAPIService.QuoteAirOffers")
	if err != nil {
		return localVarReturnValue, nil, NewGenericOpenAPIErrorWithError(err.Error())
	}

	localVarPath := localBasePath + "/v1/shopping/flight-offers/pricing"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.XHTTPMethodOverride == nil {
		return localVarReturnValue, nil, ReportError("xHTTPMethodOverride is required and must be specified")
	}
	if r.PriceFlightOffersBody == nil {
		return localVarReturnValue, nil, ReportError("priceFlightOffersBody is required and must be specified")
	}

	if r.Include != nil {
		ParameterAddToHeaderOrQuery(localVarQueryParams, "include", r.Include, "csv")
	}
	if r.ForceClass != nil {
		ParameterAddToHeaderOrQuery(localVarQueryParams, "forceClass", r.ForceClass, "")
	} else {
		var defaultValue bool = false
		r.ForceClass = &defaultValue
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.amadeus+json"}

	// set Content-Type header
	localVarHTTPContentType := SelectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.amadeus+json"}

	// set Accept header
	localVarHTTPHeaderAccept := SelectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	ParameterAddToHeaderOrQuery(localVarHeaderParams, "X-HTTP-Method-Override", r.XHTTPMethodOverride, "")
	// body params
	localVarPostBody = r.PriceFlightOffersBody
	req, err := a.Client.PrepareRequest(r.Context, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
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
		newErr := NewGenericOpenAPIErrorWithErrorAndBody(localVarHTTPResponse.Status, localVarBody)
		if localVarHTTPResponse.StatusCode == 400 {
			var v models.Error400
			err = a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.SetError(err.Error())
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.SetError(FormatErrorMessage(localVarHTTPResponse.Status, &v))
			newErr.SetModel(v)
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		var v models.Error500
		err = a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.SetError(err.Error())
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.SetError(FormatErrorMessage(localVarHTTPResponse.Status, &v))
		newErr.SetModel(v)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.Client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := NewGenericOpenAPIErrorWithErrorAndBody(err.Error(), localVarBody)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
