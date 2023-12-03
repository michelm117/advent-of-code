package day_3

import (
	"reflect"
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


