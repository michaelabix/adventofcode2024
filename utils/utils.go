package utils

import (
	"math"
	"os"
)

// linked list struct
// probably figure out how to make Val generic later
type Node struct {
	Val  int
	Prev *Node
	Next *Node
	//Metadata any
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

func ListAppend(head *Node, tail *Node, value int) (*Node, *Node) {
	node := new(Node)
	node.Val = value
	if tail == nil {
		return node, node
	}

	current := tail
	current.Next = node
	node.Prev = tail

	return head, node
}

// tail can be any node & is used to limit searching if needed
// set tail to nil if there's no limit on the search
func ListSearch(head *Node, tail *Node, value int) *Node {
	current := head
	for current.Next != nil {
		if tail != nil {
			if current.Next == tail {
				return nil
			}
		}
		if current.Next.Val == value {
			return current.Next
		}
		current = current.Next
	}
	return nil
}

// set head to nil if there's no limit on the search
func ListReverseSearch(head *Node, tail *Node, value int) *Node {
	current := tail
	for current.Prev != nil {
		if head != nil {
			if current.Prev == head {
				return nil
			}
		}
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

// update this to update tail
func ListMoveNodeAfter(tail *Node, node *Node, placeAfter *Node) *Node {
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
	node.Prev = placeAfter
	if placeAfter.Next != nil {
		// set Next
		node.Next = placeAfter.Next
		placeAfter.Next.Prev = node
	} else {
		node.Next = nil
		tail = node
	}

	placeAfter.Next = node
	return tail
}

func ListMoveNodeBefore(head *Node, node *Node, placeBefore *Node) *Node {
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
		head = node
	}

	placeBefore.Prev = node

	return head
}

func ListFindHead(node *Node) *Node {
	curr := node
	for curr.Prev != nil {
		curr = curr.Prev
	}
	head := curr
	return head
}
