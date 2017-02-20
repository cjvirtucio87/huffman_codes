package huffman

import "github.com/cjvirtucio87/bst"

// remove the first two elements of the dataset
func removeFirstTwo(dataset []bst.Datum) ([]bst.Datum, []bst.Datum) {
	if len(dataset) > 2 {
		output := dataset[0:2]
		return output, dataset[2:]
	}

	return dataset, dataset
}

// search for bst.Node with matching key
func search(searchKey rune, root *bst.Node) *bst.Node {
	// preorder traversal
	if root == nil || searchKey == root.Key {
		return root
	}

	// traverse left
	left := search(searchKey, root.Left)
	if left != nil && searchKey == left.Key {
		return left
	}

	// traverse right
	right := search(searchKey, root.Right)
	if right != nil && searchKey == right.Key {
		return right
	}

	return nil
}

// // check if the bst.Datum already has a bst.Node
// func checkVector(d bst.Datum) bool {
// 	// need to traverse huffman tree
// }
