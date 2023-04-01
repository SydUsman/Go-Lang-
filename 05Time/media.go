package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to Multiverse of madness")
	Time := time.Now()
	fmt.Println(Time)

	fmt.Println(Time.Format("01-02-2006 Monday MST 15:04:05"))
}