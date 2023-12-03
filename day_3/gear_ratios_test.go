package day_3

import "testing"

func TestSumOfParts(t *testing.T) {
	lines := []string{
		"12.......*..",
		"+.........34",
		".......-12..",
		"..78........",
		"..*....60...",
		"78.........9",
		".5.....23..$",
		"8...90*12...",
		"............",
		"2.2......12.",
		".*.........*",
		"1.1..503+.56",
	}
	sum := sumOfParts(lines)

	expectedSum := 23 + 34 + 12 + 78 + 78 + 9 + 23 + 90 + 12 + 2 + 2 + 12 + 1 + 1 + 503 + 56
	if sum != expectedSum {
		t.Errorf("Expected %d, got %d", expectedSum, sum)
	}
}
