package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var verbose = flag.Bool("v", false, "show verbose progress messages")

type fileSize struct {
	root string
	size int64
}

type rootStat struct {
	nfiles int64
	nbytes int64
}

func main() {
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}
	fileSizes := make(chan fileSize)
	wg := sync.WaitGroup{}
	for _, root := range roots {
		wg.Add(1)
		go walkDir(root, root, fileSizes, &wg)
	}

	go func() {
		wg.Wait()
		close(fileSizes)
	}()

	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(1 * time.Millisecond)
	}

	stats := make(map[string]rootStat)

loop:
	for {
		select {
		case fileSize, ok := <-fileSizes:
			if !ok {
				break loop
			}
			stats[fileSize.root] = rootStat{
				stats[fileSize.root].nfiles + 1,
				stats[fileSize.root].nbytes + fileSize.size,
			}
		case <-tick:
			printDiskUsage(stats)
		}
	}
	printDiskUsage(stats)
}

func printDiskUsage(stats map[string]rootStat) {
	for root, stat := range stats {
		fmt.Printf("%s: %d files %.1f KB\n", root, stat.nfiles, float64(stat.nbytes)/1e3)
	}
	fmt.Println()
}

func walkDir(root string, dir string, fileSizes chan<- fileSize, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, entry := range dirEnts(dir) {
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			wg.Add(1)
			walkDir(root, subdir, fileSizes, wg)
		} else {
			fileSizes <-  fileSize{root, entry.Size() }
		}
	}
}

var sema = make(chan struct{}, 20)

func dirEnts(dir string) []os.FileInfo {
	sema <- struct{}{}
	defer func() { <-sema }()
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return nil
	}
	return entries
}
