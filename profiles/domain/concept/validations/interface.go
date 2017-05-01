package validations

import (
	e "github.com/laux-development/micro_app/profiles/domain/entity"
)

// ValidationProvider provides validation behaviour
type Provider interface {
	Validate() error
	ValidEntity() (*e.Profile, error)
}
