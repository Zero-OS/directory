package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/zero-os/0-directory/directory/db/pool"
)

// ListResourcePools is the handler for GET /providers/{providerid}/resource_pools
func (api ProvidersAPI) ListResourcePools(w http.ResponseWriter, r *http.Request) {
	mgr := pool.NewManager(r)
	r.ParseForm()
	pools, err := mgr.List(pool.NewPoolQuery(r.Form))
	if err != nil {
		http.Error(w, fmt.Sprintf("Error listing the resource pools :%v", err), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&pools)
}
