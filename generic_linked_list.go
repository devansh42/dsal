package dsal

//This implements Generic Linked List

//GenericSingleLinkedList, is the skeleton for linked list
type GenericSingleLinkedList interface {

	//Push, pushes key to the node
	Push(interface{})
	//Pop, pops something from linked list
	Pop() interface{}
	//Length, returns length of the linked list
	Length() int
	//IncLength, increases length of linked list
	IncLength(int)
	//DecLength, decreases length of linked list
	DecLength(int)
	//Null, Returns null node
	Sentinel() *ListSNode
}
