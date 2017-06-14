package pool

import (
	"gopkg.in/validator.v2"
)

type Pricing struct {
	Cu            int `json:"cu" validate:"nonzero"`
	Min_nr_months int `json:"min_nr_months" validate:"nonzero"`
	Su            int `json:"su" validate:"nonzero"`
	Tu            int `json:"tu" validate:"nonzero"`
}

func (s Pricing) Validate() error {

	return validator.Validate(s)
}
