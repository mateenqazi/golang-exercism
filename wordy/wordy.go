package main

import (
	"fmt"
	"regexp"
	"strconv"

	"golang.org/x/exp/slices"
)

// Parse and evaluate simple math word problems returning the answer as an integer.

func main() {
	str := "What is 5 plus 13?"
	OPERATIONS_REGEX := `(-?\d+|plus|minus|multiplied by|divided by)`
	OPERATOR_STRING := []string{"plus", "minus", "multiplied by", "divided by"}
	var interResult int64

	re := regexp.MustCompile(OPERATIONS_REGEX).FindAllString(str, -1)

	fmt.Println(re)

	if len(re) > 1 {
		for index, value := range re {
			if slices.Contains(OPERATOR_STRING, value) {
				fmt.Printf("Operator name %v  %v %v \n", value, re[index-1], re[index+1])

				y, err := strconv.ParseInt(re[index-1], 10, 64)
				if err != nil {
					fmt.Println("error occurs in parse int")
					break
				}

				z, errValue := strconv.ParseInt(re[index+1], 10, 64)
				if errValue != nil {
					fmt.Println("error occurs in parse int")
					break
				}

				var resultBefore int64
				if interResult != 0 {
					resultBefore = interResult
				} else {
					resultBefore = y
				}

				interResult = operationExecuateOnString(value, resultBefore, z)
			}
		}
	} else {
		interResult, _ = strconv.ParseInt(re[0], 10, 64)
	}

	fmt.Println(interResult)
}

func plus(val1, val2 int64) int64 {
	return val1 + val2
}

func minus(val1, val2 int64) int64 {
	return val1 - val2
}

func multiplied(val1, val2 int64) int64 {
	return val1 * val2
}

func divided(val1, val2 int64) int64 {
	return val1 / val2
}

func operationExecuateOnString(operator string, val1, val2 int64) int64 {
	switch operator {

	case "plus":
		return plus(val1, val2)
	case "minus":
		return minus(val1, val2)
	case "multiplied by":
		return multiplied(val1, val2)
	case "divided by":
		return divided(val1, val2)
	default:
		return -1
	}
}
