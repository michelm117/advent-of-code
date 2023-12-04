package day_4

import (
	"reflect"
	"testing"
)

func TestGetGameSets(t *testing.T) {
	line := "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53"
	actualWinning, actualScratchpad, err := getGameSets(line)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	expectedWinningSet := map[string]bool{
		"17": true, "41": true, "48": true, "83": true, "86": true,
	}
	if !reflect.DeepEqual(expectedWinningSet, actualWinning) {
		t.Errorf("Expected %v, got %v", expectedWinningSet, actualWinning)
	}

	expectedScratchSet := map[string]bool{
		"6": true, "9": true, "17": true, "31": true, "48": true, "53": true, "83": true, "86": true,
	}
	if !reflect.DeepEqual(expectedScratchSet, actualScratchpad) {
		t.Errorf("Expected %v, got %v", expectedScratchSet, actualScratchpad)
	}

	line = "Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19"
	actualWinning, actualScratchpad, err = getGameSets(line)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	expectedWinningSet = map[string]bool{
		"13": true, "16": true, "20": true, "32": true, "61": true,
	}
	if !reflect.DeepEqual(expectedWinningSet, actualWinning) {
		t.Errorf("Expected %v, got %v", expectedWinningSet, actualWinning)
	}

	expectedScratchSet = map[string]bool{
		"17": true, "19": true, "24": true, "30": true, "32": true, "61": true, "68": true, "82": true,
	}
	if !reflect.DeepEqual(expectedScratchSet, actualScratchpad) {
		t.Errorf("Expected %v, got %v", expectedScratchSet, actualScratchpad)
	}



	line = "Card    2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19"
	actualWinning, actualScratchpad, err = getGameSets(line)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	expectedWinningSet = map[string]bool{
		"13": true, "16": true, "20": true, "32": true, "61": true,
	}
	if !reflect.DeepEqual(expectedWinningSet, actualWinning) {
		t.Errorf("Expected %v, got %v", expectedWinningSet, actualWinning)
	}

	expectedScratchSet = map[string]bool{
		"17": true, "19": true, "24": true, "30": true, "32": true, "61": true, "68": true, "82": true,
	}
	if !reflect.DeepEqual(expectedScratchSet, actualScratchpad) {
		t.Errorf("Expected %v, got %v", expectedScratchSet, actualScratchpad)
	}
}

func TestGetWinningNumbersFromScratchpad(t *testing.T) {
	winningSet := map[string]bool{
		"17": true, "41": true, "48": true, "83": true, "86": true,
	}
	scratchSet := map[string]bool{
		"6": true, "9": true, "17": true, "31": true, "48": true, "53": true, "83": true, "86": true,
	}
	actual := getWinningNumbersFromScratchpad(winningSet, scratchSet)

	expected := []string{"17", "48", "83", "86"}
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestCalculateGamePoints(t *testing.T) {
	winningNumbers := []string{"48", "83", "17", "86"}
	actual := calculateGamePoints(winningNumbers)
	expected := 8
	if expected != actual {
		t.Errorf("Expected %v, got %v", expected, actual)
	}

	winningNumbers = []string{"32", "61"}
	actual = calculateGamePoints(winningNumbers)
	expected = 2
	if expected != actual {
		t.Errorf("Expected %v, got %v", expected, actual)
	}

	winningNumbers = []string{"84"}
	actual = calculateGamePoints(winningNumbers)
	expected = 1
	if expected != actual {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestCalculateTotalPoints(t *testing.T) {
	lines := []string{
		"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
		"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19",
		"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1",
		"Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83",
		"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36",
		"Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11"}

	actual := calculateTotalPoints(lines)
	expected := 13
	if expected != actual {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}
