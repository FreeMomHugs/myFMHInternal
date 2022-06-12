package internal

import (
	"bytes"
	"context"
	"fmt"
	"google.golang.org/appengine/v2/datastore"
	"log"
)

func GetUserByEmail(ctx context.Context, Email string) (*MyFMHUser, *datastore.Key) {
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

func GetChapter(ctx context.Context, id int) *Chapter {
	key := datastore.NewKey(ctx, "chapter", "", id, nil)
	var x Chapter
	err := datastore.Get(ctx, key, &x)
	if err != nil {
		log.Println("unable to retrieve key " + err.Error())
	}
	return &x

}

func GetUser(ctx context.Context, UserID string) (*MyFMHUser, *datastore.Key) {
	q := datastore.NewQuery("myFMHUser").
		Filter("UserID=", UserID)
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
	Affiliation     int
	Title           string
	ProfilePicture  string
	AboutMe         string
	RecoveryEmail   string
	StripeAccountID string
	LocalSlackID    string
}

type Chapter struct {
	ChapterName  string
	State        string
	quickBooksID string
}

type QBAccount struct {
	Name                          string  `json:"Name"`
	SubAccount                    bool    `json:"SubAccount"`
	FullyQualifiedName            string  `json:"FullyQualifiedName"`
	Active                        bool    `json:"Active"`
	Classification                string  `json:"Classification"`
	AccountType                   string  `json:"AccountType"`
	AccountSubType                string  `json:"AccountSubType"`
	CurrentBalance                float64 `json:"CurrentBalance"`
	CurrentBalanceWithSubAccounts float64 `json:"CurrentBalanceWithSubAccounts"`
	CurrencyRef                   struct {
		Value string `json:"value"`
		Name  string `json:"name"`
	} `json:"CurrencyRef"`
	Domain    string `json:"domain"`
	Sparse    bool   `json:"sparse"`
	ID        string `json:"Id"`
	SyncToken string `json:"SyncToken"`
	MetaData  struct {
		CreateTime      string `json:"CreateTime"`
		LastUpdatedTime string `json:"LastUpdatedTime"`
	} `json:"MetaData"`
	ParentRef struct {
		Value string `json:"value"`
	} `json:"ParentRef,omitempty"`
}
