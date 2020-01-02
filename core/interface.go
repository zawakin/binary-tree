package core

// IndexTree is interface of index tree.
type IndexTree interface {
	Search(Key int) (*string, int)
	Insert(Key int, Value string)
	Delete(Key int)
}
