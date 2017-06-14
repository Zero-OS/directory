package api

import (
	"fmt"
	"net/http"

	log "github.com/Sirupsen/logrus"

	"github.com/gorilla/mux"
	"github.com/zero-os/0-directory/directory/db/pool"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// DeletePool is the handler for DELETE /providers/resource_pools/{poolid}
func (api ProvidersAPI) DeletePool(w http.ResponseWriter, r *http.Request) {
	poolID := mux.Vars(r)["poolid"]

	if !bson.IsObjectIdHex(poolID) {
		http.Error(w, "pool id is not valid", http.StatusBadRequest)
		return
	}

	mgr := pool.NewManager(r)
	pool, err := mgr.Get(poolID)
	if err != nil {
		if err == mgo.ErrNotFound {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		http.Error(w, fmt.Sprintf("Error deleting the resource pool with id %s :%v", poolID, err), http.StatusInternalServerError)
		return
	}

	org, err := getOrganization(r)
	if err != nil {
		http.Error(w, "No scope user:memberof present in JWT", http.StatusBadRequest)
		return
	}

	log.Debug(pool.Organization, org)
	if pool.Organization != org {
		http.Error(w, "This resource pool doesn't belong to you. You can not delete it", http.StatusForbidden)
		return
	}

	err = mgr.Delete(poolID)
	if err != nil && err != mgo.ErrNotFound {
		http.Error(w, fmt.Sprintf("Error deleting the resource pool with id %s :%v", poolID, err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
