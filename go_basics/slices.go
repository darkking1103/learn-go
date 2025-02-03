package main

import (
	"fmt"
)

/* array
	arrays are non-resizable arrays
	arrays cannot be initialize with a variable like this: "arr := [size_var]string{"one", "two"}
*/


/* slices
	slices are resizable arrays
*/

func main() {

	// arrays
	arr := [3]string{"one", "two", "three"}
	fmt.Println(arr, "\n")


	// Slices
	slice := []string{"first", "second", "third"}
	slice = append(slice, "fourth")

	slice2 := arr[:] // initialize 'slice2' with 'slice' as the same underlying array
	slice2[1] = "one"

	slice3 := append(slice, "five", "six")

	slice4 := make([]string, 3) // create slice with a given size
	slice4[0] = "hello"
	slice4[1] = "world"
	slice4[2] = "!"
	// slice4[3] = "." // index out of range
	fmt.Println(slice4)

	fmt.Println(slice)  // [first one third fourth]
	fmt.Println(slice2) // [first one third fourth]
	fmt.Println(slice3, "\n") // [first one third fourth five six]

	copy(slice2, slice3) // copy slice2 into slice3


}
