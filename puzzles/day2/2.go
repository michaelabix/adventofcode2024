package day2

import (
	"bytes"
	"log/slog"
	"math"
	"strconv"
	"strings"

	"github.com/michaelabix/adventofcode2024/utils"
)

func Solve() {
	inputPath := "puzzles/day2/input.txt"
	data, dataErr := utils.ReadFile(&inputPath)
	if dataErr != nil {
		slog.Error(dataErr.Error())
	}
	input := parse(data)

	// part 1
	slog.Info("The answer to part 1 is: " + strconv.Itoa(part1(&input)))

	// part 2
	slog.Info("The answer to part 2 is: " + strconv.Itoa(part2(&input)))

}

func part1(data *[][]int) int {
	numSafe := 0
	for i := range len(*data) {
		if processLine(&(*data)[i]) {
			numSafe += 1
		}
	}
	return numSafe
}

func part2(data *[][]int) int {
	numSafe := 0
	for i := range len(*data) { // this could be a go routine
		if processLine(&(*data)[i]) {
			numSafe += 1
		} else { // check if one mistake exists
			for j := range len((*data)[i]) {
				newArray := (*data)[i]
				var newSlice []int
				// this is not the most efficient way of doing this
				// i should have messed with linked lists sooner
				if j == 0 {
					newSlice = newArray[j+1:]
				} else if j == len(newArray)-1 {
					newSlice = newArray[:j-1]
				} else {
					newSlice = newArray
					newSlice = append(newSlice[:j], newSlice[j+1:]...)
				}
				if processLine(&newSlice) {
					numSafe += 1
					break
				}
			}
		}
	}

	return numSafe
}

func processLine(line *[]int) bool {
	var isSafe bool
	// check if starting values are increasing, decreasing, or unsafe
	increasing := checkIncreasing((*line)[0], (*line)[1])

	// loop through remaining values and determine safety
	for i := 1; i < len(*line); i++ {
		isSafe = checkSafety(increasing, (*line)[i-1], (*line)[i])
		if !isSafe {
			break
		}
	}
	return isSafe
}

func checkIncreasing(left int, right int) int {
	increasing := -2
	if right > left {
		increasing = 1
	} else if right < left {
		increasing = 0
	} else {
		increasing = -1
	}
	return increasing
}

func checkSafety(increasing int, left int, right int) bool {
	if increasing == -1 || (increasing == 0 && right >= left) || (increasing == 1 && right <= left) || math.Abs(float64(right-left)) > 3 {
		return false
	}
	return true
}

func parse(data []byte) [][]int {
	// split bytes at newline
	split := []byte("\n")
	separated := bytes.SplitAfter(data, split)

	// initialize 2D array
	integers := make([][]int, len(separated))

	for i := 0; i < len(separated); i++ {
		// split at whitespace
		fields := strings.Fields(string(separated[i]))

		// initialize inner array
		integers[i] = make([]int, len(fields))

		// convert to int
		for j := range len(fields) {
			converted, err := strconv.Atoi(fields[j])
			if err != nil {
				slog.Warn(err.Error())
			}
			integers[i][j] = converted
		}
	}
	return integers
}
