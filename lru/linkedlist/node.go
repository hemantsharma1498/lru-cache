package dll

type node struct {
	key  int
	prev *node
	next *node
}

func newNode(key int) *node {
	return &node{key: key, prev: nil, next: nil}
}
