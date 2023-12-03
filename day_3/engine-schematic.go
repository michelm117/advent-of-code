package day_3

import (
	"fmt"
	"regexp"
	"strings"
)

type Coord struct {
	x, y int
}

type NumberCoord struct {
	number string
	coords []Coord
}

type EngineSchematic struct {
	visual [][]string
}

func NewEngine(lines []string) *EngineSchematic {
	engine := EngineSchematic{}
	engine.visual = engine.GenerateVisual(lines)
	return &engine
}

func (e *EngineSchematic) Print() {
	fmt.Printf("+%s+\n", strings.Repeat("-", len(e.visual[0])))
	for _, line := range e.visual {
		fmt.Printf("|%s|\n", strings.Join(line, ""))
	}
	fmt.Printf("+%s+\n", strings.Repeat("-", len(e.visual[0])))
}

func (e *EngineSchematic) GenerateVisual(lines []string) [][]string {
	var mapOfLines [][]string
	for _, line := range lines {
		mapOfLines = append(mapOfLines, splitStringToCharArray(line))
	}
	return mapOfLines
}

func (e *EngineSchematic) getAdjacentSymbols(x, y int) []string {
	symbols := []string{}
	coords := []Coord{
		{-1, 0}, {1, 0}, {0, -1}, {0, 1}, // left, right, up, down
		{-1, -1}, {1, -1}, {-1, 1}, {1, 1}, // diagonals
	}

	for _, coord := range coords {
		newX, newY := x+coord.x, y+coord.y
		if newX >= 0 && newX < len(e.visual[y]) && newY >= 0 && newY < len(e.visual) {
			symbol := string(e.visual[newY][newX])
			// ignore numbers and dots
			if symbol != "." && !regexp.MustCompile(`\d`).MatchString(symbol) {
				symbols = append(symbols, symbol)
			}
		}
	}

	return symbols
}

func (e *EngineSchematic) CollectNumbers() []NumberCoord {
	allNumbers := []NumberCoord{}
	re := regexp.MustCompile(`\d`)
	for y, line := range e.visual {
		tmpNumber := NumberCoord{}
		numberIsAttached := false
		for x, char := range line {
			isNumber := re.MatchString(char)

			if isNumber && x == len(line)-1 {
				fmt.Println("End of line", char)
				numberIsAttached = false
				allNumbers = append(allNumbers, tmpNumber)
				continue
			}

			// start of a new number
			if isNumber && !numberIsAttached {
				numberIsAttached = true
				tmpNumber = NumberCoord{string(char), []Coord{{x, y}}}
				continue
			}

			// continue a number
			if isNumber && numberIsAttached {
				tmpNumber.number += string(char)
				tmpNumber.coords = append(tmpNumber.coords, Coord{x, y})
				continue
			}

			// end of a number
			if !isNumber && numberIsAttached {
				numberIsAttached = false
				allNumbers = append(allNumbers, tmpNumber)
				continue
			}

		}
	}
	return allNumbers
}

func splitStringToCharArray(line string) []string {
	// append each character to a slice
	var splitLine []string
	for _, char := range line {
		splitLine = append(splitLine, string(char))
	}
	return splitLine
}
