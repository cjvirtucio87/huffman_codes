package huffman

import "testing"

func TestCreateNode(t *testing.T) {
	data := map[string]int{"a": 5, "d": 10, "e": 3}
	var newNode node

	// create slice of nodes
	for k, v := range data {
		newNode = createNode(k, v, map[string]*node{"left": nil, "right": nil})
		if newNode.key != k {
			t.Error("Expected ", k, ", got ", newNode.key)
		} else if newNode.val != v {
			t.Error("Expected ", v, ", got ", newNode.val)
		}
	}
}

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
