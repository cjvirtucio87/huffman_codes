package huffman

import "testing"

func TestRemoveFirstTwo(t *testing.T) {
	data := []datum{
		{key: rune('a'), val: 3},
		{key: rune('b'), val: 5},
		{key: rune('c'), val: 7},
	}

	sliced, orig := removeFirstTwo(data)

	if sliced[0].key != rune('a') {
		t.Error("Expected ", rune('a'), ", got ", sliced[0].key)
	} else if sliced[1].key != rune('b') {
		t.Error("Expected ", rune('b'), ", got ", sliced[1].key)
	} else if len(orig) != 3 {
		// must not mutate array pointed to by data
		t.Error("Expected ", 3, ", got ", len(orig))
	}
}

func TestSearch(t *testing.T) {
	root := node{
		key:    rune('a'),
		val:    1,
		parent: nil,
		left:   nil,
		right:  nil,
	}

	data := []datum{
		{key: rune('b'), val: 3},
		{key: rune('c'), val: 5},
		{key: rune('d'), val: 7},
	}

	for _, d := range data {
		newNode := node{
			key:    d.key,
			val:    d.val,
			parent: &root,
			left:   nil,
			right:  nil,
		}

		if d.val < root.val {
			root.left = &newNode
		} else if d.val > root.val {
			root.right = &newNode
		}

		root = newNode
	}

	result := search(rune('b'), &root)

	if result == nil {
		t.Error("Expected ", rune('b'), ", got ", result)
	}
}
