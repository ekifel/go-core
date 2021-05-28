package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"go-core/gosearch/pkg/crawler"
	"go-core/gosearch/pkg/crawler/spider"
	"go-core/gosearch/pkg/index"
	"io"
	"log"
	"os"
	"sort"
)

func main() {
	depth := 2
	urls := []string{"https://go.dev", "https://golang.org"}
	var docs []crawler.Document

	word := flag.String("s", "", "Word for search")
	flag.Parse()

	if *word == "" {
		log.Printf("Вы не указали слово для поиска. Сделайте это используя флаг -s\n")
		os.Exit(1)
	}
	index := index.New()

	ds := "./dataStorage.json"
	if _, err := os.Stat("/path/to/whatever"); os.IsNotExist(err) {
		docs = scanUrls(urls, depth, index)
		err = writeToFile(ds, docs)
		if err != nil {
			log.Println(err)
		}
	} else {
		docs, err = loadFromFile(ds)
		if err != nil {
			log.Fatal(err)
		}
		for _, doc := range docs {
			index.Add(doc.Title, len(docs))
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

func scanUrls(urls []string, depth int, index *index.Index) []crawler.Document {
	s := spider.New()

	var docs []crawler.Document
	for _, url := range urls {
		data, err := s.Scan(url, depth)
		if err != nil {
			log.Println(err)
			continue
		}
		for _, doc := range data {
			index.Add(doc.Title, len(docs))
			docs = append(docs, doc)
		}
	}

	return docs
}

func writeToFile(path string, docs []crawler.Document) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	for _, d := range docs {
		err = putDocInFile(f, d)
		if err != nil {
			return err
		}
	}
	return nil
}

func putDocInFile(w io.Writer, d crawler.Document) error {
	b, err := json.Marshal(d)
	if err != nil {
		return err
	}
	_, err = w.Write(append(b, '\n'))
	if err != nil {
		return err
	}
	return nil
}

func getDocsFromFile(r io.Reader) ([]crawler.Document, error) {
	docs := []crawler.Document{}
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		res := crawler.Document{}
		err := json.Unmarshal(scanner.Bytes(), &res)
		if err != nil {
			return nil, err
		}
		docs = append(docs, res)
	}
	return docs, scanner.Err()
}

func loadFromFile(path string) ([]crawler.Document, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return getDocsFromFile(f)
}
