package main

import "fmt"

func main() {
	// // Method Number 1
	// fmt.Println("Lets Start Slices")
	// var mySlice = []string{"Apple","Orange","Guava"}
	// fmt.Println(len(mySlice))

	// mySlice = append(mySlice, "Mango","Banana")
	// fmt.Println(mySlice)

	// fmt.Println(mySlice[1:3])

	// //Memory Management in GoLang

	// var newSlice = make([]int,4)
	// newSlice[0] = 23
	// newSlice[1] = 64
	// newSlice[2] = 85
	// newSlice[3] = 32

	// //newSlice[4] = 87 // Cause an Error So use below Method
	// newSlice = append(newSlice, 43,11,92)
	// sort.Ints(newSlice)
	// fmt.Println(sort.IntsAreSorted(newSlice))
	// fmt.Println(newSlice)

	// How to remove Element from lice via Index
	var fruits = []string{"Apple", "Orange", "Peach", "Mango"}
	index := 1
	fruits = append(fruits[:index], fruits[index+1:]...)
	fmt.Println(fruits)


}