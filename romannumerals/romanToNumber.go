package main

import "fmt"

var romanNum = map[int]string{
	1000: "M", 2000: "MM", 3000: "MMM",
	100: "C", 200: "CC", 300: "CCC", 400: "CD", 500: "D", 600: "DC", 700: "DCC", 800: "DCCC", 900: "CM",
	10: "X", 20: "XX", 30: "XXX", 40: "XL", 50: "L", 60: "LX", 70: "LXX", 80: "LXXX", 90: "XC",
	1: "I", 2: "II", 3: "III", 4: "IV", 5: "V", 6: "VI", 7: "VII", 8: "VIII", 9: "IX",
}

// Write a function to convert from normal numbers to Roman Numerals.
func main() {

	var number int

	fmt.Printf("Enter Numeric value 1 to 3999: ")
	fmt.Scan(&number)

	if number <= 0 || number >= 4000 {
		fmt.Printf("Only numbers greater than 0 and less than 4000 can be converted into roman numerals \n")
	} else {
		romanValue := ""

		k := number / 1000
		romanValue += romanNum[k*1000]

		c := (number / 100) % 10
		romanValue += romanNum[c*100]

		d := (number / 10) % 10
		romanValue += romanNum[d*10]

		u := (number / 1) % 10
		romanValue += romanNum[u*1]

		fmt.Println(romanValue)
	}
}
