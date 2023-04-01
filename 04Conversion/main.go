package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to Fiverr")
	fmt.Println("Please Leave your rating between 1 and 5: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for your rating of ", input)

	conveRate, err := strconv.ParseFloat(strings.TrimSpace(input) ,64)

	if err != nil {
		fmt.Println("Error: ", err)
	} else{
		fmt.Println("Thanks for your rating of ", conveRate + 0.5)
	}
	




}