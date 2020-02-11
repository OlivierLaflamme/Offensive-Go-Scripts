package main

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"
)

func main() {
	var wait sync.WaitGroup
	var urls []string

	stat, err := os.Stdin.Stat()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[!] Usage: cat urls.txt | online", err)
		os.Exit(3)
	}

	if (stat.Mode() & os.ModeCharDevice) == 0 {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			urls = append(urls, scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintf(os.Stderr, "[!] Usage: cat urls.txt | online", err)
			os.Exit(3)
		}
	}

	wait.Add(len(urls))

	for _, url := range urls {
		go getHost(url, &wait)
	}
	wait.Wait()
}
func getHost(url string, wait *sync.WaitGroup) {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	timeout := time.Duration(2 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	_, err := client.Get(url)
	if err == nil {
		fmt.Println(url)
	}
	wait.Done()
}
