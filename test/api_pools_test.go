/*
E621

Testing PoolsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi-e621-go

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func Test_openapi-e621-go_PoolsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test PoolsAPIService AddPostToPool", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.PoolsAPI.AddPostToPool(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PoolsAPIService CreatePool", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.PoolsAPI.CreatePool(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PoolsAPIService DeletePool", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id float32

		httpRes, err := apiClient.PoolsAPI.DeletePool(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PoolsAPIService EditPool", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id float32

		httpRes, err := apiClient.PoolsAPI.EditPool(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PoolsAPIService GetPool", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id float32

		resp, httpRes, err := apiClient.PoolsAPI.GetPool(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PoolsAPIService RemovePostFromPool", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.PoolsAPI.RemovePostFromPool(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PoolsAPIService RevertPool", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id float32

		httpRes, err := apiClient.PoolsAPI.RevertPool(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PoolsAPIService SearchPools", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.PoolsAPI.SearchPools(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
