//@author Devansh Gupta
//This file contains code for Disjoint Set Implementation in Golang
//This implements tree representation of disjoint sets
package dsal

type DisjointSet struct {
	parent *DisjointSet //Parent Element of the current Set
	value  interface{}  //Value of the item
	rank   int          //This decides the rank of current set
}

//NewDisjointSet, returns a pointer to the set containing given values
func NewDisjointSet(values ...interface{}) *DisjointSet {
	set := new(DisjointSet)
	set.parent = set
	set.rank = 0
	if len(values) > 1 {
		set.value = values[0] //Creation of new set
		for _, v := range values[1:] {
			seti := new(DisjointSet)
			seti.value = v
			seti.rank = 0
			seti.parent = seti
			findparent(set).Union(findparent(seti))
		}
	}
	return set
}

//linksets, links two sets , first is finds the set
func linksets(a, b *DisjointSet) {
	findparent(a).Union(findparent(b))
}

//findparent, This Find's Parent of a disjoint set
func findparent(item *DisjointSet) *DisjointSet {
	if item != item.parent {
		item.parent = findparent(item.parent)

	}
	return item
}

//Union, unions two sets
func (d *DisjointSet) Union(item *DisjointSet) {

	if d.rank > item.rank {
		item.parent = d
	} else { // d have <= rank of that of item
		d.parent = item
		if d.rank == item.rank {
			d.rank++ //Increasing rank of set arbitrary
		}
	}
}

func (d *DisjointSet) Add(item interface{}) {
	x := new(DisjointSet)
	x.parent = x
	x.rank = 0
	x.value = item
	linksets(x, d)
}

//8661 9353 1169
