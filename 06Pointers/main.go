package main

import "fmt"

func main() {
	// var ptr *int
	// fmt.Println("Value of Pointer is ", ptr)

	myPointer := 45
	var ptr = &myPointer
	fmt.Println("Value is ,", ptr)
	fmt.Println("Value is ,", *ptr)

	// POinter dont make copy
	*ptr = *ptr + 2
	fmt.Println("New Value is, ", myPointer)


}