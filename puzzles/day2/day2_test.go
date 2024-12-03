package day2

import (
	"reflect"
	"testing"

	"github.com/michaelabix/adventofcode2024/utils"
)

// parse
func TestParse(t *testing.T) {
	expected := [][]int{{7, 6, 4, 2, 1},
		{1, 2, 7, 8, 9},
		{9, 7, 6, 2, 1},
		{1, 3, 2, 4, 5},
		{8, 6, 4, 4, 1},
		{1, 3, 6, 7, 9}}

	// because I'm lazy and don't want to write a byte array by hand right now
	inputFile := "sample.txt"
	data, _ := utils.ReadFile(&inputFile)

	integers := parse(data)
	if !reflect.DeepEqual(expected, integers) {
		t.Fatalf("parse: parsed data does not match expected output")
	}

}

// probably test if the data contents can't be converted to int
func TestParseBadType(t *testing.T) {

}

func TestCheckIncreasing(t *testing.T) {

}
