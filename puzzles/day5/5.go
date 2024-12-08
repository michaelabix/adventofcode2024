package day5

import (
	"bytes"
	"log/slog"
	"strconv"

	"github.com/michaelabix/adventofcode2024/utils"
)

func Solve() {
	inputPath := "puzzles/day5/sample.txt"
	data, dataErr := utils.ReadFile(&inputPath)
	if dataErr != nil {
		slog.Error(dataErr.Error())
	}

	instructions, updates := parse(&data)

	answer1, indexes := part1(&instructions, updates)
	slog.Info("The answer to part 1 is: " + strconv.Itoa(answer1))

	slog.Info("The answer to part 2 is: " + strconv.Itoa(part2(&instructions, updates, &indexes)))

}

func part1(instructions *map[int]map[int]bool, updates []*utils.Node) (int, []int) {
	num := 0
	var brokenIndexes []int
	for i := range len(updates) {
		ruleBroken := checkRules(instructions, updates[i])
		// find middle if rules are never broken
		if !ruleBroken {
			middle := utils.ListFindMiddle(updates[i])
			num += middle.Val
		} else {
			brokenIndexes = append(brokenIndexes, i)
		}
	}

	return num, brokenIndexes
}

func part2(instructions *map[int]map[int]bool, updates []*utils.Node, indexes *[]int) int {
	num := 0
	for i := range len(*indexes) {
		curr := updates[(*indexes)[i]]
		// loop through list

		head := sort(instructions, updates[(*indexes)[i]], curr)

		// print to make sure
		middle := utils.ListFindMiddle(head)
		num += middle.Val
	}

	return num
}

func checkRules(instructions *map[int]map[int]bool, update *utils.Node) bool {
	ruleBroken := false
	curr := *update
	// loop through list
	for curr.Next != nil {
		ruleBroken, _ = process(instructions, &curr)
		if ruleBroken {
			break
		}
		curr = *curr.Next
	}
	// process last node
	if !ruleBroken {
		ruleBroken, _ = process(instructions, &curr)
	}
	return ruleBroken
}

func sort(instructions *map[int]map[int]bool, head *utils.Node, curr *utils.Node) *utils.Node {
	for curr.Next != nil {
		ruleBroken, moveBefore := process(instructions, curr)
		if ruleBroken {
			utils.ListMoveNodeBefore(head, curr, moveBefore)

		} else {
			curr = curr.Next
		}
	}
	ruleBroken, moveBefore := process(instructions, curr)
	head = utils.ListFindHead(head)
	if ruleBroken {
		utils.ListMoveNodeBefore(head, curr, moveBefore)
		if curr.Next != nil {
			head = sort(instructions, curr, curr)
		}
	}
	return head
}

func process(instructions *map[int]map[int]bool, curr *utils.Node) (bool, *utils.Node) {
	// check if the number has before instructions
	if _, ok := (*instructions)[curr.Val]; ok {
		// check if previous nodes contain values that must be after the current value
		for k := range (*instructions)[curr.Val] {
			node := utils.ListReverseSearch(curr, k)
			if node != nil {
				return true, node
			}
		}
	}

	return false, nil
}

func parse(data *[]byte) (map[int]map[int]bool, []*utils.Node) {
	split := bytes.Split(*data, []byte("\n\n"))
	instruct := bytes.Split(split[0], []byte("\n"))
	lines := bytes.Split(split[1], []byte("\n"))
	var lists []*utils.Node
	beforeInstructions := make(map[int]map[int]bool)
	for i := range len(lines) {
		temp := bytes.Split(lines[i], []byte(","))
		var head *utils.Node
		for j := range len(temp) {
			num, _ := strconv.Atoi(string(temp[j]))
			head = utils.ListAppend(head, num)
		}
		lists = append(lists, head)
	}

	for i := range len(instruct) {
		nums := bytes.Split(instruct[i], []byte("|"))
		left, _ := strconv.Atoi(string(nums[0]))
		right, _ := strconv.Atoi(string(nums[1]))
		if _, ok := beforeInstructions[left]; ok {
			beforeInstructions[left][right] = true
		} else {
			beforeInstructions[left] = make(map[int]bool)
			beforeInstructions[left][right] = true
		}
	}
	return beforeInstructions, lists
}
