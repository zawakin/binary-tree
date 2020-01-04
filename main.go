package main

import (
	"math/rand"
	"os"

	"github.com/zawawahoge/binary-tree/bst"
	"github.com/zawawahoge/binary-tree/core"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func createTree() core.IndexTree {
	binarySearchTree := bst.NewBinarySearchTree()

	N := 1000

	for i := 0; i < N; i++ {
		k := rand.Intn(10 * N)
		v := randomHashString(4)
		binarySearchTree.Insert(k, v)
	}
	return binarySearchTree
}

func main() {
	graphWrapper := core.NewGraphWrapper("G")

	graphWrapper.MustAddAttr("layout", "dot")
	graphWrapper.MustAddAttr("nodesep", "0.2")
	graphWrapper.MustAddAttr("ranksep", "0.8")

	// Node設定
	nodeAttrs := graphWrapper.NodeAttrs
	// nodeAttrs["colorscheme"] = "rdylgn11"
	// nodeAttrs["style"] = "\"solid,filled\""
	// nodeAttrs["fontcolor"] = "6"
	// nodeAttrs["fontname"] = "\"Migu 1M\""
	nodeAttrs["fontname"] = "\"Arial\""
	// nodeAttrs["color"] = "7"
	// nodeAttrs["fillcolor"] = "11"
	nodeAttrs["shape"] = "point"

	// Edge設定
	edgeAttrs := graphWrapper.EdgeAttrs
	edgeAttrs["color"] = "black"

	tree := createTree()
	tree.PrintTree(graphWrapper)

	// dotファイル出力
	s := graphWrapper.G.String()
	if _, err := os.Stat("output"); os.IsNotExist(err) {
		os.Mkdir("output", 0777)
	}
	file, err := os.Create(`output/graph.dot`)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	file.Write([]byte(s))

}

func randomHashString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
