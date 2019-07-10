//@author Devansh Gupta
//SingleList is the singly linked list implementation
package dsal

func NewSingleList() *SingleList {
	b := new(SingleList)
	b.null = new(ListSNode)
	b.tail = b.null
	b.null.next = b.null
	//sentinel node is initially header and tail also for the linked list
	return b
}

type SingleList struct {
	null, tail *ListSNode
	pot        *ListSNode //Previous of Tail
	length     int
}

func (b *SingleList) Pop() interface{} {
	if b.length > 0 {
		v := b.tail.key
		b.pot.next = b.null //Previous of tail to sentinel node
		b.tail = nil        //Deleting tail node
		b.tail = b.pot      //Making previous tail new tail
		b.length--
		return v
	}
	return nil
}

//PrePend, insert a node at the head of list
func (s *SingleList) PrePend(v interface{}) {
	n := new(ListSNode)
	n.key = v
	n.next = s.null.next
	s.null.next = n
	s.length++
}

//Push, Pushes a new node at the end of List
func (b *SingleList) Push(v interface{}) {
	n := new(ListSNode)
	n.key = v
	n.next = b.null
	b.tail.next = n
	b.pot = b.tail //Point to prevoius tail
	b.tail = n
	b.length++
}

func (b SingleList) Length() int {
	return b.length
}

func (b SingleList) find(key interface{}) *ListSNode {
	var x *ListSNode
	for x = b.null.next; x.key != key && x != b.null; x = x.next {

	}

	return x
}

//Find, returns node, that contain given key, returns nil if not found
func (b SingleList) Find(key interface{}) ListSNode {
	var n *ListSNode
	nx := b.find(key)
	if nx == b.null {
		return *n
	}
	return *nx
}

//Head, returns Head of the List as a value for immutiblity reasons
func (b SingleList) Head() ListSNode {
	return *b.null.next
}

//Tail, returns Tail of the List as a value for immutibilty reasons
func (b SingleList) Tail() ListSNode {
	return *b.tail
}

//At, returns key value at given index, returns nil if index out of bound
func (b SingleList) at(i int) *ListSNode {
	if i < 0 {
		return nil
	}
	if i < b.length {
		var x *ListSNode
		j := 0
		for x = b.null.next; j < i; x = x.next {
			j++
		}
		return x
	}
	return nil
}
func (b SingleList) At(i int) interface{} {
	x := b.at(i)
	if x != nil {
		return x.key
	}
	return nil
}

//Sentinel, returns the sentinel node used in list implementation
func (b SingleList) Sentinel() *ListSNode {
	return b.null
}

//DecLength, decreases length of linked list by 'v'
func (b *SingleList) DecLength(v int) {
	b.length -= v
}

//IncLength, increases length of linked list by 'v'
func (b *SingleList) IncLength(v int) {
	b.length += v
}
