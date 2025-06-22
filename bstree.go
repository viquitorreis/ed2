package main

import "fmt"

func main() {
	bst := NewBSTree()
	bst.Insert(4).Insert(2).Insert(1).Insert(10)
	bst.PrintInorder(bst.GetRoot())
}

type BSTree interface {
	Insert(d int) *bst
	PrintInorder(n *Node)
	GetRoot() *Node
}

type bst struct {
	root *Node
}

type Node struct {
	data  int
	left  *Node
	right *Node
}

func NewBSTree() BSTree {
	return &bst{}
}

func (t *bst) GetRoot() *Node {
	return t.root
}

func (t *bst) Insert(d int) *bst {
	println("inserting data:", d)
	if t.root == nil {
		t.root = &Node{d, nil, nil}
	} else {
		t.root.insertRec(d)
	}

	return t
}

func (n *Node) insertRec(d int) {
	if n == nil {
		return
	}

	if d < n.data {
		if n.left == nil {
			n.left = &Node{d, nil, nil}
		} else {
			n.left.insertRec(d)
		}
	} else {
		if n.right == nil {
			n.right = &Node{d, nil, nil}
		} else {
			n.left.insertRec(d)
		}
	}
}

func (t *bst) PrintInorder(n *Node) {
	if n == nil {
		return
	}

	t.PrintInorder(n.left)
	fmt.Printf("%d ", n.data)
	t.PrintInorder(n.right)
}
