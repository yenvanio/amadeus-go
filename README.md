# Go API client for openapi

Before using this API, we recommend you read our **[Authorization Guide](https://developers.amadeus.com/self-service/apis-docs/guides/authorization-262)** for more information on how to generate an access token. 

Please also be aware that our test environment is based on a subset of the production, if you are not returning any results try with big cities/airports like LON (London) or NYC (New-York).

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.3.0
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import openapi "github.com/GIT_USER_ID/GIT_REPO_ID"
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

All URIs are relative to *https://test.api.amadeus.com/v1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*ShoppingAPI* | [**QuoteAirOffers**](docs/ShoppingAPI.md#quoteairoffers) | **Post** /shopping/flight-offers/pricing | Confirm pricing of given flightOffers.


## Documentation For Models

 - [AdditionalService](docs/AdditionalService.md)
 - [AdditionalServiceType](docs/AdditionalServiceType.md)
 - [AdditionalServicesRequest](docs/AdditionalServicesRequest.md)
 - [Address](docs/Address.md)
 - [AircraftEquipment](docs/AircraftEquipment.md)
 - [AllotmentDetails](docs/AllotmentDetails.md)
 - [BaggageAllowance](docs/BaggageAllowance.md)
 - [Bags](docs/Bags.md)
 - [BaseName](docs/BaseName.md)
 - [BookingRequirements](docs/BookingRequirements.md)
 - [ChargeableCheckdBags](docs/ChargeableCheckdBags.md)
 - [ChargeableSeat](docs/ChargeableSeat.md)
 - [Co2Emission](docs/Co2Emission.md)
 - [Contact](docs/Contact.md)
 - [ContactDictionary](docs/ContactDictionary.md)
 - [ContactPurpose](docs/ContactPurpose.md)
 - [CreditCardFee](docs/CreditCardFee.md)
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
 - [ExtendedPrice](docs/ExtendedPrice.md)
 - [FareDetailsBySegment](docs/FareDetailsBySegment.md)
 - [FareRules](docs/FareRules.md)
 - [Fee](docs/Fee.md)
 - [FeeType](docs/FeeType.md)
 - [FlightEndPoint](docs/FlightEndPoint.md)
 - [FlightOffer](docs/FlightOffer.md)
 - [FlightOfferPricingIn](docs/FlightOfferPricingIn.md)
 - [FlightOfferPricingOut](docs/FlightOfferPricingOut.md)
 - [FlightOfferSource](docs/FlightOfferSource.md)
 - [FlightSegment](docs/FlightSegment.md)
 - [FlightStop](docs/FlightStop.md)
 - [GetPriceQuery](docs/GetPriceQuery.md)
 - [IdentityDocument](docs/IdentityDocument.md)
 - [IncludedResourcesMap](docs/IncludedResourcesMap.md)
 - [Issue](docs/Issue.md)
 - [IssueSource](docs/IssueSource.md)
 - [Itineraries](docs/Itineraries.md)
 - [LocationValue](docs/LocationValue.md)
 - [LoyaltyProgram](docs/LoyaltyProgram.md)
 - [Name](docs/Name.md)
 - [OperatingFlight](docs/OperatingFlight.md)
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
 - [TravelerPricing](docs/TravelerPricing.md)
 - [TravelerPricingFareOption](docs/TravelerPricingFareOption.md)
 - [TravelerType](docs/TravelerType.md)


## Documentation For Authorization

Endpoints do not require authorization.


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



