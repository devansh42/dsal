package dsal

//ForwardIterator, is read only iterator for forward traversal of linear data structure
type ForwardIterator interface {
	Next() (interface{}, error)
	HasNext() bool
}

//BackwardIterator, is read only  iterator for backward traversal of linear data structure
type BackwardIterator interface {
	Prev() (interface{}, error)
	HasPrev() bool
}

//Iterator, is read only iterator for bi directional traversal of linear data structure
type Iterator interface {
	ForwardIterator
	BackwardIterator
}

type arrayIterableAttrs struct {
	beg, cur, end uint
	array         *[]interface{}
}
