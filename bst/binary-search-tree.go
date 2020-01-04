package bst

import (
	"fmt"
	"strconv"

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
	node, dep, err := search(t.root, key, 1)
	if err != nil {
		return nil, 0, err
	}
	if node == nil {
		return nil, 0, nil
	}
	return &node.value, dep, nil
}

func search(node *binarySearchTreeNode, key int, dep int) (*binarySearchTreeNode, int, error) {
	if key < node.key && node.left != nil {
		return search(node.left, key, dep+1)
	} else if key > node.key && node.right != nil {
		return search(node.right, key, dep+1)
	}
	if key == node.key {
		return node, dep, nil
	}
	return nil, dep, nil
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

func (t *binarySearchTree) PrintTree(gw *core.GraphWrapper) {
	depthTree := make(map[int]int)
	if t.root == nil {
		panic("no node exists")
	}
	print(t.root, 1, depthTree, gw)
	fmt.Println(depthTree)
}

func print(node *binarySearchTreeNode, depth int, depthTree map[int]int, gw *core.GraphWrapper) {
	depthTree[depth]++
	// gw.MustAddNode(fmt.Sprintf("'key=%d value=%s'", node.key, node.value))
	attrs := map[string]string{
		"label": strconv.Itoa(node.key),
	}
	gw.MustAddNode(node.value, attrs)
	fmt.Println(node.value)

	// fmt.Printf("%d key=%d value=%s\n", depth, node.key, node.value)
	if node.left != nil {
		src := node.value
		dst := node.left.value
		gw.MustAddEdge(src, dst, true)
		print(node.left, depth+1, depthTree, gw)
	}
	if node.right != nil {
		src := node.value
		dst := node.right.value
		gw.MustAddEdge(src, dst, true)
		print(node.right, depth+1, depthTree, gw)
	}
}
