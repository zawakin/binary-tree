package main

import (
	"fmt"
	"github.com/zawawahoge/binary-tree/bst"
)

func main() {
	fmt.Println("Index tree")
	binarySearchTree := bst.NewBinarySearchTree()
	data := map[int]string{
		1: "A",
		2: "B",
		3: "C",
	}
	keys := []int{3, 2, 1}
	for _, k := range keys {
		binarySearchTree.Insert(k, data[k])
	}
	fmt.Printf("tree = %#v\n", binarySearchTree)
	for _, k := range keys {
		fmt.Printf("key=%d, want=%s  ", k, data[k])
		got, cnt := binarySearchTree.Search(k)
		if got == nil {
			fmt.Printf("not found; cnt=%d\n", cnt)
		} else {
			fmt.Printf("got=%s, cnt=%d\n", *got, cnt)
		}
	}
	fmt.Println(binarySearchTree)
}
