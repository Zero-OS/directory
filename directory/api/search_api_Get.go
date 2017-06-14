package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/zero-os/0-directory/directory/db/pool"
)

// Get is the handler for GET /search
// Search for resource pools.
func (api SearchAPI) Get(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	mgr := pool.NewManager(r)

	pools, err := mgr.List(pool.NewPoolQuery(r.Form))
	if err != nil {
		http.Error(w, fmt.Sprintf("Error listing the resource pools :%v", err), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&pools)
}
