package main

type binarySearchTreeNode struct {
	key   int
	value string
	left  *binarySearchTreeNode
	right *binarySearchTreeNode
}

type binarySearchTree struct {
	IndexTree
	root *binarySearchTreeNode
}

// NewBinarySearchTree is a constructor of binary search tree.
func NewBinarySearchTree() IndexTree {
	return &binarySearchTree{
		root: nil,
	}
}

// Search is a method to search node with key by binary search.
func (t *binarySearchTree) Search(key int) (*string, int) {
	cnt := 0
	pointer := t.root
	for pointer != nil {
		cnt++
		if key == pointer.key {
			return &pointer.value, cnt
		}
		if key < pointer.key && pointer.left != nil {
			pointer = pointer.left
		} else if key > pointer.key && pointer.right != nil {
			pointer = pointer.right
		}
	}
	return nil, cnt
}

func newBinarySearchTreeNode(key int, value string) *binarySearchTreeNode {
	return &binarySearchTreeNode{
		key:   key,
		value: value,
		left:  nil,
		right: nil,
	}
}

// Insert is a method to insert a node by binary search.
func (t *binarySearchTree) Insert(key int, value string) {
	if t.root == nil {
		t.root = newBinarySearchTreeNode(key, value)
		return
	}
	pointer := t.root
	for {
		if key <= pointer.key {
			if pointer.left == nil {
				pointer.left = newBinarySearchTreeNode(key, value)
				return
			}
			pointer = pointer.left
		} else {
			if pointer == nil {
				pointer.right = newBinarySearchTreeNode(key, value)
			}
			pointer = pointer.right
		}
	}
}

// Delete is a method to delete a node by binary search.
func (t *binarySearchTree) Delete(key int) {
	panic("not implemented yet")
}
