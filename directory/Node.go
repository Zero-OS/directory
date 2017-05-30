package main

import (
	"gopkg.in/validator.v2"
)

type Node struct {
	Description string `json:"description" validate:"nonzero"`
	Nr_cu       int    `json:"nr_cu" validate:"nonzero"`
	Nr_su       int    `json:"nr_su" validate:"nonzero"`
	Nr_tu       int    `json:"nr_tu" validate:"nonzero"`
	Remarks     string `json:"remarks" validate:"nonzero"`
	Uid         string `json:"uid" validate:"regexp=^\w+$,nonzero"`
}

func (s Node) Validate() error {

	return validator.Validate(s)
}
