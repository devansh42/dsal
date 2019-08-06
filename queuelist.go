//@author Devansh Gupta
package dsal

import "errors"

//QueueList, is the queue implemented by single link list
type QueueList struct {
	slist *SingleList
}

func NewQueueList() *QueueList {
	x := new(QueueList)
	x.slist = NewSingleList()
	return x
}

func (s QueueList) Length() int {
	return s.slist.Length()
}

func (s QueueList) Enqueue(v interface{}) {
	s.slist.Push(v)
}

func (s QueueList) Dequeue() (interface{}, error) {
	if s.Length() > 0 {
		v := s.slist.Pop()
		return v, nil
	}
	return nil, errors.New("Queue underflow")
}
