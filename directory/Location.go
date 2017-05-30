package main

import (
	"gopkg.in/validator.v2"
)

type Location struct {
	Address     string  `json:"address" validate:"nonzero"`
	Description string  `json:"description" validate:"nonzero"`
	Lattitude   float64 `json:"lattitude" validate:"nonzero"`
	Longitute   float64 `json:"longitute" validate:"nonzero"`
	Name        string  `json:"name" validate:"regexp=^\w+$,nonzero"`
	Remarks     string  `json:"remarks" validate:"nonzero"`
}

func (s Location) Validate() error {

	return validator.Validate(s)
}
