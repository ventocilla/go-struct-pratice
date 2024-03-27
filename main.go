package main

import (
	"example.com/note/note"
	"fmt"
)

func main() {
	title, content := getNoteData()

	userNote, err := Note.New(title, content)

	fmt.Println(title)
	fmt.Println(content)
}

func getNoteData() (string, string) {
	title := getUserInput("Note title:")
	content := getUserInput("Note content")
	return title, content
}

func getUserInput(prompt string) string {
	fmt.Println(prompt)
	var value string
	fmt.Scanln(&value)
	return value
}
