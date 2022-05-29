package internal

import (
	"bytes"
	"context"
	"fmt"
	"google.golang.org/appengine/v2/datastore"
)

func GetUserByEmail(Email string, ctx context.Context) (*MyFMHUser, *datastore.Key) {
	q := datastore.NewQuery("myFMHUser").
		Filter("Email=", Email)
	b := new(bytes.Buffer)
	var x MyFMHUser
	for t := q.Run(ctx); ; {
		key, err := t.Next(&x)
		if err == datastore.Done {
			break
		}
		if err != nil {
			panic(err)
			return &x, nil
		}
		fmt.Fprintf(b, "Key=%v\nWidget=%#v\n\n", key, x)
		return &x, key
	}
	return nil, nil
}

type MyFMHUser struct {
	UserID          string
	Email           string
	Name            string
	FirstName       string
	LastName        string
	Pronouns        string
	Phone           string
	Address1        string
	Address2        string
	City            string
	State           string
	ZipCode         string
	Affiliation     string
	Title           string
	ProfilePicture  string
	AboutMe         string
	RecoveryEmail   string
	StripeAccountID string
	LocalSlackID    string
}
