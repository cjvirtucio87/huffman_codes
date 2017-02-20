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
	var dataset []bst.Datum

	// remove the first two twice
	firstTwo, dataset := removeFirstTwo(data)
	_, dataset = removeFirstTwo(dataset)

	if firstTwo[0].Key != rune('a') {
		t.Error("Expected ", rune('a'), ", got ", firstTwo[0].Key)
	} else if firstTwo[1].Key != rune('b') {
		t.Error("Expected ", rune('b'), ", got ", firstTwo[1].Key)
	} else if len(dataset) != 1 {
		// must not mutate array pointed to by data
		t.Error("Expected ", 1, ", got ", len(dataset))
	}
}

func TestSearch(t *testing.T) {
	root := bst.CreateTree(data)

	result := search(rune('b'), &root)

	if result == nil {
		t.Error("Expected ", rune('b'), ", got ", result)
	}
}
