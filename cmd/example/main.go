package main

import (
	"fmt"

	"github.com/ez-connect/go-sonic/sonic"
)

func main() {

	ingester, err := sonic.NewIngester("localhost", 1491, "SecretPassword")
	if err != nil {
		panic(err)
	}

	// I will ignore all errors for demonstration purposes

	_ = ingester.BulkPush("movies", "general", 3, []sonic.IngestBulkRecord{
		{"id:6ab56b4kk3", "Star wars"},
		{"id:5hg67f8dg5", "Spider man"},
		{"id:1m2n3b4vf6", "Batman"},
		{"id:68d96h5h9d0", "This is another movie"},
	}, "eng")

	search, err := sonic.NewSearch("localhost", 1491, "SecretPassword")
	if err != nil {
		panic(err)
	}

	results, _ := search.Query("movies", "general", "man", 10, 0, "eng")

	fmt.Println(results)
}
