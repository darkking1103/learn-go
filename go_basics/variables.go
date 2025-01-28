package main

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

	fmt.Printf("Type of variable is \"%T\"\n", variable)
}

/* Primitive types
	int along with variants like uint8, uint64, int8, int64, uintptr
	bool
	float along with float32 and float64
	complex
	string
*/

/* Compound types
	Array
	Slices
	Maps
	Structs
	Pointers

	Functions
	Channels
*/