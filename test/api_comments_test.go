/*
E621

Testing CommentsAPIService

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

func Test_e621_CommentsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test CommentsAPIService CreateComment", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.CommentsAPI.CreateComment(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CommentsAPIService DeleteComment", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id int32

		httpRes, err := apiClient.CommentsAPI.DeleteComment(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CommentsAPIService EditComment", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id int32

		httpRes, err := apiClient.CommentsAPI.EditComment(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CommentsAPIService GetComment", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id int32

		resp, httpRes, err := apiClient.CommentsAPI.GetComment(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CommentsAPIService HideComment", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id int32

		resp, httpRes, err := apiClient.CommentsAPI.HideComment(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CommentsAPIService MarkComment", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id int32

		resp, httpRes, err := apiClient.CommentsAPI.MarkComment(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CommentsAPIService SearchComments", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.CommentsAPI.SearchComments(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CommentsAPIService UnhideComment", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id int32

		resp, httpRes, err := apiClient.CommentsAPI.UnhideComment(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
