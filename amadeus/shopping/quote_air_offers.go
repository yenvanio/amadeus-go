package shopping

import (
	"bytes"
	"context"
	api "github.com/yenvanio/amadeus-go/amadeus"
	"github.com/yenvanio/amadeus-go/amadeus/models"
	"io"
	"net/http"
	"net/url"
)

type ApiQuoteAirOffersRequest struct {
	ctx                   context.Context
	ApiService            *APIService
	xHTTPMethodOverride   *string
	priceFlightOffersBody *models.GetPriceQuery
	include               *[]string
	forceClass            *bool
}

// XHTTPMethodOverride the HTTP method to apply
func (r ApiQuoteAirOffersRequest) XHTTPMethodOverride(xHTTPMethodOverride string) ApiQuoteAirOffersRequest {
	r.xHTTPMethodOverride = &xHTTPMethodOverride
	return r
}

// PriceFlightOffersBody list of criteria to confirm the price of a dedicated flight-offers
func (r ApiQuoteAirOffersRequest) PriceFlightOffersBody(priceFlightOffersBody models.GetPriceQuery) ApiQuoteAirOffersRequest {
	r.priceFlightOffersBody = &priceFlightOffersBody
	return r
}

// Include Sub-resources to be included:  * **credit-card-fees** to get the credit card fee applied on the booking  * **bags** to get extra bag options  * **other-services** to get services options  * **detailed-fare-rules** to get detailed fare rules options
func (r ApiQuoteAirOffersRequest) Include(include []string) ApiQuoteAirOffersRequest {
	r.include = &include
	return r
}

// ForceClass parameter to force the usage of booking class for pricing - **true**, to for pricing with the specified booking class - **false**, to get the best available price
func (r ApiQuoteAirOffersRequest) ForceClass(forceClass bool) ApiQuoteAirOffersRequest {
	r.forceClass = &forceClass
	return r
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
		ctx:        ctx,
	}
}

// QuoteAirOffersExecute Execute executes the request
//
//	@return SuccessPricing
func (a *APIService) QuoteAirOffersExecute(r ApiQuoteAirOffersRequest) (*models.SuccessPricing, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []api.FormFile
		localVarReturnValue *models.SuccessPricing
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "ShoppingAPIService.QuoteAirOffers")
	if err != nil {
		return localVarReturnValue, nil, api.NewGenericOpenAPIErrorWithError(err.Error())
	}

	localVarPath := localBasePath + "/shopping/flight-offers/pricing"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.xHTTPMethodOverride == nil {
		return localVarReturnValue, nil, api.ReportError("xHTTPMethodOverride is required and must be specified")
	}
	if r.priceFlightOffersBody == nil {
		return localVarReturnValue, nil, api.ReportError("priceFlightOffersBody is required and must be specified")
	}

	if r.include != nil {
		api.ParameterAddToHeaderOrQuery(localVarQueryParams, "include", r.include, "csv")
	}
	if r.forceClass != nil {
		api.ParameterAddToHeaderOrQuery(localVarQueryParams, "forceClass", r.forceClass, "")
	} else {
		var defaultValue bool = false
		r.forceClass = &defaultValue
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.amadeus+json"}

	// set Content-Type header
	localVarHTTPContentType := api.SelectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.amadeus+json"}

	// set Accept header
	localVarHTTPHeaderAccept := api.SelectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	api.ParameterAddToHeaderOrQuery(localVarHeaderParams, "X-HTTP-Method-Override", r.xHTTPMethodOverride, "")
	// body params
	localVarPostBody = r.priceFlightOffersBody
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v models.Error400
			err = a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.SetError(err.Error())
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.SetError(api.FormatErrorMessage(localVarHTTPResponse.Status, &v))
			newErr.SetModel(v)
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		var v models.Error500
		err = a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.SetError(err.Error())
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.SetError(api.FormatErrorMessage(localVarHTTPResponse.Status, &v))
		newErr.SetModel(v)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.Client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := api.NewGenericOpenAPIErrorWithErrorAndBody(err.Error(), localVarBody)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
