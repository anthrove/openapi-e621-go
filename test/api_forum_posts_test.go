/*
E621

Testing ForumPostsAPIService

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

func Test_e621_ForumPostsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ForumPostsAPIService CreateForumPost", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ForumPostsAPI.CreateForumPost(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ForumPostsAPIService DeleteForumPost", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id float32

		httpRes, err := apiClient.ForumPostsAPI.DeleteForumPost(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ForumPostsAPIService EditForumPost", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id float32

		httpRes, err := apiClient.ForumPostsAPI.EditForumPost(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ForumPostsAPIService GetForumPost", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id float32

		resp, httpRes, err := apiClient.ForumPostsAPI.GetForumPost(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ForumPostsAPIService HideForumPost", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id float32

		resp, httpRes, err := apiClient.ForumPostsAPI.HideForumPost(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ForumPostsAPIService MarkForumPost", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id float32

		resp, httpRes, err := apiClient.ForumPostsAPI.MarkForumPost(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ForumPostsAPIService SearchForumPosts", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ForumPostsAPI.SearchForumPosts(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ForumPostsAPIService UnhideForumPost", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id float32

		resp, httpRes, err := apiClient.ForumPostsAPI.UnhideForumPost(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
