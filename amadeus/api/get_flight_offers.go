package api

import (
	"bytes"
	"context"
	"github.com/yenvanio/amadeus-go/amadeus/models"
	"io"
	"net/http"
	"net/url"
)

type ApiGetFlightOffersRequest struct {
	Context                 context.Context
	ApiService              *APIService
	OriginLocationCode      *string
	DestinationLocationCode *string
	DepartureDate           *string
	Adults                  *int32
	ReturnDate              *string
	Children                *int32
	Infants                 *int32
	TravelClass             *string
	IncludedAirlineCodes    *string
	ExcludedAirlineCodes    *string
	NonStop                 *bool
	CurrencyCode            *string
	MaxPrice                *int32
	Max                     *int32
}

func (r ApiGetFlightOffersRequest) Execute() (*models.FlightOfferSuccess, *http.Response, error) {
	return r.ApiService.GetFlightOffersExecute(r)
}

/*
GetFlightOffers Return list of Flight Offers based on searching criteria.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetFlightOffersRequest
*/
func (a *APIService) GetFlightOffers(ctx context.Context) ApiGetFlightOffersRequest {
	return ApiGetFlightOffersRequest{
		ApiService: a,
		Context:    ctx,
	}
}

// GetFlightOffersExecute Execute executes the request
//
//	@return Success
func (a *APIService) GetFlightOffersExecute(r ApiGetFlightOffersRequest) (*models.FlightOfferSuccess, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []FormFile
		localVarReturnValue *models.FlightOfferSuccess
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.Context, "ShoppingAPIService.GetFlightOffers")
	if err != nil {
		return localVarReturnValue, nil, NewGenericOpenAPIErrorWithError(err.Error())
	}

	localVarPath := localBasePath + "/v2/shopping/flight-offers"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.OriginLocationCode == nil {
		return localVarReturnValue, nil, ReportError("originLocationCode is required and must be specified")
	}
	if r.DestinationLocationCode == nil {
		return localVarReturnValue, nil, ReportError("destinationLocationCode is required and must be specified")
	}
	if r.DepartureDate == nil {
		return localVarReturnValue, nil, ReportError("departureDate is required and must be specified")
	}
	if r.Adults == nil {
		return localVarReturnValue, nil, ReportError("adults is required and must be specified")
	}
	if *r.Adults < 1 {
		return localVarReturnValue, nil, ReportError("adults must be greater than 1")
	}
	if *r.Adults > 9 {
		return localVarReturnValue, nil, ReportError("adults must be less than 9")
	}

	ParameterAddToHeaderOrQuery(localVarQueryParams, "originLocationCode", r.OriginLocationCode, "")
	ParameterAddToHeaderOrQuery(localVarQueryParams, "destinationLocationCode", r.DestinationLocationCode, "")
	ParameterAddToHeaderOrQuery(localVarQueryParams, "departureDate", r.DepartureDate, "")
	if r.ReturnDate != nil {
		ParameterAddToHeaderOrQuery(localVarQueryParams, "returnDate", r.ReturnDate, "")
	}
	ParameterAddToHeaderOrQuery(localVarQueryParams, "adults", r.Adults, "")
	if r.Children != nil {
		ParameterAddToHeaderOrQuery(localVarQueryParams, "children", r.Children, "")
	}
	if r.Infants != nil {
		ParameterAddToHeaderOrQuery(localVarQueryParams, "infants", r.Infants, "")
	}
	if r.TravelClass != nil {
		ParameterAddToHeaderOrQuery(localVarQueryParams, "travelClass", r.TravelClass, "")
	}
	if r.IncludedAirlineCodes != nil {
		ParameterAddToHeaderOrQuery(localVarQueryParams, "includedAirlineCodes", r.IncludedAirlineCodes, "")
	}
	if r.ExcludedAirlineCodes != nil {
		ParameterAddToHeaderOrQuery(localVarQueryParams, "excludedAirlineCodes", r.ExcludedAirlineCodes, "")
	}
	if r.NonStop != nil {
		ParameterAddToHeaderOrQuery(localVarQueryParams, "nonStop", r.NonStop, "")
	} else {
		var defaultValue bool = false
		r.NonStop = &defaultValue
	}
	if r.CurrencyCode != nil {
		ParameterAddToHeaderOrQuery(localVarQueryParams, "currencyCode", r.CurrencyCode, "")
	}
	if r.MaxPrice != nil {
		ParameterAddToHeaderOrQuery(localVarQueryParams, "maxPrice", r.MaxPrice, "")
	}
	if r.Max != nil {
		ParameterAddToHeaderOrQuery(localVarQueryParams, "max", r.Max, "")
	} else {
		var defaultValue int32 = 250
		r.Max = &defaultValue
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

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

type ApiSearchFlightOffersRequest struct {
	Context             context.Context
	ApiService          *APIService
	XHTTPMethodOverride *string
	GetFlightOffersBody *models.GetFlightOffersQuery
}

func (r ApiSearchFlightOffersRequest) Execute() (*models.FlightOfferSuccess, *http.Response, error) {
	return r.ApiService.SearchFlightOffersExecute(r)
}

/*
SearchFlightOffers Return list of Flight Offers based on posted searching criteria.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiSearchFlightOffersRequest
*/
func (a *APIService) SearchFlightOffers(ctx context.Context) ApiSearchFlightOffersRequest {
	return ApiSearchFlightOffersRequest{
		ApiService: a,
		Context:    ctx,
	}
}

// SearchFlightOffersExecute Execute executes the request
//
//	@return Success1
func (a *APIService) SearchFlightOffersExecute(r ApiSearchFlightOffersRequest) (*models.FlightOfferSuccess, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []FormFile
		localVarReturnValue *models.FlightOfferSuccess
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.Context, "ShoppingAPIService.SearchFlightOffers")
	if err != nil {
		return localVarReturnValue, nil, NewGenericOpenAPIErrorWithError(err.Error())
	}

	localVarPath := localBasePath + "/v2/shopping/flight-offers"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.XHTTPMethodOverride == nil {
		return localVarReturnValue, nil, ReportError("xHTTPMethodOverride is required and must be specified")
	}
	if r.GetFlightOffersBody == nil {
		return localVarReturnValue, nil, ReportError("getFlightOffersBody is required and must be specified")
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
	localVarPostBody = r.GetFlightOffersBody
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
