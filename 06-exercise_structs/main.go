package main

import (
	"bufio"
	"example.com/exercise-struct/note"
	"example.com/exercise-struct/todo"
	"fmt"
	"os"
	"strings"
)

type saver interface { // If interface has 1 func, should declare variable follow by that one func name and add "er"
	Save() error
}

func main() {
	title, content := getNoteData()
	todoText := getData("Todo content: ")

	todo, err := todo.New(todoText)
	if err != nil {
		fmt.Println(err)
		return
	}

	todo.Display()
	err = saveData(todo)
	if err != nil {
		return
	}

	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}

	userNote.Display()
	err = saveData(userNote)
	if err != nil {
		return
	}

	fmt.Println("Saving the not successded!")

	//jsonData, err := json.Marshal(&userNote)
	//if err != nil {
	//	fmt.Println(errors.New("Cannot mashall json."))
	//}
	//err = ioutil.WriteFile("learn_go.json", jsonData, 0644)
	//if err != nil {
	//	fmt.Println(errors.New("Cannot write file"))
	//}
}

func saveData(data saver) error {
	err := data.Save()
	if err != nil {
		fmt.Println("Saving the note failed.")
		return err
	}

	return nil
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
