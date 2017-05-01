package datastore

import (
	"testing"

	"cloud.google.com/go/datastore"

	"golang.org/x/net/context"

	e "github.com/laux-development/micro_app/profiles/domain/entity"
)

func TestWithdrawLowBal(t *testing.T) {
	ctx := context.Background()

	entity := &e.Profile{
		Address:    "a nice house in a nice street, anicepostcode",
		Company:    "Laux Development",
		DOB:        "20/01/2017",
		FirstName:  "James",
		Gender:     "male",
		GithubName: "bleeppurple",
		JobRole:    "Developer",
		LastName:   "Laux",
	}

	emptyEntity := &e.Profile{}

	profile := NewProfile(ctx, entity)
	err := profile.Save()

	if err != nil {
		t.Errorf("expected profile to save but got an error %s", err)
	}

	client, err := profile.client.Connection(profile.context)

	// return if client fails
	if err != nil {
		t.Errorf("datastore client wasnt created properly %s", err)
	}

	profile.key, err = client.Put(profile.context, datastore.IncompleteKey(entityName, nil), entity)

	if err != nil {
		t.Errorf("datastore client did not save properly %s", profile.key)
	}
	client.Get(profile.context, profile.key, emptyEntity)

	// return if client fails
	if emptyEntity.FirstName != "James" {
		t.Errorf("datastore client did not query properly %+v", emptyEntity)
	}

}
