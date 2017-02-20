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

func getMid(lb int, ub int) int {
	return (lb + ((ub - lb) / 2)) + 1
}

// search insertion point
func findInsertionPoint(key int, lb int, ub int, dataset []datum) int {
	mid := getMid(lb, ub)

	// while we haven't found the insertion point
	if !(mid <= 1 || key == dataset[mid].val) {
		if key < dataset[mid].val {
			return findInsertionPoint(key, lb, mid, dataset)
		} else if key > dataset[mid].val {
			return findInsertionPoint(key, mid+1, ub, dataset)
		}
	}

	return mid
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
