package main

import (
	"fmt"
	"strconv"
)

// Convert a number, represented as a sequence of digits in one base, to any other base

func main() {
	inputBase := 2
	digits := "101011"
	outputBase := 3

	if inputBase < 2 || outputBase < 2 {
		fmt.Println("Base must be >= 2")
	} else {
		var sum int

		for _, d := range digits {
			num, _ := strconv.Atoi(string(d))
			sum = sum*inputBase + num
		}

		str := strconv.FormatInt(int64(sum), outputBase) // convert into specific base
		fmt.Println(str)
	}
}
