// fetches urls concurrently and reports the process
// in the following pattern: [time elapsed] [num of bytes] [url]
package main

import (
	"fmt"
	"net/http"
	"os"
	"io"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	
	filename := "fetchall_report.txt" 
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetchall: reading file%s: %v\n", filename, err)
		os.Exit(1)
	}

	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}
	for range os.Args[1:] {
		fmt.Fprintln(file, <-ch)
	}
	fmt.Fprintf(file, "%.2f elapsed.\n\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	if !strings.HasPrefix(url, "http://") {
		if !strings.HasPrefix(url, "https://") {
			url = "http://" + url
		}
	}
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: recieving from %s: %v\n", url, err)
		ch <- fmt.Sprintf("%.2f\terror\t%s", time.Since(start).Seconds(), url)
	}
	nbytes, err := io.Copy(io.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: copying response to Stdout: %v\n", err)
		ch <- fmt.Sprintf("%.2f\terror\t%s", time.Since(start).Seconds(), url)
	}
	ch <- fmt.Sprintf("%.2f\t%7d\t%s", time.Since(start).Seconds(), nbytes, url)
}
