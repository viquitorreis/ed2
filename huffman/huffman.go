package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	str := "helloworld!"

	Encode(str)
}

func Encode(s string) []rune {
	freq := make(map[rune]int, utf8.RuneCountInString(s))

	for _, r := range s {
		freq[r] += 1
	}

	for r, frq := range freq {
		fmt.Printf("%c: %d\n", r, frq)
	}

	ordered := make([]rune, 0, len(freq))
	for k, _ := range freq {
		val := ordered[k]
		if val == 0 {
			ordered[k] = k
		}
	}

	return []rune{}
}

func swapMap(mp *map[rune]int, lo, hi *rune) {

}
