package main

import (
	"fmt"
	"golangbook/chap2/tempcov"
	"os"
	"strconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		f := tempcov.Fahrenheit(t)
		c := tempcov.Celsius(t)
		fmt.Printf("%s = %s, %s = %s\n", f, tempcov.FToC(f), c, tempcov.CToF(c))
	}
}
