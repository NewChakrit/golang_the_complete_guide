package main

import (
	"encoding/json"
	"errors"
	"example.com/exercise-struct/note"
	"fmt"
	"io/ioutil"
)

func main() {
	title, content := getNoteData()

	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}

	userNote.Display()

	jsonData, err := json.Marshal(&userNote)
	if err != nil {
		fmt.Println(errors.New("Cannot mashall json."))
	}
	err = ioutil.WriteFile("learn_go.json", jsonData, 0644)
	if err != nil {
		fmt.Println(errors.New("Cannot write file"))
	}
}

func getNoteData() (string, string) {
	title := getData("Note title: ")
	content := getData("Note content: ")

	return title, content
}

func getData(promptText string) string {
	var value string
	fmt.Print(promptText)
	fmt.Scanln(&value)

	if value == "" {
		return ""
	}

	return value
}
