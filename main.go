package main

import (
	"log/slog"
	"os"
	"strconv"
	"time"

	"github.com/michaelabix/adventofcode2024/puzzles/day1"
	"github.com/michaelabix/adventofcode2024/puzzles/day2"
	"github.com/michaelabix/adventofcode2024/puzzles/day3"
	"github.com/michaelabix/adventofcode2024/puzzles/day4"
	"github.com/michaelabix/adventofcode2024/puzzles/day5"
	"github.com/michaelabix/adventofcode2024/puzzles/day7"
)

func main() {
	start := time.Now()
	if len(os.Args) > 1 && isInt(os.Args[1]) {
		day, _ := strconv.Atoi(os.Args[1])
		switchDay(day)
	} else {
		slog.Error("First argument is not an integer from 1 to 25")
	}
	slog.Info("Execution time: " + time.Since(start).String())
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
		day1.Solve()
	case 2:
		day2.Solve()
	case 3:
		day3.Solve()
	case 4:
		day4.Solve()
	case 5:
		day5.Solve()
	case 7:
		day7.Solve()
	default:
		slog.Info("Has the puzzle been solved yet?")
	}
}
