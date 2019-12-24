package dsal

//ListSNode, node for single linked list
type ListSNode struct {
	next *ListSNode
	key  interface{}
}

//Next, returns next node
func (l ListSNode) Next() *ListSNode {
	return l.next
}

//Key, returns key for the node
func (l ListSNode) Key() interface{} {
	return l.key
}

func (l *ListSNode) SetNext(n *ListSNode) {
	l.next = n
}

func (l *ListSNode) SetKey(k interface{}) {
	l.key = k
}

//ListNode, node for doubly linked list
type ListNode struct {
	next, prev *ListNode
	key        interface{}
}

func (l *ListNode) SetNext(n *ListNode) {
	l.next = n
}
func (l *ListNode) SetPrev(n *ListNode) {
	l.prev = n
}
func (l *ListNode) SetKey(k interface{}) {
	l.key = k
}

func (l ListNode) Next() *ListNode {
	return l.next
}
func (l ListNode) Prev() *ListNode {
	return l.prev
}

func (l ListNode) Key() interface{} {
	return l.key
}

//Node, is any body which have a key inside of it
type Node interface {
	Key() interface{}
}
