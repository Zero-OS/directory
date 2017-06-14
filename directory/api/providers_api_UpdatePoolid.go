package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	log "github.com/Sirupsen/logrus"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/gorilla/mux"
	"github.com/zero-os/0-directory/directory/db/pool"
)

// UpdatePoolid is the handler for PUT /providers/{providerid}/resource_pools/{poolid}
// Update poolid
func (api ProvidersAPI) UpdatePoolid(w http.ResponseWriter, r *http.Request) {
	poolID := mux.Vars(r)["poolid"]

	if !bson.IsObjectIdHex(poolID) {
		http.Error(w, "pool id is not valid", http.StatusBadRequest)
		return
	}

	org, err := getOrganization(r)
	if err != nil {
		http.Error(w, "No scope user:memberof present in JWT", http.StatusBadRequest)
		return
	}

	var reqBody pool.ResourcePoolCreate
	// decode request
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		log.Errorf("Error decoding create resource pool request: %v\n", err)
		http.Error(w, "Can't decode json", http.StatusBadRequest)
		return
	}

	respBody := &pool.ResourcePool{
		ResourcePoolCreate: reqBody,
		Organization:       org,
		Provider:           org,
		UID:                bson.ObjectIdHex(poolID),
	}

	poolMgr := pool.NewManager(r)
	if err := poolMgr.Update(respBody); err != nil {
		if err == mgo.ErrNotFound {
			http.Error(w, fmt.Sprintf("No pool with id (%s) found\n", poolID), http.StatusNotFound)
		} else {
			http.Error(w, fmt.Sprintf("Error updating pool with id (%s) :%v", poolID, err), http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&respBody)
}
