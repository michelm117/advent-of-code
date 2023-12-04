package utils_test

import (
	"reflect"
	"testing"

	"github.com/michelm117/advent-of-code/utils"
)

func TestReverseString(t *testing.T) {
	actual := utils.ReverseString("12345")
	expected := "54321"
	if expected != actual {
		t.Errorf("Expected %v, got %v", expected, actual)
	}

	actual = utils.ReverseString("")
	expected = ""
	if expected != actual {
		t.Errorf("Expected %v, got %v", expected, actual)
	}

	actual = utils.ReverseString("-1")
	expected = "1-"
	if expected != actual {
		t.Errorf("Expected %v, got %v", expected, actual)
	}

	actual = utils.ReverseString("-")
	expected = "-"
	if expected != actual {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestSetFromArray(t *testing.T) {
	digits := []string{"1", "2", "3", "4", "5"}
	actual := utils.SetFromArray(digits)

	expected := map[string]bool{
		"1": true, "2": true, "3": true, "4": true, "5": true,
	}
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}
