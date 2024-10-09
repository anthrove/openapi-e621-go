/*
E621

Testing TagAliasesAPIService

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

func Test_e621_TagAliasesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test TagAliasesAPIService ApproveTagAlias", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id float32

		httpRes, err := apiClient.TagAliasesAPI.ApproveTagAlias(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TagAliasesAPIService CreateTagAlias", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.TagAliasesAPI.CreateTagAlias(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TagAliasesAPIService EditTagAlias", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id float32

		httpRes, err := apiClient.TagAliasesAPI.EditTagAlias(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TagAliasesAPIService GetTagAlias", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id float32

		resp, httpRes, err := apiClient.TagAliasesAPI.GetTagAlias(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TagAliasesAPIService RejectTagAlias", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id float32

		httpRes, err := apiClient.TagAliasesAPI.RejectTagAlias(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TagAliasesAPIService SearchTagAliases", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.TagAliasesAPI.SearchTagAliases(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
