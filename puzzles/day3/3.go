package day3

import (
	"errors"
	"log/slog"
	"regexp"
	"strconv"
	"strings"

	"github.com/michaelabix/adventofcode2024/utils"
)

func Solve() {
	inputPath := "puzzles/day3/sample.txt"
	data, dataErr := utils.ReadFile(&inputPath)
	if dataErr != nil {
		slog.Error(dataErr.Error())
	}
	input := parseMul(string(data))

	// part 1
	slog.Info("The answer to part 1 is: " + strconv.Itoa(answer(&input)))

	input = parseDo(string(data))

	// part 2
	slog.Info("The answer to part 2 is: " + strconv.Itoa(answer(&input)))
}

func answer(data *[]string) int {
	answer := 0
	for p := range len(*data) {
		product, err := multiply((*data)[p])
		if err == nil {
			answer += product
		}
	}
	return answer
}

func multiply(data string) (int, error) {
	split := strings.Split((data), ",")
	val1, err1 := strconv.Atoi(split[0])
	val2, err2 := strconv.Atoi(split[1])

	if err1 != nil || err2 != nil {
		slog.Error("Could not parse integers from string")
		return 0, errors.New("multiply: could not parse integers from string " + data)
	}

	return val1 * val2, nil
}

func parseMul(data string) []string {
	var pairs []string
	r := regexp.MustCompile(`mul\([0-9]{1,3}\,[0-9]{1,3}\)`)
	matches := r.FindAllStringSubmatch(data, -1)
	for i := range len(matches) {
		for j := range len(matches[i]) {
			pairs = append(pairs, matches[i][j][4:len(matches[i][j])-1])
		}
	}
	return pairs
}

func parseDo(data string) []string {
	var pairs []string
	instructions := strings.Split(data, "do()")
	for i := range len(instructions) {
		newString := (strings.Split(instructions[i], "don't()"))
		results := parseMul(newString[0])
		pairs = append(pairs, results...)
	}
	return pairs
}
