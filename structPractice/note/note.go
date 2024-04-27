package note

import (
	"encoding/json"
	"errors"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func New(title, content string) (*Note, error) {
	if title == "" || content == "" {
		return &Note{}, errors.New("title and content cannot be empty")
	}

	return &Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}

func Display(note *Note) {
	println("Title:", note.Title)
	println("Content:", note.Content)
	println("Created At:", note.CreatedAt.Format("2006-01-02 15:04:05"))
}

func (note *Note) SaveFile() error {
	fileName := strings.ReplaceAll(note.Title, " ", "_")
	fileName = strings.ToLower(fileName)
	json, err := json.Marshal(note)

	if err != nil {
		return err
	}

	return os.WriteFile(fileName+".json", json, 0644)
}
