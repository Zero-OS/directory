package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/gorilla/mux"
	"github.com/zero-os/0-directory/directory/db/pool"
)

// GetPool is the handler for GET /providers/resource_pools/{poolid}
func (api ProvidersAPI) GetPool(w http.ResponseWriter, r *http.Request) {
	poolID := mux.Vars(r)["poolid"]

	if !bson.IsObjectIdHex(poolID) {
		http.Error(w, "pool id is not valid", http.StatusBadRequest)
		return
	}

	mgr := pool.NewManager(r)
	pool, err := mgr.Get(poolID)
	if err != nil {
		if err == mgo.ErrNotFound {
			http.Error(w, fmt.Sprintf("Not resource pool with this id (%s) found", poolID), http.StatusNotFound)
		} else {
			http.Error(w, fmt.Sprintf("Error retrieving the resource pool with id %s :%v", poolID, err), http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&pool)
}
