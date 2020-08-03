package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fileNames := os.Args[1:]
	lineMap := make(map[string][]string)
	for _, fileName := range fileNames {
		updateLineMap(fileName, lineMap)
	}

	for line, files := range lineMap {
		if len(files) > 1 {
			fmt.Printf("Dup line: %s \t %v", line, files)
		}
	}
}

func updateLineMap(fileName string, mapping map[string][]string) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %s\n", fileName)
		os.Exit(1)
	}
	sc := bufio.NewScanner(file)
	for sc.Scan() {
		line := sc.Text()
		mapping[line] = append(mapping[line], fileName)
	}
}
