package twothree

import (
	"testing"
)

func TestInsert(t *testing.T) {
	root := New("9", "")
	root = root.Insert("5", "")
	root = root.Insert("8", "")
	root = root.Insert("3", "")
	root = root.Insert("2", "")
	root = root.Insert("4", "")
	root = root.Insert("7", "")

	actual := root.String()
	expected := "(5/2 (3/2 (2) (4)) (8/2 (7) (9)))"
	if actual != expected {
		t.Errorf("tree mismatch\n  actual:    %s\n  expected:  %s", actual, expected)
	}
}
