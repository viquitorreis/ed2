package main

type HuffmanMinHeap []HuffmanTree

// simples. Poderiamos checar se o char unicode específico é menor também
// e/ou se a subárvore i tem menos leafs que a j
func (h HuffmanMinHeap) Less(i, j int) bool {
	return h[i].Freq() < h[j].Freq()
}

func (h HuffmanMinHeap) Len() int {
	return len(h)
}

func (h HuffmanMinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *HuffmanMinHeap) Push(d any) {
	*h = append(*h, d.(HuffmanTree))
}

func (h *HuffmanMinHeap) Pop() (poppedElement any) {
	poppedElement = (*h)[h.Len()-1]

	*h = (*h)[:len(*h)-1]

	return
}
