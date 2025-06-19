package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"example.com/exercise-struct/note"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
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
	fmt.Printf("%v ", promptText)

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n") // TrimSuffix คือ ลบข้อความท้ายของ string
	text = strings.TrimSuffix(text, "\r")

	return text
}
