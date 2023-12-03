package day_3

import (
	"regexp"
	"strconv"

	"github.com/michelm117/advent-of-code/utils"
)

func Solve(filePath string) (int, int) {
	lines := utils.ReadInputFileAsArray(filePath)
	engine := NewEngine(lines)
	sumOfParts := sumOfParts(engine, lines)

	gearRatio := gearRatios(engine, lines)
	return sumOfParts, gearRatio
}

func sumOfParts(engine *EngineSchematic, lines []string) int {
	numbers := []string{}
	symbolsOnly := regexp.MustCompile(`[^0-9.]`)
	symbols := engine.CollectSymbols(symbolsOnly)
	for _, symbol := range symbols {
		numbers = append(numbers, engine.FindNumbers(symbol.coord)...)
	}

	sum := 0
	for _, number := range numbers {
		if number == "" {
			continue
		}
		value, _ := strconv.Atoi(number)

		sum += value
	}
	return sum
}

func gearRatios(engine *EngineSchematic, lines []string) int {
	symbolsOnly := regexp.MustCompile(`\*`)
	symbols := engine.CollectSymbols(symbolsOnly)
	sum := 0
	for _, symbol := range symbols {
		numbers := engine.FindNumbers(symbol.coord)
		if len(numbers) != 2 {
			continue
		}

		num1, _ := strconv.Atoi(numbers[0])
		num2, _ := strconv.Atoi(numbers[1])
		sum += num1 * num2

	}

	return sum
}
