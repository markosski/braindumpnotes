package store

import (
	"errors"

	notes "braindumpnotes.com/notes/pkg"
)

type InMemAccountStore struct {
	data []notes.Account
}

func CreateInMemAccountStore(accounts ...notes.Account) *InMemAccountStore {
	data := []notes.Account{}
	data = append(data, accounts...)

	return &InMemAccountStore{
		data: data,
	}
}

func (store *InMemAccountStore) Get(email notes.Email) (notes.Account, error) {
	data := &store.data
	for _, acc := range *data {
		if acc.Email == email {
			return acc, nil
		}
	}
	return notes.Account{}, errors.New("could not find note")
}

func (store *InMemAccountStore) Create(account notes.Account) error {
	data := &store.data
	*data = append(*data, account)
	return nil
}
