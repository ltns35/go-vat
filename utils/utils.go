package utils

import (
	"strconv"
)

func IntAt(str string, index int) int {
	char := StringAt(str, index)
	result, _ := strconv.Atoi(char)

	return result
}

func StringAt(str string, index int) string {
	chars := []rune(str)

	return string(chars[index])
}
