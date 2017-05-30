package main

import (
	"encoding/json"
	"net/http"
)

// ListNodes is the handler for GET /providers/{providerid}/resource_pools/nodes
// List all nodes
func (api ProvidersAPI) ListNodes(w http.ResponseWriter, r *http.Request) { // page := req.FormValue("page")// per_page := req.FormValue("per_page")
	var respBody []Node
	json.NewEncoder(w).Encode(&respBody)
}
