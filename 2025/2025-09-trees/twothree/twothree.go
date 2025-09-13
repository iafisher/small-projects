package twothree

import (
	"fmt"
	"strings"
)

type TwoNode struct {
	Key   string
	Value string
	Left  Node
	Right Node
}

type ThreeNode struct {
	Key1   string
	Key2   string
	Value1 string
	Value2 string
	Left   Node
	Middle Node
	Right  Node
}

type Node interface {
	insert(string, string) (Node, *TwoNode)
	Insert(string, string) Node
	String() string
}

func New(key string, value string) Node {
	return TwoNode{Key: key, Value: value, Left: nil, Right: nil}
}

func (n TwoNode) insert(key string, value string) (Node, *TwoNode) {
	cmp := strings.Compare(key, n.Key)

	var promoted *TwoNode
	if cmp < 0 {
		if n.Left == nil {
			return ThreeNode{
				Key1:   key,
				Key2:   n.Key,
				Value1: value,
				Value2: n.Value,
				Left:   nil,
				Middle: nil,
				Right:  nil,
			}, nil
		} else {
			n.Left, promoted = n.Left.insert(key, value)
			if promoted != nil {
				return ThreeNode{
					Key1:   promoted.Key,
					Key2:   n.Key,
					Value1: promoted.Value,
					Value2: n.Value,
					Left:   promoted.Left,
					Middle: promoted.Right,
					Right:  n.Right,
				}, nil
			} else {
				return n, nil
			}
		}
	} else if cmp == 0 {
		n.Value = value
		return n, nil
	} else {
		if n.Right == nil {
			return ThreeNode{
				Key1:   n.Key,
				Key2:   key,
				Value1: n.Value,
				Value2: value,
				Left:   nil,
				Middle: nil,
				Right:  nil,
			}, nil
		} else {
			n.Right, promoted = n.Right.insert(key, value)
			if promoted != nil {
				return ThreeNode{
					Key1:   n.Key,
					Key2:   promoted.Key,
					Value1: n.Value,
					Value2: promoted.Value,
					Left:   n.Left,
					Middle: promoted.Left,
					Right:  promoted.Right,
				}, nil
			} else {
				return n, nil
			}
		}
	}
}

func (n TwoNode) Insert(key string, value string) Node {
	r, promoted := n.insert(key, value)
	if promoted != nil {
		return promoted
	} else {
		return r
	}
}

func (n ThreeNode) insert(key string, value string) (Node, *TwoNode) {
	var promoted *TwoNode
	cmp1 := strings.Compare(key, n.Key1)
	if cmp1 < 0 {
		if n.Left == nil {
			// TODO: `n` return value is just ignored; any way to do this more elegantly?
			return n, &TwoNode{
				Key:   n.Key1,
				Value: n.Value1,
				Left:  New(key, value),
				Right: New(n.Key2, n.Value2),
			}
		} else {
			n.Left, promoted = n.Left.insert(key, value)
			if promoted != nil {
				return n, &TwoNode{
					Key:   n.Key1,
					Value: n.Value1,
					Left:  promoted,
					Right: TwoNode{
						Key:   n.Key2,
						Value: n.Value2,
						Left:  n.Middle,
						Right: n.Right,
					},
				}
			} else {
				return n, nil
			}
		}
	} else if cmp1 == 0 {
		n.Value1 = value
		return n, nil
	}

	cmp2 := strings.Compare(key, n.Key2)
	if cmp2 < 0 {
		if n.Middle == nil {
			// TODO: `n` return value is just ignored; any way to do this more elegantly?
			return n, &TwoNode{
				Key:   key,
				Value: value,
				Left:  New(n.Key1, n.Value1),
				Right: New(n.Key2, n.Value2),
			}
		} else {
			n.Middle, promoted = n.Middle.insert(key, value)
			if promoted != nil {
				return n, &TwoNode{
					Key:   promoted.Key,
					Value: promoted.Value,
					Left: TwoNode{
						Key:   n.Key1,
						Value: n.Value1,
						Left:  n.Left,
						Right: promoted.Left,
					},
					Right: TwoNode{
						Key:   n.Key2,
						Value: n.Value2,
						Left:  promoted.Right,
						Right: n.Right,
					},
				}
			} else {
				return n, nil
			}
		}
	} else if cmp2 == 0 {
		n.Value2 = value
		return n, nil
	} else {
		if n.Right == nil {
			// TODO: `n` return value is just ignored; any way to do this more elegantly?
			return n, &TwoNode{
				Key:   n.Key2,
				Value: n.Value2,
				Left:  New(n.Key1, n.Value1),
				Right: New(key, value),
			}
		} else {
			n.Right, promoted = n.Right.insert(key, value)
			if promoted != nil {
				return n, &TwoNode{
					Key:   n.Key2,
					Value: n.Value2,
					Left: TwoNode{
						Key:   n.Key1,
						Value: n.Value1,
						Left:  n.Left,
						Right: n.Middle,
					},
					Right: promoted,
				}
			} else {
				return n, nil
			}
		}
	}
}

func (n ThreeNode) Insert(key string, value string) Node {
	r, promoted := n.insert(key, value)
	if promoted != nil {
		return promoted
	} else {
		return r
	}
}

func (n TwoNode) String() string {
	if n.Left == nil {
		return fmt.Sprintf("(%s)", n.Key)
	} else {
		return fmt.Sprintf("(%s/2 %s %s)", n.Key, n.Left.String(), n.Right.String())
	}
}

func (n ThreeNode) String() string {
	if n.Left == nil {
		return fmt.Sprintf("(%s,%s)", n.Key1, n.Key2)
	} else {
		return fmt.Sprintf("(%s,%s/3 %s %s %s)", n.Key1, n.Key2, n.Left.String(), n.Middle.String(), n.Right.String())
	}
}
