/*
E621

Testing UserFeedbacksAPIService

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

func Test_e621_UserFeedbacksAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test UserFeedbacksAPIService CreateUserFeedback", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.UserFeedbacksAPI.CreateUserFeedback(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserFeedbacksAPIService DeleteUserFeedback", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id float32

		httpRes, err := apiClient.UserFeedbacksAPI.DeleteUserFeedback(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserFeedbacksAPIService DestroyUserFeedback", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id float32

		httpRes, err := apiClient.UserFeedbacksAPI.DestroyUserFeedback(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserFeedbacksAPIService EditUserFeedback", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id float32

		httpRes, err := apiClient.UserFeedbacksAPI.EditUserFeedback(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserFeedbacksAPIService GetUserFeedback", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id float32

		resp, httpRes, err := apiClient.UserFeedbacksAPI.GetUserFeedback(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserFeedbacksAPIService SearchUserFeedbacks", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.UserFeedbacksAPI.SearchUserFeedbacks(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserFeedbacksAPIService UndeleteUserFeedback", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id float32

		httpRes, err := apiClient.UserFeedbacksAPI.UndeleteUserFeedback(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}