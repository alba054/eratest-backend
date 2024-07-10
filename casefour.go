package main

import "fmt"

func insertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1

		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j = j - 1
		}

		arr[j+1] = key
	}
}

func main() {
	arr := []int{5, 1, 5, 2, 3, 5, 4, 2, 10, 9, 0, 10, 13}
	fmt.Println(arr)
	insertionSort(arr)
	fmt.Println(arr)
}
