package huffman

type datum struct {
	key rune
	val int
}

type node struct {
	key    rune
	val    int
	parent *node
	left   *node
	right  *node
}

// remove the first two elements of the dataset
func removeFirstTwo(dataset []datum) ([]datum, []datum) {
	if len(dataset) > 2 {
		output := dataset[0:2]
		return output, dataset[2:]
	}

	return dataset, dataset
}

// search insertion point
func findInsertionPoint(key int) {

}

// insert new parent's k/v pair
// func insertSum(sum) {

// }

// search for node with matching key
func findNode(searchKey rune, root *node) *node {
	// preorder traversal
	if root == nil || searchKey == root.key {
		return root
	}

	// traverse left
	left := findNode(searchKey, root.left)
	if left != nil && searchKey == left.key {
		return left
	}

	// traverse right
	right := findNode(searchKey, root.right)
	if right != nil && searchKey == right.key {
		return right
	}

	return nil
}

// // check if the datum already has a node
// func checkVector(d datum) bool {
// 	// need to traverse huffman tree
// }
