package tree

// BinaryTreeNode creates a node of Binary Tree
type BinaryTreeNode struct {
	Data  int
	Left  *BinaryTreeNode
	Right *BinaryTreeNode
}

// LCA find lowest common ancestor of two nodes in a binary tree.
func LCA(node *BinaryTreeNode, data1 int, data2 int) *BinaryTreeNode {

	if node == nil {
		return nil
	}

	if node.Data == data1 || node.Data == data2 {
		return node
	}

	leftNode := LCA(node.Left, data1, data2)
	rightNode := LCA(node.Right, data1, data2)

	if leftNode != nil && rightNode != nil {
		return node
	}

	if leftNode != nil {
		return leftNode
	}

	return rightNode
}
