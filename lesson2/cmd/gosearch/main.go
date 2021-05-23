package main

import (
	"flag"
	"fmt"
	"go-core/lesson2/pkg/crawler"
	"go-core/lesson2/pkg/crawler/spider"
	"log"
	"os"
	"strings"
)

func main() {
	s := spider.New()
	depth := 2
	urls := []string{"https://go.dev", "https://golang.org"}
	var docs []crawler.Document

	word := flag.String("s", "", "Word for search")
	flag.Parse()

	if *word == "" {
		log.Println("Вы не указали слово для поиска. Сделайте это используя флаг -s")
		os.Exit(1)
	}

	for _, url := range urls {
		data, err := s.Scan(url, depth)
		if err != nil {
			log.Println(err)
			continue
		}
		docs = append(docs, data...)
	}

	if *word != "" {
		fmt.Println("Results:")

		for _, d := range docs {
			if strings.Contains(strings.ToLower(d.Title), strings.ToLower(*word)) ||
				strings.Contains(strings.ToLower(d.Body), strings.ToLower(*word)) {
				fmt.Println(d.URL, d.Title)
			}
		}
	}
}
