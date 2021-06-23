package store

import (
	"errors"
	"fmt"
	"log"

	note "braindumpnotes.com/notes/pkg"
	uuid "github.com/google/uuid"
)

type InMemStore struct {
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

func CreateInMemNoteStore(notes ...note.Note) *InMemStore {
	data := []note.Note{}
	data = append(data, notes...)

	return &InMemStore{
		data: data,
	}
}

func (store *InMemStore) List() ([]note.Note, error) {
	return store.data, nil
}

func (store *InMemStore) Get(name string) (note.Note, error) {
	return findNote(&store.data, name)
}

func (store *InMemStore) Create(inputNote note.InputNote) error {
	_, err := findNote(&store.data, inputNote.Name)
	if err == nil {
		return fmt.Errorf("note already exists: %s", inputNote.Name)
	}

	newNote := note.Note{
		Id:      uuid.NewString(),
		Name:    inputNote.Name,
		Content: inputNote.Content,
	}

	store.data = append(store.data, newNote)
	return nil
}

func (store *InMemStore) Update(fname string, updateNote note.UpdateNote) error {
	updated := false
	for i, n := range store.data {
		if n.Name == fname {
			if updateNote.Name != nil {
				n.Name = *updateNote.Name
			}
			if updateNote.Content != nil {
				n.Content = *updateNote.Content
			}
			store.data[i] = n
			log.Printf("Update note: %s", n)
			updated = true
		}
	}

	log.Printf("Update note: %s", store.data)

	if !updated {
		return fmt.Errorf("could not find note to update: %s", updateNote.Name)
	}
	return nil
}

func (store *InMemStore) Delete(noteName string) error {
	deleted := false
	result := []note.Note{}

	for i, n := range store.data {
		if n.Name == noteName {
			log.Printf("Delete note: %s", n)
		} else {
			result = append(result, store.data[i])
			deleted = true
		}
	}

	if !deleted {
		return fmt.Errorf("could not find note to delete: %s", noteName)
	}
	return nil
}
