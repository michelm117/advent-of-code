package day_3

import (
	"reflect"
	"regexp"
	"testing"
)

func TestSplitStringToCharArray(t *testing.T) {
	line := "467..*114.."
	splitLine := splitStringToCharArray(line)
	if len(splitLine) != len(line) {
		t.Errorf("Expected %d, got %d", len(line), len(splitLine))
	}

	expected := []string{"4", "6", "7", ".", ".", "*", "1", "1", "4", ".", "."}
	if !reflect.DeepEqual(splitLine, expected) {
		t.Errorf("Expected %v, got %v", expected, splitLine)
	}
}

func TestGenerateVisual(t *testing.T) {
	lines := []string{
		"467..*114..",
		"......755.",
	}
	engine := EngineSchematic{}
	visualMap := engine.GenerateVisual(lines)

	if len(visualMap) != len(lines) {
		t.Errorf("Expected %d, got %d", len(lines), len(visualMap))
	}

	for i, line := range visualMap {
		if len(line) != len(lines[i]) {
			t.Errorf("Expected %d, got %d", len(lines[0]), len(line))
		}
	}
}

func TestCollectSymbols(t *testing.T) {
	lines := []string{
		"467..*114..",
		"......755.",
		"$.....755+",
	}
	engine := NewEngine(lines)
	symbolsOnly := regexp.MustCompile(`[^0-9.]`)
	symbols := engine.CollectSymbols(symbolsOnly)

	expectedLength := 3
	if len(symbols) != expectedLength {
		t.Errorf("Expected %d, got %d", expectedLength, len(symbols))
	}

	expected := []Symbol{
		{"*", Coord{5, 0}},
		{"$", Coord{0, 2}},
		{"+", Coord{9, 2}},
	}

	if !reflect.DeepEqual(symbols, expected) {
		t.Errorf("Expected %v, got %v", expected, symbols)
	}
}

func TestGetAdjacentMatches(t *testing.T) {
	lines := []string{
		"467..*114..",
		"......755.",
		"$.....755+",
	}
	engine := NewEngine(lines)

	isNumber := regexp.MustCompile(`[0-9]`)

	symbols := engine.getAdjacentMatches(5, 0, isNumber)
	expectedLength := 2
	if len(symbols) != expectedLength {
		t.Errorf("Expected %d, got %d", expectedLength, len(symbols))
	}

	expected := []Symbol{
		{"1", Coord{6, 0}},
		{"7", Coord{6, 1}},
	}

	if !reflect.DeepEqual(symbols, expected) {
		t.Errorf("Expected %v, got %v", expected, symbols)
	}

	// Test that it doesn't go out of bounds in a corner
	symbols = engine.getAdjacentMatches(9, 2, isNumber)
	expectedLength = 2
	if len(symbols) != expectedLength {
		t.Errorf("Expected %d, got %d", expectedLength, len(symbols))
	}

	expected = []Symbol{
		{"5", Coord{8, 2}},
		{"5", Coord{8, 1}},
	}

	if !reflect.DeepEqual(symbols, expected) {
		t.Errorf("Expected %v, got %v", expected, symbols)
	}
}

func TestGetSymbolAt(t *testing.T) {
	lines := []string{
		"467..*114..",
		"......755.",
		"$.....755+",
	}
	engine := NewEngine(lines)
	symbol := engine.GetSymbolAt(5, 0)
	expected := "*"
	if symbol != expected {
		t.Errorf("Expected %s, got %s", expected, symbol)
	}
}

func TestFindNumbers(t *testing.T) {
	lines := []string{
		"467..*114..",
		"......755.",
		"$.....755+",
	}
	engine := NewEngine(lines)
	numbers := engine.FindNumbers(Coord{5, 0})

	expectedLength := 3
	if len(numbers) != expectedLength {
		t.Errorf("Expected %d, got %d", expectedLength, len(numbers))
	}

	expected := []Symbol{
		{"4", Coord{0, 0}},
		{"7", Coord{1, 0}},
		{"5", Coord{6, 1}},
	}

	if !reflect.DeepEqual(numbers, expected) {
		t.Errorf("Expected %v, got %v", expected, numbers)
	}
}
