package main

import (
	"fmt"
	"math/rand"
	"os"

	"github.com/zawawahoge/binary-tree/bst"
	"github.com/zawawahoge/binary-tree/core"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// NumberOfNodes is a number of nodes drawn.
const NumberOfNodes int = 100

func main() {

	if _, err := os.Stat("output"); os.IsNotExist(err) {
		os.Mkdir("output", 0777)
	}

	tree := bst.NewBinarySearchTree()
	// binarySearchTree := balancedbst.NewbalancedBinarySearchTree()

	for i := 0; i < NumberOfNodes; i++ {
		k := rand.Intn(NumberOfNodes * 10)
		v := randomHashString(4)
		tree.Insert(k, v)

		graphWrapper := newDefaultGraphWrapper()
		tree.PrintTree(graphWrapper)

		s := graphWrapper.G.String()
		file, err := os.Create(fmt.Sprintf("output/%04d.dot", i))
		if err != nil {
			panic(err)
		}
		defer file.Close()
		file.Write([]byte(s))
	}

}

func newDefaultGraphWrapper() *core.GraphWrapper {
	graphWrapper := core.NewGraphWrapper("G")

	graphWrapper.MustAddAttr("layout", "dot")
	graphWrapper.MustAddAttr("nodesep", "0.2")
	graphWrapper.MustAddAttr("ranksep", "0.8")

	// Node設定
	nodeAttrs := graphWrapper.NodeAttrs
	nodeAttrs["colorscheme"] = "ylgn9"
	nodeAttrs["style"] = "\"solid,filled\""
	// nodeAttrs["fontcolor"] = "6"
	// nodeAttrs["fontname"] = "\"Migu 1M\""
	// nodeAttrs["fontname"] = "\"Arial\""
	// nodeAttrs["color"] = "7"
	// nodeAttrs["fillcolor"] = "11"
	// nodeAttrs["shape"] = "point"

	// Edge設定
	edgeAttrs := graphWrapper.EdgeAttrs
	edgeAttrs["color"] = "black"

	return graphWrapper
}

func randomHashString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
