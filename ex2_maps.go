package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	str := strings.Fields(s)
	ret := make(map[string]int)
	for i := 0; i < len(str); i++ {
		ret[str[i]]
		i++
	}
	return ret
}

// return a map of the counts of each word in the string
func main() {
	wc.Test(WordCount)
}