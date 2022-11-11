package main

import (
	"bufio"
	"fmt"
	"os"
)

// You can use os.Open to open a external file with the path.
// In this exercise you can give the path of multiple files through the arguments
// (if you don't know how use the arguments, go to 1.2 section)
// and then, you will count the reapeted lines like in dup1.go


// Golang does not have a try catch, how would you handle errors?
// many of the functions return an error type value, if this is nil, all is good.

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			// f.Close() is really important, you can't forget this
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}
