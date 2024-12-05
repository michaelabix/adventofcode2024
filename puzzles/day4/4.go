package day4

import (
	"bytes"
	"log/slog"
	"strconv"

	"github.com/michaelabix/adventofcode2024/utils"
)

func Solve() {
	inputPath := "puzzles/day4/input.txt"
	data, dataErr := utils.ReadFile(&inputPath)
	if dataErr != nil {
		slog.Error(dataErr.Error())
	}

	// split at newline to make 2D array
	parsed := bytes.Split(data, []byte("\n"))

	slog.Info("The answer to part 1 is: " + strconv.Itoa(part1(&parsed)))
	slog.Info("The answer to part 1 is: " + strconv.Itoa(part2(&parsed)))
}

func part1(data *[][]byte) int {
	xmas := []byte{88, 77, 65, 83}
	num := evaluate1(data, &xmas)
	return num
}

func part2(data *[][]byte) int {
	mas := []byte{65, 77, 83}
	num := evaluate2(data, &mas)
	return num
}

// this is super gross and is what i get for writing code after 10pm
// this will get fixed later
func evaluate1(data *[][]byte, letters *[]byte) int {
	num := 0
	for i := 0; i < len(*data); i++ {
		for j := 0; j < len((*data)[i]); j++ {
			// find x
			if check(data, i, j, (*letters)[0]) {
				// check up
				if i > 2 {
					if check(data, i-1, j, (*letters)[1]) {
						if check(data, i-2, j, (*letters)[2]) {
							if check(data, i-3, j, (*letters)[3]) {
								num += 1
							}
						}
					}
				}

				// check down
				if i < len(*data)-3 {
					if check(data, i+1, j, (*letters)[1]) {
						if check(data, i+2, j, (*letters)[2]) {
							if check(data, i+3, j, (*letters)[3]) {
								num += 1
							}
						}
					}
				}

				// check left
				if j > 2 {
					if check(data, i, j-1, (*letters)[1]) {
						if check(data, i, j-2, (*letters)[2]) {
							if check(data, i, j-3, (*letters)[3]) {
								num += 1
							}
						}
					}
				}

				// check right
				if j < len((*data)[i])-3 {
					if check(data, i, j+1, (*letters)[1]) {
						if check(data, i, j+2, (*letters)[2]) {
							if check(data, i, j+3, (*letters)[3]) {
								num += 1
							}
						}
					}
				}

				// check diagonal top left
				if i > 2 && j > 2 {
					if check(data, i-1, j-1, (*letters)[1]) {
						if check(data, i-2, j-2, (*letters)[2]) {
							if check(data, i-3, j-3, (*letters)[3]) {
								num += 1
							}
						}
					}
				}

				// check diagonal top right
				if i > 2 && j < len((*data)[i])-3 {
					if check(data, i-1, j+1, (*letters)[1]) {
						if check(data, i-2, j+2, (*letters)[2]) {
							if check(data, i-3, j+3, (*letters)[3]) {
								num += 1
							}
						}
					}
				}

				// check diagonal bottom left
				if i < len(*data)-3 && j > 2 {
					if check(data, i+1, j-1, (*letters)[1]) {
						if check(data, i+2, j-2, (*letters)[2]) {
							if check(data, i+3, j-3, (*letters)[3]) {
								num += 1
							}
						}
					}
				}

				// check diagonal bottom right
				if i < len(*data)-3 && j < len((*data)[i])-3 {
					if check(data, i+1, j+1, (*letters)[1]) {
						if check(data, i+2, j+2, (*letters)[2]) {
							if check(data, i+3, j+3, (*letters)[3]) {
								num += 1
							}
						}
					}
				}
			}
		}
	}
	return num
}

func evaluate2(data *[][]byte, letters *[]byte) int {
	num := 0
	for i := 0; i < len(*data); i++ {
		for j := 0; j < len((*data)[i]); j++ {
			// find A
			if check(data, i, j, (*letters)[0]) {
				half := false
				if i > 0 && j > 0 && i < len(*data)-1 && j < len((*data)[i])-1 {
					// check diagonal top left and right bottom
					if check(data, i-1, j-1, (*letters)[1]) {
						if check(data, i+1, j+1, (*letters)[2]) {
							half = true
						}

					} else if check(data, i-1, j-1, (*letters)[2]) {
						if check(data, i+1, j+1, (*letters)[1]) {
							half = true
						}
					}

					// check diagonal bottom left and top right
					if check(data, i+1, j-1, (*letters)[1]) && half {
						if check(data, i-1, j+1, (*letters)[2]) {
							num += 1
						}
					} else if check(data, i+1, j-1, (*letters)[2]) && half {
						if check(data, i-1, j+1, (*letters)[1]) {
							num += 1
						}
					}
				}
			}
		}
	}
	return num
}

func check(data *[][]byte, y int, x int, letter byte) bool {
	return (*data)[y][x] == letter
}
