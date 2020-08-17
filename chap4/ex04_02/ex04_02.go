package main

import (
	"bufio"
	"crypto/sha256"
	sha5122 "crypto/sha512"
	"flag"
	"fmt"
	"os"
)

var sha384 = flag.Bool("sha384", false, "Use SHA384 instead")
var sha512 = flag.Bool("sha512", false, "Use SHA512 instead")

func main() {
	flag.Parse()
	if *sha384 && *sha512 {
		fmt.Fprintf(os.Stderr, "Choose one algorithm only")
		os.Exit(1)
	}

	// Read input
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter input for hashing:")
	scanner.Scan()
	input := scanner.Text()

	// Print hash based on selected algorithm
	switch {
	case *sha384:
		fmt.Printf("%x\n", sha5122.Sum384([]byte(input)))
	case *sha512:
		fmt.Printf("%x\n", sha5122.Sum512([]byte(input)))
	default:
		fmt.Printf("%x\n", sha256.Sum256([]byte(input)))
	}
}
