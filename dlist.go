//@author Devansh gupta
//This is a implementaion of a doubly linked with a dummy sentinel to simplify things
package dsal

func NewBasicList() *BasicList {
	b := new(BasicList)
	b.null = new(ListNode)
	b.null.next = b.null
	b.null.prev = b.null
	return b
}

type ListNode struct {
	next, prev *ListNode
	key        interface{}
}

type BasicList struct {
	null   *ListNode
	length int
}

//Push, Pushes a new node at the end of List
func (b *BasicList) Push(v interface{}) {
	n := new(ListNode)
	n.key = v
	n.next = b.null
	n.prev = b.null.prev
	b.null.prev.next = n
	b.null.prev = n
	b.length++
}

func (b BasicList) Length() int {
	return b.length
}

//Remove, Removes a node from BasicList, returns true if operation was success and returns false if not
func (b *BasicList) Remove(key interface{}) {
	n := b.find(key)
	if n != b.null {
		n.prev.next = n.next
		n.next.prev = n.prev
		n = nil
		b.length--
	}
}
func (b BasicList) find(key interface{}) *ListNode {
	var x *ListNode
	for x = b.null.next; x.key != key && x != b.null; x = x.next {

	}

	return x
}

//Find, returns node, that contain given key, returns nil if not found
func (b BasicList) Find(key interface{}) ListNode {
	var n *ListNode
	nx := b.find(key)
	if nx == b.null {
		return *n
	}
	return *nx
}

//Head, returns Head of the List as a value for immutiblity reasons
func (b BasicList) Head() ListNode {
	return *b.null.next
}

//Tail, returns Tail of the List as a value for immutibilty reasons
func (b BasicList) Tail() ListNode {
	return *b.null.prev
}

//At, returns key value at given index, returns nil if index out of bound
func (b BasicList) at(i int) *ListNode {
	if i < 0 {
		return nil
	}
	if i < b.length {
		var x *ListNode
		j := 0
		for x = b.null.next; j < i; x = x.next {
			j++
		}
		return x
	}
	return nil
}
func (b BasicList) At(i int) interface{} {
	x := b.at(i)
	if x != nil {
		return x.key
	}
	return nil
}

//Append, appends a new node after a given node, if node doesn't found on place, it pushes node at the end
func (b *BasicList) Append(i int, key interface{}) {
	b.insert(i, key, true)
}

//Prepend, prepend a new node before a given node, if node doesn't found at place it will prepend node before list head
func (b *BasicList) Prepend(i int, key interface{}) {
	b.insert(i, key, false)
}

//Insert, inserts a node at a given index, if index is out of bound and then it push the node at the end
func (b *BasicList) insert(i int, key interface{}, after bool) {
	x := b.at(i)
	if x != nil {
		n := new(ListNode)
		n.key = key
		if after {
			n.next = x.next
			n.prev = x
			x.next.prev = n
			x.next = n
		} else {
			n.next = x
			n.prev = x.prev
			x.prev.next = n
			x.prev = n

		}
		b.length++
		return
	}
	if after {

		b.Push(key) //Pushing the element

	} else {

		n := new(ListNode)
		n.key = key
		n.next = b.null.next
		n.prev = b.null
		b.null.next.prev = n
		b.null.next = n
		b.length++
	}
}

//Sentinel, returns the sentinel node used in list implementation
func (b BasicList) Sentinel() ListNode {
	return *b.null
}
