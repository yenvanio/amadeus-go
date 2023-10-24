# ShoppingAPI

All URIs are relative to *https://test.api.amadeus.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetFlightOffers**](ShoppingAPI.md#GetFlightOffers) | **Get** /shopping/flight-offers | Return list of Flight Offers based on searching criteria.
[**SearchFlightOffers**](ShoppingAPI.md#SearchFlightOffers) | **Post** /shopping/flight-offers | Return list of Flight Offers based on posted searching criteria.
[**QuoteAirOffers**](ShoppingAPI.md#QuoteAirOffers) | **Post** /shopping/flight-offers/pricing | Confirm pricing of given flightOffers.

## QuoteAirOffers

> SuccessPricing QuoteAirOffers(ctx).XHTTPMethodOverride(xHTTPMethodOverride).PriceFlightOffersBody(priceFlightOffersBody).Include(include).ForceClass(forceClass).Execute()

Confirm pricing of given flightOffers.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/yenvanio/amadeus-go"
)

func main() {
    xHTTPMethodOverride := "xHTTPMethodOverride_example" // string | the HTTP method to apply (default to "GET")
    priceFlightOffersBody := *openapiclient.NewGetPriceQuery(*openapiclient.NewFlightOfferPricingIn("flight-offer-pricing", []openapiclient.FlightOffer{*openapiclient.NewFlightOffer("flight-offer", "1")})) // GetPriceQuery | list of criteria to confirm the price of a dedicated flight-offers
    include := []string{"Include_example"} // []string | Sub-resources to be included:  * **credit-card-fees** to get the credit card fee applied on the booking  * **bags** to get extra bag options  * **other-services** to get services options  * **detailed-fare-rules** to get detailed fare rules options  (optional)
    forceClass := true // bool | parameter to force the usage of booking class for pricing - **true**, to for pricing with the specified booking class - **false**, to get the best available price  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ShoppingAPI.QuoteAirOffers(context.Background()).XHTTPMethodOverride(xHTTPMethodOverride).PriceFlightOffersBody(priceFlightOffersBody).Include(include).ForceClass(forceClass).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ShoppingAPI.QuoteAirOffers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `QuoteAirOffers`: SuccessPricing
    fmt.Fprintf(os.Stdout, "Response from `ShoppingAPI.QuoteAirOffers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQuoteAirOffersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xHTTPMethodOverride** | **string** | the HTTP method to apply | [default to &quot;GET&quot;]
 **priceFlightOffersBody** | [**GetPriceQuery**](GetPriceQuery.md) | list of criteria to confirm the price of a dedicated flight-offers | 
 **include** | **[]string** | Sub-resources to be included:  * **credit-card-fees** to get the credit card fee applied on the booking  * **bags** to get extra bag options  * **other-services** to get services options  * **detailed-fare-rules** to get detailed fare rules options  | 
 **forceClass** | **bool** | parameter to force the usage of booking class for pricing - **true**, to for pricing with the specified booking class - **false**, to get the best available price  | [default to false]

### Return type

[**SuccessPricing**](SuccessPricing.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/vnd.amadeus+json
- **Accept**: application/vnd.amadeus+json

## GetFlightOffers

> Success GetFlightOffers(ctx).OriginLocationCode(originLocationCode).DestinationLocationCode(destinationLocationCode).DepartureDate(departureDate).Adults(adults).ReturnDate(returnDate).Children(children).Infants(infants).TravelClass(travelClass).IncludedAirlineCodes(includedAirlineCodes).ExcludedAirlineCodes(excludedAirlineCodes).NonStop(nonStop).CurrencyCode(currencyCode).MaxPrice(maxPrice).Max(max).Execute()

Return list of Flight Offers based on searching criteria.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    originLocationCode := "SYD" // string | city/airport [IATA code](http://www.iata.org/publications/Pages/code-search.aspx) from which the traveler will depart, e.g. BOS for Boston
    destinationLocationCode := "BKK" // string | city/airport [IATA code](http://www.iata.org/publications/Pages/code-search.aspx) to which the traveler is going, e.g. PAR for Paris
    departureDate := time.Now() // string | the date on which the traveler will depart from the origin to go to the destination. Dates are specified in the [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) YYYY-MM-DD format, e.g. 2017-12-25
    adults := int32(56) // int32 | the number of adult travelers (age 12 or older on date of departure).  The total number of seated travelers (adult and children) can not exceed 9. (default to 1)
    returnDate := time.Now() // string | the date on which the traveler will depart from the destination to return to the origin. If this parameter is not specified, only one-way itineraries are found. If this parameter is specified, only round-trip itineraries are found. Dates are specified in the [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) YYYY-MM-DD format, e.g. 2018-02-28 (optional)
    children := int32(56) // int32 | the number of child travelers (older than age 2 and younger than age 12 on date of departure) who will each have their own separate seat. If specified, this number should be greater than or equal to 0  The total number of seated travelers (adult and children) can not exceed 9. (optional)
    infants := int32(56) // int32 | the number of infant travelers (whose age is less or equal to 2 on date of departure). Infants travel on the lap of an adult traveler, and thus the number of infants must not exceed the number of adults. If specified, this number should be greater than or equal to 0 (optional)
    travelClass := "travelClass_example" // string | most of the flight time should be spent in a cabin of this quality or higher. The accepted travel class is economy, premium economy, business or first class. If no travel class is specified, the search considers any travel class (optional)
    includedAirlineCodes := "includedAirlineCodes_example" // string | This option ensures that the system will only consider these airlines. This can not be cumulated with parameter excludedAirlineCodes.  Airlines are specified as [IATA airline codes](http://www.iata.org/publications/Pages/code-search.aspx) and are comma-separated, e.g. 6X,7X,8X  (optional)
    excludedAirlineCodes := "excludedAirlineCodes_example" // string | This option ensures that the system will ignore these airlines. This can not be cumulated with parameter includedAirlineCodes.  Airlines are specified as [IATA airline codes](http://www.iata.org/publications/Pages/code-search.aspx) and are comma-separated, e.g. 6X,7X,8X  (optional)
    nonStop := true // bool | if set to true, the search will find only flights going from the origin to the destination with no stop in between (optional) (default to false)
    currencyCode := "currencyCode_example" // string | the preferred currency for the flight offers. Currency is specified in the [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) format, e.g. EUR for Euro (optional)
    maxPrice := int32(56) // int32 | maximum price per traveler. By default, no limit is applied. If specified, the value should be a positive number with no decimals (optional)
    max := int32(56) // int32 | maximum number of flight offers to return. If specified, the value should be greater than or equal to 1 (optional) (default to 250)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ShoppingAPI.GetFlightOffers(context.Background()).OriginLocationCode(originLocationCode).DestinationLocationCode(destinationLocationCode).DepartureDate(departureDate).Adults(adults).ReturnDate(returnDate).Children(children).Infants(infants).TravelClass(travelClass).IncludedAirlineCodes(includedAirlineCodes).ExcludedAirlineCodes(excludedAirlineCodes).NonStop(nonStop).CurrencyCode(currencyCode).MaxPrice(maxPrice).Max(max).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ShoppingAPI.GetFlightOffers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFlightOffers`: Success
    fmt.Fprintf(os.Stdout, "Response from `ShoppingAPI.GetFlightOffers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFlightOffersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **originLocationCode** | **string** | city/airport [IATA code](http://www.iata.org/publications/Pages/code-search.aspx) from which the traveler will depart, e.g. BOS for Boston |
 **destinationLocationCode** | **string** | city/airport [IATA code](http://www.iata.org/publications/Pages/code-search.aspx) to which the traveler is going, e.g. PAR for Paris |
 **departureDate** | **string** | the date on which the traveler will depart from the origin to go to the destination. Dates are specified in the [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) YYYY-MM-DD format, e.g. 2017-12-25 |
 **adults** | **int32** | the number of adult travelers (age 12 or older on date of departure).  The total number of seated travelers (adult and children) can not exceed 9. | [default to 1]
 **returnDate** | **string** | the date on which the traveler will depart from the destination to return to the origin. If this parameter is not specified, only one-way itineraries are found. If this parameter is specified, only round-trip itineraries are found. Dates are specified in the [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) YYYY-MM-DD format, e.g. 2018-02-28 |
 **children** | **int32** | the number of child travelers (older than age 2 and younger than age 12 on date of departure) who will each have their own separate seat. If specified, this number should be greater than or equal to 0  The total number of seated travelers (adult and children) can not exceed 9. |
 **infants** | **int32** | the number of infant travelers (whose age is less or equal to 2 on date of departure). Infants travel on the lap of an adult traveler, and thus the number of infants must not exceed the number of adults. If specified, this number should be greater than or equal to 0 |
 **travelClass** | **string** | most of the flight time should be spent in a cabin of this quality or higher. The accepted travel class is economy, premium economy, business or first class. If no travel class is specified, the search considers any travel class |
 **includedAirlineCodes** | **string** | This option ensures that the system will only consider these airlines. This can not be cumulated with parameter excludedAirlineCodes.  Airlines are specified as [IATA airline codes](http://www.iata.org/publications/Pages/code-search.aspx) and are comma-separated, e.g. 6X,7X,8X  |
 **excludedAirlineCodes** | **string** | This option ensures that the system will ignore these airlines. This can not be cumulated with parameter includedAirlineCodes.  Airlines are specified as [IATA airline codes](http://www.iata.org/publications/Pages/code-search.aspx) and are comma-separated, e.g. 6X,7X,8X  |
 **nonStop** | **bool** | if set to true, the search will find only flights going from the origin to the destination with no stop in between | [default to false]
 **currencyCode** | **string** | the preferred currency for the flight offers. Currency is specified in the [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) format, e.g. EUR for Euro |
 **maxPrice** | **int32** | maximum price per traveler. By default, no limit is applied. If specified, the value should be a positive number with no decimals |
 **max** | **int32** | maximum number of flight offers to return. If specified, the value should be greater than or equal to 1 | [default to 250]

### Return type

[**Success**](Success.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.amadeus+json

## SearchFlightOffers

> Success1 SearchFlightOffers(ctx).XHTTPMethodOverride(xHTTPMethodOverride).GetFlightOffersBody(getFlightOffersBody).Execute()

Return list of Flight Offers based on posted searching criteria.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    xHTTPMethodOverride := "xHTTPMethodOverride_example" // string | the HTTP method to apply (default to "GET")
    getFlightOffersBody := *openapiclient.NewGetFlightOffersQuery([]openapiclient.OriginDestination{*openapiclient.NewOriginDestination()}, []openapiclient.ExtendedTravelerInfo{*openapiclient.NewExtendedTravelerInfo("1", openapiclient.TravelerType("ADULT"))}, []openapiclient.FlightOfferSource{openapiclient.FlightOfferSource("GDS")}) // GetFlightOffersQuery | list of criteria to retrieve a list of flight offers

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ShoppingAPI.SearchFlightOffers(context.Background()).XHTTPMethodOverride(xHTTPMethodOverride).GetFlightOffersBody(getFlightOffersBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ShoppingAPI.SearchFlightOffers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchFlightOffers`: Success1
    fmt.Fprintf(os.Stdout, "Response from `ShoppingAPI.SearchFlightOffers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchFlightOffersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xHTTPMethodOverride** | **string** | the HTTP method to apply | [default to &quot;GET&quot;]
 **getFlightOffersBody** | [**GetFlightOffersQuery**](GetFlightOffersQuery.md) | list of criteria to retrieve a list of flight offers |

### Return type

[**Success**](Success.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

