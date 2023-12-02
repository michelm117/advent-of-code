package day_1

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func Calibrate(filePath string) int {
	lines := readInputFileAsArray(filePath)
	return getCalibrationValueSum(lines)
}

func getCalibrationValueSum(lines []string) int {
	sum := 0
	for _, line := range lines {
		val, err := getCalibrationValue(line)
		if err != nil {
			fmt.Println("Error getting calibration value: " + err.Error())
			fmt.Println("Skipping line: " + line)
			continue
		}
		sum += val
	}
	return sum
}

func getCalibrationValue(line string) (int, error) {
	leftPos := -1
	rightPos := -1
	for i := 0; i < len(line); i++ {
		char := rune(line[i])
		if unicode.IsDigit(char) {
			leftPos = i
			break
		}
	}

	for i := len(line) - 1; i >= 0; i-- {
		char := rune(line[i])
		if unicode.IsDigit(char) {
			rightPos = i
			break
		}
	}

	if leftPos == -1 || rightPos == -1 {
		return 0, errors.New("Line does not contain digit: " + line)
	}

	calibrationValue, err := concatToInt(string(line[leftPos]), string(line[rightPos]))
	if err != nil {
		return 0, err
	}

	return calibrationValue, nil
}

func concatToInt(leftValue string, rightValue string) (int, error) {
	concatenatedValue := leftValue + rightValue
	value, err := strconv.Atoi(concatenatedValue)
	if err != nil {
		return 0, errors.New("Could not convert string to int: " + concatenatedValue)
	}
	return value, nil
}

func readInputFileAsArray(filePath string) []string {
	data, err := os.ReadFile(filePath)
	if err != nil {
		panic("Error reading file: " + err.Error())
	}
	return strings.Split(string(data), "\n")
}
