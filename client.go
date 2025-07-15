package e621

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"

	"golang.org/x/time/rate"
)

type ClientConfig struct {
	Username string
	APIKey   string
}

func WithRateLimiter() ClientOption {
	return func(c *Client) error {
		limiter := rate.NewLimiter(1, 2)

		c.RequestEditors = append(c.RequestEditors, func(ctx context.Context, req *http.Request) error {
			// Wait for rate limiter
			err := limiter.Wait(context.Background())
			if err != nil {
				return err
			}
			return nil
		})
		return nil
	}

}

func WithAuthorization(username, apiKey string) ClientOption {
	return func(c *Client) error {
		c.RequestEditors = append(c.RequestEditors, func(ctx context.Context, req *http.Request) error {
			if username != "" && apiKey != "" {
				authHeader := "Basic " + base64.StdEncoding.EncodeToString([]byte(username+":"+apiKey))
				req.Header.Set("Authorization", authHeader)
			}
			return nil
		})
		return nil
	}
}

func WithUserAgent(username string) ClientOption {
	return func(c *Client) error {
		c.RequestEditors = append(c.RequestEditors, func(ctx context.Context, req *http.Request) error {
			if len(username) == 0 {
				return errors.New("username must be set")
			}

			req.Header.Set("User-Agent", fmt.Sprintf("Go-e621-SDK used by %s | (made by the Anthrove Team)", username))

			return nil
		})
		return nil
	}
}
