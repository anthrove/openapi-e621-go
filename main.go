package main

import (
	"context"
	"strconv"

	"github.com/goforj/godump"
)

func main() {
	c, err := NewClientWithResponses("https://e621.net")
	if err != nil {
		panic(err)
	}
	userID := 136501

	resp, err := c.GetUserWithResponse(context.Background(), strconv.Itoa(userID))
	if err != nil {
		panic(err)
	}

	godump.Dump(resp.JSON200.AsFullUser())
}
