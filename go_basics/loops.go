package main

import (
	"fmt"
	"os"
	"strconv"
)

var foo string = "hello world"

/* :=
The := symbol is part of
a short variable declaration, a statement that declares one or more variables and gives them
appropriate types based on the initializer values
*/

/* slice
s[m:n] yields a slice that refers to elements m thro ugh n-1, so the
elements we need for our next example are those in the slice os.Args[1:len(os.Args)]. If m
or n is omitt ed, it defaults to 0 or len(s) respectively, so we can abbreviate the desired slice as
os.Args[1:]
*/

func main() {
	// i++ is a statement not an expression unlike in other langs like C
	for i := 0; i < 5; i++ {
		fmt.Println("Hello", i)
	}
	fmt.Println()

	// while loop
	var i = 1
	
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run loops.go <integer-arg>")
		return
	}
	
	var n, err = strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Usage: go run loops.go <integer-arg>")
		return
	}

	fmt.Println("Table of", n)
	for i <= 10 {
		fmt.Println(n, "x", i, "=", n*i)
		i++
	}

	/* range based loop
	One idea would be to assign the index to an obviously temporary variable like temp and ignore its value, but Go
	does not permit unused local variables, so this would result in a compilation error.
	The solution is to use the blank identifier, whose name is _
	*/
	var arr = [4]string{"Hello", ",", " World", "!\n"}
	for _, word := range arr {
		fmt.Print(word)
	}

}
