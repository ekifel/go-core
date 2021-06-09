package main

import (
	"flag"
	"fmt"
	"go-core/gosearch_file/pkg/crawler"
	"go-core/gosearch_file/pkg/crawler/spider"
	"go-core/gosearch_file/pkg/index"
	"go-core/gosearch_file/pkg/storage"
	"log"
	"os"
	"sort"
)

func main() {
	word := flag.String("s", "", "Word for search")
	flag.Parse()

	if *word == "" {
		log.Printf("Вы не указали слово для поиска. Сделайте это используя флаг -s\n")
		os.Exit(1)
	}

	index := index.New()
	path := "./dataStorage.json"
	docs, err := loadData(path)
	if err != nil {
		log.Print(err)
	}
	index.Create(&docs)
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

func scanUrls(urls []string) ([]crawler.Document, error) {
	s := spider.New()
	var docs []crawler.Document
	depth := 2

	for _, url := range urls {
		d, err := s.Scan(url, depth)
		if err != nil {
			return nil, err
		}
		docs = append(docs, d...)
	}

	for _, doc := range docs {
		doc.ID = len(docs)
	}

	return docs, nil
}

func loadData(path string) ([]crawler.Document, error) {
	urls := []string{"https://go.dev", "https://golang.org"}

	if storage.CheckStorage(path) {
		store, err := os.Open(path)
		if err != nil {
			return nil, err
		}
		defer store.Close()
		docs, err := storage.Load(store)
		if err != nil {
			return nil, err
		}
		return docs, nil
	}

	docs, err := scanUrls(urls)
	if err != nil {
		return nil, err
	}
	if len(docs) > 0 {
		err = storage.Save(path, docs)
		if err != nil {
			return docs, err
		}
	}

	return docs, nil
}
