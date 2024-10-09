/*
E621

Testing PostApprovalsAPIService

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

func Test_e621_PostApprovalsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test PostApprovalsAPIService ApprovePost", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.PostApprovalsAPI.ApprovePost(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PostApprovalsAPIService SearchPostApprovals", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.PostApprovalsAPI.SearchPostApprovals(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PostApprovalsAPIService UnapprovePost", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.PostApprovalsAPI.UnapprovePost(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
