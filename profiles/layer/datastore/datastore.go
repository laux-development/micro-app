package datastore

import (
	"golang.org/x/net/context"

	ds "cloud.google.com/go/datastore"

	v "github.com/laux-development/micro_app/profiles/domain/concept/validations"
	e "github.com/laux-development/micro_app/profiles/domain/entity"
	client "github.com/laux-development/micro_app/profiles/layer/datastore_client"
)

const (
	entityName = "Profile"
	projectID  = "test"
)

// TODO: I think you could probably recude this code to just take an entity and produce a valid entity
// at the moment you have 2 pointers to the same entity

// Profile
// takes a request context (which is required by GAE)
// a Profile entity
// and Profile validator
type Profile struct {
	key     *ds.Key
	context context.Context
	// datastore client
	client client.ConnectionProvider
	// Validator interface for entity
	validator v.Provider
}

// Profile
// Returns pointer to Profile
// takes an request context and an Profile entity
func NewProfile(ctx context.Context, entity *e.Profile) (p *Profile) {
	p = &Profile{}

	p.context = ctx
	p.client = client.NewClient(projectID)
	p.validator = v.Profile{entity}

	return
}
