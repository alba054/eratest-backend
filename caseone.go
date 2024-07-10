package main

import (
	"fmt"
	"strconv"
)

func isDivisibleByThree(number int16) bool {
	if number%3 == 0 {
		return true
	}

	return false
}

func endsWithThree(number int16) bool {
	numberInString := strconv.Itoa(int(number))

	if numberInString[len(numberInString)-1] == byte('3') {
		return true
	}

	return false
}

func hateThree(number int16) int16 {
	var k int16 = 0
	var sequence int16 = 0

	for k < number {
		sequence++
		if !endsWithThree(sequence) && !isDivisibleByThree(sequence) {
			k++
		}
	}

	return sequence
}

func main() {
	var t byte

	fmt.Scanln(&t)

	for i := 0; i < int(t); i++ {
		var x int16
		fmt.Scanln(&x)
		fmt.Println(hateThree(int16(x)))
	}
}
