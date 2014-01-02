package main

import (
	"regexp"
	"sort"
)

type Words []*Word
type Word struct {
	name       string
	occurences int
}

func (w Words) Len() int {
	return len(w)
}
func (w Words) Swap(i, j int) {
	w[i], w[j] = w[j], w[i]
}
func (w Words) Less(i, j int) bool {
	return w[i].occurences > w[j].occurences
}

func WordCount(s string) Words {
	var pairs Words
	r, _ := regexp.Compile("\\s+")
	words := make(map[string]int)
	for _, word := range r.Split(s, -1) {
		words[word] += 1
	}
	for key, value := range words {
		pairs = append(pairs, &Word{key, value})
	}
	sort.Sort(pairs)
	return pairs
}
