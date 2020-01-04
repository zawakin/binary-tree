package core

import "github.com/awalterschulze/gographviz"

// IndexTree is interface of index tree.
type IndexTree interface {
	Search(Key int) (*string, int, error)
	Insert(Key int, Value string) error
	InsertAll(data map[int]string) error
	Delete(Key int) error
	PrintTree(gw *GraphWrapper)
}

// GraphWrapper is a wrapper struct of gographviz.Graph.
type GraphWrapper struct {
	G         *gographviz.Graph
	NodeAttrs map[string]string
	EdgeAttrs map[string]string
}

// MustAddNode is a method to add node with no error.
func (gw *GraphWrapper) MustAddNode(name string) {
	if err := gw.G.AddNode(gw.G.Name, name, gw.NodeAttrs); err != nil {
		panic(err)
	}
}

// MustAddEdge is a method to add edge with no error.
func (gw *GraphWrapper) MustAddEdge(src string, dst string, directed bool) {
	if err := gw.G.AddEdge(src, dst, directed, gw.EdgeAttrs); err != nil {
		panic(err)
	}
}
