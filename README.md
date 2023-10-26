# Go API client for Amadeus

Put the package under your project folder and add the following in import:

```golang
import openapi "github.com/yenvanio/amadeus-go"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), openapi.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), openapi.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), openapi.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), openapi.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://test.api.amadeus.com*

Class | Method | HTTP request                                   | Description
------------ | ------------- |------------------------------------------------| -------------
*OAuth2AccessTokenAPI* | [**GetOauth2TokenInfo**](docs/OAuth2AccessTokenAPI.md#getoauth2tokeninfo) | **Get** /v1/security/oauth2/token/{access_token} | The OAuth 2.0 token info endpoint
*OAuth2AccessTokenAPI* | [**Oauth2Token**](docs/OAuth2AccessTokenAPI.md#oauth2token) | **Post** /v1/security/oauth2/token   | The OAuth 2.0 token endpoint
*ShoppingAPI* | [**GetFlightOffers**](docs/ShoppingAPI.md#getflightoffers) | **Get** /v2/shopping/flight-offers         | Return list of Flight Offers based on searching criteria.
*ShoppingAPI* | [**SearchFlightOffers**](docs/ShoppingAPI.md#searchflightoffers) | **Post** /v2/shopping/flight-offers     | Return list of Flight Offers based on posted searching criteria.
*ShoppingAPI* | [**QuoteAirOffers**](docs/ShoppingAPI.md#quoteairoffers) | **Post** /v1/shopping/flight-offers/pricing | Confirm pricing of given flightOffers.


## Documentation For Authorization

All endpoints are required to be authenticated by an access token. An access token can be generated with a Client ID and Secret.
These must be obtained by signing up with Amadeus.

See the example below to generate and use the access token.

```go
	// Initialize Client
	amadeusClient := api.NewAPIClient(api.NewConfiguration())
	
	// Generate Access Token
	token, _, err := amadeusClient.OAuth2AccessTokenAPI.Oauth2TokenExecute(api.ApiOauth2TokenRequest{
		Context:      ctx,
		ApiService:   amadeusClient.OAuth2AccessTokenAPI,
		GrantType:    &GrantType,
		ClientId:     &ClientId,
		ClientSecret: &ClientSecret,
	})
	
	// Add Auth Token to Header
	headers := amadeusClient.Cfg.DefaultHeader
	headers["Authorization"] = fmt.Sprintf("Bearer %s", *token.AccessToken)
	amadeusClient.Cfg.DefaultHeader = headers
```


## Documentation For Models
 - [AdditionalInformation](docs/AdditionalInformation.md)
 - [AdditionalService](docs/AdditionalService.md)
 - [AdditionalServiceType](docs/AdditionalServiceType.md)
 - [AdditionalServicesRequest](docs/AdditionalServicesRequest.md)
 - [Address](docs/Address.md)
 - [AircraftEquipment](docs/AircraftEquipment.md)
 - [AllotmentDetails](docs/AllotmentDetails.md)
 - [AmadeusOAuth2Token](docs/AmadeusOAuth2Token.md)
 - [BaggageAllowance](docs/BaggageAllowance.md)
 - [Bags](docs/Bags.md)
 - [BaseName](docs/BaseName.md)
 - [BookingRequirements](docs/BookingRequirements.md)
 - [CabinRestriction](docs/CabinRestriction.md)
 - [CarrierRestrictions](docs/CarrierRestrictions.md)
 - [ChargeableCheckdBags](docs/ChargeableCheckdBags.md)
 - [ChargeableSeat](docs/ChargeableSeat.md)
 - [Co2Emission](docs/Co2Emission.md)
 - [CollectionLinks](docs/CollectionLinks.md)
 - [CollectionMeta](docs/CollectionMeta.md)
 - [CollectionMetaLink](docs/CollectionMetaLink.md)
 - [ConnectionRestriction](docs/ConnectionRestriction.md)
 - [Contact](docs/Contact.md)
 - [ContactDictionary](docs/ContactDictionary.md)
 - [ContactPurpose](docs/ContactPurpose.md)
 - [Coverage](docs/Coverage.md)
 - [CreditCardFee](docs/CreditCardFee.md)
 - [DateTimeRange](docs/DateTimeRange.md)
 - [DateTimeType](docs/DateTimeType.md)
 - [Description](docs/Description.md)
 - [DetailedFareRules](docs/DetailedFareRules.md)
 - [Dictionaries](docs/Dictionaries.md)
 - [Discount](docs/Discount.md)
 - [DiscountTravelerType](docs/DiscountTravelerType.md)
 - [DiscountType](docs/DiscountType.md)
 - [Document](docs/Document.md)
 - [DocumentType](docs/DocumentType.md)
 - [ElementaryPrice](docs/ElementaryPrice.md)
 - [EmergencyContact](docs/EmergencyContact.md)
 - [Error400](docs/Error400.md)
 - [Error500](docs/Error500.md)
 - [ErrorResponse](docs/ErrorResponse.md)
 - [ExtendedCabinRestriction](docs/ExtendedCabinRestriction.md)
 - [ExtendedPrice](docs/ExtendedPrice.md)
 - [ExtendedPricingOptions](docs/ExtendedPricingOptions.md)
 - [ExtendedTravelerInfo](docs/ExtendedTravelerInfo.md)
 - [FareDetailsBySegment](docs/FareDetailsBySegment.md)
 - [FareRules](docs/FareRules.md)
 - [Fee](docs/Fee.md)
 - [FeeType](docs/FeeType.md)
 - [FlightEndPoint](docs/FlightEndPoint.md)
 - [FlightFilters](docs/FlightFilters.md)
 - [FlightOffer](docs/FlightOffer.md)
 - [FlightOfferPricingIn](docs/FlightOfferPricingIn.md)
 - [FlightOfferPricingOut](docs/FlightOfferPricingOut.md)
 - [FlightOfferSource](docs/FlightOfferSource.md)
 - [FlightSegment](docs/FlightSegment.md)
 - [FlightStop](docs/FlightStop.md)
 - [GetFlightOffersQuery](docs/GetFlightOffersQuery.md)
 - [GetPriceQuery](docs/GetPriceQuery.md)
 - [IdentityDocument](docs/IdentityDocument.md)
 - [IncludedResourcesMap](docs/IncludedResourcesMap.md)
 - [Issue](docs/Issue.md)
 - [IssueSource](docs/IssueSource.md)
 - [Itineraries](docs/Itineraries.md)
 - [LocationValue](docs/LocationValue.md)
 - [LoyaltyProgram](docs/LoyaltyProgram.md)
 - [Name](docs/Name.md)
 - [OneWayCombinations](docs/OneWayCombinations.md)
 - [OperatingFlight](docs/OperatingFlight.md)
 - [OriginDestination](docs/OriginDestination.md)
 - [OriginDestinationLight](docs/OriginDestinationLight.md)
 - [OriginalFlightEndPoint](docs/OriginalFlightEndPoint.md)
 - [OriginalFlightStop](docs/OriginalFlightStop.md)
 - [OtherServices](docs/OtherServices.md)
 - [PassengerConditions](docs/PassengerConditions.md)
 - [Payment](docs/Payment.md)
 - [PaymentBrand](docs/PaymentBrand.md)
 - [Phone](docs/Phone.md)
 - [PhoneDeviceType](docs/PhoneDeviceType.md)
 - [Price](docs/Price.md)
 - [PricingOptions](docs/PricingOptions.md)
 - [Segment](docs/Segment.md)
 - [ServiceName](docs/ServiceName.md)
 - [SliceDiceIndicator](docs/SliceDiceIndicator.md)
 - [Stakeholder](docs/Stakeholder.md)
 - [StakeholderGender](docs/StakeholderGender.md)
 - [SuccessPricing](docs/SuccessPricing.md)
 - [Tax](docs/Tax.md)
 - [TermAndCondition](docs/TermAndCondition.md)
 - [TravelClass](docs/TravelClass.md)
 - [Traveler](docs/Traveler.md)
 - [TravelerInfo](docs/TravelerInfo.md)
 - [TravelerPricing](docs/TravelerPricing.md)
 - [TravelerPricingFareOption](docs/TravelerPricingFareOption.md)
 - [TravelerType](docs/TravelerType.md)


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



