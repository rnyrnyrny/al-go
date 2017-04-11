package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	fileCounts := make(map[string]map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, fileCounts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2:	%v\n", err)
				continue
			}
			countLines(f, counts, fileCounts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
			for file, n := range fileCounts[line] {
				fmt.Printf("\t%s[%d]\n", file, n)
			}
		}
	}
}

func countLines(f *os.File, counts map[string]int, fileCounts map[string]map[string]int) {
	input := bufio.NewScanner(f)
	var line string
	for input.Scan() {
		line = input.Text()
		counts[line]++
		if counts[line] == 1 {
			fileCounts[line] = make(map[string]int)
		}
		fileCounts[line][f.Name()]++
	}
	// NOTE: ignoring potential errors from input.Err()
}
