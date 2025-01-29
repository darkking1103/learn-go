package main

import "fmt"

/* -- this is package level
	All the variables declared here will be visible to all the files
	of this package

	Variable declared at package level with the first letter as capital letter becomes public variable,
	like Println int fmt package fmt.Println()
*/
const PublicVariable string = "Value of public variable"


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

	const EMAIL = "example@example.com"
	// EMAIL = "some other value" // error cannot change a constant variable

	const char_a byte = 'a';
	fmt.Println(char_a) // prints 97 ascii value of 'a'

	// string
	const str1 = "thisisastring"
	fmt.Println(str1[0]) // 116 for 't'
	fmt.Println(str1[4]) // 105 for 'i'
	fmt.Println(str1[0:]) // "thisisastring" from index 0->len(str1)
	fmt.Println(str1[0:4]) // "this" from index 0->3
}


/* Primitive types
	int (default value 0 if not initialized with a value)
		- uint8, uint64, int8, int64, uintptr
		- An int type will default to be 64 bits when used on a 64 bit system, 32 bits on a 32 bit system

	bool
	float along with float32 and float64
	complex along with complex64, complex128
	string (default "")
	byte (represents a single ASCII character like char type in other langs)
*/

/* Compound types
	Array
	Slices
	Maps
	Structs
	Pointers

	Functions
	Channels
	...
*/