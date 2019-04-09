//@author Devansh Gupta
//This is the the heap implementation in golang it support max and min both
package dsal

//NewHeap, returns new build heap
//max => is heap is a max heap or min heap
//ar=> slice for building heap
func NewHeap(ar []int, max bool) *Heap {
	x := new(Heap)
	x.ar = ar
	x.length = len(ar)
	x.ismax = max
	x.buildheap() //Build the heap
	return x
}

type Heap struct {
	ar     []int
	ismax  bool
	length int
}

//swap, this swaps the element at the given indexs
func (h *Heap) swap(x, y int) {
	if x != y {
		h.ar[x] = h.ar[x] ^ h.ar[y]
		h.ar[y] = h.ar[x] ^ h.ar[y]
		h.ar[x] = h.ar[x] ^ h.ar[y]
	}
}

//heapify, heapify an index in Heap according to it's heap property
func (h *Heap) heapify(i int) {

	if i < h.l() && i > -1 && i < h.l()>>1 { //runs only when index is smaller than length of heap array and non negative and less than middle index
		var t int = i
		if h.comp(h.ar[h.lt(i)], h.ar[t]) {
			t = h.lt(i)
		}
		if h.comp(h.ar[h.rt(i)], h.ar[t]) {
			t = h.rt(i)
		}
		//now t have max or min of all
		h.swap(i, t)
		if i != t {
			h.heapify(t) //Heapify the new one
		}
	}
}

//buildheap, builds the heap by calling maximum heapify function on heap array
func (h *Heap) buildheap() {
	var m int = h.l() >> 1 //get the index of the array
	for ; m >= 0; m-- {
		h.heapify(m)
	}
}
func (h *Heap) GetHeap() []int {
	return h.ar
}

//l, returns size of heap
func (h Heap) l() int {
	return h.length
}

func (h Heap) pt(i int) int {
	return i >> 1 //i/2
}
func (h Heap) lt(i int) int {
	return i << 1 //2i
}
func (h Heap) rt(i int) int {
	return 1 + i<<1 //2i+1
}

//comp, Function to compare things depeds on type of heap either max or min
func (h Heap) comp(x, y int) bool {
	if h.ismax {
		return x >= y
	}
	return x <= y
}

//sort, this sorts the heap with heapsort algorithm
//if heap is max heap than it will be in increasing order else decreasing
func (h Heap) sort() {
	for h.length != 0 {
		h.swap(0, h.length-1)
		h.length-- //Decreasing heap length
		h.buildheap()
	}
}
