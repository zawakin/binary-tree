package main

import (
	"math/rand"

	"github.com/zawawahoge/binary-tree/bst"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func main() {
	binarySearchTree := bst.NewBinarySearchTree()

	N := 1000

	for i := 0; i < N; i++ {
		k := rand.Intn(N)
		v := randomHashString(20)
		binarySearchTree.Insert(k, v)
	}
	// for k, v := range data {
	// 	binarySearchTree.Insert(k, v)
	// }
	binarySearchTree.PrintTree()
}

func randomHashString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
