package avl

import (
	"fmt"
	"strings"
)

type Node struct {
	Key     string
	Value   string
	Balance int
	Left    *Node
	Right   *Node
}

func New(key string, value string) *Node {
	return &Node{Key: key, Value: value, Balance: 0, Left: nil, Right: nil}
}

// precondition: n.Balance == 2
// precondition: n.Left.Balance == 1
func (n *Node) rotateRight() *Node {
	newRoot := n.Left
	n.Left = newRoot.Right
	newRoot.Right = n
	return newRoot
}

// precondition: n.Balance == -2
// precondition: n.Right.Balance == -1
func (n *Node) rotateLeft() *Node {
	newRoot := n.Right
	n.Right = newRoot.Left
	newRoot.Left = n
	return newRoot
}

// precondition: n.Balance == 2
// precondition: n.Left.Balance == -1
func (n *Node) rotateLeftRight() *Node {
	n.Left = n.Left.rotateLeft()
	return n.rotateRight()
}

// precondition: n.Balance == -2
// precondition: n.Right.Balance == 1
func (n *Node) rotateRightLeft() *Node {
	n.Right = n.Right.rotateRight()
	return n.rotateLeft()
}

// returns `(node, heightChange)`
// postcondition: `node` is a balanced AVL tree
// postcondition: `heightChange` is -1, 0, or 1
func (n *Node) insert(key string, value string) (*Node, int) {
	if n == nil {
		return New(key, value), 1
	}

	myHeightChange := 0
	var childHeightChange int
	cmp := strings.Compare(key, n.Key)
	if cmp < 0 {
		n.Left, childHeightChange = n.Left.insert(key, value)
		if childHeightChange != 0 && n.Balance >= 0 {
			myHeightChange = childHeightChange
		}
		n.Balance += childHeightChange
	} else if cmp == 0 {
		n.Value = value
		return n, 0
	} else {
		n.Right, childHeightChange = n.Right.insert(key, value)
		if childHeightChange != 0 && n.Balance <= 0 {
			myHeightChange = childHeightChange
		}
		n.Balance -= childHeightChange
	}

	if n.Balance == -2 {
		if n.Right.Balance == 1 {
			b := n.Right.Left.Balance
			fmt.Printf("rotateRightLeft at key=%s\n", n.Key)
			r := n.rotateRightLeft()
			r.Balance = 0
			if b == -1 {
				r.Left.Balance = 1
				r.Right.Balance = 0
			} else if b == 0 {
				r.Left.Balance = 0
				r.Right.Balance = 0
			} else {
				r.Left.Balance = 0
				r.Right.Balance = -1
			}
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
			b := n.Left.Right.Balance
			fmt.Printf("rotateLeftRight at key=%s\n", n.Key)
			r := n.rotateLeftRight()
			r.Balance = 0
			if b == -1 {
				r.Left.Balance = 1
				r.Right.Balance = 0
			} else if b == 0 {
				r.Left.Balance = 0
				r.Right.Balance = 0
			} else {
				r.Left.Balance = 0
				r.Right.Balance = -1
			}
			return r, 0
		} else {
			fmt.Printf("rotateRight at key=%s\n", n.Key)
			r := n.rotateRight()
			r.Balance = 0
			r.Right.Balance = 0
			return r, 0
		}
	}

	return n, myHeightChange
}

func (n *Node) Insert(key string, value string) *Node {
	r, _ := n.insert(key, value)
	return r
}

func (n *Node) Retrieve(key string) string {
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

func (n *Node) printRec() {
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

func (n Node) Print() {
	n.printRec()
	fmt.Println("")
}
