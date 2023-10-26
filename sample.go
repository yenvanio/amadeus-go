package main

import (
	"context"
	"fmt"
	"github.com/yenvanio/amadeus-go/amadeus/api"
	"github.com/yenvanio/amadeus-go/amadeus/models"
	"time"
)

var (
	ClientId            = "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
	ClientSecret        = "xxxxxxxxxxxxxxxx"
	GrantType           = "client_credentials"
	XHTTPMethodOverride = "GET"
)

func main() {
	// Initialize Client
	amadeusClient := api.NewAPIClient(api.NewConfiguration())
	ctx := context.Background()

	// Generate Access Token
	token, _, err := amadeusClient.OAuth2AccessTokenAPI.Oauth2TokenExecute(api.ApiOauth2TokenRequest{
		Context:      ctx,
		ApiService:   amadeusClient.OAuth2AccessTokenAPI,
		GrantType:    &GrantType,
		ClientId:     &ClientId,
		ClientSecret: &ClientSecret,
	})
	if err != nil {
		fmt.Println(fmt.Errorf("error generating access token %w", err))
		return
	}

	// Add Auth Token to Header
	headers := amadeusClient.Cfg.DefaultHeader
	headers["Authorization"] = fmt.Sprintf("Bearer %s", *token.AccessToken)
	amadeusClient.Cfg.DefaultHeader = headers

	// Initialize Flight Search Data
	originLocationCode := "YYZ"
	destinationLocationCode := "YVR"
	departureDate := time.Now().Add(time.Hour * 24 * 7 * 2).Format(time.DateOnly)
	var adults int32 = 2
	var maxResults int32 = 1

	// Get Flight Offers
	offers, _, err := amadeusClient.ShoppingAPI.GetFlightOffersExecute(api.ApiGetFlightOffersRequest{
		Context:                 ctx,
		ApiService:              amadeusClient.ShoppingAPI,
		OriginLocationCode:      &originLocationCode,
		DestinationLocationCode: &destinationLocationCode,
		DepartureDate:           &departureDate,
		Adults:                  &adults,
		Max:                     &maxResults,
	})

	if err != nil {
		fmt.Println(fmt.Errorf("error fetching flight offers %w", err))
		return
	}

	//Get Update Flight Offers Price
	prices, _, err := amadeusClient.ShoppingAPI.QuoteAirOffersExecute(api.ApiQuoteAirOffersRequest{
		Context:             ctx,
		ApiService:          amadeusClient.ShoppingAPI,
		XHTTPMethodOverride: &XHTTPMethodOverride,
		PriceFlightOffersBody: &models.GetPriceQuery{
			Data: models.FlightOfferPricingIn{
				Type:         "flight-offers-pricing",
				FlightOffers: offers.GetData(),
			},
		},
	})
	if err != nil {
		fmt.Println(fmt.Errorf("error fetching flight prices %w", err))
		return
	}

	fmt.Println(fmt.Sprintf("Base Fare: %v", prices.Data.FlightOffers[0].Price.GetBase()))
}
