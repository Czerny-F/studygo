package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	ret := make(map[string]int)
	for _, w := range strings.Fields(s) {
		ret[w] += 1
	}
	return ret
}

func main() {
	wc.Test(WordCount)
}
