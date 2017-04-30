package datastore

import (
	"context"

	"cloud.google.com/go/datastore"

	v "app/profiles/domain/concept/validations"
	e "app/profiles/domain/entity"
)

const (
	entityName = "Profile"
)

// Profile
// takes a request context (which is required by GAE)
// a Profile entity
// and Profile validator
type Profile struct {
	key     *datastore.Key
	context context.Context
	entity  *e.Profile

	// Validator interface for entity
	validator v.Provider
}

// Profile
// Returns pointer to Profile
// takes an request context and an Profile entity
func NewProfile(ctx context.Context, entity *e.Profile) (p *Profile) {
	p = &Profile{}

	p.context = ctx
	p.entity = entity
	p.validator = v.Profile{entity}

	return
}
