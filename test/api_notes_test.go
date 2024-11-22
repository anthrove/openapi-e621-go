/*
E621

Testing NotesAPIService

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

func Test_e621_NotesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test NotesAPIService CreateNote", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.NotesAPI.CreateNote(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NotesAPIService DeleteNote", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id int32

		httpRes, err := apiClient.NotesAPI.DeleteNote(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NotesAPIService EditNote", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id int32

		httpRes, err := apiClient.NotesAPI.EditNote(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NotesAPIService GetNote", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id int32

		resp, httpRes, err := apiClient.NotesAPI.GetNote(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NotesAPIService RevertNote", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id int32

		httpRes, err := apiClient.NotesAPI.RevertNote(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NotesAPIService SearchNotes", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.NotesAPI.SearchNotes(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
