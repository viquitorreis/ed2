package main

import "fmt"

func main() {
	bst := NewBSTree()
	bst.Insert(4).Insert(2).Insert(1).Insert(10)
	bst.PrintInorder(bst.GetRoot())

	node := bst.GetNode(1)

	// fmt.Printf("node data: %d. Left: %d, Right: %d \n", node.data, node.left.data, node.right.data)

	depth := bst.Depth(node)

	println("node depth: ", depth)

	println("quantidade de folhas:", bst.CountLeaves())
}

type BSTree interface {
	GetRoot() *Node
	GetNode(d int) *Node
	Insert(d int) *bst
	PrintInorder(n *Node)
	Depth(n *Node) int
	CountLeaves() int
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
			n.right.insertRec(d)
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

func (t *bst) GetNode(d int) *Node {
	if t.root == nil {
		return nil
	} else {
		return t.root.getNodeRec(t.GetRoot(), d)
	}
}

func (nd *Node) getNodeRec(n *Node, d int) *Node {
	if n.data == d {
		return n
	}

	if d < n.data {
		if n.left == nil {
			return n
		} else {
			return n.getNodeRec(n.left, d)
		}
	} else {
		if n.right == nil {
			return n
		} else {
			return n.getNodeRec(n.right, d)
		}
	}
}

func (t *bst) Depth(n *Node) int {
	if t == nil || n == nil {
		return -1
	}

	return t.root.depthRec(t.GetRoot(), n, 0)
}

func (nd *Node) depthRec(current, target *Node, currDepth int) int {
	if current == nil {
		return -1
	}

	if current == target {
		return currDepth
	}

	if target.data < current.data {
		return nd.depthRec(current.left, target, currDepth+1)
	} else {
		return nd.depthRec(current.right, target, currDepth+1)
	}
}

func (t *bst) CountLeaves() int {
	if t.root == nil {
		return 0
	}

	return t.GetRoot().countLeavesRec(t.GetRoot())
}

func (nd *Node) countLeavesRec(n *Node) int {
	if n == nil {
		return 0
	}

	// é folha
	if n.left == nil && n.right == nil {
		return 1
	}

	return nd.countLeavesRec(n.left) + nd.countLeavesRec(n.right)
}
