package pool

import (
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/validator.v2"
)

type ResourcePool struct {
	ResourcePoolCreate
	UID          bson.ObjectId `json:"UID" validate:"nonzero" bson:"_id,omitempty"`
	Organization string        `json:"organization" validate:"nonzero"`
	Provider     string        `json:"provider" validate:"nonzero"`
	Quality      PoolQuality   `json:"quality" validate:"nonzero"`
}

func (s ResourcePool) Validate() error {

	return validator.Validate(s)
}
