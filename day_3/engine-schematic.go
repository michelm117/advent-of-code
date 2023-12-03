package day_3

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Coord struct {
	x, y int
}

type Symbol struct {
	value string
	coord Coord
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

func (e *EngineSchematic) GetSymbolAt(x, y int) string {
	return e.visual[y][x]
}

func (e *EngineSchematic) getAdjacentMatches(x, y int, re *regexp.Regexp) []Symbol {
	symbols := []Symbol{}
	coords := []Coord{
		{-1, 0}, {1, 0}, {0, -1}, {0, 1}, // left, right, up, down
		{-1, -1}, {1, -1}, {-1, 1}, {1, 1}, // diagonals
	}

	for _, coord := range coords {
		newX, newY := x+coord.x, y+coord.y
		if newX >= 0 && newX < len(e.visual[y]) && newY >= 0 && newY < len(e.visual) {
			symbol := e.GetSymbolAt(newX, newY)
			if re.MatchString(symbol) {
				symbols = append(symbols, Symbol{symbol, Coord{newX, newY}})
			}
		}
	}
	return symbols
}

func (e *EngineSchematic) CollectSymbols(re *regexp.Regexp) []Symbol {
	allSymbols := []Symbol{}
	for y, line := range e.visual {
		for x, char := range line {
			if re.MatchString(char) {
				allSymbols = append(allSymbols, Symbol{char, Coord{x, y}})
			}
		}
	}
	return allSymbols
}

func (e *EngineSchematic) FindNumbers(start Coord) []string {
	isNumber := regexp.MustCompile(`[0-9]`)
	tmpNumbers := e.getEmptyVisual()
	for _, symbol := range e.getAdjacentMatches(start.x, start.y, isNumber) {
		tmpNumbers =
			e.findNumbersRecursive(isNumber, symbol.coord.x, symbol.coord.y, tmpNumbers)
	}
	numbers := []string{}
	for _, line := range tmpNumbers {
		numberStrings := getNumberStrings(line)
		numbers = append(numbers, numberStrings...)
	}
	return numbers
}

func getNumberStrings(arr []string) []string {
	result := []string{}
	tmpNumber := ""
	appending := false

	for _, str := range arr {
		if str != "" {
			num, err := strconv.Atoi(str)
			if err == nil && num >= 0 && num <= 9 {
				if !appending {
					appending = true
					tmpNumber = strconv.Itoa(num)
				} else {
					tmpNumber += strconv.Itoa(num)
				}
			} else {
				if appending {
					appending = false
					result = append(result, tmpNumber)
					tmpNumber = ""
				}
			}
		} else {
			if appending {
				appending = false
				result = append(result, tmpNumber)
				tmpNumber = ""
			}
		}
	}

	if appending {
		result = append(result, tmpNumber)
	}

	return result
}

func (e *EngineSchematic) findNumbersRecursive(isNumber *regexp.Regexp, x, y int, number [][]string) [][]string {
	symbol := e.GetSymbolAt(x, y)
	if !isNumber.MatchString(symbol) {
		return number
	}
	if number[y][x] != "" {
		return number
	}
	number[y][x] = symbol
	tmpOne := make([]string, len(e.visual[y]))
	if x+1 < len(e.visual[y]) {
		tmpOne = e.findNumbersRecursive(isNumber, x+1, y, number)[y]
	}

	tmpTwo := make([]string, len(e.visual[y]))
	if x-1 >= 0 {
		tmpTwo = e.findNumbersRecursive(isNumber, x-1, y, number)[y]
	}

	//merge tmpOne and tmpTwo
	merged := make([]string, len(e.visual[y]))
	for i := range tmpOne {
		if tmpOne[i] == "" && tmpTwo[i] == "" {
			merged[i] = ""
		} else if tmpOne[i] == tmpTwo[i] {
			merged[i] = tmpOne[i]
		} else if tmpOne[i] == "" && tmpTwo[i] != "" {
			merged[i] = tmpTwo[i]
		} else if tmpOne[i] != "" && tmpTwo[i] == "" {
			merged[i] = tmpOne[i]
		} else {
			panic("Different values in same position found")
		}
	}
	number[y] = merged
	return number
}

func (e *EngineSchematic) getEmptyVisual() [][]string {
	emptyVisual := make([][]string, len(e.visual))
	for i := range emptyVisual {
		emptyVisual[i] = make([]string, len(e.visual[i]))
	}
	return emptyVisual
}

func splitStringToCharArray(line string) []string {
	// append each character to a slice
	var splitLine []string
	for _, char := range line {
		splitLine = append(splitLine, string(char))
	}
	return splitLine
}
