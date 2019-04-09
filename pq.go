//@author Devansh Gupta
//This is the Priority queue implementation in golang
//It Supports both Max and Min Priority Queue
package dsal

import "errors"

//PriorityQueue, it is a type of Heap
type PriorityQueue struct {
	*Heap
}

//NewPriorityQueue, returns new priority key instance
func NewPriorityQueue(ar []int, max bool) *PriorityQueue {
	x := new(PriorityQueue)
	x.Heap = NewHeap(ar, max)
	return x
}

//IncreaseKey, increases key of a max pq
func (p *PriorityQueue) IncreaseKey(i, k int) error {
	if i < p.length && i >= 0 {
		if p.ismax {
			if k > p.ar[i] {
				p.changekey(i, k)

			}
			return nil
		}
		return errors.New("This method is not allowed")

	}
	return nil
}

//DecreaseKey, decreases key of a min pq
func (p *PriorityQueue) DecreaseKey(i, k int) error {
	if i < p.length && i >= 0 {
		if !p.ismax {
			if k < p.ar[i] {
				p.changekey(i, k)
			}
			return nil
		}
		return errors.New("This method is not allowed")
	}
	return nil
}

//changekey, changes key of an index
func (p *PriorityQueue) changekey(i, k int) {

	if p.ar[i] != k && i < p.length && i >= 0 {
		p.ar[i] = k //Setting the key
		for i = p.pt(i); i != 0; i = p.pt(i) {
			p.heapify(i)
		}
		p.heapify(i) //Heapify Root
	}

}

//InsertKey, inserts a new key into the priority
func (p *PriorityQueue) InsertKey(k int) {
	p.ar = append(p.ar, -1) //Appending new key
	p.length++              //Increasing pq key
	p.changekey(p.length-1, k)
}

//ExtractExtreme, extracts the extreme from the priority queue
func (p *PriorityQueue) ExtractExtreme() (int, error) {
	if p.length < 0 {
		return 0, errors.New("Priority Queue underflow")
	}
	key := p.ar[0]
	p.swap(0, p.length-1)
	p.length--
	p.heapify(0)
	return key, nil
}

//GetExtreme, returns extreme key from the Priority Queue, incase of Max-Pq it is maximum else minimum
func (p PriorityQueue) GetExtreme() int {
	if p.length < 0 {
		panic("Priority Queue underflow")
	}
	return p.ar[0]

}
