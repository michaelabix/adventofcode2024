package day7

import (
	"bytes"
	"log/slog"
	"strconv"

	"github.com/michaelabix/adventofcode2024/utils"
)

func Solve() {
	inputPath := "puzzles/day7/sample.txt"
	data, dataErr := utils.ReadFile(&inputPath)
	if dataErr != nil {
		slog.Error(dataErr.Error())
	}

	parsed := parse(&data)
	slog.Info("The answer to part 1 is: " + strconv.Itoa(part1(parsed)))
	slog.Info("The answer to part 2 is: " + strconv.Itoa(part2(parsed)))

}

func part1(data *[][]int) int {
	num := 0
	for i := range len((*data)) {
		// initialize to match answer
		possibilities := []int{(*data)[i][0]}

		// loop through values to find possible paths
		for j := len((*data)[i]) - 1; j > 1; j-- {
			possibilities = process(possibilities, (*data)[i][j], false)
		}

		// loop through results to find if any matched
		for k := range len(possibilities) {
			if possibilities[k] == (*data)[i][1] {
				num += (*data)[i][0]
				break
			}
		}
	}

	return num
}

func part2(data *[][]int) int {
	num := 0
	for i := range len((*data)) {
		// initialize to match answer
		possibilities := []int{(*data)[i][0]}
		// loop through values to find possible paths
		for j := len((*data)[i]) - 1; j > 1; j-- {
			possibilities = process(possibilities, (*data)[i][j], true)
		}

		// loop through results to find if any matched
		for k := range len(possibilities) {
			if possibilities[k] == (*data)[i][1] {
				num += (*data)[i][0]
				break
			}
		}
	}

	return num
}

func process(possibilities []int, val int, part3 bool) []int {
	var newVals []int
	for i := range len(possibilities) {
		// only add additional path if divisible
		if tryMod(possibilities[i], val) {
			newVals = append(newVals, possibilities[i]/val)
		}
		if part3 {
			newVals = append(newVals, tryOr(possibilities[i], val))
		}
		// always subtract
		newVals = append(newVals, possibilities[i]-val)
	}
	return newVals
}

func tryOr(num int, val int) int {
	divisor := FindDivisor(val)
	if num%divisor == val {
		return num / divisor
	} else {
		return 0
	}
}

func FindDivisor(val int) int {
	count := 1
	for val != 0 {
		val /= 10
		count *= 10
	}
	return count
}

func tryMod(dividend int, divisor int) bool {
	return dividend%divisor == 0
}

func parse(data *[]byte) *[][]int {
	var parsed [][]int
	split := bytes.Split(*data, []byte("\n"))
	for i := range len(split) {
		splitAgain := bytes.Split(split[i], []byte(" "))
		answer := bytes.Split(splitAgain[0], []byte(":"))
		ans, _ := strconv.Atoi(string(answer[0]))
		values := []int{ans}
		for j := 1; j < len(splitAgain); j++ {
			val, _ := strconv.Atoi(string(splitAgain[j]))
			values = append(values, val)
		}
		parsed = append(parsed, values)
	}
	return &parsed
}
