package main

import "fmt"

func main() {
	fmt.Println("We are learning Struct")
	User1 := User{"Usman","usman@google.dev",true,16}
	fmt.Printf("Details of User1 is: %+v\n", User1)


}

type User struct{
	Name string
	Email string
	Status bool
	Age int
}