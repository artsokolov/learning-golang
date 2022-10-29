package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	linesInFiles := make(map[string][]string)
	files := os.Args[1:]

	if len(files) > 0 {
		for _, file := range files {
			f, err := os.Open(file)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%v\n", err)
				continue
			}

			countLinesInFiles(f, linesInFiles)

			f.Close()
		}
	}

	for str, fileNames := range linesInFiles {
		fmt.Printf("%q: %s\n", str, strings.Join(fileNames, ", "))
	}
}

func countLinesInFiles(f *os.File, lines map[string][]string) {
	foundStrings := make(map[string]bool)
	input := bufio.NewScanner(f)

	for input.Scan() {
		foundStrings[input.Text()] = true
	}

	for str := range foundStrings {
		lines[str] = append(lines[str], f.Name())
	}
}
