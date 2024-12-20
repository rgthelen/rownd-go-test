package main

import (
    "context"
    "flag"
    "fmt"
    "log"
    "os"
    
    "github.com/rgthelen/rownd-go-test/pkg/rownd"
)

func main() {
    // Parse command line flags
    var (
        appKey    = flag.String("app-key", os.Getenv("ROWND_APP_KEY"), "Rownd app key")
        appSecret = flag.String("app-secret", os.Getenv("ROWND_APP_SECRET"), "Rownd app secret")
        token     = flag.String("token", "", "Rownd authentication token")
    )
    flag.Parse()

    // Validate required flags
    if *appKey == "" || *appSecret == "" {
        log.Fatal("app-key and app-secret are required")
    }
    if *token == "" {
        log.Fatal("token is required")
    }

    // Create client with timeout context
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    client, err := rownd.NewClient(&rownd.Config{
        AppKey:    *appKey,
        AppSecret: *appSecret,
    })
    if err != nil {
        log.Fatalf("Failed to create client: %v", err)
    }

    // Run the validation
    if err := run(ctx, client, *token); err != nil {
        log.Fatal(err)
    }
}

func run(ctx context.Context, client *rownd.Client, token string) error {
    // Validate token
    tokenInfo, err := client.ValidateToken(ctx, token)
    if err != nil {
        return fmt.Errorf("validate token: %w", err)
    }

    // Get user profile
    user, err := client.GetUser(ctx, tokenInfo.UserID)
    if err != nil {
        return fmt.Errorf("get user: %w", err)
    }

    fmt.Printf("User profile: %+v\n", user)
    return nil
}