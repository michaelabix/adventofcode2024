package utils

import (
	"strconv"
	"testing"
)

// FindFirstOccurrence
func TestFindFirstOccurrence(t *testing.T) {
	list := []int{0, 0, 0, 1, 1, 4, 4, 4, 4, 7, 8, 2345}
	index := FindFirstOccurrence(&list, 0, 12, 0)
	if index != 0 {
		t.Fatalf("FindFirstOccurrence: index returned was " + strconv.Itoa(index) + ", expected 0")
	}
	index = FindFirstOccurrence(&list, 0, 12, 7)
	if index != 9 {
		t.Fatalf("FindFirstOccurrence: index returned was " + strconv.Itoa(index) + ", expected 9")
	}
	index = FindFirstOccurrence(&list, 0, 12, 2345)
	if index != 11 {
		t.Fatalf("FindFirstOccurrence: index returned was " + strconv.Itoa(index) + ", expected 11")
	}
}

func TestFindFirstOccurrenceNonexistent(t *testing.T) {
	list := []int{0, 0, 0, 1, 1, 4, 4, 4, 4, 7, 8, 2345}
	index := FindFirstOccurrence(&list, 0, 12, 5)
	if index != -1 {
		t.Fatalf("FindFirstOccurrence: index returned was " + strconv.Itoa(index) + ", expected -1")
	}
}

// FindLastOccurrence
func TestFindLastOccurrence(t *testing.T) {
	list := []int{0, 0, 0, 1, 1, 4, 4, 4, 4, 7, 8, 2345}
	index := FindLastOccurrence(&list, 0, 12, 4)
	if index != 8 {
		t.Fatalf("FindFirstOccurrence: index returned was " + strconv.Itoa(index) + ", expected 8")
	}
	index = FindLastOccurrence(&list, 0, 12, 0)
	if index != 2 {
		t.Fatalf("FindFirstOccurrence: index returned was " + strconv.Itoa(index) + ", expected 2")
	}
	index = FindLastOccurrence(&list, 0, 12, 2345)
	if index != 11 {
		t.Fatalf("FindFirstOccurrence: index returned was " + strconv.Itoa(index) + ", expected 11")
	}
}

func TestFindLastOccurrenceNonexistent(t *testing.T) {
	list := []int{0, 0, 0, 1, 1, 4, 4, 4, 4, 7, 8, 2345}
	index := FindLastOccurrence(&list, 0, 12, 5)
	if index != -1 {
		t.Fatalf("FindFirstOccurrence: index returned was " + strconv.Itoa(index) + ", expected -1")
	}
}

// ReadFile
func TestReadFile(t *testing.T) {
	file := "sample.txt"
	data, err := ReadFile(&file)
	if err != nil {
		t.Fatalf("ReadFile: something went wrong reading valid file")
	}
	if string(data) != "this is a test file" {
		t.Fatalf("Readfile: file contents do not match expected string")
	}
}

func TestReadFileNonexistent(t *testing.T) {
	file := "doesn'texist.txt"
	_, err := ReadFile(&file)
	if err == nil {
		t.Fatalf("ReadFile: did not throw error for nonexistent file")
	}
}
