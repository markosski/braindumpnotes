package env

import (
	"fmt"

	note "braindumpnotes.com/notes/pkg"
	store "braindumpnotes.com/notes/pkg/store"
)

type Env struct {
	Store note.NoteStore
}

func MakeLiveEnv() Env {
	return Env{}
}

func MakeInMemEnv() Env {
	fmt.Println("Constructing Environment")
	// var nstore note.NoteStore

	nstore := store.Create(
		note.Note{
			Name:    "first",
			Content: "First Note",
			Format:  note.NOTE_FORMAT_EXT_MD},
	)

	return Env{
		Store: nstore,
	}
}
