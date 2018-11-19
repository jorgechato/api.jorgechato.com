package api

import (
	"net/http"
)

// Build the needed endpoints
func Build() (mux *http.ServeMux) {
	// Register handlers to routes.
	mux = http.NewServeMux()
	mux.Handle("/", graphqlHandler())
	mux.Handle("/health", healthHandler())

	return
}