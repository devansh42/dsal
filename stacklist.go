//@author Devansh Gupta
package dsal

import "errors"

//StackList, is the stack implemented by single linked list
type StackList struct {
	slist *SingleList
}

func (s *StackList) Push(v interface{}) {
	s.slist.PrePend(v)
}
func (s *StackList) Pop() (interface{}, error) {
	v := s.slist.Pop()
	if v == nil {
		return v, errors.New("Stack underflow")
	}
	return v, nil
}
func (s StackList) Length() int {
	return s.slist.Length()
}

//DropStack, drops all the elements of the "src" into the "dest" stack
//And makes "src" empty
func DropStack(dest, src *StackList) {
	for src.Length() > 0 {
		i, _ := src.Pop()
		dest.Push(i)
	}

}

//NewStackList returns pointer to new stack list
func NewStackList() *StackList {
	x := new(StackList)
	x.slist = NewSingleList()
	return x
}
