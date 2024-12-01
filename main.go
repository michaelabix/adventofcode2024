package main

import (
	"log/slog"
	"os"
	"strconv"

	"github.com/michaelabix/adventofcode2024/day1"
)

func main() {
	if len(os.Args) > 1 && isInt(os.Args[1]) {
		// fix this so it's more dynamic later
		day1.Solve()
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
