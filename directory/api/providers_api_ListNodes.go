package api

import (
	"encoding/json"
	"net/http"

	"github.com/zero-os/0-directory/directory/db/pool"
)

// ListNodes is the handler for GET /providers/{providerid}/resource_pools/nodes
// List all nodes
func (api ProvidersAPI) ListNodes(w http.ResponseWriter, r *http.Request) { // page := req.FormValue("page")// per_page := req.FormValue("per_page")
	var respBody []pool.Node
	json.NewEncoder(w).Encode(&respBody)
}
