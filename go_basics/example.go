package main

import (
	"fmt"
	"os"
)

// print the name of the program or executable
func example1() {
	var arg = os.Args[0]
	var ans = ""
	for i := len(arg) - 1; i >= 0 && arg[i] != '/'; i-- {
		ans = string(arg[i]) + ans // string and byte concatenation
	}
	fmt.Println("program or executable name is:", ans);

}

func main() {
	example1()
}
