package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	for idx, arg := range args {
		fmt.Println(idx, arg)
	}
}
