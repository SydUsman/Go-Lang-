package main

import "fmt"

var Aage int = 22
var Age string = "Twenty Two"

func main() {
	var username float64 = 244.654335345
	fmt.Println("Hello", username, "Welcome to GoLang")
	fmt.Println("Type of username is", fmt.Sprintf("%T", username))

	// implicit style
	var name = "UsmanNaqvi"
	fmt.Println("Hello", name, "Welcome to GoLang")

	// no var Style
	myname := "Usman"
	fmt.Println("Hello", myname, "Welcome to GoLang")

}