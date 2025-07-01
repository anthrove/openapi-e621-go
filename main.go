package main

import (
	"context"
	"github.com/goforj/godump"
)

//type AccessDeniedReason string
//type AccessDeniedSuccess string
//type MessageErrorSuccess string
//type NotFoundReason string
//type NotFoundSuccess string
//type WarningRecordType string

func main() {
	c, err := NewClientWithResponses("https://e621.net")
	if err != nil {
		panic(err)
	}
	userID := 136501

	params := &ListFavoritesParams{
		UserId: &userID,
	}

	resp, err := c.ListFavoritesWithResponse(context.Background(), params)
	if err != nil {
		panic(err)
	}

	godump.Dump(resp.JSON200.Posts[0])
}
