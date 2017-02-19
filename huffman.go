package huffman

type node struct {
	key    string
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

// create node with children config
func createNode(k string, v int, c map[string]*node) node {
	return node{key: k, val: v, left: c["left"], right: c["right"]}
}

// remove the first two elements of the dataset
func removeFirstTwo(dataset []datum) ([]datum, []datum) {
	output := dataset[0:2]
	return output, dataset
}
