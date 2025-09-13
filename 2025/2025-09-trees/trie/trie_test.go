package trie

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	trie := Node{}
	trie.Insert("hello")
	trie.Insert("goodbye")

	assertIn(t, trie, "hello")
	assertIn(t, trie, "goodbye")

	assertNotIn(t, trie, "good")
	assertNotIn(t, trie, "hell")
	assertNotIn(t, trie, "helloween")

	trie.Insert("good")
	assertIn(t, trie, "good")

	trie.Insert("hen")
	assertIn(t, trie, "hen")
	assertNotIn(t, trie, "he")

	assertStrEq(t, fmt.Sprintf("%+v", trie.AllMatches("he")), "[hello hen]")
}

func assertStrEq(t *testing.T, actual string, expected string) {
	t.Helper()
	if actual != expected {
		t.Errorf("tree mismatch\n  actual:    %s\n  expected:  %s", actual, expected)
	}
}

func assertIn(t *testing.T, trie Node, s string) {
	t.Helper()
	if !trie.Retrieve(s) {
		t.Errorf("\"%s\" expected to be in trie", s)
	}
}

func assertNotIn(t *testing.T, trie Node, s string) {
	t.Helper()
	if trie.Retrieve(s) {
		t.Errorf("\"%s\" expected _not_ to be in trie", s)
	}
}
