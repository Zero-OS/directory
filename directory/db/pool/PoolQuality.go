package pool

import (
	"gopkg.in/validator.v2"
)

type PoolQuality struct {
	Average_restauration_time int          `json:"average_restauration_time" validate:"nonzero"`
	Down_periods              []DownPeriod `json:"down_periods" validate:"nonzero"`
	Uptime                    int          `json:"uptime" validate:"nonzero"`
}

func (s PoolQuality) Validate() error {

	return validator.Validate(s)
}
