package main

import (
	"log/slog"
	"os"
	"strconv"

	"github.com/michaelabix/adventofcode2024/puzzles"
)

func main() {
	if len(os.Args) > 1 && isInt(os.Args[1]) {
		day, _ := strconv.Atoi(os.Args[1])
		switchDay(day)
	} else {
		slog.Error("First argument is not an integer from 1 to 25")
	}
}

func isInt(digit string) bool {
	integer, err := strconv.Atoi(digit)
	if err == nil && integer > 0 && integer < 26 {
		return true
	} else {
		return false
	}
}

func switchDay(day int) {
	switch day {
	case 1:
		puzzles.SolveDay1()
	case 2:
		puzzles.SolveDay2
	default:
		slog.Info("Has the puzzle been solved yet?")
	}
}
