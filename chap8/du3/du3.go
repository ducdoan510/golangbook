package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"sync"
)

func main() {
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}
	fileSizes := make(chan int64)
	wg := sync.WaitGroup{}

	for _, root := range roots {
		wg.Add(1)
		go walkDirs(root, fileSizes, &wg)
	}

	go func() {
		wg.Wait()
		close(fileSizes)
	}()
	var nfiles, nbytes int64
	for size := range fileSizes {
		nfiles++
		nbytes += size
	}
	printDiskUsage(nfiles, nbytes)
}

func printDiskUsage(nfiles int64, nbytes int64) {
	fmt.Printf("%d files %.1f KB\n", nfiles, float64(nbytes)/1e3)
}

func walkDirs(dir string, fileSizes chan<- int64, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, entry := range dirEnts(dir) {
		if entry.IsDir() {
			subDir := path.Join(dir, entry.Name())
			wg.Add(1)
			walkDirs(subDir, fileSizes, wg)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

func dirEnts(dir string) []os.FileInfo {
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		return nil
	}
	return entries
}