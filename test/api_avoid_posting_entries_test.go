/*
E621

Testing AvoidPostingEntriesAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	openapiclient "github.com/anthrove/openapi-e621-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_openapi_AvoidPostingEntriesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test AvoidPostingEntriesAPIService CreateAvoidPosting", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.AvoidPostingEntriesAPI.CreateAvoidPosting(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AvoidPostingEntriesAPIService DeleteAvoidPosting", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var idOrArtistName string

		httpRes, err := apiClient.AvoidPostingEntriesAPI.DeleteAvoidPosting(context.Background(), idOrArtistName).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AvoidPostingEntriesAPIService DestroyAvoidPosting", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var idOrArtistName string

		httpRes, err := apiClient.AvoidPostingEntriesAPI.DestroyAvoidPosting(context.Background(), idOrArtistName).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AvoidPostingEntriesAPIService EditAvoidPosting", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var idOrArtistName string

		resp, httpRes, err := apiClient.AvoidPostingEntriesAPI.EditAvoidPosting(context.Background(), idOrArtistName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AvoidPostingEntriesAPIService GetAvoidPosting", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var idOrArtistName string

		resp, httpRes, err := apiClient.AvoidPostingEntriesAPI.GetAvoidPosting(context.Background(), idOrArtistName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AvoidPostingEntriesAPIService SearchAvoidPostings", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.AvoidPostingEntriesAPI.SearchAvoidPostings(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AvoidPostingEntriesAPIService UndeleteAvoidPosting", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var idOrArtistName string

		httpRes, err := apiClient.AvoidPostingEntriesAPI.UndeleteAvoidPosting(context.Background(), idOrArtistName).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
