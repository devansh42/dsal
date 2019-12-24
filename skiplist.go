package dsal

import (
	"log"
	"math"
	"math/rand"
	"time"
)

//This implements Skip List Data Structure

type float float32

type Comparable interface {
	Score() float
}

//SkipListNode, is the regular node for SkipList
type SkipListNode struct {
	//Key, data to be held with in the node
	//This key type must be compareable
	Key Comparable

	//Levels, is a list of level links for the node
	Levels *SingleList
}

//SkipList, is the skiplist
type SkipList struct {
	length int
	levels int
	header *SkipListNode
}

//NewSkipList, returns a new skiplist having comparing function 'f'
func NewSkipList() *SkipList {

	x := new(SkipList)
	h := new(SkipListNode)
	h.Levels = NewSingleList()
	x.header = h
	return x
}

func (s *SkipList) Contains(key Comparable) bool {
	_, _, err := s.findOrCreate(key, false)
	switch err {
	case nil:
		return true
	default:
		return false
	}

}

//Insert, inserts a node in list, returns true if node is inserted else return false (means node exists previously)
func (s *SkipList) Insert(key Comparable) bool {
	_, created, _ := s.findOrCreate(key, true)
	return created
}

//findOrCreate, finds or create a new node in the skip list
func (s *SkipList) findOrCreate(key Comparable, create bool) (*SkipListNode, bool, error) {

	var change *StackList
	if create {

		change = NewStackList() //It remebers the points at which we made a switch
	}
	lookingNode := s.header
	for {
		levels := lookingNode.Levels
		var currentLevel = levels.Head()

		for ; currentLevel != levels.Sentinel(); currentLevel = currentLevel.Next() { //Traversing level one by one
			n := currentLevel.Key().((*SkipListNode))
			if n.Key.Score() > key.Score() { //bigger node
				//Let's low down
				if create { //remembering points we just switch to the lower node
					change.Push(currentLevel.Key())
				}
				continue
			} else if n.Key.Score() <= key.Score() {
				//Let's go to next node
				lookingNode = n
				break
			}

		}
		log.Print(lookingNode)
		if currentLevel == levels.Sentinel() {
			//It means we are at sentinel node
			//looking node is the just small than our key
			//Write place to insert node
			if create {
				nn := new(SkipListNode)
				nn.Key = key
				nlevels := NewSingleList()
				randomLevel := s.getRandomLevel()
				if randomLevel > s.levels {
					s.header.Levels.PrePend(nn)
				}
				for i := 0; i < randomLevel; i++ {
					link, _ := change.Pop()
					ptrLink := link.(*ListSNode)
					nlevels.PrePend(ptrLink.Key()) //setting next node link to new node
					ptrLink.SetKey(nn)             //setting previous node link to point to this new node 'nn'
				}
				return nn, true, nil
			}
			return lookingNode, false, nocontent{"no node found"}
		}
		if lookingNode.Key.Score() == key.Score() { //We have found the desired node
			//Let's return it
			return lookingNode, false, nil
		}

	}

	return new(SkipListNode), false, nil

}

//Returns some random level for node
func (s SkipList) getRandomLevel() int {
	t := time.Now()
	rand.Seed(t.UnixNano())
	ln2 := int(math.Log2(float64(1 + s.Length())))

	return 1 + rand.Intn(ln2+1) //+1 is added in case ln2 is 0 which is a common scenario

}

//Length, returns length of the skiplist
func (s SkipList) Length() int {
	return s.length
}

type nocontent struct {
	Err string
}

func (n nocontent) Error() string { return n.Err }

func (n nocontent) nocontent() {

}

type nocontentErr interface {
	Error() string
	nocontent()
}
