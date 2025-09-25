package main

import (
	"fmt"
	"net/url"
	"os"
	"strconv"
	"sync"
)

type config struct {
	baseURL            *url.URL
	pages              map[string]PageData
	mu                 *sync.Mutex
	concurrencyControl chan struct{}
	wg                 *sync.WaitGroup
	maxPages           int
}

func main() {
	args := os.Args
	maxConcurrency := 5
	maxPagesDefault := 10
	filename := "report.csv"

	// Verifica se existem mais argumentos.
	if len(args) < 2 {
		fmt.Println("no website provided")
		os.Exit(1)
	}
	if len(args) > 5 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}
	if len(args) == 3 {
		num, err := strconv.Atoi(args[2])
		if err != nil {
			fmt.Println("Error converting string to int in second arg:", err)
			os.Exit(1)
		}
		maxConcurrency = num
	}
	if len(args) == 4 {
		num, err := strconv.Atoi(args[3])
		if err != nil {
			fmt.Println("Error converting string to int in third arg:", err)
			os.Exit(1)
		}
		maxPagesDefault = num
	}
	if len(args) == 5 {
		filename = args[4]
	}
	var cfg config
	BASE_URL := args[1]
	cfg.pages = make(map[string]PageData)
	cfg.concurrencyControl = make(chan struct{}, maxConcurrency)
	cfg.mu = &sync.Mutex{}
	cfg.wg = &sync.WaitGroup{}
	cfg.maxPages = maxPagesDefault
	baseURL, err := url.Parse(BASE_URL)
	if err != nil {
		fmt.Printf("error parsing base url, %v\n", err)
		os.Exit(1)
	}
	cfg.baseURL = baseURL

	fmt.Printf("starting crawl on website: %s\n", BASE_URL)
	//fmt.Println("agora vai começar a parte recursiva")
	cfg.wg.Add(1)
	cfg.concurrencyControl <- struct{}{}
	go cfg.crawlPage(BASE_URL)
	<-cfg.concurrencyControl
	cfg.wg.Wait()
	//fmt.Println("agora termina a parte recursiva")
	err = writeCSVReport(cfg.pages, filename)
	if err != nil {
		fmt.Printf("error writing csv file, %v\n", err)
		os.Exit(1)
	}
}
