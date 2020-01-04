package balancedbst

import (
	"fmt"

	"github.com/zawawahoge/binary-tree/core"
)

type balancedBSTNode struct {
	key   int
	value string
	left  *balancedBSTNode
	right *balancedBSTNode
}

type balancedBST struct {
	core.IndexTree
	root *balancedBSTNode
}

// NewBalancedBST is a constructor of binary search tree.
func NewBalancedBST() core.IndexTree {
	return &balancedBST{
		root: nil,
	}
}

func newBalancedBSTNode(key int, value string) *balancedBSTNode {
	return &balancedBSTNode{
		key:   key,
		value: value,
		left:  nil,
		right: nil,
	}
}

// Search is a method to search node with key by binary search.
func (t *balancedBST) Search(key int) (*string, int, error) {
	node, dep, err := search(t.root, key, 1)
	if err != nil {
		return nil, 0, err
	}
	if node == nil {
		return nil, 0, nil
	}
	return &node.value, dep, nil
}

func search(node *balancedBSTNode, key int, dep int) (*balancedBSTNode, int, error) {
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
func (t *balancedBST) Insert(key int, value string) error {
	if t.root == nil {
		t.root = newBalancedBSTNode(key, value)
		return nil
	}
	err := insert(t.root, key, value)
	if err != nil {
		return err
	}
	return nil
}

func insert(node *balancedBSTNode, key int, value string) error {
	if key < node.key {
		if node.left == nil {
			node.left = newBalancedBSTNode(key, value)
			return nil
		}
		if err := insert(node.left, key, value); err != nil {
			return err
		}
	} else {
		if node.right == nil {
			node.right = newBalancedBSTNode(key, value)
			return nil
		}
		if err := insert(node.right, key, value); err != nil {
			return err
		}
	}
	return nil
}

func (t *balancedBST) InsertAll(data map[int]string) error {
	for k, v := range data {
		err := t.Insert(k, v)
		if err != nil {
			return err
		}
	}
	return nil
}

// Delete is a method to delete a node by binary search.
func (t *balancedBST) Delete(key int) error {
	panic("not implemented yet")
}

func (t *balancedBST) PrintTree(gw *core.GraphWrapper) {
	depthTree := make(map[int]int)
	if t.root == nil {
		panic("no node exists")
	}
	print(t.root, 1, depthTree, gw)
	fmt.Println(depthTree)
}

func print(node *balancedBSTNode, depth int, depthTree map[int]int, gw *core.GraphWrapper) {
	depthTree[depth]++
	gw.MustAddNode(node.value, nil)
	fmt.Println(node.value)

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
