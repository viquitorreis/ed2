package main

import (
	"container/heap"
	"fmt"
	"unicode/utf8"
)

// https://medium.com/@anshulkumar1615/compression-tool-in-golang-part1-understanding-huffman-tree-and-its-implementation-6d5cea448831
// https://github.com/minhajuddinkhan/huffman/blob/master/huffman.go
// https://golangprojectstructure.com/understanding-huffman-coding-in-go/

func main() {
	str := "helloworld!"

	tree := buildTree(str)

	printCodes(tree, []byte{})
}

type HuffmanTree interface {
	Freq() int
	LeafCount() int
}

type HuffmanLeaf struct {
	freq int
	val  rune
}

// Nó Intermediário -> nao tem valor e sim a frequencia... É uma subarvore/arvore
type HuffmanNode struct {
	freq        int
	leafCount   int
	left, right HuffmanTree
}

func (l *HuffmanLeaf) Freq() int {
	return l.freq
}

func (l *HuffmanLeaf) LeafCount() int {
	return 1
}

func (n *HuffmanNode) Freq() int {
	return n.freq
}

func (n *HuffmanNode) LeafCount() int {
	return n.leafCount
}

func buildTree(s string) HuffmanTree {
	// 1: Criamos hash table com o caractere e sua frequencia
	freq := make(map[rune]int, utf8.RuneCountInString(s))
	for _, r := range s {
		freq[r] += 1
	}

	// 2: Criar a min Heap
	minHeap := make(HuffmanMinHeap, 0)
	for c, f := range freq {
		minHeap = append(minHeap, &HuffmanLeaf{freq: f, val: c})
	}

	heap.Init(&minHeap)

	// 3: arvore é montada a partir do elemento de menor frequencia, até chegar na root (ultimo node e de maior frequencia)
	for minHeap.Len() > 1 {
		left := heap.Pop(&minHeap).(HuffmanTree)
		right := heap.Pop(&minHeap).(HuffmanTree)

		heap.Push(&minHeap, &HuffmanNode{
			freq:      left.Freq() + right.Freq(),
			leafCount: left.LeafCount() + right.LeafCount(),
			left:      left,
			right:     right,
		})
	}

	// retorna ultimo que sobrou no minHeap (root)
	return heap.Pop(&minHeap).(HuffmanTree)
}

func printCodes(tree HuffmanTree, prefix []byte) {
	switch i := tree.(type) {
	case *HuffmanLeaf:
		fmt.Printf("%c\t%d\t%s\n", i.val, i.freq, string(prefix))
	case *HuffmanNode:
		// traverse filho da esquerda
		prefix = append(prefix, '0')
		printCodes(i.left, prefix)
		prefix = prefix[:len(prefix)-1]

		// traverse filho da direita
		prefix = append(prefix, '1')
		printCodes(i.right, prefix)
		prefix = prefix[:len(prefix)-1]
	}
}
