package main

import (
	"math/rand"
	"os"

	"github.com/awalterschulze/gographviz"
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
	graphWrapper := &core.GraphWrapper{
		G:         gographviz.NewGraph(),
		NodeAttrs: make(map[string]string),
		EdgeAttrs: make(map[string]string),
	}
	g := graphWrapper.G
	if err := g.SetName("G"); err != nil {
		panic(err)
	}
	// 有向グラフか
	if err := g.SetDir(true); err != nil {
		panic(err)
	}

	// グラフ全体の設定
	// if err := g.AddAttr("G", "bgcolor", "\"#343434\""); err != nil {
	// 	panic(err)
	// }
	if err := g.AddAttr("G", "layout", "dot"); err != nil {
		panic(err)
	}
	if err := g.AddAttr("G", "nodesep", "0.2"); err != nil {
		panic(err)
	}
	if err := g.AddAttr("G", "ranksep", "0.8"); err != nil {
		panic(err)
	}

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
	s := g.String()
	file, err := os.Create(`./binary-search-tree.dot`)
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
