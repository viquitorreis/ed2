package main

import "fmt"

func main() {
	avl := NewAVLTree()
	avl.Insert(10).Insert(30).Insert(4)
	avl.PrintInorder()

	fmt.Printf("%+v", avl.Search(4))
}

// avl -> binary search tree balanceada
// Bounce Factor (fator de balanceamento) > Height left - Height right
// Nenhum Bounce Factor pode ser bf > 1 || bf < -1 --> se tiver, deve fazer rotation
// Ou seja, bounce factor deve estar no intervalo [-1,0,1]
// rotation right -> bf do nó é +2 --> subarvore ESQUERDA é mais alta e o bf do filho esquerdo é +1 ou 0
// rotation left -> bf do nó é -2 subarvore DIREITA mais alta e o bf do filho direito é -1 ou 0

type avltree struct {
	root *Node
}

type Node struct {
	data   int
	height int
	left   *Node
	right  *Node
}

type IAVLTree interface {
	Insert(d int) *avltree
	PrintInorder()
	Search(d int) *Node
	Contains(d int) bool
}

func NewAVLTree() IAVLTree {
	return &avltree{}
}

// funcao que atualiza a altura de um nó apos o re-balanceamento.
// Deve ser chamada apenas se todos os filhos estiverem com a altura atualizada
func (n *Node) updateHeight() {
	left, right := -1, -1
	if n.left != nil {
		left = n.left.height
	}

	if n.right != nil {
		right = n.right.height
	}

	if left > right {
		n.height = left + 1
	} else {
		n.height = right + 1
	}
}

func (n *Node) rightRotation() *Node {
	child := n.left
	n.left = child.right
	child.right = n

	n.updateHeight()
	child.updateHeight()

	return child
}

func (n *Node) leftRotation() *Node {
	child := n.right
	n.right = child.left
	child.left = n

	n.updateHeight()
	child.updateHeight()

	return child
}

func (n *Node) balanceFactor() int {
	left, right := -1, -1
	if n.left != nil {
		left = n.left.height
	}
	if n.right != nil {
		right = n.right.height
	}

	return left - right
}

func (n *Node) rebalance() *Node {
	bf := n.balanceFactor()
	if bf > 1 {
		// LR -> left right. Desbalanceamento na subärvore da DIREITA do filho esquerdo
		if n.left != nil && n.left.balanceFactor() < 0 {
			n.left = n.left.leftRotation()
		}
		// ll -> left left. Desbalanceamento na subarvore ESQUERDA do filho esquerdo
		// rotaçao simples ä direita
		n = n.rightRotation()
	} else if bf < -1 {
		// RL -> right left. Desbalanceamento na subarvore ESQUERDA do filho direito
		// rotacao dupla (left do filho + right do nó)
		if n.right != nil && n.right.balanceFactor() > 0 {
			n.right = n.right.rightRotation()
		}
		// RR -> rikght right. desbalanceamento na subarvore da DIREITA do filho direito.
		n = n.leftRotation()
	}

	return n
}

func (t *avltree) Insert(d int) *avltree {
	if t.root == nil {
		t.root = &Node{data: d, height: 0, left: nil, right: nil}
		return t
	}

	t.root.insertRec(t.root, d)

	return t
}

func (n *Node) insertRec(nd *Node, d int) *Node {
	if nd == nil {
		return &Node{data: d, height: 0, left: nil, right: nil}
	}

	if d < nd.data {
		nd.left = n.insertRec(nd.left, d)
	} else if d > nd.data {
		nd.right = n.insertRec(nd.right, d)
	} else {
		return nd
	}

	nd.updateHeight()
	return nd.rebalance()
}

func (t *avltree) PrintInorder() {
	t.root.printInorderRec(t.root)
}

func (t *Node) printInorderRec(n *Node) {
	if n != nil {
		t.printInorderRec(n.left)
		fmt.Printf("%d ", n.data)
		t.printInorderRec(n.right)
	}
}

func (n *Node) search(d int) *Node {
	for current := n; current != nil; {
		if d < current.data {
			current = current.left
		} else if d > current.data {
			current = current.right
		} else {
			return current
		}
	}

	return nil
}

func (t *avltree) Search(d int) *Node {
	if t.root == nil {
		return nil
	}

	return t.root.search(d)
}

func (t *avltree) Height() int {
	if t.root == nil {
		return -1
	}
	return t.root.height
}

func (t *avltree) Contains(d int) bool {
	return t.Search(d) != nil
}

func (t *avltree) findMin(n *Node) *Node {
	for n.left != nil {
		n = n.left
	}
	return n
}

func (t *avltree) Remove(d int) *avltree {
	if t.root == nil {
		return t
	}
	t.root = t.removeRec(t.root, d)
	return t
}

func (t *avltree) removeRec(nd *Node, d int) *Node {
	if nd == nil {
		return nil
	}

	if d < nd.data {
		nd.left = t.removeRec(nd.left, d)
	} else if d > nd.data {
		nd.right = t.removeRec(nd.right, d)
	} else {
		if nd.left == nil {
			return nd.right
		} else if nd.right == nil {
			return nd.left
		}

		// nó com dois filhos sucessor menor da subarvore direita
		successor := t.findMin(nd.right)
		nd.data = successor.data
		nd.right = t.removeRec(nd.right, successor.data)
	}

	nd.updateHeight()
	return nd.rebalance()
}
