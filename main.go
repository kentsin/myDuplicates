package main

import "flag"

var (
	singleThread  bool   = false
	filenameMatch string = "*"
)

func main() {
	flag.stringVar(&filenameMatch, "name", "*", "filename pattern")
	flag.BoolVar(&singleThread, "single thread", false, "use single thread only")
}
