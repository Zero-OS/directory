package pool

import (
	"gopkg.in/validator.v2"
)

type ResourcePoolCreate struct {
	Cu_max                    int                                    `json:"cu_max" validate:"nonzero"`
	Cu_planned                int                                    `json:"cu_planned" validate:"nonzero"`
	Datacenter_tier           EnumResourcePoolCreateDatacenter_tier  `json:"datacenter_tier" validate:"nonzero"`
	Datacenter_uptime_sla_min float64                                `json:"datacenter_uptime_sla_min" validate:"nonzero"`
	Ipv4_enabled              bool                                   `json:"ipv4_enabled"`
	Ipv6_enabled              bool                                   `json:"ipv6_enabled"`
	Ipv4_nr_pub               int                                    `json:"ipv4_nr_pub" validate:"nonzero"`
	Network_redundant         bool                                   `json:"network_redundant"`
	Network_speed             int                                    `json:"network_speed" validate:"nonzero"`
	Pricing_currency          EnumResourcePoolCreatePricing_currency `json:"pricing_currency" validate:"nonzero"`
	Su_max                    int                                    `json:"su_max" validate:"nonzero"`
	Su_planned                int                                    `json:"su_planned" validate:"nonzero"`
	Tu_max                    int                                    `json:"tu_max" validate:"nonzero"`
	Tu_planned                int                                    `json:"tu_planned" validate:"nonzero"`
	Pricings                  []Pricing                              `json:pricings validate:"nonzero"`
	Nodes                     []Node                                 `json:"nodes" validate:"nonzero"`
	Location                  Location                               `json:"Location"`
}

func (s ResourcePoolCreate) Validate() error {

	return validator.Validate(s)
}
