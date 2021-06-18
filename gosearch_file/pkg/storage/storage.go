package storage

import (
	"bufio"
	"encoding/json"
	"go-core/gosearch_file/pkg/crawler"
	"os"
)

func Save(path string, docs []crawler.Document) error {
	store, err := os.Create(path)
	if err != nil {
		return err
	}
	defer store.Close()

	for _, doc := range docs {
		d, err := json.Marshal(doc)
		if err != nil {
			return err
		}
		_, err = store.Write(append(d, '\n'))
		if err != nil {
			return err
		}
	}

	return nil
}

func Load(f *os.File) ([]crawler.Document, error) {
	var docs []crawler.Document
	s := bufio.NewScanner(f)

	for s.Scan() {
		res := crawler.Document{}
		err := json.Unmarshal(s.Bytes(), &res)
		if err != nil {
			return nil, err
		}
		docs = append(docs, res)
	}
	return docs, s.Err()
}

func CheckStorage(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}

	return true
}
