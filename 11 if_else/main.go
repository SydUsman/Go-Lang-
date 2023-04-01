package main

import (
	"fmt"
)

func main() {
	fmt.Println("Enter your Pasword")
	var pass string
  	fmt.Scan(&pass)
	
	if pass == "admin" {
		fmt.Println("Success")
	} else{
		fmt.Println("Failure")
	}

}