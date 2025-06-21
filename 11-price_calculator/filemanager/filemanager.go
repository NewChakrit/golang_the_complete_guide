package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)

func Readlins(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, errors.New("Could not open file!")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file) // todo เป็นการเปิดเผยการใช้เนื้อหา

	var lines []string

	for scanner.Scan() { // todo อ่านค่าใน file แบบ line by line จึงใช้ for เพื่อ scan ทุกบรรทัด (response type: bool)
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err() // scanner.Scan() บอกแค่ true false จึงต้องใช้ scanner.Err() ในการเช็ค err
	if err != nil {
		file.Close()
		return nil, errors.New("Reading the file content failed.")
	}

	file.Close()
	return lines, nil
}

func WriteJSON(path string, data interface{}) error {
	file, err := os.Create(path)
	if err != nil {
		return errors.New("Failed to create file.")
	}

	encoder := json.NewEncoder(file) // todo แปลงข้อความเป็นรูปแบบ JSON
	err = encoder.Encode(data)
	if err != nil {
		file.Close()
		return errors.New("Failed to convert data to JSON.")
	}

	file.Close()
	return nil
}
