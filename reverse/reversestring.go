package main

import "fmt"

// reverse the string
func main() {

	str := "checking"
	reverseValue := ""

	for _, value := range str {
		reverseValue = string(value) + reverseValue
	}

	fmt.Println(reverseValue)
}
