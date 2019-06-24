package dsal

//Bit Array Implementation in Go

type BitArray struct {
	ar   []uint8
	size int64
	cap  int64
}

//NewBitArray, returns bit array of given length
func NewBitArray(size int64) *BitArray {
	//Making an array  of 8 bit integers
	x := BitArray{make([]uint8, 1+(size/8)), 0, size}

	return &x
}

//Size, returns size of a bit array
func (b BitArray) Size() int64 {
	return b.size
}

//Append, appends a bit at the last of the bit array
func (b *BitArray) Append(bit bool) {

	if bit {

		var i = b.size / 8
		x := b.ar[i]           //Accessing corrosponding bit block
		nn := 7 - (b.size % 8) //actually this is 8 - b.size%8 -1
		var n = uint8(nn)
		x |= 1 << n
		b.ar[i] = x //Setting the bit
	}
	b.size++
}

//IsSet, returns true if bit at an index is set to 1 or not
func (b BitArray) IsSet(i int64) bool {
	return get(b, i)
}

//get, retrives bit at a given position
func get(b BitArray, i int64) bool {

	var x = b.ar[i/8]
	nn := 7 - (i % 8) //actually this is 8 - b.size%8 -1
	var n = uint8(nn)
	var f uint8 = 1 << n
	var c uint8 = f
	c = c & x

	return c == f //observing changes in c

}

const (
	and  = 1
	or   = 2
	xor  = 3
	not  = 4
	flip = 5
)

//operates different operations on bitarray
func (a BitArray) operate(b BitArray, o int) *BitArray {
	var size int64
	if a.size < b.size {
		size = a.size
	} else {
		size = b.size
	}
	x := NewBitArray(size)
	var i int64
	for i = 0; i < 1+size/8; i++ {

		k := a.ar[i]

		l := b.ar[i]
		switch o {
		case and:
			k &= l
		case or:
			k |= l
		case xor:
			k ^= l
		}
		x.ar[i] = k
	}
	return x
}

//setunset, set or unset bit a particular bit
func (a *BitArray) setunset(i int64, set bool) {
	if i >= a.cap { //array out of bound
		return
	}
	x := a.ar[i/8]
	var nn = uint8(7 - i%8)
	var f uint8 = 1 << nn
	var n = f
	if set {
		x |= n //setting bit to one
	} else {
		n &= x      //Knowing if bit is 1 or 0
		if n == f { //bit was 1
			x ^= n //inverting bit setting to 0
		}
	}
	a.ar[i/8] = x //persisting changes
	//changing size of array
	if i >= a.size {
		a.size = i + 1
	}
}

//Set, sets bit at a given index
func (a *BitArray) Set(i int64) {
	a.setunset(i, true)
}

//Unset, unsets bit at a given index
func (a *BitArray) Unset(i int64) {
	a.setunset(i, false)
}

//OR, performs OR over the bitarrays
func (a BitArray) OR(b BitArray) *BitArray {
	return a.operate(b, or)
}

//AND, performs AND over the bitarrays
func (a BitArray) AND(b BitArray) *BitArray {
	return a.operate(b, and)
}

//XOR, performs XOR over the bitarrays
func (a BitArray) XOR(b BitArray) *BitArray {
	return a.operate(b, xor)
}
