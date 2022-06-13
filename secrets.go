package internal

import (
	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"context"
	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"
	"log"
)

func AccessSecretVersion(ctx context.Context, name string) []byte {
	// Create the client.
	cachedResult, success := GetFromCache(ctx, name)
	if success {
		return cachedResult
	}
	client, err := secretmanager.NewClient(ctx)
	if err != nil {
		log.Println("failed to create secretmanager client: " + err.Error())
	}
	defer client.Close()

	// Build the request.
	req := &secretmanagerpb.AccessSecretVersionRequest{
		Name: name,
	}

	// Call the API.
	result, err := client.AccessSecretVersion(ctx, req)
	if err != nil {
		log.Println("failed to access secret version: " + err.Error())
	}

	AddToCache(ctx, name, result.Payload.Data)

	return result.Payload.Data
}
