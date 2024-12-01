package day1

import (
	"slices"
	"strconv"
	"testing"
)

func TestPart1(t *testing.T) {
	list1 := []int{1, 2, 3, 3, 3, 4}
	list2 := []int{3, 3, 3, 4, 5, 9}
	answer, err := part1(&list1, &list2)
	if err != nil {
		t.Fatalf("part1: returned error " + err.Error())
	}
	if answer != 11 {
		t.Fatalf("part1: answer returned was " + strconv.Itoa(answer) + ", but expected 11")
	}
}

func TestPart1BadListFormat(t *testing.T) {
	list1 := []int{3, 4, 2, 1, 3}
	list2 := []int{4, 3, 5, 3, 9, 3}
	_, err := part1(&list1, &list2)
	if err == nil {
		t.Fatalf("part1: failed to handle two different array lengths")
	}
}

func TestPart2(t *testing.T) {
	list1 := []int{1, 2, 3, 3, 3, 4}
	list2 := []int{3, 3, 3, 4, 5, 9}
	ans := part2(&list1, &list2)
	if ans != 31 {
		t.Fatalf("part2: failed to calculate correct answer. Expected 31, got " + strconv.Itoa(ans))
	}
}

func TestParse(t *testing.T) {
	inputPath := "sample.txt"
	list1 := []int{3, 4, 2, 1, 3, 3}
	list2 := []int{4, 3, 5, 3, 9, 3}

	listOne, listTwo, err := parse(&inputPath)
	if err != nil {
		t.Fatalf("parse: error occurred")
	}
	if !slices.Equal(list1, listOne) || !slices.Equal(list2, listTwo) {
		t.Fatalf("parse: unexpected data in lists")
	}
}

func TestParseNonexistentPath(t *testing.T) {
	inputPath := "day1/doesn'texist.txt"
	_, _, err := parse(&inputPath)
	if err == nil {
		t.Fatalf("parse: somehow read nonexistent path")
	}
}
