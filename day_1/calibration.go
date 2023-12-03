package day_1

import (
	"errors"
	"fmt"
	"strconv"

	helper "github.com/michelm117/advent-of-code/utils"
)

func Calibrate(filePath string) int {
	t := initTree()
	mapping := map[string]string{
		"ONE":   "1",
		"TWO":   "2",
		"THREE": "3",
		"FOUR":  "4",
		"FIVE":  "5",
		"SIX":   "6",
		"SEVEN": "7",
		"EIGHT": "8",
		"NINE":  "9",
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",

		"ENO":   "1",
		"OWT":   "2",
		"EERHT": "3",
		"RUOF":  "4",
		"EVIF":  "5",
		"XIS":   "6",
		"NEVES": "7",
		"THGIE": "8",
		"ENIN":  "9",
		"eno":   "1",
		"owt":   "2",
		"eerht": "3",
		"ruof":  "4",
		"evif":  "5",
		"xis":   "6",
		"neves": "7",
		"thgie": "8",
		"enin":  "9",

		"1": "1",
		"2": "2",
		"3": "3",
		"4": "4",
		"5": "5",
		"6": "6",
		"7": "7",
		"8": "8",
		"9": "9",
	}
	lines := helper.ReadInputFileAsArray(filePath)
	return getCalibrationValueSum(t, mapping, lines)
}

func initTree() *Tree {
	t := Tree{}
	patterns := []string{"ONE", "TWO", "THREE", "FOUR", "FIVE", "SIX", "SEVEN", "EIGHT", "NINE"}
	for _, pattern := range patterns {
		t.AddPattern(pattern)
	}
	// add 1 to 9 to the tree
	for i := 1; i < 10; i++ {
		t.AddPattern(strconv.FormatInt(int64(i), 10))
	}
	return &t
}

func getCalibrationValueSum(t *Tree, mapping map[string]string, lines []string) int {
	rt := t.Reverse()

	sum := 0
	for _, line := range lines {
		l, r, err := getCalibrationValue(line, t, rt, mapping)
		if err != nil {
			fmt.Println(err, "\nSkipping line")
			continue
		}

		sum += concatAndTransformToNumber(l, r)
	}
	return sum
}

func getCalibrationValue(line string, t *Tree, reversedTree *Tree, mapping map[string]string) (string, string, error) {
	// find left match
	start := 0
	end := 1
	leftMatch := ""
	for start < len(line) {
		pattern := line[start:end]
		if !t.IsPrefix(pattern) {
			start++
			end = start + 1
			continue
		}
		if mapping[pattern] != "" {
			leftMatch = mapping[pattern]
			break
		}
		end++
	}

	if leftMatch == "" {
		return "", "", errors.New("No left match found in " + line)
	}

	// find right match
	line = helper.ReverseString(line)
	start = 0
	end = 1
	rightMatch := ""
	for start < len(line) {
		pattern := line[start:end]
		if !reversedTree.IsPrefix(pattern) {
			start++
			end = start + 1
			continue
		}
		if mapping[pattern] != "" {
			rightMatch = mapping[pattern]
			break
		}
		end++
	}

	if rightMatch == "" {
		return "", "", errors.New("No right match found in " + line)
	}

	return leftMatch, rightMatch, nil
}

func concatAndTransformToNumber(leftValue string, rightValue string) int {
	concatenatedValue := leftValue + rightValue

	value, err := strconv.Atoi(concatenatedValue)
	if err != nil {
		panic("Could not convert string to int: " + err.Error())
	}
	return value
}
