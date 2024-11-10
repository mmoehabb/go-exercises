package main

import (
	"fmt"
	"os"
	"strings"
)

type Key struct {
	line, filepath string
}

func main() {
	counts := make(map[string]int)
	line_files := make(map[Key]int)
	for _, filepath := range os.Args[1:] {
		data, err := os.ReadFile(filepath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
			line_files[Key{line, filepath}]++
		}
	}
	delete(counts, "")
	for obj, n := range line_files {
		if n > 0 {
			fmt.Printf("%d\t%s\n", line_files[Key{obj.line, obj.filepath}], obj.line)
		}
	}
}
