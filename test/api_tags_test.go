/*
E621

Testing TagsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi-e621-go

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/anthrove/openapi-e621-go"
)

func Test_openapi-e621-go_TagsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test TagsAPIService CorrectTag", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id float32

		httpRes, err := apiClient.TagsAPI.CorrectTag(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TagsAPIService EditTag", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id float32

		httpRes, err := apiClient.TagsAPI.EditTag(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TagsAPIService GetTag", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id GetArtistIdOrNameParameter

		resp, httpRes, err := apiClient.TagsAPI.GetTag(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TagsAPIService PreviewTags", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.TagsAPI.PreviewTags(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TagsAPIService SearchTags", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.TagsAPI.SearchTags(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}