package validation

import (
	"errors"
	"strings"
)

func CheckEqualData(data string, validData []string) (string, error) {
	inputData := strings.ToLower(data)

	isValidData := false
	for _, data := range validData {
		if inputData == strings.ToLower(data) {
			isValidData = true
			break
		}
	}

	if !isValidData {
		return "", errors.New("data yang diinput tidak valid")
	}

	return inputData, nil
}
