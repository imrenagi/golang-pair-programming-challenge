package tree

// NewNode creates new tree node
func NewNode(val int, left, right *Node) *Node {
	return &Node{
		Val:   val,
		Left:  left,
		Right: right,
	}
}

// Node stores tree structure like data
type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

// IsBinarySearchTree ...
func IsBinarySearchTree(node *Node) bool {
	return false
}
