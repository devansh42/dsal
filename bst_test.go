//@author Devansh Gupta
//Test for BST
package dsal

import "testing"

func TestBST(t *testing.T) {
	b := NewBST()
	ar := []int{13, 14, 6, 5, 7, 2, 12}
	for _, v := range ar {
		b.Insert(v)
	}
	prin(b.root, t)
	t.Log("Min", b.Minimum())
	t.Log("Max", b.Maximum())
	t.Log(5, b.Search(5))
	t.Log(89, b.Search(89))
	t.Log(5, b.Search(5))
	//b.Remove(2) //case 1
	//	b.Remove(5) //case 1
	//b.Remove(8) //case 1
	//b.Remove(7) //case 2
	//	b.Remove(6) //case 3a
	//b.Remove()
	prin(b.root, t)

}

func TestBSTRemove(t *testing.T) {
	b := NewBST()
	ar := []int{8, 10, 3, 1, 6, 7, 4, 14, 13}
	for _, v := range ar {
		b.Insert(v)
	}
	prin(b.root, t)
	b.Remove(3)
	//t.Log(b.findMin(b.searchnode(b.root, 3).rightc).value)
	prin(b.root, t)
}

func prin(node *BSTNode, t *testing.T) {
	if node != nil {
		prin(node.leftc, t)
		t.Log(node.value)
		prin(node.rightc, t)
	}
}
