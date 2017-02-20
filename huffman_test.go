package huffman

import (
	"testing"

	"github.com/cjvirtucio87/bst"
)

var data = []bst.Datum{
	{Key: rune('a'), Val: 3},
	{Key: rune('b'), Val: 5},
	{Key: rune('c'), Val: 7},
	{Key: rune('d'), Val: 5},
	{Key: rune('e'), Val: 2},
}

func TestRemoveFirstTwo(t *testing.T) {
	sliced, orig := removeFirstTwo(data)

	if sliced[0].Key != rune('a') {
		t.Error("Expected ", rune('a'), ", got ", sliced[0].Key)
	} else if sliced[1].Key != rune('b') {
		t.Error("Expected ", rune('b'), ", got ", sliced[1].Key)
	} else if len(orig) != 5 {
		// must not mutate array pointed to by data
		t.Error("Expected ", 5, ", got ", len(orig))
	}
}

func TestSearch(t *testing.T) {
	root := bst.CreateTree(data)

	result := search(rune('b'), &root)

	if result == nil {
		t.Error("Expected ", rune('b'), ", got ", result)
	}
}
