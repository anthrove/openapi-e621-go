/*
E621

Testing ArtistsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package e621

import (
	"context"
	openapiclient "github.com/anthrove/openapi-e621-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_e621_ArtistsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ArtistsAPIService CreateArtist", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ArtistsAPI.CreateArtist(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ArtistsAPIService DeleteArtist", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var idOrName GetArtistIdOrNameParameter

		httpRes, err := apiClient.ArtistsAPI.DeleteArtist(context.Background(), idOrName).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ArtistsAPIService EditArtist", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var idOrName GetArtistIdOrNameParameter

		httpRes, err := apiClient.ArtistsAPI.EditArtist(context.Background(), idOrName).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ArtistsAPIService GetArtist", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var idOrName GetArtistIdOrNameParameter

		resp, httpRes, err := apiClient.ArtistsAPI.GetArtist(context.Background(), idOrName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ArtistsAPIService RevertArtist", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var idOrName GetArtistIdOrNameParameter

		httpRes, err := apiClient.ArtistsAPI.RevertArtist(context.Background(), idOrName).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ArtistsAPIService SearchArtists", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ArtistsAPI.SearchArtists(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}