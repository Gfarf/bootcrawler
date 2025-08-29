package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args

	// Verifica se existem mais argumentos.
	if len(args) < 2 {
		fmt.Println("no website provided")
		os.Exit(1)
	}
	if len(args) > 2 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}
	BASE_URL := args[1]
	fmt.Printf("starting crawl on website: %s", BASE_URL)
}
