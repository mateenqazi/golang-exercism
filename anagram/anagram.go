package main

import (
	"fmt"
	"sort"
	"strings"
)

// check two words are anagram or not
func main() {

	input1 := sortedString(getInputFromUser("first"))
	input2 := sortedString(getInputFromUser("second"))

	fmt.Printf("Both string are anagram %v \n", isAnagram(input1, input2))
}

func getInputFromUser(s string) string {
	input := ""

	fmt.Printf("Please Enter %v string value : ", s)
	fmt.Scan(&input)

	return strings.ToLower(input)
}

func sortedString(s string) string {
	runes := []rune(s)

	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	sortedString := string(runes)

	return sortedString
}

func isAnagram(first string, second string) bool {
	return first == second
}
