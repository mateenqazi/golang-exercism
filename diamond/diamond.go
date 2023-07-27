package main

import (
	"fmt"
	"strings"
)

// make an daimond

func main() {
	char := 'C'
	if char == 'A' {
		fmt.Println("A")
	}

	var result string
	for i := 65; i <= int(char); i++ {
		diff := int(char) - i
		midDiff := i - 65

		if midDiff > 1 {
			midDiff = midDiff*2 - 1
		}

		if i == 65 {
			result += strings.Repeat(" ", diff) + string(rune(i)) + strings.Repeat(" ", diff) + "\n"
		} else {
			result += strings.Repeat(" ", diff) + string(rune(i)) + strings.Repeat(" ", midDiff) + string(rune(i)) + strings.Repeat(" ", diff) + "\n"
		}
	}
	for i := int(char) - 1; i >= 65; i-- {
		diff := int(char) - i
		midDiff := i - 65

		if midDiff > 1 {
			midDiff = midDiff*2 - 1
		}

		if i == 65 {
			result += strings.Repeat(" ", diff) + string(rune(i)) + strings.Repeat(" ", diff) + "\n"
		} else {
			result += strings.Repeat(" ", diff) + string(rune(i)) + strings.Repeat(" ", midDiff) + string(rune(i)) + strings.Repeat(" ", diff) + "\n"
		}
	}

	fmt.Println(result)
}
