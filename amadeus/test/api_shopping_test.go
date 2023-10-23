package openapi

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/yenvanio/amadeus-go/amadeus"
	"testing"
)

func Test_openapi_ShoppingAPIService(t *testing.T) {

	configuration := amadeus.NewConfiguration()
	apiClient := amadeus.NewAPIClient(configuration)

	t.Run("Test ShoppingAPIService GetFlightOffers", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ShoppingAPI.GetFlightOffers(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ShoppingAPIService SearchFlightOffers", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ShoppingAPI.SearchFlightOffers(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
