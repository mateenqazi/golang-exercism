package main

import (
	"fmt"
	"strings"
	"unicode"
)

// you have a string. Please the frequency of words in that string.
func main() {
	wordCount := map[string]int{}
	str := "That's the password: 'PASSWORD 123'!\", cried the Special Agent.\nSo I fled."
	newValue := ""
	var newString []string

	for _, value := range strings.ToLower(str) {
		if notWordRune(value) {
			newValue += string(value)
		}
	}
	newString = strings.Split(newValue, " ")

	for _, value := range newString {
		val, ok := wordCount[value]
		if ok {
			wordCount[value] += 1
		} else {
			wordCount[value] = 1
		}
		fmt.Println(val, ok)
	}

	fmt.Println(wordCount)

}

func notWordRune(c rune) bool {
	return (unicode.IsLetter(c) || unicode.IsDigit(c) || unicode.IsSpace(c))
}
