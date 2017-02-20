package huffman

type node struct {
	key    rune
	val    int
	parent *node
	left   *node
	right  *node
}

type datum struct {
	// key is the char code for the string
	key rune
	val int
}

// remove the first two elements of the dataset
func removeFirstTwo(dataset []datum) ([]datum, []datum) {
	output := dataset[0:2]
	return output, dataset
}

// search for node with matching key
func search(searchKey rune, root *node) *node {
	// base case
	if root == nil {
		return root
	}

	// postorder traversal
	if searchKey == root.key {
		return root
	}

	// traverse left
	left := search(searchKey, root.left)
	if left != nil && searchKey == left.key {
		return left
	}

	// traverse right
	right := search(searchKey, root.right)
	if right != nil && searchKey == right.key {
		return right
	}

	return nil
}

// // check if the datum already has a node
// func checkVector(d datum) bool {
// 	// need to traverse huffman tree
// }
