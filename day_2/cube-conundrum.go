package day_2

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/michelm117/advent-of-code/utils"
)

func Solve(filePath string) (int, int) {
	lines := utils.ReadInputFileAsArray(filePath)
	result := make(map[int]bool)
	minCubeResults := make(map[int]map[string]int)
	for _, line := range lines {
		gameNbr, rest := extractGameNumber(line)
		rounds := parseLineToRound(rest)
		maxColorNbrMap := getMaxNbrOfCubes(rounds)
		minCubeResults[gameNbr] = maxColorNbrMap
		result[gameNbr] = true
		for _, round := range rounds {
			for color, nbr := range round {
				if !isPossibleToSolve(color, nbr) {
					result[gameNbr] = false
				}
			}
		}
	}

	return sumOfPossibleGames(result), calculateSumOfPowerResults(minCubeResults)
}

func calculatePowerOfResults(minCubeResults map[int]map[string]int) map[int]int {
	powerOfResults := make(map[int]int)
	for gameNbr, maxColorNbrMap := range minCubeResults {
		power := 1
		for _, nbr := range maxColorNbrMap {
			power *= nbr
		}
		powerOfResults[gameNbr] = power
	}
	return powerOfResults
}

func calculateSumOfPowerResults(minCubeResults map[int]map[string]int) int {
	sum := 0
	for _, power := range calculatePowerOfResults(minCubeResults) {
		sum += power
	}
	return sum
}

func getMaxNbrOfCubes(rounds []map[string]int) map[string]int {
	maxColorNbr := map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}

	for _, round := range rounds {
		for color, nbr := range round {
			if nbr > maxColorNbr[color] {
				maxColorNbr[color] = nbr
			}
		}
	}
	return maxColorNbr
}

func sumOfPossibleGames(games map[int]bool) int {
	sum := 0
	for gameNbr, isPossible := range games {
		if isPossible {
			sum += gameNbr
		}
	}
	return sum
}

func isPossibleToSolve(color string, nbr int) bool {
	switch color {
	case "red":
		return nbr <= 12
	case "green":
		return nbr <= 13
	case "blue":
		return nbr <= 14
	default:
		panic("Unknown color: " + color)
	}
}

func parseLineToRound(line string) []map[string]int {
	groups := strings.Split(line, "; ")
	rounds := make([]map[string]int, 0)
	for _, group := range groups {
		colors := strings.Split(group, ", ")
		colorMap := make(map[string]int)
		for _, color := range colors {
			colorParts := strings.Split(color, " ")
			cubeColor := colorParts[1]
			cubeNbr := colorParts[0]
			nbr, err := strconv.Atoi(cubeNbr)
			if err != nil {
				panic(err)
			}
			colorMap[cubeColor] = nbr
		}
		rounds = append(rounds, colorMap)
	}

	return rounds
}

func extractGameNumber(line string) (int, string) {
	re := regexp.MustCompile(`Game (\d+): `)
	matches := re.FindStringSubmatch(line)
	if matches == nil {
		panic("Could not determine game number: " + line)
	}
	gameNbr, err := strconv.Atoi(matches[1])
	if err != nil {
		panic("Could not convert game number to int: " + err.Error())
	}
	return gameNbr, re.ReplaceAllString(line, "")
}
