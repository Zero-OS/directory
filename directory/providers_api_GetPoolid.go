package main

import (
	"encoding/json"
	"net/http"
)

// GetPoolid is the handler for GET /providers/{providerid}/resource_pools/{poolid}
// Get detail view about poolid
func (api ProvidersAPI) GetPoolid(w http.ResponseWriter, r *http.Request) {
	var respBody ResourcePool
	json.NewEncoder(w).Encode(&respBody)
}
