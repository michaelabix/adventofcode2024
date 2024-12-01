package utils

import (
	"math"
	"os"
)

func ReadFile(path *string) ([]byte, error) {
	data, err := os.ReadFile(*path)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// binary searches for fun
// returns -1 if not found
func FindFirstOccurrence(list *[]int, low int, high int, value int) int {
	index := -1
	for low < high {
		mid := int(math.Floor((float64(high + low)) / float64(2)))
		if (*list)[mid] < value {
			low = mid + 1
		} else {
			high = mid
			if (*list)[mid] == value {
				index = mid
			}
		}
	}
	return index
}

func FindLastOccurrence(list *[]int, low int, high int, value int) int {
	index := -1
	for low < high {
		mid := int(math.Floor((float64(high + low)) / float64(2)))
		if (*list)[mid] > value {
			high = mid
		} else {
			low = mid + 1
			if (*list)[mid] == value {
				index = mid
			}
		}
	}
	return index
}
