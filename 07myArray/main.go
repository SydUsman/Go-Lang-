package main

import "fmt"

func main() {
	fmt.Println("Arrays Starts")

	// Method 1 of Declaring an Array
	var arr1 [4]string
	arr1[0] = "Usman"
	arr1[1] = "Hassan"
	arr1[2] = "Naqvi"
	arr1[3] = "Syed"

	fmt.Println("Your name is, ", arr1)
	fmt.Println("Array length is, ", len(arr1))


	// Method 2 of declaring and using an Array
	 var arr2 = [3]string {"Usman","Hassan","Naqvi"}
	 fmt.Println("Your Arr 2 is", arr2)

}