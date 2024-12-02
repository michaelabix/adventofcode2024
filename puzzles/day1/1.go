package day1

import (
	"errors"
	"log/slog"
	"sort"
	"strconv"
	"strings"

	"github.com/michaelabix/adventofcode2024/utils"
)

func Solve() {
	// read input and parse
	inputPath := "puzzles/day1/sample.txt"
	list1, list2, err := parse(&inputPath)
	if err != nil {
		slog.Error(err.Error())
		return
	}

	// sort here so it's clearer what's happening
	sort.Ints(list1)
	sort.Ints(list2)

	// part 1
	ans, part1Err := part1(&list1, &list2)
	if part1Err != nil {
		slog.Error(part1Err.Error())
	} else {
		slog.Info("The answer to part1 is: " + strconv.Itoa(ans))
	}

	// part 2
	ans2 := part2(&list1, &list2)
	slog.Info("The answer to part2 is: " + strconv.Itoa(ans2))
}

// expects sorted array
func part1(list1 *[]int, list2 *[]int) (int, error) {
	if len(*list1) != len(*list2) {
		err := errors.New("part1: list1 and list2 are not the same length")
		return 0, err
	}
	answer := 0
	for i := 0; i < len(*list1); i++ {
		// using an if avoids importing math and using floats to find absolute value
		if (*list1)[i] > (*list2)[i] {
			answer += (*list1)[i] - (*list2)[i]
		} else {
			answer += (*list2)[i] - (*list1)[i]
		}
	}
	return answer, nil
}

// expects sorted array
func part2(list1 *[]int, list2 *[]int) int {
	// first and last are used to describe the first and last occurrence of a number
	answer, lastNum, lastAns, first, last := 0, 0, 0, 0, 0
	end := len(*list2) - 1

	// loop through list1
	for i := 0; i < len(*list1); i++ {
		// if the last number and the current number are not the same
		if lastNum != (*list1)[i] {
			lastNum = (*list1)[i]
			// check that we're not exceeding the length of list2
			if last == end {
				break
			}
			// find in list2
			first = utils.FindFirstOccurrence(list2, last, end, lastNum)
			if first != -1 {
				// set last to most recent occurrence to make finding the last occurrence quicker
				last = first

				// set top to actual top
				last = utils.FindLastOccurrence(list2, last, end, lastNum)

				// set lastAnswer to current answer in case the number in list1 repeats
				lastAns = (last - (first - 1)) * lastNum

				// add result to answer
				answer += lastAns
			}
			// else if the last number and the current number are the same
		} else {
			answer += lastAns
		}
	}
	return answer
}

func parse(puzzleInputPath *string) ([]int, []int, error) {
	// read file
	data, err := utils.ReadFile(puzzleInputPath)
	if err != nil {
		slog.Error(err.Error())
		return nil, nil, err
	}
	var evens []int
	var odds []int
	// split at whitespace
	fields := strings.Fields(string(data))

	// loop through result
	for i := 0; i < len(fields); i++ {
		// convert to integer
		value, err := strconv.Atoi(fields[i])
		if err != nil {
			slog.Warn("value of " + fields[i] + " cannot be converted to integer")
		}

		// sort into even and odd indexes
		if i%2 == 0 {
			evens = append(evens, value)
		} else {
			odds = append(odds, value)
		}
	}

	return evens, odds, nil
}
