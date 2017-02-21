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

var evenData = append(data, datum{key: rune('f'), val: 35})

func TestRemoveFirstTwo(t *testing.T) {
	var dataset []datum

	// remove the first two twice
	firstTwo, dataset := removeFirstTwo(data)
	_, dataset = removeFirstTwo(dataset)

	(func(firstTwo []datum) {
		fmt.Println("huffman.removeFirstTwo should grab the first two elements in the dataset")
		if firstTwo[0].key != rune('a') {
			t.Error("Expected ", rune('a'), ", got ", firstTwo[0].key)
		} else if firstTwo[1].key != rune('b') {
			t.Error("Expected ", rune('b'), ", got ", firstTwo[1].key)
		}
	})(firstTwo)

	(func(dataset []datum) {
		fmt.Println("huffman.removeFirstTwo should return the new dataset, less the first two elements.")
		if len(dataset) != 1 {
			t.Error("Expected ", 1, ", got ", len(dataset))
		}
	})(dataset)
}

// helper for TestGetMid
func testMids(dataset []datum, t *testing.T) {
	// with original dataset
	fullMid := getMid(0, len(dataset))
	leftMid := getMid(0, getMid(0, len(dataset)))
	rightMid := getMid(getMid(0, len(dataset)), len(dataset))

	if len(dataset)%2 != 0 {
		if fullMid != 2 {
			t.Error("Expected ", 2, ", got ", fullMid)
		} else if leftMid != 1 {
			t.Error("Expected ", 1, ", got ", leftMid)
		} else if rightMid != 3 {
			t.Error("Expected ", 3, ", got ", rightMid)
		}
	} else {
		if fullMid != 3 {
			t.Error("Expected ", 3, ", got ", fullMid)
		} else if leftMid != 1 {
			t.Error("Expected ", 1, ", got ", leftMid)
		} else if rightMid != 4 {
			t.Error("Expected ", 4, ", got ", rightMid)
		}
	}
}

func TestGetMid(t *testing.T) {
	fmt.Println("huffman.getMid should calculate the mid index in the range of data.")

	testMids(data, t)
	testMids(evenData, t)
}

func testFind(dataset []datum, t *testing.T) {
	length := len(dataset)
	seven := findInsertionPoint(7, 0, length, dataset)
	fourteen := findInsertionPoint(14, 0, length, dataset)
	two := findInsertionPoint(2, 0, length, dataset)
	five := findInsertionPoint(5, 0, length, dataset)

	if seven != 2 {
		t.Error("Expected ", 2, ", got ", seven)
	} else if fourteen != 4 {
		t.Error("Expected ", 4, ", got ", fourteen)
	} else if two != 0 {
		t.Error("Expected ", 0, ", got ", two)
	} else if five != 1 {
		t.Error("Expected ", 1, ", got ", five)
	}
}

func TestFindInsertionPoint(t *testing.T) {
	fmt.Println("huffman.findInsertionPoint should return the proper index for inserting the key.")

	testFind(data, t)
	testFind(evenData, t)
}

func TestInsert(t *testing.T) {
	d := datum{key: rune('h'), val: 6}
	newData := insert(d, 2, data)

	(func(d datum, nd []datum) {
		fmt.Println("huffman.insert should return a new dataset with increased length.")
		diff := len(newData) - len(data)
		if diff != 1 {
			t.Error("Expected ", 1, ", got ", diff)
		}
	})(d, newData)

	(func(d datum, nd []datum) {
		fmt.Println("huffman.insert should place the datum at the specified index.")
		if newData[2].val != d.val {
			t.Error("Expected ", d.val, ", got ", newData[2].val)
		}
	})(d, newData)
}

func TestCreateParent(t *testing.T) {
	fmt.Println("huffman.createParent should create a new node with the sum of the first two values.")

	firstTwo, _ := removeFirstTwo(data)

	firstTwoNodes := (func(firstTwo []datum) []node {
		var acc []node
		for _, v := range firstTwo {
			acc = append(
				acc,
				node{
					key:    v.key,
					val:    v.val,
					parent: nil,
					left:   nil,
					right:  nil,
				},
			)
		}

		return acc
	})(firstTwo)

	parent := createParent(firstTwoNodes)

	if parent.val != 8 {
		t.Error("Expected ", 8, ", got ", parent.val)
	}
}

// func TestFindNode(t *testing.T) {
// 	root := bst.CreateTree(data)

// 	result := findNode(rune('b'), &root)

// 	if result == nil {
// 		t.Error("Expected ", rune('b'), ", got ", result)
// 	}
// }
