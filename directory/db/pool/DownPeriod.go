package pool

import (
	"github.com/zero-os/0-directory/directory/goraml"
	"gopkg.in/validator.v2"
)

type DownPeriod struct {
	Begin    goraml.DateTime `json:"begin" validate:"nonzero"`
	Duration int             `json:"duration" validate:"nonzero"`
}

func (s DownPeriod) Validate() error {

	return validator.Validate(s)
}
