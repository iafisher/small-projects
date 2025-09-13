package main

import (
	"fmt"
	"strings"
)

type AvlNode struct {
	Key     string
	Value   string
	Balance int
	Left    *AvlNode
	Right   *AvlNode
}

type StackItem struct {
	Node   *AvlNode
	IsLeft bool
}

// Algorithm:
//
//   Recursively traverse tree until insertion point is reached.
//   Retrace steps, keeping track of height change.
//   At each node, calculate new balance
//   If |balance| == 2, apply transformation

func New(key string, value string) *AvlNode {
	return &AvlNode{Key: key, Value: value, Balance: 0, Left: nil, Right: nil}
}

func buildStack(n *AvlNode, key string, value string) []StackItem {
	stack := []StackItem{StackItem{Node: n, IsLeft: false}}
	for {
		it := stack[len(stack)-1].Node
		cmp := strings.Compare(key, it.Key)
		if cmp < 0 {
			if it.Left == nil {
				it.Left = New(key, value)
				// TODO: messy
				return append(stack, StackItem{Node: it.Left, IsLeft: true})
			} else {
				stack = append(stack, StackItem{Node: it.Left, IsLeft: true})
			}
		} else if cmp == 0 {
			it.Value = value
			return []StackItem{}
		} else {
			if it.Right == nil {
				it.Right = New(key, value)
				// TODO: messy
				return append(stack, StackItem{Node: it.Right, IsLeft: false})
			} else {
				stack = append(stack, StackItem{Node: it.Right, IsLeft: false})
			}
		}
	}
}

func updateBalances(stack []StackItem) {
	if len(stack) < 2 {
		panic(
			fmt.Sprintf(
				"precondition violated (updateBalances): stack should have at least 2 items (actual: %d)",
				len(stack)))
	}

	isLeft := stack[len(stack)-1].IsLeft
	for i := len(stack) - 2; i >= 0; i-- {
		it := stack[i]
		if isLeft {
			it.Node.Balance += 1
		} else {
			it.Node.Balance -= 1
		}

		if it.Node.Balance == 0 {
			break
		}

		isLeft = it.IsLeft
	}
}

func (n *AvlNode) Insert(key string, value string) *AvlNode {
	if n == nil {
		return New(key, value)
	}

	// How is balance updated when we insert a new node X?
	//
	// Suppose X was inserted as left child of Y.
	// If balance(Y) was 0, it is now 1.
	// If balance(Y) was -1, it is now 0.
	//
	// Suppose X was inserted as right child of Y.
	// If balance(Y) was 0, it is now -1.
	// If balance(Y) was 1, it is now 0.

	stack := buildStack(n, key, value)
	if len(stack) == 0 {
		return n
	}

	updateBalances(stack)
	return n
}

func (n *AvlNode) Retrieve(key string) string {
	if n == nil {
		return ""
	}

	cmp := strings.Compare(key, n.Key)
	if cmp < 0 {
		return n.Left.Retrieve(key)
	} else if cmp == 0 {
		return n.Value
	} else {
		return n.Left.Retrieve(key)
	}
}

func (n *AvlNode) printRec() {
	if n == nil {
		fmt.Print("()")
	} else if n.Left == nil && n.Right == nil {
		fmt.Printf("%s:%d", n.Key, n.Balance)
	} else {
		fmt.Printf("(%s:%d ", n.Key, n.Balance)
		n.Left.printRec()
		fmt.Print(" ")
		n.Right.printRec()
		fmt.Print(")")
	}
}

func (n AvlNode) Print() {
	n.printRec()
	fmt.Println("")
}

func main() {
	root := New("50", "")
	root = root.Insert("60", "")
	root = root.Insert("70", "")
	root = root.Insert("30", "")
	root = root.Insert("20", "")
	root = root.Insert("40", "")
	root = root.Insert("45", "")
	root = root.Insert("35", "")
	root = root.Insert("10", "")
	root.Print()
}
