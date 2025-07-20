package main

import (
	"fmt"
	"math"
)

func main() {
	unorderedArr := []int{30, 25, 20, 15, 10, 5}

	heap := NewMaxHeap(unorderedArr, 100)
	heap.Print()

	heap.Insert(21)
	heap.Print()
}

type IMaxHeap interface {
	Insert(d int) *maxHeap
	Push(d int)
	Print()
}

type maxHeap struct {
	arr     []int
	lastIdx int
}

func NewMaxHeap(arr []int, size int) IMaxHeap {
	array := make([]int, 0, size) // length 0, capacidade -> size
	array = append(array, arr...)

	return &maxHeap{
		arr:     array,
		lastIdx: len(arr) - 1,
	}
}

func (h *maxHeap) Insert(d int) *maxHeap {
	if len(h.arr) == 0 {
		h.arr = make([]int, 1)
		h.arr = append(h.arr, d)
		h.lastIdx++
		return h
	}

	h.Push(d)
	h.BubbleUp()
	return h
}

func (h *maxHeap) BubbleUp() {
	last := h.arr[h.lastIdx]
	parent, parentIdx := h.GetParent(h.lastIdx)

	if !h.Less(last, parent) {
		h.swap(h.lastIdx, parentIdx)
		h.BubbleUp()
	} else {
		return
	}
}

func (h *maxHeap) GetParent(childIdx int) (int, int) {
	parentIdx := int(math.Floor(float64(childIdx)/2)) - 1
	return h.arr[parentIdx], parentIdx
}

func (h *maxHeap) Less(i, j int) bool {
	return i < j
}

func (h *maxHeap) swap(i, j int) {
	h.arr[i], h.arr[j] = h.arr[j], h.arr[i]
}

func (h *maxHeap) Push(d int) {
	h.lastIdx++
	h.arr = append(h.arr, d)
}

func (h *maxHeap) Print() {
	for i := range h.arr {
		fmt.Printf("%d ", h.arr[i])
	}
	fmt.Printf("\n")
}
