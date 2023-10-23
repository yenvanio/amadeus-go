# \ShoppingAPI

All URIs are relative to *https://test.api.amadeus.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

