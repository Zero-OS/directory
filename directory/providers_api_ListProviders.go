package main

import (
	"encoding/json"
	"net/http"
)

// ListProviders is the handler for GET /providers
// List all providers
func (api ProvidersAPI) ListProviders(w http.ResponseWriter, r *http.Request) { // page := req.FormValue("page")// per_page := req.FormValue("per_page")
	var respBody []Provider
	json.NewEncoder(w).Encode(&respBody)
}
