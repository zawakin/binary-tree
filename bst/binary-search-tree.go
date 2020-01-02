package bst

import (
	"github.com/zawawahoge/binary-tree/core"
)

type binarySearchTreeNode struct {
	key   int
	value string
	left  *binarySearchTreeNode
	right *binarySearchTreeNode
}

type binarySearchTree struct {
	core.IndexTree
	root *binarySearchTreeNode
}

// NewBinarySearchTree is a constructor of binary search tree.
func NewBinarySearchTree() core.IndexTree {
	return &binarySearchTree{
		root: nil,
	}
}

func newBinarySearchTreeNode(key int, value string) *binarySearchTreeNode {
	return &binarySearchTreeNode{
		key:   key,
		value: value,
		left:  nil,
		right: nil,
	}
}

// Search is a method to search node with key by binary search.
func (t *binarySearchTree) Search(key int) (*string, int) {
	node := search(t.root, key)
	if node == nil {
		return nil, 0
	}
	return &node.value, 0
}

func search(node *binarySearchTreeNode, key int) *binarySearchTreeNode {
	if key < node.key && node.left != nil {
		return search(node.left, key)
	} else if key > node.key && node.right != nil {
		return search(node.right, key)
	}
	if key == node.key {
		return node
	}
	return nil
}

// Insert is a method to insert a node by binary search.
func (t *binarySearchTree) Insert(key int, value string) {
	if t.root == nil {
		t.root = newBinarySearchTreeNode(key, value)
		return
	}
	insert(t.root, key, value)
}

func insert(node *binarySearchTreeNode, key int, value string) {
	if key < node.key {
		if node.left == nil {
			node.left = newBinarySearchTreeNode(key, value)
			return
		}
		insert(node.left, key, value)
	} else {
		if node.right == nil {
			node.right = newBinarySearchTreeNode(key, value)
			return
		}
		insert(node.right, key, value)
	}
}

// Delete is a method to delete a node by binary search.
func (t *binarySearchTree) Delete(key int) {
	panic("not implemented yet")
}
