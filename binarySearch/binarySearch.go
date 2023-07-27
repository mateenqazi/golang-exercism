package main

import (
	"fmt"
	"sort"
)

// implementation of binary search
func main() {
	list := []int{1, 4, 5, 6, 10, 22, 3, 45, 200, 119, 233}

	valueFind := 45
	sort.Ints(list)

	l, r := 0, len(list)-1

	result := false
	for l <= r {
		mid := (l + r) / 2
		if list[mid] == valueFind {
			result = true
			break
		} else if valueFind > list[mid] {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	fmt.Println("#####################")
	fmt.Println("value exit in array ", result)
	fmt.Println("#####################")
}
