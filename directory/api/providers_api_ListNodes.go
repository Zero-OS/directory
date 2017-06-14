package api

import (
	"encoding/json"
	"net/http"

	"gopkg.in/mgo.v2/bson"

	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
	"github.com/zero-os/0-directory/directory/db/pool"
)

// ListNodes is the handler for GET /providers/{providerid}/resource_pools/{poolid}/nodes
// List all nodes
func (api ProvidersAPI) ListNodes(w http.ResponseWriter, r *http.Request) { // page := req.FormValue("page")// per_page := req.FormValue("per_page")
	poolID := mux.Vars(r)["poolid"]

	if !bson.IsObjectIdHex(poolID) {
		http.Error(w, "pool id is not valid", http.StatusBadRequest)
		return
	}

	mgr := pool.NewManager(r)
	nodes, err := mgr.ListNodes(poolID)
	if err != nil {
		http.Error(w, "error retrieving nodes: %v\n", http.StatusInternalServerError)
		return
	}
	log.Debugf("nodes: %+v", nodes)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&nodes)
}
