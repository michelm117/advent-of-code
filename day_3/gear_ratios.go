package day_3

import (
	"regexp"
	"strconv"

	"github.com/michelm117/advent-of-code/utils"
)

func Solve(filePath string) int {
	lines := utils.ReadInputFileAsArray(filePath)
	engine := NewEngine(lines)
	sumOfParts := sumOfParts(engine, lines)
	return sumOfParts
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

// func gearRatios(engine *EngineSchematic, lines []string) []string {
// 	// gearRatios := []string{}
// 	// symbols := engine.CollectSymbols()
// 	// for _, symbol := range symbols {
// 	// 	gearRatios = append(gearRatios, engine.FindGearRatio(symbol.coord)...)
// 	// }
// 	// return gearRatios
// }
