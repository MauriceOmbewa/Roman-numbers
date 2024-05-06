package main

import (
	"fmt"
)

func intToRoman(num int) string {
	var roman string = ""
	var numbers = []int{1, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000}
	var romans = []string{"I", "IV", "V", "IX", "X", "XL", "L", "XC", "C", "CD", "D", "CM", "M"}
	var index = len(romans) - 1

	for num > 0 {
		for numbers[index] <= num { // looping numbers from 1000 going downwords(right to left)
			roman += romans[index]
			num -= numbers[index]
		}
		index -= 1 // reducing the index by 1 since we are looping from left to right
	}

	return roman
}

func main() {
	fmt.Println(intToRoman(354))
	fmt.Println(intToRoman(45))
	fmt.Println(intToRoman(1238))
	fmt.Println(intToRoman(3259))
}
