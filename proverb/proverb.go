package main

import "fmt"

// Given a list of inputs, generate the relevant proverb
func main() {

	list1 := []string{"nail", "shoe", "horse", "rider", "message", "battle", "kingdom"}

	for _, value := range list1 {
		fmt.Printf("For want of a %v the shoe was lost.\n", value)
	}
}
