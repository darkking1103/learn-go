package go_basics

import "fmt"

/* -- this is package level
All the variables declared here will be visible to all the files
of this package
*/

// create uninitialized variables
var n int
var b bool
var str string

// create initialized variables
var a int = 4
var c = 5 // automatic type deduction


func main() {
	variable := 3; // short variable declaration can only be done inside blocks
	fmt.Println(variable)
}
