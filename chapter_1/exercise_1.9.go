// copies the content found in url to the standard output
// and prints and the end each url with its res status code
package main

import (
	"fmt"
	"net/http"
	"os"
	"io"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") {
			if !strings.HasPrefix(url, "https://") {
				url = "http://" + url
			}
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: recieving from %s: %v\n", url, err)
			os.Exit(1)
		}
		_, err = io.Copy(os.Stdout, resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: copying response to Stdout: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Request URL: %s\nResponse Status Code: %s\n", url, resp.Status)
		resp.Body.Close()
	}
}
