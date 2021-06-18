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
	word := flag.String("s", "", "Укажите слово для поиска")
	flag.Parse()

	if *word == "" {
		flag.PrintDefaults()
		return
	}

	urls := []string{"https://go.dev", "https://golang.org"}
	index := index.New()
	path := "./dataStorage.json"
	var docs []crawler.Document

	if storage.CheckStorage(path) {
		d, err := loadData(path)
		if err != nil {
			log.Printf("%s\n", err)
		} else {
			docs = append(docs, d...)
		}
	}
	if docs == nil {
		d, err := scanUrls(urls)
		if err != nil {
			log.Fatalf("%s\n", err)
		}
		docs = append(docs, d...)

		if len(docs) > 0 {
			err = storage.Save(path, docs)
			if err != nil {
				log.Printf("%s\n", err)
			}
		}
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
		for _, doc := range d {
			doc.ID = len(docs)
			docs = append(docs, doc)
		}
	}

	return docs, nil
}

func loadData(path string) ([]crawler.Document, error) {
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
