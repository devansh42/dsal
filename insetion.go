package dsal

//This code implements insertion sort algorithm in go

//InsertionSort, Sorts arrays using insertion sort algorithm
//This algorithm must be used on small set of number may be having 20 to 40 elements
func InsertionSort(c []int64) {

	var i, j int
	i = 1
	for i = 1; i < len(c); i++ {
		for j = i - 1; j >= 0 && c[j] > c[j+1]; j-- {
			swap64(c, j+1, j)
		}
	}

}

//swap64, swaps array of 64 bit integer
func swap64(c []int64, x, y int) {
	//Swaps element of an array
	v := c[x]
	c[x] = c[y]
	c[y] = v
}
