package main

import (
	"fmt"
	"github.com/maciej-gol/word-suggest/internal/query"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Missing input query.")
		return
	}
	q := query.NewQuery(os.Args[1])
	word, error := q.Execute()
	if error != nil {
		fmt.Println(error)
	}
	fmt.Println(word)
}
