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

//type displayer interface {
//	Display()
//}

type outputtable interface {
	saver
	Display()
}

func main() {
	printSomething(1)
	printSomething(1.5)
	printSomething("Hello")

	title, content := getNoteData()
	todoText := getData("Todo content: ")

	todo, err := todo.New(todoText)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = outPutData(todo)
	if err != nil {
		return
	}

	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = outPutData(userNote)
	if err != nil {
		return
	}

	fmt.Println("Saving the not successded!")

}

func printSomething(value interface{}) { // type interface{} <= หมายความว่ารับทุกค่า
	intVal, ok := value.(int) // .(int) เป็นการ check type ท่าใหม่แหะ!
	if ok {
		fmt.Printf("Integer:", intVal)
		return
	}

	floatVal, ok := value.(float64)
	if ok {
		fmt.Printf("Float:", floatVal)
		return
	}

	stringVal, ok := value.(string)
	if ok {
		fmt.Printf("String:", stringVal)
		return
	}

	//switch value.(type) { // ท่าใหม่แหะ! Wow
	//case int:
	//	fmt.Println("Integer: ", value)
	//case float64:
	//	fmt.Println("Float: ", value)
	//case string:
	//	fmt.Println(value)
	//}
}

func outPutData(data outputtable) error {
	data.Display()
	return saveData(data)
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
