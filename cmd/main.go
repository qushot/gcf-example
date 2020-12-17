package main

import (
	"context"
	"gcfexample"
	"log"
	"os"

	"github.com/GoogleCloudPlatform/functions-framework-go/funcframework"
)

func main() {
	ctx := context.Background()

	// HTTPEntrypoint
	if err := funcframework.RegisterHTTPFunctionContext(ctx, "/http-entry-point", gcfexample.HTTPEntryPoint); err != nil {
		log.Fatalf("funcframework.RegisterHTTPFunctionContext: %v\n", err)
	}

	// BackgroundPubSubEntryPoint
	if err := funcframework.RegisterEventFunctionContext(ctx, "/background-pubsub-entry-point", gcfexample.BackgroundPubSubEntryPoint); err != nil {
		log.Fatalf("funcframework.RegisterEventFunctionContext: %v\n", err)
	}

	// Use PORT environment variable, or default to 8080.
	port := "8080"
	if envPort := os.Getenv("PORT"); envPort != "" {
		port = envPort
	}
	if err := funcframework.Start(port); err != nil {
		log.Fatalf("funcframework.Start: %v\n", err)
	}
}
