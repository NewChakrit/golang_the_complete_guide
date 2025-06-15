package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func WriteFloatTofile(fileName string, value float64) {
	valueStr := fmt.Sprint(value)
	err := os.WriteFile(fileName, []byte(valueStr), 0644)
	if err != nil {
		panic(err)
	}
}

func GetFloatFromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return 10000, errors.New("[ERROR] Failed to find file.\nSet default balance to  .\n")
	}

	balStr := string(data)
	balance, err := strconv.ParseFloat(balStr, 64)
	if err != nil {
		return 10000, errors.New("Cannot parse stored balance value.\n")
	}

	return balance, nil
}
