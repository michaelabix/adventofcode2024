package main

import "testing"

func TestIsInt(t *testing.T) {
	if !isInt("001") && !isInt("12") && !isInt("25") {
		t.Fatalf("isInt: returned false for valid value")
	}
}

func TestIsIntOutOfRange(t *testing.T) {
	if isInt("0") || isInt("26") || isInt("10000000000000000000") {
		t.Fatalf("isInt: returned true for values out of range")
	}
}
