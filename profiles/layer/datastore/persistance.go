package datastore

import "cloud.google.com/go/datastore"

// Save
// This persists a StockTransaction entity to the google datastore
// Writes a key on save to the struct
// and returns an error if it fails
func (store *Profile) Save() (err error) {
	// create client
	// create datastore client
	client, err := store.client.Connection(store.context)

	// return if client fails
	if err != nil {
		return err
	}

	// run validator, return valid entity
	entity, err := store.validator.ValidEntity()

	// return if validation fails
	if err != nil {
		return err
	}

	// create new key for use in datastore
	// then upload entity with that key
	_, err = client.Put(store.context, datastore.IncompleteKey(entityName, nil), entity)

	return

}
