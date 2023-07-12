package main

import (
	"fmt"
	"math"
)

func main() {
	booking := []int{2, 2, 3, 4, 4, 4, 5, 5}
	discounts := map[int]int{
		1: 0,
		2: 5,
		3: 10,
		4: 20,
		5: 25,
	}

	mapValue := make(map[int]float64)
	for _, value := range booking {
		mapValue[value]++
	}
	totalAmount := 0.0
	for _, element := range mapValue {

		bookPrint := 8 * element
		addDiscount := 1.0

		if float64(discounts[int(element)]) != 0 {
			addDiscount = ((float64(discounts[int(element)])) / 100)
		}
		totalAmount += bookPrint - addDiscount
	}

	fmt.Println(roundFloat(totalAmount, 2))
}

func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}
