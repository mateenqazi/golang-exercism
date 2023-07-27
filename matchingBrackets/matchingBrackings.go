package main

import "fmt"

// check balance brackets
func main() {
	value := "{([{(())}])}"
	list := []string{}
	isBracketMatched := true

	for _, val := range value {
		val1 := string(val)
		switch val1 {
		case "(", "{", "[":
			list = append(list, val1)

		case ")":
			if len(list) == 0 || list[len(list)-1] != "(" {
				isBracketMatched = false
				break
			}

			list = list[:len(list)-1]

		case "}":
			if len(list) == 0 || list[len(list)-1] != "{" {
				isBracketMatched = false
				break
			}

			list = list[:len(list)-1]

		case "]":
			if len(list) == 0 || list[len(list)-1] != "[" {
				isBracketMatched = false
				break
			}

			list = list[:len(list)-1]
		}

	}
	if len(list) != 0 {
		isBracketMatched = false
	}

	fmt.Println("Backet matched : ", isBracketMatched)
}
