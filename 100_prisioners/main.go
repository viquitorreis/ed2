package main

import (
	"math/rand"
)

func main() {
	quantity := 100

	boxes := genIntArr(quantity)
	prisioners := genPrisioners(quantity)

	println("liberdade?", find(prisioners, boxes))
}

type prisioner struct {
	val int
}

func find(pris *[]prisioner, boxes *[]int) bool {
	for pIdx := 0; pIdx < len(*pris); pIdx++ {
		currPris := (*pris)[pIdx]
		currBoxIdx := currPris.val - 1 // começa no numero do prisionerio
		found := false

		for range 50 {
			boxContent := (*boxes)[currBoxIdx]

			if currPris.val == boxContent {
				found = true
				break
			} else {
				currBoxIdx = boxContent - 1
			}
		}

		if !found {
			return false
		}
	}

	return true
}

func shuffle(arr *[]int) {
	for i := range len(*arr) {
		rd := rand.Intn(len(*arr))
		temp := (*arr)[rd] // valor alvo pro swap
		(*arr)[rd] = (*arr)[i]
		(*arr)[i] = temp
	}
}

func genIntArr(n int) *[]int {
	arr := make([]int, n)

	for i := range n {
		arr[i] = i + 1
	}

	shuffle(&arr)

	return &arr
}

func genPrisioners(n int) *[]prisioner {
	nums := genIntArr(n)

	prisioners := make([]prisioner, n)
	for i := range prisioners {
		prisioners[i].val = (*nums)[i]
	}

	return &prisioners
}
