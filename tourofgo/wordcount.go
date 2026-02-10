package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	wordcount := make(map[string]int) 

	for _, v := range strings.Fields(s){
		wordcount[v] += 1
	}

	return wordcount
}

func main() {
	wc.Test(WordCount)
}
