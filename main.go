package main

import (
	"fmt"
	"net/url"
	"os"
	"sync"
)

type config struct {
	baseURL            *url.URL
	pages              map[string]PageData
	mu                 *sync.Mutex
	concurrencyControl chan struct{}
	wg                 *sync.WaitGroup
}

func main() {
	args := os.Args
	BUFFER := 1

	// Verifica se existem mais argumentos.
	if len(args) < 2 {
		fmt.Println("no website provided")
		os.Exit(1)
	}
	if len(args) > 2 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}
	var cfg config
	BASE_URL := args[1]
	cfg.concurrencyControl = make(chan struct{}, BUFFER)
	baseURL, err := url.Parse(BASE_URL)
	if err != nil {
		fmt.Printf("error parsing base url, %v\n", err)
		os.Exit(1)
	}
	cfg.baseURL = baseURL

	fmt.Printf("starting crawl on website: %s\n", BASE_URL)
	//fmt.Println("agora vai começar a parte recursiva")
	go cfg.crawlPage(BASE_URL)
	cfg.wg.Wait()
	//fmt.Println("agora termina a parte recursiva")
	for page, i := range cfg.pages {
		fmt.Printf("page %s, coung %d\n", page, i)
	}
}
