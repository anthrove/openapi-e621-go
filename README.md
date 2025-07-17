# e621 Golang Client API

This SDK is generated with the [OpenAPI specification](https://github.com/DonovanDMC/E621OpenAPI) by [DonovanDMC](https://furry.cool/)

## Building

To build the SDK yourself you need Go 1.24.4 and git.
1. `go mod tidy`
2. `go generate`
3. Profit

## Usage

```go
func main() {
    ctx := context.Background()
    
        // Step 1: Define server information and configuration
        server := "https://e621.net"
        config := ClientConfig{
            Username: "my-user-name", // Replace with actual username
            APIKey:   "my-api-key",   // Replace with actual API key
        }
    
        // Step 2: Initialize the API client
        client, err := NewClientWithResponses(
            server,
            WithUserAgent(config.Username),
            WithAuthorization(config.Username, config.APIKey),
            WithRateLimiter(),
        )
        if err != nil {
            log.Fatalf("failed to create client: %v", err)
        }
    
        // Step 3: Fetch user details using the client
        userResp, err := client.GetUserWithResponse(ctx, "duke")
        if err != nil {
            log.Fatalf("failed to get user: %v", err)
        }
    
        // Step 4: Parse the user response
        // Convert the response into a usable form to access user data
        user, err := userResp.JSON200.AsFullUser()
        if err != nil {
            log.Fatalf("failed to parse user response: %v", err)
        }
    
        // Step 5: Set parameters for listing favorite posts
        favParams := &ListFavoritesParams{
            Limit:  nil,     // No limit on the number of results (could be customized)
            Page:   nil,     // Defaults to first page of results
            UserId: &user.Id, // Use the user ID obtained previously
        }
    
        // Step 6: Request favorite posts
        favoritesResp, err := client.ListFavoritesWithResponse(ctx, favParams)
        if err != nil {
            // Log and terminate the program if the request for favorites fails
            log.Fatalf("failed to list favorites: %v", err)
        }
    
        // Step 7: Verify response status code
        if favoritesResp.StatusCode() != 200 {
            errorMessage := fmt.Sprintf("unexpected status code received: %d", favoritesResp.StatusCode())
            log.Fatal(errors.New(errorMessage))
        }
    
        // Step 8: Access the retrieved favorite posts
        favorites := favoritesResp.JSON200.Posts
    
        // Step 9: Output the first favorite post ID, if available
        if len(favorites) > 0 {
            fmt.Printf("First Favorite Post ID: %d\n", favorites[0].Id)
        } else {
            log.Println("No favorite posts found.")
        }
}
```