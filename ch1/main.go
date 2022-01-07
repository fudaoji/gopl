package main

import "os"

func main() {
	var s, sep string
	for _, item := range os.Args[0:] {
		s += sep + item
		sep = "\t"
	}
	println(s)
}
