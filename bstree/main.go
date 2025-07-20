package main

import (
	"fmt"
)

func main() {
	bst := NewBSTree()
	bst.Insert(4).Insert(2).Insert(1).Insert(10).Insert(8).Insert(7).Insert(9)
	bst.PrintInorder(bst.GetRoot())

	println("height is", bst.Height())

	parent, child := bst.GetNodeAndParent(1)
	fmt.Printf("child is: %+v\n", child)
	fmt.Printf("parent is: %+v\n", parent)

	min := bst.GetMin(bst.GetRoot())
	fmt.Printf("min node: %+v", min)

	bst.PrintInorder(bst.GetRoot())
	println()
	bst.Delete(4)
	bst.PrintInorder(bst.GetRoot())
}

type bst struct {
	root *Node
}

type Node struct {
	data  int
	left  *Node
	right *Node
}

type IBSTree interface {
	Insert(d int) *bst
	GetRoot() *Node
	GetMin(n *Node) *Node
	GetNode(d int) *Node
	GetNodeAndParent(d int) (parent, node *Node)
	PrintInorder(n *Node)
	Depth(n *Node) int
	CountLeaves() int
	Height() int
	Delete(d int) *Node
	Contains(d int) bool
}

func NewBSTree() IBSTree {
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

func (t *bst) GetNodeAndParent(d int) (parent, node *Node) {
	if t.root == nil {
		return nil, nil
	}

	return t.root.getNodeAndParentRec(t.root, d)
}

func (nd *Node) getNodeAndParentRec(root *Node, d int) (prt, node *Node) {
	if root == nil {
		return nil, nil
	}
	// achou o nó desejado
	if root.data == d {
		return nil, root
	}

	var parent, targetNode *Node

	if d < root.data {
		parent, targetNode = nd.getNodeAndParentRec(root.left, d)
	} else {
		parent, targetNode = nd.getNodeAndParentRec(root.right, d)
	}

	// ja passou pelo base case e achou o nó desejado
	// entao o nó atual (ao voltar a recursão) será o pai (que por enquanto é nulo)
	if parent == nil && targetNode != nil {
		return root, targetNode
	}

	return parent, targetNode
}

func (t *bst) GetMin(n *Node) *Node {
	if t.root == nil || n.left == nil {
		return n
	}

	return t.GetMin(n.left)
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

func (t *bst) Height() int {
	if t.root == nil || t == nil {
		return 0
	}

	return t.root.HeightRec(t.root)
}

func (n *Node) HeightRec(nd *Node) int {
	if nd == nil {
		return 0
	}

	leftSum := n.HeightRec(nd.left)
	rightSum := n.HeightRec(nd.right)

	if rightSum > leftSum {
		return rightSum + 1
	}

	return leftSum + 1
}

func (t *bst) Delete(d int) *Node {
	if t == nil || t.root == nil {
		return nil
	}

	parent, node := t.GetNodeAndParent(d)
	if node == nil {
		return t.root
	}

	// caso 1: 0 filhos
	if node.left == nil && node.right == nil {
		if parent == nil {
			t.root = nil
			return node
		}

		if parent.left == node {
			parent.left = nil
		} else {
			parent.right = nil
		}

		return node
	}

	// caso 2: 1 filho
	if node.left != nil && node.right == nil {
		if parent == nil {
			t.root = node.left
			return node
		}
		if parent.left == node {
			parent.left = node.left
		} else {
			parent.right = node.left
		}
		return node
	}

	if node.left == nil && node.right != nil {
		if parent == nil {
			t.root = node.right
			return node
		}
		if parent.left == node {
			parent.left = node.right
		} else {
			parent.right = node.right
		}
		return node
	}

	// caso 3: 2 filhos
	// sucessor -> elemento da arvore da direita, com o menor valor
	successor := t.GetMin(node.right)
	// a subarvore da esquerda vai se lidar ao novo pai...
	successor.left = node.left
	// a subarvore da direita sobe pra ocupar o lugar do nó deletado
	if parent == nil {
		t.root = node.right
		return node.right
	}

	if parent.left == node {
		parent.left = node.right // conecta o pai na subarvore da direita
	} else {
		parent.right = node.right
	}

	return node
}

func (t *bst) Contains(d int) bool {
	return t.Search(d) != nil
}

func (t *bst) Search(d int) *Node {
	if t.root == nil {
		return nil
	}
	return t.root.searchRec(d)
}

func (n *Node) searchRec(d int) *Node {
	if n == nil {
		return nil
	}
	if n.data == d {
		return n
	}
	if d < n.data {
		return n.left.searchRec(d)
	}
	return n.right.searchRec(d)
}
