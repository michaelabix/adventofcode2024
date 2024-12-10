package day9

import (
	"log/slog"
	"strconv"

	"github.com/michaelabix/adventofcode2024/utils"
)

func Solve() {
	inputPath := "puzzles/day9/sample.txt"
	data, dataErr := utils.ReadFile(&inputPath)
	if dataErr != nil {
		slog.Error(dataErr.Error())
	}

	head, tail := parse(&data)

	slog.Info("The answer to part 1 is: " + strconv.Itoa(part1(head, tail)))
}

func part1(head *utils.Node, tail *utils.Node) int {
	num := 0

	prev := tail

	for prev.Prev != nil {
		if prev.Val != -1 {
			found := utils.ListSearch(head, prev, -1)
			if found != nil {
				before := prev.Prev
				head = utils.ListMoveNodeBefore(head, prev, found)
				tail = utils.ListMoveNodeAfter(tail, found, before)
				prev = tail
			} else {
				prev = prev.Prev
			}
			if prev.Prev != nil {
				prev = prev.Prev
			}
		} else {
			prev = prev.Prev
		}
	}
	// do some math
	i := 0
	// reset curr to head
	curr := head
	for curr.Next != nil {
		if curr.Val != -1 {
			num += i * curr.Val
		}
		i += 1
		curr = curr.Next
		// we never care about tail because we already know it's -1 after sorting
	}
	return num
}

func part2(head *utils.Node, tail *utils.Node) int {
	return 0
}

func parse(data *[]byte) (*utils.Node, *utils.Node) {
	var tail *utils.Node
	var head *utils.Node
	nextID := 0
	for i := range len(*data) {
		if i%2 == 0 {
			num, _ := strconv.Atoi(string((*data)[i]))
			for j := 0; j < num; j++ {
				head, tail = utils.ListAppend(head, tail, nextID)
			}
			nextID += 1
		} else {
			num, _ := strconv.Atoi(string((*data)[i]))
			for j := 0; j < num; j++ {
				head, tail = utils.ListAppend(head, tail, -1)
			}
		}
	}

	return head, tail

}
