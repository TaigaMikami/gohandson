package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	count := map[string]int{}

	for _, word := range strings.Fields(s) {
		count[word]++
	}
	return count
}

func main() {
	wc.Test(WordCount)
}
