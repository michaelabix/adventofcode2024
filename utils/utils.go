package utils

import (
	"math"
	"os"
)

// linked list struct
type Node struct {
	Val  int
	Prev *Node
	Next *Node
}

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

func ListAppend(head *Node, value int) *Node {
	node := new(Node)
	node.Val = value
	if head == nil {
		return node
	}

	current := head
	prev := head
	for current.Next != nil {
		current = current.Next
		prev = current
	}
	current.Next = node
	node.Prev = prev

	return head
}

func ListSearch(head *Node, value int) *Node {
	current := head
	for current.Next != nil {
		if current.Next.Val == value {
			return current.Next
		}
		current = current.Next
	}
	return nil
}

func ListReverseSearch(tail *Node, value int) *Node {
	current := tail
	for current.Prev != nil {
		if current.Prev.Val == value {
			return current.Prev
		}
		current = current.Prev
	}
	return nil
}

func ListFindMiddle(head *Node) *Node {
	current := head
	count := 0
	middleNode := head
	for current.Next != nil {
		count += 1
		current = current.Next
	}
	middle := count / 2
	for i := 0; i < middle; i++ {
		middleNode = middleNode.Next
	}
	return middleNode
}

func ListMoveNodeAfter(head *Node, node *Node, placeAfter *Node) {
	// pop
	if node.Prev != nil {
		// set previous node's next to next node
		node.Prev.Next = node.Next
	}
	if node.Next != nil {
		// set next's previous to previous node
		node.Next.Prev = node.Prev
	}

	// set links for current node
	node.Next = placeAfter.Next
	if placeAfter.Next != nil {
		// set previous
		node.Prev = placeAfter.Next.Prev
		placeAfter.Next.Prev = node
	} else {
		node.Prev = nil
	}

	placeAfter.Next = node
}

func ListMoveNodeBefore(head *Node, node *Node, placeBefore *Node) {
	// pop
	if node.Prev != nil {
		// set previous node's next to next node
		node.Prev.Next = node.Next
	}
	if node.Next != nil {
		// set next's previous to previous node
		node.Next.Prev = node.Prev
	}

	// set links for current node
	node.Next = placeBefore
	if placeBefore.Prev != nil {
		// set previous
		node.Prev = placeBefore.Prev
		placeBefore.Prev.Next = node
	} else {
		node.Prev = nil
	}

	placeBefore.Prev = node
}

func ListFindHead(node *Node) *Node {
	curr := node
	for curr.Prev != nil {
		curr = curr.Prev
	}
	head := curr
	return head
}
