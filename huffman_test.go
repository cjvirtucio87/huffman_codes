package huffman

import (
	"testing"
)

var data = []datum{
	{key: rune('a'), val: 3},
	{key: rune('b'), val: 5},
	{key: rune('c'), val: 7},
	{key: rune('d'), val: 5},
	{key: rune('e'), val: 2},
}

func TestRemoveFirstTwo(t *testing.T) {
	var dataset []datum

	// remove the first two twice
	firstTwo, dataset := removeFirstTwo(data)
	_, dataset = removeFirstTwo(dataset)

	if firstTwo[0].key != rune('a') {
		t.Error("Expected ", rune('a'), ", got ", firstTwo[0].key)
	} else if firstTwo[1].key != rune('b') {
		t.Error("Expected ", rune('b'), ", got ", firstTwo[1].key)
	} else if len(dataset) != 1 {
		// must not mutate array pointed to by data
		t.Error("Expected ", 1, ", got ", len(dataset))
	}
}

// func TestFindNode(t *testing.T) {
// 	root := bst.CreateTree(data)

// 	result := findNode(rune('b'), &root)

// 	if result == nil {
// 		t.Error("Expected ", rune('b'), ", got ", result)
// 	}
// }
