package avl

import "testing"

func TestInsert(t *testing.T) {
	root := New("5", "")
	root = root.Insert("6", "")
	root = root.Insert("8", "")
	root = root.Insert("3", "")
	root = root.Insert("2", "")
	root = root.Insert("4", "")
	root = root.Insert("7", "")
	root.Check()

	actual := root.String()
	expected := "(5:0 (3:0 2:0 4:0) (7:0 6:0 8:0))"
	if actual != expected {
		t.Errorf("tree mismatch\n  actual:    %s\n  expected:  %s", actual, expected)
	}
}
