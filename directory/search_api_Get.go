package main

import (
	"encoding/json"
	"net/http"
)

// Get is the handler for GET /search
// Search for resource pools.
func (api SearchAPI) Get(w http.ResponseWriter, r *http.Request) { // address := req.FormValue("address")// currency := req.FormValue("currency")// datacenter_tier := req.FormValue("datacenter_tier")// lattitude := req.FormValue("lattitude")// longitute := req.FormValue("longitute")// min_cu := req.FormValue("min_cu")// min_su := req.FormValue("min_su")// min_tu := req.FormValue("min_tu")// network_redundant := req.FormValue("network_redundant")// network_speed := req.FormValue("network_speed")// page := req.FormValue("page")// per_page := req.FormValue("per_page")// provider := req.FormValue("provider")// uptime := req.FormValue("uptime")
	var respBody ResourcePool
	json.NewEncoder(w).Encode(&respBody)
}
