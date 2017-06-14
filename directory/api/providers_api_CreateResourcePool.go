package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/zaibon/identityserver/db"
	"github.com/zero-os/0-directory/directory/db/pool"
	"github.com/zero-os/0-directory/directory/db/provider"
)

// CreateResourcePool is the handler for POST /providers/resource_pools
func (api ProvidersAPI) CreateResourcePool(w http.ResponseWriter, r *http.Request) {

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
	}

	poolMgr := pool.NewManager(r)
	providerMgr := provider.NewManager(r)

	p := &provider.Provider{Name: org}
	if err := providerMgr.Create(p); err != nil && err != db.ErrDuplicate {
		http.Error(w, fmt.Sprintf("error creating provider : %v\n", err), http.StatusInternalServerError)
		return
	}

	if err := poolMgr.Create(respBody); err != nil {
		log.Error(err)
		http.Error(w, fmt.Sprintf("error creating pool : %v\n", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&respBody)
}
