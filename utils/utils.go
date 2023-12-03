package utils

import (
	"os"
	"strings"
)

func ReadInputFileAsArray(filePath string) []string {
	data, err := os.ReadFile(filePath)
	if err != nil {
		panic("Error reading file: " + err.Error())
	}
	return strings.Split(string(data), "\n")
}

func ReverseString(s string) string {
	var reversed string
	for _, c := range s {
		reversed = string(c) + reversed
	}
	return reversed
}
