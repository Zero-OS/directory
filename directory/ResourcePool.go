package main

import (
	"gopkg.in/validator.v2"
)

type ResourcePool struct {
	ResourcePoolCreate
	Provider string      `json:"provider" validate:"nonzero"`
	Quality  PoolQuality `json:"quality" validate:"nonzero"`
}

func (s ResourcePool) Validate() error {

	return validator.Validate(s)
}
