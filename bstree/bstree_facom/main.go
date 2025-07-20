package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	bst := NewBSTree()
	bst.Insert(4).Insert(2).Insert(1).Insert(10).Insert(8).Insert(7).Insert(9)
	bst.PrintInorder()
	fmt.Printf("\nAltura da árvore: %d\n", bst.Height())

	fmt.Printf("Buscar 7: %+v\n", bst.Search(7))
	fmt.Printf("Contém 7: %t\n", bst.Contains(7))
	fmt.Printf("Contém 99: %t\n", bst.Contains(99))

	fmt.Print("Removendo 4. Árvore: ")
	bst.Remove(4)
	bst.PrintInorder()
	fmt.Printf("\nAltura após remoção: %d\n", bst.Height())

	fmt.Println("\n=== Benchmarks ===")
	fmt.Println("Benchmark Inserção 100 elementos:")
	benchmarkInsercao(100)

	fmt.Println("Benchmark Inserção 1K elementos:")
	benchmarkInsercao(1000)

	fmt.Println("Benchmark Inserção 10K elementos:")
	benchmarkInsercao(10000)

	fmt.Println("Benchmark Busca 1K elementos:")
	benchmarkBusca(1000)

	fmt.Println("Benchmark Remoção 1K elementos:")
	benchmarkRemocao(1000)

	fmt.Println("\nTodos os testes concluídos!")
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
	PrintInorder()
	Search(d int) *Node
	Contains(d int) bool
	Remove(d int) *bst
	Height() int
}

func NewBSTree() IBSTree {
	return &bst{}
}

func (t *bst) Height() int {
	if t.root == nil {
		return -1
	}
	return t.root.heightRec()
}

func (n *Node) heightRec() int {
	if n == nil {
		return -1
	}
	leftHeight := n.left.heightRec()
	rightHeight := n.right.heightRec()
	if leftHeight > rightHeight {
		return leftHeight + 1
	}
	return rightHeight + 1
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
	} else if d > n.data {
		if n.right == nil {
			n.right = &Node{d, nil, nil}
		} else {
			n.right.insertRec(d)
		}
	}
}

func (t *bst) PrintInorder() {
	t.printInorderRec(t.root)
}

func (t *bst) printInorderRec(n *Node) {
	if n == nil {
		return
	}
	t.printInorderRec(n.left)
	fmt.Printf("%d ", n.data)
	t.printInorderRec(n.right)
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

func (t *bst) Contains(d int) bool {
	return t.Search(d) != nil
}

func (t *bst) Remove(d int) *bst {
	if t.root == nil {
		return t
	}
	t.root = t.removeRec(t.root, d)
	return t
}

func (t *bst) removeRec(n *Node, d int) *Node {
	if n == nil {
		return nil
	}

	if d < n.data {
		n.left = t.removeRec(n.left, d)
	} else if d > n.data {
		n.right = t.removeRec(n.right, d)
	} else {
		if n.left == nil {
			return n.right
		} else if n.right == nil {
			return n.left
		}

		successor := t.findMin(n.right)
		n.data = successor.data
		n.right = t.removeRec(n.right, successor.data)
	}

	return n
}

func (t *bst) findMin(n *Node) *Node {
	for n.left != nil {
		n = n.left
	}
	return n
}

func generateRandomSlice(size int) []int {
	rand.Seed(time.Now().UnixNano())
	slice := make([]int, size)
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(size * 10)
	}
	return slice
}

func benchmarkInsercao(size int) {
	data := generateRandomSlice(size)
	start := time.Now()

	bst := NewBSTree()
	for _, v := range data {
		bst.Insert(v)
	}

	duration := time.Since(start)
	fmt.Printf("  Inserção de %d elementos: %v\n", size, duration)
}

func benchmarkBusca(size int) {
	bst := NewBSTree()
	data := generateRandomSlice(size)
	for _, v := range data {
		bst.Insert(v)
	}

	start := time.Now()
	for i := 0; i < 1000; i++ {
		bst.Search(data[i%len(data)])
	}
	duration := time.Since(start)
	fmt.Printf("  1000 buscas em árvore de %d elementos: %v\n", size, duration)
}

func benchmarkRemocao(size int) {
	data := generateRandomSlice(size)
	bst := NewBSTree()
	for _, v := range data {
		bst.Insert(v)
	}

	start := time.Now()
	for _, v := range data {
		bst.Remove(v)
	}
	duration := time.Since(start)
	fmt.Printf("  Remoção de %d elementos: %v\n", size, duration)
}
