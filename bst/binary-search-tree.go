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
	root    *binarySearchTreeNode
	history treeHistory
}

type treeHistory struct {
	lastOperatedNode []*binarySearchTreeNode
}

func newTreeHistory() *treeHistory {
	nodes := make([]*binarySearchTreeNode, 1000)
	return &treeHistory{
		lastOperatedNode: nodes,
	}
}

func (hist *treeHistory) append(node *binarySearchTreeNode) {
	fmt.Println(node)
	hist.lastOperatedNode = append(hist.lastOperatedNode, node)
}

func (t *binarySearchTree) lastAddedNode() *binarySearchTreeNode {
	return t.history.lastOperatedNode[len(t.history.lastOperatedNode)-1]
}

// NewBinarySearchTree is a constructor of binary search tree.
func NewBinarySearchTree() core.IndexTree {
	return &binarySearchTree{
		root:    nil,
		history: treeHistory{},
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
		t.history.append(t.root)
		return nil
	}
	node, err := insert(t.root, key, value)
	if err != nil {
		return err
	}
	t.history.append(node)
	return nil
}

func insert(node *binarySearchTreeNode, key int, value string) (*binarySearchTreeNode, error) {
	var err error
	var lastAddedNode *binarySearchTreeNode
	if key < node.key {
		if node.left == nil {
			node.left = newBinarySearchTreeNode(key, value)
			return node.left, nil
		}
		lastAddedNode, err = insert(node.left, key, value)
		if err != nil {
			return nil, err
		}
	} else {
		if node.right == nil {
			node.right = newBinarySearchTreeNode(key, value)
			return node.right, nil
		}
		lastAddedNode, err = insert(node.right, key, value)
		if err != nil {
			return nil, err
		}
	}
	return lastAddedNode, nil
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
	if t.root == nil {
		panic("no node exists")
	}
	t.print(t.root, 1, gw)
}

func (t *binarySearchTree) print(node *binarySearchTreeNode, depth int, gw *core.GraphWrapper) (int, int) {
	var hl, hr int
	if node.left != nil {
		src := node.value
		dst := node.left.value
		gw.MustAddEdge(src, dst, true)
		l, r := t.print(node.left, depth+1, gw)
		if l < r {
			l = r
		}
		hl = l + 1
	} else {
		hl = 0
	}
	if node.right != nil {
		src := node.value
		dst := node.right.value
		gw.MustAddEdge(src, dst, true)
		l, r := t.print(node.right, depth+1, gw)
		if l > r {
			r = l
		}
		hr = r + 1
	} else {
		hr = 0
	}

	coef := (hl - hr)
	if coef < 0 {
		coef = -coef
	}
	coef++
	if coef > 9 {
		coef = 9
	}

	attrs := map[string]string{
		"label":     fmt.Sprintf("\"%d\n%d\"", node.key, hl-hr),
		"fillcolor": fmt.Sprintf("%d", coef),
	}
	if node == t.lastAddedNode() {
		fmt.Println("fill")
		attrs["fillcolor"] = "black"
		attrs["fixedsize"] = "true"
		attrs["width"] = "1.1"
		attrs["height"] = "1.1"
	}
	gw.MustAddNode(node.value, attrs)
	return hl, hr
}
