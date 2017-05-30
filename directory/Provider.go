package main

import (
	"gopkg.in/validator.v2"
)

type Provider struct {
	Description string `json:"description" validate:"nonzero"`
	Email       string `json:"email" validate:"nonzero"`
	Name        string `json:"name" validate:"regexp=^\w+$,nonzero"`
	Telephone   string `json:"telephone" validate:"nonzero"`
	Website_url string `json:"website_url" validate:"nonzero"`
}

func (s Provider) Validate() error {

	return validator.Validate(s)
}
