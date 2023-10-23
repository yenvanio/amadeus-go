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

type ApiGetFlightOffersRequest struct {
	ctx                     context.Context
	ApiService              *APIService
	originLocationCode      *string
	destinationLocationCode *string
	departureDate           *string
	adults                  *int32
	returnDate              *string
	children                *int32
	infants                 *int32
	travelClass             *string
	includedAirlineCodes    *string
	excludedAirlineCodes    *string
	nonStop                 *bool
	currencyCode            *string
	maxPrice                *int32
	max                     *int32
}

// OriginLocationCode city/airport [IATA code](http://www.iata.org/publications/Pages/code-search.aspx) from which the traveler will depart, e.g. BOS for Boston
func (r ApiGetFlightOffersRequest) OriginLocationCode(originLocationCode string) ApiGetFlightOffersRequest {
	r.originLocationCode = &originLocationCode
	return r
}

// DestinationLocationCode city/airport [IATA code](http://www.iata.org/publications/Pages/code-search.aspx) to which the traveler is going, e.g. PAR for Paris
func (r ApiGetFlightOffersRequest) DestinationLocationCode(destinationLocationCode string) ApiGetFlightOffersRequest {
	r.destinationLocationCode = &destinationLocationCode
	return r
}

// DepartureDate the date on which the traveler will depart from the origin to go to the destination. Dates are specified in the [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) YYYY-MM-DD format, e.g. 2017-12-25
func (r ApiGetFlightOffersRequest) DepartureDate(departureDate string) ApiGetFlightOffersRequest {
	r.departureDate = &departureDate
	return r
}

// Adults the number of adult travelers (age 12 or older on date of departure).  The total number of seated travelers (adult and children) can not exceed 9.
func (r ApiGetFlightOffersRequest) Adults(adults int32) ApiGetFlightOffersRequest {
	r.adults = &adults
	return r
}

// ReturnDate the date on which the traveler will depart from the destination to return to the origin. If this parameter is not specified, only one-way itineraries are found. If this parameter is specified, only round-trip itineraries are found. Dates are specified in the [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) YYYY-MM-DD format, e.g. 2018-02-28
func (r ApiGetFlightOffersRequest) ReturnDate(returnDate string) ApiGetFlightOffersRequest {
	r.returnDate = &returnDate
	return r
}

// Children the number of child travelers (older than age 2 and younger than age 12 on date of departure) who will each have their own separate seat. If specified, this number should be greater than or equal to 0  The total number of seated travelers (adult and children) can not exceed 9.
func (r ApiGetFlightOffersRequest) Children(children int32) ApiGetFlightOffersRequest {
	r.children = &children
	return r
}

// Infants the number of infant travelers (whose age is less or equal to 2 on date of departure). Infants travel on the lap of an adult traveler, and thus the number of infants must not exceed the number of adults. If specified, this number should be greater than or equal to 0
func (r ApiGetFlightOffersRequest) Infants(infants int32) ApiGetFlightOffersRequest {
	r.infants = &infants
	return r
}

// TravelClass most of the flight time should be spent in a cabin of this quality or higher. The accepted travel class is economy, premium economy, business or first class. If no travel class is specified, the search considers any travel class
func (r ApiGetFlightOffersRequest) TravelClass(travelClass string) ApiGetFlightOffersRequest {
	r.travelClass = &travelClass
	return r
}

// IncludedAirlineCodes This option ensures that the system will only consider these airlines. This can not be cumulated with parameter excludedAirlineCodes.  Airlines are specified as [IATA airline codes](http://www.iata.org/publications/Pages/code-search.aspx) and are comma-separated, e.g. 6X,7X,8X
func (r ApiGetFlightOffersRequest) IncludedAirlineCodes(includedAirlineCodes string) ApiGetFlightOffersRequest {
	r.includedAirlineCodes = &includedAirlineCodes
	return r
}

// ExcludedAirlineCodes This option ensures that the system will ignore these airlines. This can not be cumulated with parameter includedAirlineCodes.  Airlines are specified as [IATA airline codes](http://www.iata.org/publications/Pages/code-search.aspx) and are comma-separated, e.g. 6X,7X,8X
func (r ApiGetFlightOffersRequest) ExcludedAirlineCodes(excludedAirlineCodes string) ApiGetFlightOffersRequest {
	r.excludedAirlineCodes = &excludedAirlineCodes
	return r
}

// NonStop if set to true, the search will find only flights going from the origin to the destination with no stop in between
func (r ApiGetFlightOffersRequest) NonStop(nonStop bool) ApiGetFlightOffersRequest {
	r.nonStop = &nonStop
	return r
}

// CurrencyCode the preferred currency for the flight offers. Currency is specified in the [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) format, e.g. EUR for Euro
func (r ApiGetFlightOffersRequest) CurrencyCode(currencyCode string) ApiGetFlightOffersRequest {
	r.currencyCode = &currencyCode
	return r
}

// MaxPrice maximum price per traveler. By default, no limit is applied. If specified, the value should be a positive number with no decimals
func (r ApiGetFlightOffersRequest) MaxPrice(maxPrice int32) ApiGetFlightOffersRequest {
	r.maxPrice = &maxPrice
	return r
}

// Max maximum number of flight offers to return. If specified, the value should be greater than or equal to 1
func (r ApiGetFlightOffersRequest) Max(max int32) ApiGetFlightOffersRequest {
	r.max = &max
	return r
}

func (r ApiGetFlightOffersRequest) Execute() (*models.Success, *http.Response, error) {
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
		ctx:        ctx,
	}
}

// GetFlightOffersExecute Execute executes the request
//
//	@return Success
func (a *APIService) GetFlightOffersExecute(r ApiGetFlightOffersRequest) (*models.Success, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []api.FormFile
		localVarReturnValue *models.Success
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "ShoppingAPIService.GetFlightOffers")
	if err != nil {
		return localVarReturnValue, nil, api.NewGenericOpenAPIErrorWithError(err.Error())
	}

	localVarPath := localBasePath + "/shopping/flight-offers"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.originLocationCode == nil {
		return localVarReturnValue, nil, api.ReportError("originLocationCode is required and must be specified")
	}
	if r.destinationLocationCode == nil {
		return localVarReturnValue, nil, api.ReportError("destinationLocationCode is required and must be specified")
	}
	if r.departureDate == nil {
		return localVarReturnValue, nil, api.ReportError("departureDate is required and must be specified")
	}
	if r.adults == nil {
		return localVarReturnValue, nil, api.ReportError("adults is required and must be specified")
	}
	if *r.adults < 1 {
		return localVarReturnValue, nil, api.ReportError("adults must be greater than 1")
	}
	if *r.adults > 9 {
		return localVarReturnValue, nil, api.ReportError("adults must be less than 9")
	}

	api.ParameterAddToHeaderOrQuery(localVarQueryParams, "originLocationCode", r.originLocationCode, "")
	api.ParameterAddToHeaderOrQuery(localVarQueryParams, "destinationLocationCode", r.destinationLocationCode, "")
	api.ParameterAddToHeaderOrQuery(localVarQueryParams, "departureDate", r.departureDate, "")
	if r.returnDate != nil {
		api.ParameterAddToHeaderOrQuery(localVarQueryParams, "returnDate", r.returnDate, "")
	}
	api.ParameterAddToHeaderOrQuery(localVarQueryParams, "adults", r.adults, "")
	if r.children != nil {
		api.ParameterAddToHeaderOrQuery(localVarQueryParams, "children", r.children, "")
	}
	if r.infants != nil {
		api.ParameterAddToHeaderOrQuery(localVarQueryParams, "infants", r.infants, "")
	}
	if r.travelClass != nil {
		api.ParameterAddToHeaderOrQuery(localVarQueryParams, "travelClass", r.travelClass, "")
	}
	if r.includedAirlineCodes != nil {
		api.ParameterAddToHeaderOrQuery(localVarQueryParams, "includedAirlineCodes", r.includedAirlineCodes, "")
	}
	if r.excludedAirlineCodes != nil {
		api.ParameterAddToHeaderOrQuery(localVarQueryParams, "excludedAirlineCodes", r.excludedAirlineCodes, "")
	}
	if r.nonStop != nil {
		api.ParameterAddToHeaderOrQuery(localVarQueryParams, "nonStop", r.nonStop, "")
	} else {
		var defaultValue bool = false
		r.nonStop = &defaultValue
	}
	if r.currencyCode != nil {
		api.ParameterAddToHeaderOrQuery(localVarQueryParams, "currencyCode", r.currencyCode, "")
	}
	if r.maxPrice != nil {
		api.ParameterAddToHeaderOrQuery(localVarQueryParams, "maxPrice", r.maxPrice, "")
	}
	if r.max != nil {
		api.ParameterAddToHeaderOrQuery(localVarQueryParams, "max", r.max, "")
	} else {
		var defaultValue int32 = 250
		r.max = &defaultValue
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

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

type ApiSearchFlightOffersRequest struct {
	ctx                 context.Context
	ApiService          *APIService
	xHTTPMethodOverride *string
	getFlightOffersBody *models.GetFlightOffersQuery
}

// XHTTPMethodOverride the HTTP method to apply
func (r ApiSearchFlightOffersRequest) XHTTPMethodOverride(xHTTPMethodOverride string) ApiSearchFlightOffersRequest {
	r.xHTTPMethodOverride = &xHTTPMethodOverride
	return r
}

// GetFlightOffersBody list of criteria to retrieve a list of flight offers
func (r ApiSearchFlightOffersRequest) GetFlightOffersBody(getFlightOffersBody models.GetFlightOffersQuery) ApiSearchFlightOffersRequest {
	r.getFlightOffersBody = &getFlightOffersBody
	return r
}

func (r ApiSearchFlightOffersRequest) Execute() (*models.Success, *http.Response, error) {
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
		ctx:        ctx,
	}
}

// SearchFlightOffersExecute Execute executes the request
//
//	@return Success1
func (a *APIService) SearchFlightOffersExecute(r ApiSearchFlightOffersRequest) (*models.Success, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []api.FormFile
		localVarReturnValue *models.Success
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "ShoppingAPIService.SearchFlightOffers")
	if err != nil {
		return localVarReturnValue, nil, api.NewGenericOpenAPIErrorWithError(err.Error())
	}

	localVarPath := localBasePath + "/shopping/flight-offers"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.xHTTPMethodOverride == nil {
		return localVarReturnValue, nil, api.ReportError("xHTTPMethodOverride is required and must be specified")
	}
	if r.getFlightOffersBody == nil {
		return localVarReturnValue, nil, api.ReportError("getFlightOffersBody is required and must be specified")
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
	localVarPostBody = r.getFlightOffersBody
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
