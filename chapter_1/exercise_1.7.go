// copies the content found in url to the standard output
package main

import (
	"fmt"
	"net/http"
	"os"
	"io"
)

func main() {
	for _, url := range os.Args[1:] {
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
		resp.Body.Close()
	}
}
