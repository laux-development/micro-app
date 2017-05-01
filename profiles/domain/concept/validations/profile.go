package validations

import (
	"errors"

	e "github.com/laux-development/micro_app/profiles/domain/entity"
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
func (p Profile) Validate() (err error) {

	// set error to nil to initialize
	err = nil

	err = validateFirstName(p.FirstName)
	err = validateLastName(p.LastName)
	err = validateAddress(p.Address)

	return err
}

// ValidEntity validates and then returns a profile entity
func (p Profile) ValidEntity() (profile *e.Profile, err error) {

	err = p.Validate()
	if err != nil {
		return nil, err
	}
	profile = p.Profile

	return profile, nil

}

// validate firstname
func validateFirstName(thing string) error {
	if len(thing) < 2 || len(thing) > 40 {
		return errors.New(msgNameInvalid)
	}
	return nil
}

// validate lastname
func validateLastName(thing string) error {
	if len(thing) < 2 || len(thing) > 40 {
		return errors.New(msgNameInvalid)
	}
	return nil
}

// // validate address
func validateAddress(thing string) error {
	if len(thing) < 20 || len(thing) > 200 {
		return errors.New(msgAddressInvalid)
	}
	return nil
}
