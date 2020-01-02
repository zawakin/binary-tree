package bst

import "testing"

func Test_bst_Search(t *testing.T) {
	binarySearchTree := NewBinarySearchTree()
	data := map[int]string{
		3: "C",
		1: "A",
		2: "B",
	}
	err := binarySearchTree.InsertAll(data)
	if err != nil {
		t.Fatalf("InsertAll(%#v)=%v; want nil", data, err)
	}
}
