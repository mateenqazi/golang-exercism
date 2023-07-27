package main

import (
	"fmt"
	"math"
)

/*

To try and encourage more sales of different books from a popular 5 book series, a bookshop has decided to offer discounts on multiple book purchases.

One copy of any of the five books costs $8.

If, however, you buy two different books, you get a 5% discount on those two books.

If you buy 3 different books, you get a 10% discount.

If you buy 4 different books, you get a 20% discount.

If you buy all 5, you get a 25% discount.

Note that if you buy four books, of which 3 are different titles, you get a 10% discount on the 3 that form part of a set, but the fourth book still costs $8.
*/

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
