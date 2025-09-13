package avl

import (
	"fmt"
	"io"
	"os"
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

func (n *Node) rotateRight() *Node {
	newRoot := n.Left
	n.Left = newRoot.Right
	newRoot.Right = n
	return newRoot
}

func (n *Node) rotateLeft() *Node {
	newRoot := n.Right
	n.Right = newRoot.Left
	newRoot.Left = n
	return newRoot
}

func (n *Node) rotateLeftRight() *Node {
	n.Left = n.Left.rotateLeft()
	return n.rotateRight()
}

func (n *Node) rotateRightLeft() *Node {
	n.Right = n.Right.rotateRight()
	return n.rotateLeft()
}

// returns `(node, heightChanged)`
// postcondition: `node` is a balanced AVL tree
//
// Returning `heightChanged` allows us to more efficiently re-calculate the balance of
// intermediate nodes without recomputing the heights of all child trees -- at the cost
// of considerably complicating the implementation.
func (n *Node) insert(key string, value string) (*Node, bool) {
	if n == nil {
		return New(key, value), true
	}

	myHeightChanged := false
	var childHeightChanged bool
	cmp := strings.Compare(key, n.Key)
	if cmp < 0 {
		n.Left, childHeightChanged = n.Left.insert(key, value)
		if childHeightChanged {
			if n.Balance >= 0 {
				myHeightChanged = true
			}
			n.Balance += 1
		} else {
			return n, false
		}
	} else if cmp == 0 {
		n.Value = value
		return n, false
	} else {
		n.Right, childHeightChanged = n.Right.insert(key, value)
		if childHeightChanged {
			if n.Balance <= 0 {
				myHeightChanged = true
			}
			n.Balance -= 1
		} else {
			return n, false
		}
	}

	if n.Balance == -2 {
		if n.Right.Balance == 1 {
			b := n.Right.Left.Balance
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
			return r, false
		} else {
			r := n.rotateLeft()
			r.Balance = 0
			r.Left.Balance = 0
			return r, false
		}
	} else if n.Balance == 2 {
		if n.Left.Balance == -1 {
			b := n.Left.Right.Balance
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
			return r, false
		} else {
			r := n.rotateRight()
			r.Balance = 0
			r.Right.Balance = 0
			return r, false
		}
	}

	return n, myHeightChanged
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

func (n *Node) stringRec(sb *strings.Builder) {
	if n == nil {
		sb.WriteString("()")
	} else if n.Left == nil && n.Right == nil {
		sb.WriteString(fmt.Sprintf("%s:%d", n.Key, n.Balance))
	} else {
		sb.WriteString(fmt.Sprintf("(%s:%d ", n.Key, n.Balance))
		n.Left.stringRec(sb)
		sb.WriteByte(' ')
		n.Right.stringRec(sb)
		sb.WriteByte(')')
	}
}

func (n Node) String() string {
	var sb strings.Builder
	n.stringRec(&sb)
	return sb.String()
}

func (n Node) Fprint(w io.Writer) {
	fmt.Fprintf(w, "%s\n", n.String())
}

func (n Node) Print() {
	n.Fprint(os.Stdout)
}

func (n *Node) Height() int {
	if n == nil {
		return 0
	} else {
		return max(n.Left.Height(), n.Right.Height()) + 1
	}
}

func (n *Node) Check() {
	if n == nil {
		return
	}
	leftHeight := n.Left.Height()
	rightHeight := n.Right.Height()
	if n.Balance != leftHeight-rightHeight {
		fmt.Fprintf(
			os.Stderr,
			"error: balance is %d, expected %d = %d - %d\n",
			n.Balance,
			leftHeight-rightHeight,
			leftHeight,
			rightHeight)
		fmt.Fprint(os.Stderr, "  ")
		n.Fprint(os.Stderr)
	}

	n.Left.Check()
	n.Right.Check()
}
