package main

import (
	"fmt"
	"golangbook/chap4/github"
	"log"
	"time"
)

func main() {
	terms := []string{"repo:golang/go", "is:open", "json", "decoder"}
	result, err := github.SearchIssues(terms)
	if err != nil {
		log.Fatal(err)
	}
	for _, item := range result.Items {
		fmt.Printf("%-5d | ", item.Number)
		createdAt := item.CreatedAt
		diff := time.Since(createdAt)
		days := diff.Hours() / 24
		switch {
		case days < 30:
			fmt.Println("Less than a month old")
		case days < 365:
			fmt.Println("Less than a year old")
		default:
			fmt.Println("More than or equal a year old")
		}
	}
}
