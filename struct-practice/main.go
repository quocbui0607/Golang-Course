package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	note "test.com/note/note"
	todo "test.com/note/todo"
)

type Saver interface {
	Save() error
}

type Output interface {
	Saver
	Display()
}

func main() {
	title, content := getNoteData()

	toDoText := getUserInput("Todo text:")

	toDo, _ := todo.New(toDoText)
	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	outputData(toDo)
	outputData(userNote)

}

func outputData(output Output) {
	output.Display()
	saveData(output)
}

func saveData(saver Saver) {
	err := saver.Save()

	if err != nil {
		fmt.Println("Saving failed", err)
		return
	}

	fmt.Println("Saving succeeded")
}

func getNoteData() (string, string) {
	title := getUserInput("Note title: ")
	content := getUserInput("Note content: ")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
