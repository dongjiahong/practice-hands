package main

import (
	"fmt"
	"sort"
)

type info struct {
	name  string
	score int
}

type infoList []info

func (il infoList) Len() int           { return len(il) }
func (il infoList) Less(i, j int) bool { return il[j].score < il[i].score }
func (il infoList) Swap(i, j int)      { il[i], il[j] = il[j], il[i] }

func rankByWordCount(wordFrequencies map[string]int) infoList {
	il := make(infoList, len(wordFrequencies))
	i := 0
	for k, v := range wordFrequencies {
		il[i] = info{k, v}
		i++
	}
	//sort.Sort(sort.Reverse(il))
	sort.Sort(il)
	return il
}

func main() {
	words := map[string]int{
		"aa": 2,
		"cc": 2,
		"bb": 3,
		"gg": 4,
		"dd": 9,
	}
	fmt.Println(rankByWordCount(words))
}
