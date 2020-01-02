package bst

import (
	"fmt"

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
func (t *binarySearchTree) Search(key int) (*string, int, error) {
	node, err := search(t.root, key)
	if err != nil {
		return nil, 0, err
	}
	if node == nil {
		return nil, 0, nil
	}
	return &node.value, 0, nil
}

func search(node *binarySearchTreeNode, key int) (*binarySearchTreeNode, error) {
	if key < node.key && node.left != nil {
		return search(node.left, key)
	} else if key > node.key && node.right != nil {
		return search(node.right, key)
	}
	if key == node.key {
		return node, nil
	}
	return nil, nil
}

// Insert is a method to insert a node by binary search.
func (t *binarySearchTree) Insert(key int, value string) error {
	if t.root == nil {
		t.root = newBinarySearchTreeNode(key, value)
		return nil
	}
	err := insert(t.root, key, value)
	if err != nil {
		return err
	}
	return nil
}

func insert(node *binarySearchTreeNode, key int, value string) error {
	if key < node.key {
		if node.left == nil {
			node.left = newBinarySearchTreeNode(key, value)
			return nil
		}
		if err := insert(node.left, key, value); err != nil {
			return err
		}
	} else {
		if node.right == nil {
			node.right = newBinarySearchTreeNode(key, value)
			return nil
		}
		if err := insert(node.right, key, value); err != nil {
			return err
		}
	}
	return nil
}

func (t *binarySearchTree) InsertAll(data map[int]string) error {
	for k, v := range data {
		err := t.Insert(k, v)
		if err != nil {
			return err
		}
	}
	return nil
}

// Delete is a method to delete a node by binary search.
func (t *binarySearchTree) Delete(key int) error {
	panic("not implemented yet")
}

func (t *binarySearchTree) PrintTree() {
	if t.root == nil {
		fmt.Println("no node exists.")
		return
	}
	print(t.root, 1)
}

func print(node *binarySearchTreeNode, depth int) {
	fmt.Printf("%d key=%d value=%s\n", depth, node.key, node.value)
	if node.left != nil {
		print(node.left, depth+1)
	}
	if node.right != nil {
		print(node.right, depth+1)
	}
}
