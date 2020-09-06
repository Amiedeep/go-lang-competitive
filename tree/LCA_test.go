package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLCA(t *testing.T) {

	assert := assert.New(t)
	node4 := BinaryTreeNode{4, nil, nil}
	node5 := BinaryTreeNode{5, nil, nil}
	node6 := BinaryTreeNode{6, nil, nil}
	node7 := BinaryTreeNode{7, nil, nil}

	node2 := BinaryTreeNode{2, &node4, &node5}
	node3 := BinaryTreeNode{3, &node6, &node7}

	root := BinaryTreeNode{1, &node2, &node3}

	foundNode := LCA(&root, 4, 5)
	assert.Equal(&node2, foundNode)

	foundNode = LCA(&root, 4, 6)
	assert.Equal(&root, foundNode)

	foundNode = LCA(&root, 3, 4)
	assert.Equal(&root, foundNode)

	foundNode = LCA(&root, 2, 4)
	assert.Equal(&node2, foundNode)
}
