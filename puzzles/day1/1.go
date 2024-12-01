package puzzles

import (
	"errors"
	"log/slog"
	"sort"
	"strconv"
	"strings"

	"github.com/michaelabix/adventofcode2024/utils"
)

func SolveDay1() {
	inputPath := "day1/1.txt"
	list1, list2, err := day1Parse(&inputPath)
	if err != nil {
		slog.Error(err.Error())
		return
	}

	// sort here so it's clearer what's happening
	sort.Ints(list1)
	sort.Ints(list2)

	// part 1
	ans, part1Err := day1Part1(&list1, &list2)
	if part1Err != nil {
		slog.Error(part1Err.Error())
	} else {
		slog.Info("The answer to part1 is: " + strconv.Itoa(ans))
	}

	// part 2
	ans2 := day2Part2(&list1, &list2)
	slog.Info("The answer to part2 is: " + strconv.Itoa(ans2))
}

// expects sorted array
func day1Part1(list1 *[]int, list2 *[]int) (int, error) {
	if len(*list1) != len(*list2) {
		err := errors.New("part1: list1 and list2 are not the same length")
		return 0, err
	}
	answer := 0
	for i := 0; i < len(*list1); i++ {
		if (*list1)[i] > (*list2)[i] {
			answer += (*list1)[i] - (*list2)[i]
		} else {
			answer += (*list2)[i] - (*list1)[i]
		}
	}
	return answer, nil
}

// expects sorted array
func day2Part2(list1 *[]int, list2 *[]int) int {
	answer, lastNum, lastAns, bottom, top := 0, 0, 0, 0, 0
	upper := len(*list2) - 1

	// loop through list1
	for i := 0; i < len(*list1); i++ {
		// if the last number and the current number are not the same
		if lastNum != (*list1)[i] {
			lastNum = (*list1)[i]
			// check that we're not exceeding the length of list2
			if top == upper {
				break
			}
			// find in list2
			bottom = utils.FindFirstOccurrence(list2, top, upper, lastNum)
			if bottom != -1 {
				top = utils.FindLastOccurrence(list2, top, upper, lastNum)
				lastAns = (top - (bottom - 1)) * lastNum
				answer += lastAns
			}
			// else if the last number and the current number are the same
		} else {
			answer += lastAns
		}
	}
	return answer
}

func day1Parse(puzzleInputPath *string) ([]int, []int, error) {
	data, err := utils.ReadFile(puzzleInputPath)
	if err != nil {
		slog.Error(err.Error())
		return nil, nil, err
	}
	var evens []int
	var odds []int
	fields := strings.Fields(string(data))
	for i := 0; i < len(fields); i++ {
		value, err := strconv.Atoi(fields[i])
		if err != nil {
			slog.Warn("value of " + fields[i] + " cannot be converted to integer")
		}
		if i%2 == 0 {
			evens = append(evens, value)
		} else {
			odds = append(odds, value)
		}
	}

	return evens, odds, nil
}
