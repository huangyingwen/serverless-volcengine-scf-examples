package main

import (
	"context"
	"fmt"

	"github.com/volcengine/vefaas-golang-runtime/events"
	"github.com/volcengine/vefaas-golang-runtime/vefaas"
	"github.com/volcengine/vefaas-golang-runtime/vefaascontext"
)

func main() {
	// Start your vefaas function =D.
	vefaas.Start(handler)
}

// Define your handler function.
func handler(ctx context.Context, r *events.HTTPRequest) (*events.EventResponse, error) {
	fmt.Printf("received new request: %s %s, request id: %s\n", r.HTTPMethod, r.Path, vefaascontext.RequestIdFromContext(ctx))

	return &events.EventResponse{
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: []byte("Hello veFaaS!"),
	}, nil
}
