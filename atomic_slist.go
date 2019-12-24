package dsal

//This implements a thread safe Single Linked list

type SafeSingleList struct {
	null, tail *ListSNode
	pot        *ListSNode //Previous of Tail
	length     int
}

//Length, returns length of the linked list
func (s SafeSingleList) Length() int {
	return s.length
}

//DecLength, decreases length of linked list by 'v'
func (b *SafeSingleList) DecLength(v int) {
	b.length -= v
}

//IncLength, increases length of linked list by 'v'
func (b *SafeSingleList) IncLength(v int) {
	b.length += v
}
