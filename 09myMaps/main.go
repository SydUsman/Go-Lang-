package main

import "fmt"

func main() {
	fmt.Println("Maps")
	var Language =make(map[string]string)
	Language["JS"] = "Javascript"
	Language["Py"] = "Python"
	Language["Rb"] = "Ruby"

	for key, value := range Language{
		fmt.Printf("Language code %v is %v\n", key, value)
	}


}