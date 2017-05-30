package main

import (
	"encoding/json"
	"net/http"
)

// CreateProvideridresourcePool is the handler for POST /providers/{providerid}/resource_pools
// Create a new provideridresource_pool
func (api ProvidersAPI) CreateProvideridresourcePool(w http.ResponseWriter, r *http.Request) {
	var reqBody ResourcePoolCreate

	// decode request
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		w.WriteHeader(400)
		return
	}

	var respBody ResourcePool
	json.NewEncoder(w).Encode(&respBody)
}
