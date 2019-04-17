//@author Devansh Gupta
//This implements a BST in Golang
//This supports only integer values
//This implementation doesn't support any type of balancing
package dsal

import (
	"log"
)

//BSTNode, is node for bst
type BSTNode struct {
	value  int
	parent *BSTNode
	leftc  *BSTNode
	rightc *BSTNode
}

//BSTree, represents in BST
type BSTree struct {
	root *BSTNode
}

//NewBST, returns a new simple bst pointer without any configration
func NewBST() *BSTree {
	return new(BSTree)
}

func (b BSTree) Root() BSTNode {
	return *b.root
}

//Maximum, returns largest element in the tree
func (b BSTree) Maximum() int {
	return b.findMax(b.root).value
}

//Minimum, returns minimum element in the tree
func (b BSTree) Minimum() int {

	return b.findMin(b.root).value
}

//Search, searchs for a value in tree
func (b BSTree) Search(v int) bool {

	if b.searchnode(b.root, v) != nil {
		return true
	}
	return false
}

//searchnode, internal method to search in bst
func (b BSTree) searchnode(ele *BSTNode, v int) *BSTNode {
	x := ele
	for x != nil {
		if v < x.value {
			x = x.leftc
		} else if v > x.value {
			x = x.rightc
		} else {
			break //Now x is node we r looking fors
		}
	}
	return x

}
func (b BSTree) findMax(n *BSTNode) *BSTNode {
	x := n
	for x.rightc != nil {
		x = x.rightc
	}

	return x
}

func (b BSTree) findMin(n *BSTNode) *BSTNode {
	x := n
	for x.leftc != nil {
		x = x.leftc
	}
	return x
}

func (b *BSTree) Insert(v int) {
	b.insert(b.root, v)
}

//insert, internal method to insert the node
func (b *BSTree) insert(node *BSTNode, v int) {
	x := node
	if x == nil {
		x = new(BSTNode)
		x.value = v
		b.root = x //Making root of tree
		log.Print("Root created")
		return
	}
	var p *BSTNode //Remembers parent of the given node
	for x != nil {
		p = x
		if v < x.value {
			x = x.leftc
		} else if v > x.value {
			x = x.rightc
		} else {
			return //Return if value already exists

		}
	}
	x = new(BSTNode)
	x.value = v
	x.parent = p
	if v < p.value { //Setting the node pointer
		p.leftc = x
	} else {
		p.rightc = x
	}

}

//removenode, is the internal method to remove a node form bst
//We have given right node some priority
func (b *BSTree) removenode(node *BSTNode, v int) {
	if node == nil {
		return
	}
	n := b.searchnode(node, v)
	if n != nil {
		p := n.parent
		if n.leftc == nil {
			//case 1. when node is a leave
			//log.Print("case 1")
			b.transplant(p, n.rightc, v)

		} else if n.rightc == nil {
			//case 2. when node have only one child which is left child ,
			// incase it have right child as its child, then case 1 will also cover that condition
			//log.Print("case 1")

			b.transplant(p, n.leftc, v)

		} else {
			//case 3. when node have both of the child
			//Find successor of the node
			suc := b.findMin(n.rightc)
			suc.parent.leftc = suc.rightc //make suc the new left child of parent of suc
			if suc.rightc != nil {
				suc.rightc.parent = suc.parent //set new parent of right child of suc
			}
			suc.leftc = n.leftc //setting new left child of the node
			if n.leftc != nil {
				n.leftc.parent = suc

			}
			suc.rightc = n.rightc //setting new right child of the node
			if n.rightc != nil {
				n.rightc.parent = suc
			}
			b.transplant(p, suc, v)

		}
		n = nil //Getting node garbage collected
	}
}

//transplant, transplants nodes for deletation process
func (b *BSTree) transplant(p, n *BSTNode, v int) {

	if v < p.value {
		p.leftc = n
	} else {
		p.rightc = n
	}
	if n != nil {
		n.parent = p
	}
}

//Remove, removes given value from tree
func (b BSTree) Remove(v int) {
	b.removenode(b.root, v)
}

//printInorder, prints a binary inorder
func (b BSTree) printInorder(node *BSTNode) {
	if node != nil {
		b.printInorder(node.leftc)
		log.Print(node.value)
		b.printInorder(node.rightc)
	}
}
