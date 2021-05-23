package main

import (
	"flag"
	"fmt"
	"go-core/gosearch/pkg/crawler"
	"go-core/gosearch/pkg/crawler/spider"
	"go-core/gosearch/pkg/index"
	"log"
	"os"
	"sort"
)

func main() {
	s := spider.New()
	depth := 2
	urls := []string{"https://go.dev", "https://golang.org"}
	var docs []crawler.Document

	word := flag.String("s", "", "Word for search")
	flag.Parse()

	if *word == "" {
		log.Printf("Вы не указали слово для поиска. Сделайте это используя флаг -s\n")
		os.Exit(1)
	}

	index := index.Index{}
	index.New()

	for _, url := range urls {
		data, err := s.Scan(url, depth)
		if err != nil {
			log.Println(err)
			continue
		}
		for _, doc := range data {
			doc.ID = len(docs)
			index.Add(doc.Title, doc.ID)
			docs = append(docs, doc)
		}
	}

	sort.Slice(docs, func(i, j int) bool { return docs[i].ID <= docs[j].ID })

	if *word != "" {
		fmt.Printf("Результаты:\n")

		ids := index.Search(*word)
		for _, id := range ids {
			d := sort.Search(len(docs), func(i int) bool { return docs[i].ID >= id })
			if docs[d].ID == id {
				fmt.Printf("Url: %v, Title: %v\n", docs[d].URL, docs[d].Title)
			}
		}
	}
}
