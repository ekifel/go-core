package index

import (
	"strings"
)

type Index struct {
	data map[string][]int
}

func New() *Index {
	index := new(Index)
	index.data = map[string][]int{}

	return index
}

func (ind *Index) Add(title string, j int) {
	keys := strings.Split(strings.ToLower(title), " ")
	for _, key := range keys {
		if !ind.indexExists(key, j) {
			ind.data[key] = append(ind.data[key], j)
		}
	}
}

func (ind *Index) Search(word string) []int {
	return ind.data[strings.ToLower(word)]
}

// func (i *Index) Print() {
// 	for i, d := range i.data {
// 		fmt.Printf("Word: '%v' contains in docs: %v\n", i, d)
// 	}
// }

func (ind *Index) indexExists(key string, id int) bool {
	for _, d := range ind.data[key] {
		if d == id {
			return true
		}
	}

	return false
}
