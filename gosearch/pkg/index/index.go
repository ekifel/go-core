package index

import (
	"strings"
)

type Index struct {
	data map[string][]int
}

func (i *Index) New() {
	i.data = make(map[string][]int)
}

func (i *Index) Add(title string, j int) {
	keys := strings.Split(strings.ToLower(title), " ")
	for _, key := range keys {
		if !i.indexExists(key, j) {
			i.data[key] = append(i.data[key], j)
		}
	}
}

func (i *Index) Search(word string) []int {
	return i.data[strings.ToLower(word)]
}

// func (i *Index) Print() {
// 	for i, d := range i.data {
// 		fmt.Printf("Word: '%v' contains in docs: %v\n", i, d)
// 	}
// }

func (i *Index) indexExists(key string, j int) bool {
	for _, d := range i.data[key] {
		if d == j {
			return true
		}
	}

	return false
}
