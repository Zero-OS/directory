package api

import (
	"encoding/json"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/zero-os/0-directory/directory/db/provider"
)

// ListProviders is the handler for GET /providers
// List all providers
func (api ProvidersAPI) ListProviders(w http.ResponseWriter, r *http.Request) {
	// page := req.FormValue("page")
	// per_page := req.FormValue("per_page")

	mgr := provider.NewManager(r)
	providers, err := mgr.List()
	if err != nil {
		http.Error(w, "error retrieving providers: %v\n", http.StatusInternalServerError)
		return
	}
	log.Debugf("%+v", providers)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&providers)
}
