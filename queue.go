//@author Devansh Gupta
//BasicQueue is array queue implementation
package dsal

import (
	"errors"
)

type BasicQueue struct {
	i, j  int
	array []interface{}
}

func NewBasicQueue(size int) *BasicQueue {
	b := new(BasicQueue)
	b.array = make([]interface{}, size)
	return b
}

//Enqueue, insert something in Queue
func (b *BasicQueue) Enqueue(v interface{}) error {

	if b.Full() {
		return errors.New("Queue overflow")
	}
	b.array[b.j] = v
	b.j++
	if b.Full() {
		b.j = 0
	}

	return nil
}

//Dequeue, deletes first element from queue
func (b *BasicQueue) Dequeue() (interface{}, error) {
	if b.Empty() {
		return nil, errors.New("Queue underflow")
	}
	v := b.array[b.i]
	b.i++
	if b.Full() {
		b.i = 0
	}
	return v, nil
}

//Length, Returns length of the queue
func (b BasicQueue) Length() int {

	return len(b.array)
}
func (b BasicQueue) Empty() bool {
	return b.i == b.j
}
func (b BasicQueue) Full() bool {
	return len(b.array) == abs(b.i-b.j)
}
func abs(x int) int {
	if x < 0 {
		x *= -1
	}
	return x
}
