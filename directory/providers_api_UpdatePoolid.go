package main

import (
	"encoding/json"
	"net/http"
)

// UpdatePoolid is the handler for PUT /providers/{providerid}/resource_pools/{poolid}
// Update poolid
func (api ProvidersAPI) UpdatePoolid(w http.ResponseWriter, r *http.Request) {
	var respBody ResourcePool
	json.NewEncoder(w).Encode(&respBody)
}
