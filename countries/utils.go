package countries

import (
	"strconv"
)

func getIntFromStringAt(str string, index int) (int, error) {

	char := getStringAt(str, index)
	result, err := strconv.Atoi(char)

	return result, err
}

func getStringAt(str string, index int) string {

	chars := []rune(str)

	return string(chars[index])
}
