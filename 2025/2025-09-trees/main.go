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

func rotateRight(n *AvlNode) {
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
		} else if it.Node.Balance == 2 {
			rotateRight(it.Node)
		}

		isLeft = it.IsLeft
	}
}

// precondition: n.Balance == -2
// precondition: n.Right.Balance == -1
func (n *AvlNode) rotateRight() *AvlNode {
	newRoot := n.Left
	n.Left = newRoot.Right
	newRoot.Right = n
	return newRoot
}

// precondition: n.Balance == 2
// precondition: n.Right.Balance == 1
func (n *AvlNode) rotateLeft() *AvlNode {
	newRoot := n.Right
	n.Right = newRoot.Left
	newRoot.Left = n
	return newRoot
}

func (n *AvlNode) rotateLeftRight() *AvlNode {
	n.Left = n.Left.rotateLeft()
	return n.rotateRight()
}

func (n *AvlNode) rotateRightLeft() *AvlNode {
	n.Right = n.Right.rotateRight()
	return n.rotateLeft()
}

// returns `(node, heightChange)`
// postcondition: `node` is a balanced AVL tree
// postcondition: `heightChange` is -1, 0, or 1
func (n *AvlNode) insert2(key string, value string) (*AvlNode, int) {
	if n == nil {
		return New(key, value), 1
	}

	// logging := key == "2"
	logging := false

	myHeightChange := 0
	var childHeightChange int
	cmp := strings.Compare(key, n.Key)
	if cmp < 0 {
		n.Left, childHeightChange = n.Left.insert2(key, value)
		if logging {
			fmt.Printf("got %d at key=%s\n", childHeightChange, n.Key)
		}
		if childHeightChange != 0 && n.Balance >= 0 {
			myHeightChange = childHeightChange
		}
		n.Balance += childHeightChange
	} else if cmp == 0 {
		n.Value = value
		return n, 0
	} else {
		n.Right, childHeightChange = n.Right.insert2(key, value)
		if childHeightChange != 0 && n.Balance <= 0 {
			myHeightChange = childHeightChange
		}
		n.Balance -= childHeightChange
	}

	if n.Balance == -2 {
		if n.Right.Balance == 1 {
			fmt.Printf("rotateRightLeft at key=%s\n", n.Key)
			r := n.rotateRightLeft()
			r.Balance = 0
			r.Left.Balance = 0
			r.Right.Balance = 0
			return r, 0
		} else {
			fmt.Printf("rotateLeft at key=%s\n", n.Key)
			r := n.rotateLeft()
			r.Balance = 0
			r.Left.Balance = 0
			return r, 0
		}
	} else if n.Balance == 2 {
		if n.Left.Balance == -1 {
			fmt.Printf("rotateLeftRight at key=%s\n", n.Key)
			r := n.rotateLeftRight()
			r.Balance = 0
			r.Left.Balance = 0
			// TODO: this seems wrong to hard-code
			r.Right.Balance = -1
			return r, 0
		} else {
			fmt.Printf("rotateRight at key=%s\n", n.Key)
			r := n.rotateRight()
			r.Balance = 0
			// TODO: this seems inconsistent with the above
			r.Right.Balance = 0
			return r, 0
		}
	}

	if logging {
		fmt.Printf("returning %d from key=%s\n", myHeightChange, n.Key)
	}
	return n, myHeightChange
}

func (n *AvlNode) Insert2(key string, value string) *AvlNode {
	r, _ := n.insert2(key, value)
	return r
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
	// root := New("50", "")
	// root = root.Insert2("60", "")
	// root = root.Insert2("70", "")
	// root = root.Insert2("30", "")
	// root = root.Insert2("20", "")
	// root = root.Insert2("40", "")
	// root = root.Insert2("45", "")
	// root = root.Insert2("35", "")
	// root = root.Insert2("10", "")
	// root.Print()

	root := New("5", "")
	fmt.Println("insert 6")
	root = root.Insert2("6", "")
	root.Print()
	fmt.Println("insert 8")
	root = root.Insert2("8", "")
	root.Print()
	fmt.Println("insert 3")
	root = root.Insert2("3", "")
	root.Print()
	fmt.Println("insert 2")
	root = root.Insert2("2", "")
	root.Print()
	fmt.Println("insert 4")
	root = root.Insert2("4", "")
	root.Print()
	fmt.Println("insert 7")
	root = root.Insert2("7", "")
	root.Print()
}
