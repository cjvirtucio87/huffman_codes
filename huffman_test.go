package huffman

import (
	"fmt"
	"testing"
)

var data = []datum{
	{key: rune('a'), val: 3},
	{key: rune('b'), val: 5},
	{key: rune('c'), val: 7},
	{key: rune('d'), val: 12},
	{key: rune('e'), val: 30},
}

var evenData = append(data, datum{key: rune('f'), val: 4})

func TestRemoveFirstTwo(t *testing.T) {
	fmt.Println("huffman.removeFirstTwo should grab the first two elements in the dataset and return a new dataset without them.")
	var dataset []datum

	// remove the first two twice
	firstTwo, dataset := removeFirstTwo(data)
	_, dataset = removeFirstTwo(dataset)

	if firstTwo[0].key != rune('a') {
		t.Error("Expected ", rune('a'), ", got ", firstTwo[0].key)
	} else if firstTwo[1].key != rune('b') {
		t.Error("Expected ", rune('b'), ", got ", firstTwo[1].key)
	} else if len(dataset) != 1 {
		t.Error("Expected ", 1, ", got ", len(dataset))
	}
}

// helper for TestGetMid
func testMids(dataset []datum, t *testing.T) {
	// with original dataset
	fullMid := getMid(0, len(dataset))
	leftMid := getMid(0, getMid(0, len(dataset)))
	rightMid := getMid(getMid(0, len(dataset)), len(dataset))

	if len(dataset)%2 != 0 {
		if fullMid != 3 {
			t.Error("Expected ", 3, ", got ", fullMid)
		} else if leftMid != 2 {
			t.Error("Expected ", 2, ", got ", leftMid)
		} else if rightMid != 5 {
			t.Error("Expected ", 5, ", got ", rightMid)
		}
	} else {
		if fullMid != 4 {
			t.Error("Expected ", 4, ", got ", fullMid)
		} else if leftMid != 3 {
			t.Error("Expected ", 3, ", got ", leftMid)
		} else if rightMid != 6 {
			t.Error("Expected ", 6, ", got ", rightMid)
		}
	}
}

func TestGetMid(t *testing.T) {
	fmt.Println("huffman.getMid should calculate the mid index in the range of data.")

	testMids(data, t)
	testMids(evenData, t)
}

func testFind(dataset []datum, t *testing.T) {
	var seven, fourteen, two int
	length := len(dataset)

	if len(dataset)%2 == 0 {
	} else {
		seven = findInsertionPoint(7, 0, length, dataset)
		fourteen = findInsertionPoint(14, 0, length, dataset)
		two = findInsertionPoint(2, 0, length, dataset)

		if seven != 2 {
			t.Error("Expected ", 2, ", got ", seven)
		} else if fourteen != 5 {
			t.Error("Expected ", 5, ", got ", fourteen)
		} else if two != 0 {
			t.Error("Expected ", 0, ", got ", two)
		}
	}
}

func TestFindInsertionPoint(t *testing.T) {
	fmt.Println("huffman.findInsertionPoint should return the proper index for inserting the key.")

	testFind(data, t)
}

// func TestFindNode(t *testing.T) {
// 	root := bst.CreateTree(data)

// 	result := findNode(rune('b'), &root)

// 	if result == nil {
// 		t.Error("Expected ", rune('b'), ", got ", result)
// 	}
// }
