package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	singleThread  bool   = false
	filenameMatch string = "*"
)

func main() {
	flag.StringVar(&filenameMatch, "name", "*", "filename pattern")
	flag.BoolVar(&singleThread, "single thread", false, "use single thread only")
	var help = flag.Bool("h", false, "Display this message")
	flag.Parse()
	if *help {
		fmt.Println("\nduplicates is a command line tool to find duplicate files in a folder\n")
		fmt.Println("usage: duplicates [options...] path\n")
		flag.PrintDefaults()
		os.Exit(0)
	}
	fmt.printf("Using %d of threads", singleThread)
}
