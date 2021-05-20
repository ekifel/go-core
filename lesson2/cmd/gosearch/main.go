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

	searchWord := flag.String("s", "", "Word for search")
	flag.Parse()

	if *searchWord == "" {
		log.Println("You don't set any words to search. Please do it with -s flag")
		os.Exit(0)
	}

	for _, url := range urls {
		data, err := s.Scan(url, depth)

		if err != nil {
			log.Println(err)
			continue
		}
		docs = append(docs, data...)
	}

	if *searchWord != "" {
		fmt.Println("Results:")

		for _, d := range docs {
			if strings.Contains(strings.ToLower(d.Title), strings.ToLower(*searchWord)) ||
				strings.Contains(strings.ToLower(d.Body), strings.ToLower(*searchWord)) {
				fmt.Println(d.URL, d.Title)
			}
		}
	}
}
