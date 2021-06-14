package note

import (
	"fmt"
	"strings"
)

const NOTE_FORMAT_EXT_MD = "MD"
const NOTE_FORMAT_EXT_TXT = "TXT"

type InputNote struct {
	Name    string `json:"name"`
	Content string `json:"content"`
	Format  string `json:"format"`
}

type UpdateNote struct {
	Name    string  `json:"name"`
	Content *string `json:"content,omitempty"`
}

type Note struct {
	Name    string `json:"name"`
	Content string `json:"content"`
	Format  string `json:"format"`
}

type NoteStore interface {
	List() ([]Note, error)
	Get(name string) (Note, error)
	Create(newNote InputNote) error
	Update(updateNote UpdateNote) error
	Delete(name string) error
}

func GetNoteExtFromFileName(fileName string) (string, error) {
	if strings.HasSuffix(fileName, "."+NOTE_FORMAT_EXT_MD) {
		return NOTE_FORMAT_EXT_MD, nil
	} else if strings.HasSuffix(fileName, "."+NOTE_FORMAT_EXT_TXT) {
		return NOTE_FORMAT_EXT_TXT, nil
	} else {
		return "", fmt.Errorf("unrecognized extension for fileName: %s", fileName)
	}
}
