/*
E621

Testing WikiPageVersionsAPIService

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

func Test_openapi-e621-go_WikiPageVersionsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test WikiPageVersionsAPIService GetWikiPageVersion", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id float32

		resp, httpRes, err := apiClient.WikiPageVersionsAPI.GetWikiPageVersion(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WikiPageVersionsAPIService SearchWikiPageVersions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.WikiPageVersionsAPI.SearchWikiPageVersions(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
