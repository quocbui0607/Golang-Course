package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title    string    `json:"title"`
	Content  string    `json:"content"`
	CreateAt time.Time `json:"createAte"`
}

func New(title, content string) (Note, error) {
	if title == "" || content == "" {
		fmt.Println("Invalid title or content")
		return Note{}, errors.New("Invalid title/content")
	}

	return Note{
		Title:    title,
		Content:  content,
		CreateAt: time.Now(),
	}, nil
}

func (note Note) Display() {
	fmt.Printf("Your note title: %v with content:\n%v\n", note.Title, note.Content)
}

func (note Note) Save() error {
	fileName := strings.ReplaceAll(note.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"

	json, err := json.Marshal(note)

	if err != nil {
		fmt.Println(err)
		return err
	}

	return os.WriteFile(fileName, json, 0644)
}
