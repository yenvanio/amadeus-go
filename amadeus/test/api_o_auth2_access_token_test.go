package openapi

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/yenvanio/amadeus-go/amadeus/api"
	"testing"
)

func Test_openapi_OAuth2AccessTokenAPIService(t *testing.T) {

	configuration := api.NewConfiguration()
	apiClient := api.NewAPIClient(configuration)

	t.Run("Test OAuth2AccessTokenAPIService GetOauth2TokenInfo", func(t *testing.T) {
		var accessToken string

		resp, httpRes, err := apiClient.OAuth2AccessTokenAPI.GetOauth2TokenInfo(context.Background(), accessToken).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OAuth2AccessTokenAPIService Oauth2Token", func(t *testing.T) {
		resp, httpRes, err := apiClient.OAuth2AccessTokenAPI.Oauth2Token(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
