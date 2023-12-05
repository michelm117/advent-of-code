package day_4

import (
	"errors"
	"math"
	"regexp"
	"strconv"

	"github.com/michelm117/advent-of-code/utils"
)

func GetScratchpadPoints(filePath string) (int, int) {
	lines := utils.ReadInputFileAsArray(filePath)
	partOneSolution := calculateTotalPoints(lines)

	scratchCards := calculateWiningScratchCards(lines)

	scratchCardsSum := 0
	for _, value := range scratchCards {
		scratchCardsSum += value
	}
	return partOneSolution, scratchCardsSum
}

func calculateWiningScratchCards(lines []string) map[string]int {
	scratchCards := make(map[string]int)
	for i := 1; i <= len(lines); i++ {
		scratchCards[strconv.Itoa(i)] = 1
	}

	for _, line := range lines {
		gameNbr, winningSet, scratchSet, err := getGameSets(line)
		if err != nil {
			panic(err)
		}

		winningNumbers := getWinningNumbersFromScratchpad(winningSet, scratchSet)
		scratchCards = appendScratchCards(scratchCards, gameNbr, len(winningNumbers))
	}
	return scratchCards
}

func calculateTotalPoints(lines []string) int {
	points := 0
	for _, line := range lines {
		_, winningSet, scratchSet, err := getGameSets(line)
		if err != nil {
			panic(err)
		}

		winningNumbers := getWinningNumbersFromScratchpad(winningSet, scratchSet)

		gamePoints := calculateGamePoints(winningNumbers)
		points += gamePoints
	}
	return points
}

func getGameSets(line string) (string, map[string]bool, map[string]bool, error) {
	var winningSet map[string]bool
	var scratchSet map[string]bool

	groups := regexp.MustCompile(`Card\s+(\d)+: ([\d\s]+) \| ([\d\s]+)`)
	digits := regexp.MustCompile(`(\d+)`)

	matches := groups.FindStringSubmatch(line)

	if len(matches) != 4 {
		return "", nil, nil, errors.New("Could not parse line: " + line)
	}

	gameNbr := matches[1]
	winning := digits.FindAllString(matches[2], -1)
	scratchpad := digits.FindAllString(matches[3], -1)

	winningSet = utils.SetFromArray(winning)
	scratchSet = utils.SetFromArray(scratchpad)
	return gameNbr, winningSet, scratchSet, nil
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

func appendScratchCards(scratchCards map[string]int, gameNbrStr string, matchingNbrs int) map[string]int {
	if matchingNbrs == 0 {
		return scratchCards
	}

	gameNbr, err := strconv.Atoi(gameNbrStr)
	if err != nil {
		panic(err)
	}
	for i := gameNbr + 1; i <= gameNbr+matchingNbrs; i++ {
		scratchCards[strconv.Itoa(i)] += scratchCards[gameNbrStr]
	}
	return scratchCards
}
