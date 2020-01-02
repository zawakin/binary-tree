package core

// IndexTree is interface of index tree.
type IndexTree interface {
	Search(Key int) (*string, int, error)
	Insert(Key int, Value string) error
	InsertAll(data map[int]string) error
	Delete(Key int) error
	PrintTree()
}
