package day_4

import (
	"errors"
	"math"
	"regexp"

	"github.com/michelm117/advent-of-code/utils"
)

func GetScratchpadPoints(filePath string) int {
	lines := utils.ReadInputFileAsArray(filePath)
	return calculateTotalPoints(lines)
}

func getGameSets(line string) (map[string]bool, map[string]bool, error) {
	var winningSet map[string]bool
	var scratchSet map[string]bool

	groups := regexp.MustCompile(`Card\s+\d+: ([\d\s]+) \| ([\d\s]+)`)
	digits := regexp.MustCompile(`(\d+)`)

	matches := groups.FindStringSubmatch(line)
	if len(matches) == 3 {
		winning := matches[1]
		scratchpad := matches[2]

		winningSet = utils.SetFromArray(digits.FindAllString(winning, -1))
		scratchSet = utils.SetFromArray(digits.FindAllString(scratchpad, -1))
	}
	if len(winningSet) == 0 || len(scratchSet) == 0 {
		return nil, nil, errors.New("Could not parse line: " + line)
	}
	return winningSet, scratchSet, nil
}

func getWinningNumbersFromScratchpad(winningSet, scratchSet map[string]bool) []string {
	var winningNumbers []string
	for scratchNumber := range scratchSet {
		if _, ok := winningSet[scratchNumber]; ok {
			winningNumbers = append(winningNumbers, scratchNumber)
		}
	}
	return winningNumbers
}

func calculateGamePoints(winningNumbers []string) int {
	return int(math.Pow(2, float64(len(winningNumbers)-1)))
}

func calculateTotalPoints(lines []string) int {
	points := 0
	for _, line := range lines {
		winningSet, scratchSet, err := getGameSets(line)
		if err != nil {
			panic(err)
		}

		winningNumbers := getWinningNumbersFromScratchpad(winningSet, scratchSet)

		gamePoints := calculateGamePoints(winningNumbers)
		points += gamePoints
	}
	return points
}
