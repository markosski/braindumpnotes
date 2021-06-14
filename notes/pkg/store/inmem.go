package store

import (
	"errors"
	"fmt"

	note "braindumpnotes.com/notes/pkg"
)

type InMemoryStore struct {
	data []note.Note
}

func findNote(notes *[]note.Note, noteName string) (note.Note, error) {
	for _, n := range *notes {
		if n.Name == noteName {
			return n, nil
		}
	}
	return note.Note{}, errors.New("could not find note")
}

func Create(notes ...note.Note) *InMemoryStore {
	data := []note.Note{}
	data = append(data, notes...)

	return &InMemoryStore{
		data: data,
	}
}

func (store *InMemoryStore) List() ([]note.Note, error) {
	return store.data, nil
}

func (store *InMemoryStore) Get(name string) (note.Note, error) {
	return findNote(&store.data, name)
}

func (store *InMemoryStore) Create(inputNote note.InputNote) error {
	_, err := findNote(&store.data, inputNote.Name)
	if err == nil {
		return fmt.Errorf("note already exists: %s", inputNote.Name)
	}

	newNote := note.Note{
		Name:    inputNote.Name,
		Content: inputNote.Content,
		Format:  inputNote.Format,
	}

	store.data = append(store.data, newNote)
	return nil
}

func (store InMemoryStore) Update(updateNote note.UpdateNote) error {
	var foundNote *note.Note

	for _, n := range store.data {
		if n.Name == updateNote.Name {
			foundNote = &n
		}
	}

	if foundNote == nil {
		return fmt.Errorf("could not find note to update: %s", updateNote.Name)
	}

	if updateNote.Content != nil {
		foundNote.Content = *updateNote.Content
	}
	return nil
}

func (store *InMemoryStore) Delete(noteName string) error {
	return fmt.Errorf("could not find note to delete: %s", noteName)
}
