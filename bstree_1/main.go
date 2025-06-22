package main

import "fmt"

//		   6
//	      / \
//	     2   10
//	    / \  /
//	   1   4 8
//	      /
//	     3
func main() {
	t := BSTree{}
	t.Insert(6)
	t.Insert(2)
	t.Insert(1)
	t.Insert(10)
	t.Insert(4)
	t.Insert(3)
	t.Insert(8)

	t.Inorder(t.Root)
	// fmt.Println()
	// fmt.Printf("Tree depth is: %d\n", t.Depth(3))
	// fmt.Printf("Total height of tree is: %d\n", t.Height())
}

type Node struct {
	data  int
	left  *Node
	right *Node
}

type BSTree struct {
	Root *Node
}

func (t *BSTree) Insert(val int) {
	t.InsertRec(t.Root, val)
}

func (t *BSTree) InsertRec(node *Node, val int) *Node {
	if t.Root == nil {
		t.Root = &Node{val, nil, nil}
	}

	if node == nil {
		return &Node{val, nil, nil}
	}

	if val <= node.data {
		node.left = t.InsertRec(node.left, val)
	}

	if val > node.data {
		node.right = t.InsertRec(node.right, val)
	}

	return node
}

// vai printar todos em ordem
func (t *BSTree) Inorder(node *Node) {
	if node == nil {
		return
	}

	// primeiro visita todos os nos do lado esquerdo ->
	t.Inorder(node.left)
	fmt.Println(node.data, " ")
	t.Inorder(node.right)
}

func (t *BSTree) Depth(val int) int {
	if t == nil || t.Root == nil {
		return -1
	}

	return t.DepthRec(t.Root, val, 0)
}

func (t *BSTree) DepthRec(node *Node, val int, currentDepth int) int {
	if node == nil {
		return -1
	}

	fmt.Printf("Current node: %+v, val %d, depth: %d\n", node, val, currentDepth)

	if node.data == val {
		return currentDepth
	}

	if val < node.data {
		return t.DepthRec(node.left, val, currentDepth+1)
	}

	return t.DepthRec(node.right, val, currentDepth+1)
}

func (t *BSTree) Height() int {
	if t == nil || t.Root == nil {
		return -1
	}

	return t.HeightRec(t.Root)
}

func (t *BSTree) HeightRec(node *Node) int {
	if node == nil {
		return -1
	}

	leftHeight := t.HeightRec(node.left)
	fmt.Printf("  left height of node %d is: %d\n", node.data, leftHeight)

	rightHeight := t.HeightRec(node.right)
	fmt.Printf("  right height of node %d is: %d\n", node.data, rightHeight)

	if leftHeight > rightHeight {
		return leftHeight + 1
	}

	return rightHeight + 1
}
