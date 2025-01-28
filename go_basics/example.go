package main

// These example problems are taken from "The Go Programming Language" book

import (
	"fmt"
	"os"
)

// print the name of the program that invoked this program
func example1() {
	var arg = os.Args[0]
	var ans = ""
	for i := len(arg) - 1; i >= 0 && arg[i] != '/'; i-- {
		// ans += arg[i] // TODO
	}
	fmt.Println("this program name is invoked by", ans)

}

func main() {
	example1()
}
