package main

import "fmt"

func isPalindrome(words string) bool {
	for i := 0; i < len(words)/2; i++ {
		if words[i] != words[len(words)-i-1] {
			return false
		}
	}
	return true
}

func main() {
	var words string
	fmt.Scanln(&words)

	fmt.Println(isPalindrome(words))
}
