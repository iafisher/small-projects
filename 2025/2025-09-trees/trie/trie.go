package trie

type Node struct {
	Children [26]*Node
	IsLeaf   bool
}

func (n Node) index(b byte) int {
	i := int(b - 'a')
	if i >= 26 {
		panic("invariant violated: key contains character other than a-z")
	}
	return i
}

func (n *Node) findNode(key string) *Node {
	if key == "" {
		return n
	}

	child := n.Children[n.index(key[0])]
	if child == nil {
		return nil
	}
	return child.findNode(key[1:])
}

func (n *Node) Retrieve(key string) bool {
	found := n.findNode(key)
	return found != nil && found.IsLeaf
}

func (n *Node) buildStringSet(soFar []byte, acc *[]string) {
	if n.IsLeaf {
		*acc = append(*acc, string(soFar))
	} else {
		for i, c := range n.Children {
			if c != nil {
				b := byte(i) + 'a'
				c.buildStringSet(append(soFar, b), acc)
			}
		}
	}
}

func (n *Node) AllMatches(key string) []string {
	found := n.findNode(key)
	if found == nil {
		return nil
	}
	r := []string{}
	found.buildStringSet([]byte(key), &r)
	return r
}

func (n *Node) Insert(key string) {
	if key == "" {
		n.IsLeaf = true
		return
	}

	i := n.index(key[0])
	if n.Children[i] == nil {
		n.Children[i] = &Node{}
	}
	n.Children[i].Insert(key[1:])
}
