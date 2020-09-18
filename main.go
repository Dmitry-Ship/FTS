package main

import (
	"flag"
	"fmt"
	"log"
	"time"
)

func main() {
	var dumpPath, query string
	flag.StringVar(&dumpPath, "p", "enwiki-latest-abstract1.xml.gz", "wiki abstract dump path")
	flag.StringVar(&query, "q", "Small wild cat", "search query")
	flag.Parse()

	fmt.Println("loading documents...")

	loadingStart := time.Now()
	docs, err := loadDocuments(dumpPath)

	if err != nil {
		fmt.Println("error while reading document")
	}
	fmt.Println(fmt.Printf("loaded %d documents in %s ", len(docs), time.Since(loadingStart)))

	fmt.Println("indexing...")
	indexingStart := time.Now()
	indexedText := index{}
	indexedText.add(docs)
	fmt.Println(fmt.Printf("indexed  %d documents in %s ", len(docs), time.Since(indexingStart)))

	fmt.Println(fmt.Printf("searching... %s ", query))
	searchingStart := time.Now()
	result := indexedText.search(query)

	for _, docID := range result {
		doc := docs[docID]
		log.Printf("%d\t%s\n", docID, doc.Text)
	}
	fmt.Println(fmt.Printf("finished searching in %s with %d results", time.Since(searchingStart), len(result)))

}
