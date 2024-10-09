/*
E621

Testing ForumTopicsAPIService

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

func Test_e621_ForumTopicsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ForumTopicsAPIService CreateForumTopic", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ForumTopicsAPI.CreateForumTopic(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ForumTopicsAPIService DeleteForumTopic", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id float32

		httpRes, err := apiClient.ForumTopicsAPI.DeleteForumTopic(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ForumTopicsAPIService EditForumTopic", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id float32

		httpRes, err := apiClient.ForumTopicsAPI.EditForumTopic(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ForumTopicsAPIService GetForumTopic", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id float32

		resp, httpRes, err := apiClient.ForumTopicsAPI.GetForumTopic(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ForumTopicsAPIService HideForumTopic", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id float32

		resp, httpRes, err := apiClient.ForumTopicsAPI.HideForumTopic(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ForumTopicsAPIService MarkAllForumTopicsAsRead", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.ForumTopicsAPI.MarkAllForumTopicsAsRead(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ForumTopicsAPIService SearchForumTopics", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ForumTopicsAPI.SearchForumTopics(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ForumTopicsAPIService SubscribeForumTopic", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id float32

		resp, httpRes, err := apiClient.ForumTopicsAPI.SubscribeForumTopic(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ForumTopicsAPIService UnhideForumTopic", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id float32

		resp, httpRes, err := apiClient.ForumTopicsAPI.UnhideForumTopic(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ForumTopicsAPIService UnsubscribeForumTopic", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id float32

		resp, httpRes, err := apiClient.ForumTopicsAPI.UnsubscribeForumTopic(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
