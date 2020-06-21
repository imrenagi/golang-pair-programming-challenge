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
// Given a binary tree, determine if it is a valid binary search tree (BST).
// Assume a BST is defined as follows:
// * The left subtree of a node contains only nodes with keys less than the node's key.
// * The right subtree of a node contains only nodes with keys greater than the node's key.
// * Both the left and right subtrees must also be binary search trees.
// Example 1:
//  2
// / \
// 1   3
//
// Input: [2,1,3]
// Output: true
//
// Example 2:
//  5
// / \
// 1   4
// 	  / \
//   3   6
// Input: [5,1,4,null,null,3,6]
// Output: false
// Explanation: The root node's value is 5 but its right child's value is 4.
func IsBinarySearchTree(node *Node) bool {
	min := -9999
	max := 9999
	return isBinarySearchTree(node, min, max)
}

func isBinarySearchTree(root *Node, min, max int) bool {
	if root == nil {
		return true
	}

	if root.Val <= min || root.Val >= max {
		return false
	}

	return isBinarySearchTree(root.Left, min, root.Val) && isBinarySearchTree(root.Right, root.Val, max)

}
