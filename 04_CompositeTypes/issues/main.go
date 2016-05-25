package main

import (
	"fmt"
	"github.com/madmann8/GolangTraining/04_CompositeTypes/github"
	"log"
	"os"
	"time"
)

//!+
func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		if time.Since(item.CreatedAt)<2629746000000000 {
			//How to make table format:
			fmt.Printf("#%-5d %9.9s %.55s %.55s\n",
				item.Number, item.User.Login, item.Title, item.CreatedAt)
		}
		if time.Since(item.CreatedAt)<31556952000000000 {
			//How to make table format:
			fmt.Printf("#%-5d %9.9s %.55s %.55s\n",
				item.Number, item.User.Login, item.Title, item.CreatedAt)
		}
		if time.Since(item.CreatedAt)>31556952000000000 {
			//How to make table format:
			fmt.Printf("#%-5d %9.9s %.55s %.55s\n",
				item.Number, item.User.Login, item.Title)
		}
	}
}
