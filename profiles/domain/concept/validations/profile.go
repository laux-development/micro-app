package validations

import (
	"errors"

	e "app/profiles/domain/entity"
)

const (
	msgNameInvalid    string = "Name needs to be between 2 and 40 characters"
	msgAddressInvalid string = "Address needs to be between 20 and 200 characters"
)

// Provide is a wrapper which provides Validations around an entity
type Profile struct {
	*e.Profile
}

// Validate validates a profile entity
func (p Profile) Validate() error {

	// validate firstname
	if len(p.FirstName) < 2 || len(p.FirstName) > 40 {
		return errors.New(msgNameInvalid)
	}

	// validate lastname
	if len(p.LastName) < 2 || len(p.LastName) > 40 {
		return errors.New(msgNameInvalid)
	}

	// validate address
	if len(p.Address) < 20 || len(p.Address) > 200 {
		return errors.New(msgAddressInvalid)
	}

	return nil
}
