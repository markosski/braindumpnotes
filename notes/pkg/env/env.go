package env

import (
	"fmt"

	notes "braindumpnotes.com/notes/pkg"
	store "braindumpnotes.com/notes/pkg/store"
)

type Env struct {
	NoteStore    notes.NoteStore
	AccountStore notes.AccountStore
}

func MakeLiveEnv() Env {
	return Env{}
}

func MakeInMemEnv() Env {
	fmt.Println("Constructing Environment")
	// var nstore note.NoteStore

	nstore := store.CreateInMemNoteStore(
		notes.Note{
			Id:      "1a2b3c",
			Name:    "first",
			Content: "First Note",
		},
	)

	astore := store.CreateInMemAccountStore(
		notes.Account{
			Id:       "123",
			UserName: "Marcin K",
			Email:    "marcin.kossakowski@gmail.com",
			StoreIntegration: notes.StoreIntegration{
				PlatformName: notes.PLATFORM_AWS,
				Config:       nil,
			},
		},
	)

	return Env{
		NoteStore:    nstore,
		AccountStore: astore,
	}
}
