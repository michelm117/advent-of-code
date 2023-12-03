package day_3

import (
	"fmt"
	"strconv"

	"github.com/michelm117/advent-of-code/utils"
)

func Solve(filePath string) int {
	lines := utils.ReadInputFileAsArray(filePath)
	return sumOfParts(lines)

}

func sumOfParts(lines []string) int {
	engine := NewEngine(lines)
	numberCoords := engine.CollectNumbers()
	sum := 0
	for _, numberCoord := range numberCoords {
		symbols := []string{}
		for _, coord := range numberCoord.coords {
			symbols = append(symbols, engine.getAdjacentSymbols(coord.x, coord.y)...)
		}

		if len(symbols) > 0 {
			nbr, err := strconv.Atoi(numberCoord.number)
			if err != nil {
				panic(err)
			}
			sum += nbr
		} else {
			fmt.Println("No adjacent symbols for", numberCoord.number)
		}
	}

	return sum
}
