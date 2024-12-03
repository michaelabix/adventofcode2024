package day3

import (
	"reflect"
	"strconv"
	"strings"
	"testing"

	"github.com/michaelabix/adventofcode2024/utils"
)

// answer
func TestAnswer(t *testing.T) {
	values := []string{"2,3", "10,10", "5,5"}
	sum := answer(&values)
	if sum != 131 {
		t.Fatalf("Answer: returned " + strconv.Itoa(sum) + " but expected 131")
	}
}

// test one bad value - that invalid value should be ignored
func TestAnswerInvalid(t *testing.T) {
	values := []string{"a,3", "10,10", "5,5"}
	sum := answer(&values)
	if sum != 125 {
		t.Fatalf("Answer: returned " + strconv.Itoa(sum) + " but expected 125")
	}
}

func TestAnswerZero(t *testing.T) {
	values := []string{"a,b", "asdf,4", "asdf,5"}
	sum := answer(&values)
	if sum != 0 {
		t.Fatalf("Answer: returned " + strconv.Itoa(sum) + " but expected 0")
	}
}

// multiply
func TestMultiply(t *testing.T) {
	values := "5,4"
	product, err := multiply(values)
	if product != 20 {
		t.Fatalf("Multiply: answer received was " + strconv.Itoa(product) + " expected 20")
	}
	if err != nil {
		t.Fatalf("Multiply: returned error for valid value")
	}
}

func TestMultiplyNondigit(t *testing.T) {
	values := "a,b"
	_, err := multiply(values)
	if err == nil {
		t.Fatalf("Multiply: did not return error for invalid value")
	}
}

func TestMultiplyNondigitLeft(t *testing.T) {
	values := "a,2"
	_, err := multiply(values)
	if err == nil {
		t.Fatalf("Multiply: did not return error for invalid value")
	}
}

func TestMultiplyNondigitRight(t *testing.T) {
	values := "1,b"
	_, err := multiply(values)
	if err == nil {
		t.Fatalf("Multiply: did not return error for invalid value")
	}
}

func TestMultiplySpace(t *testing.T) {
	values := " ,2"
	_, err := multiply(values)
	if err == nil {
		t.Fatalf("Multiply: did not return error for invalid value")
	}
}

// parseMul
func TestParseMul(t *testing.T) {
	// because I'm lazy and don't want to write a byte array by hand right now
	inputFile := "sample.txt"
	data, _ := utils.ReadFile(&inputFile)

	pairs := parseMul(string(data))
	if !reflect.DeepEqual(pairs, []string{"2,4", "5,5", "11,8", "8,5"}) {
		t.Fatalf("parseMul: Received unexpected results. Got " + strings.Join(pairs, " "))
	}
}

// parseDo
func TestParseDo(t *testing.T) {
	inputFile := "sample.txt"
	data, _ := utils.ReadFile(&inputFile)

	pairs := parseDo(string(data))
	if !reflect.DeepEqual(pairs, []string{"2,4", "8,5"}) {
		t.Fatalf("parseMul: Received unexpected results. Got " + strings.Join(pairs, " "))
	}
}
