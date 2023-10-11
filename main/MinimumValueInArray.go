package main

import "fmt"

func main() {

	var input = []int{10, 8, 4, 20, 19, -10}

	var length = len(input)

	var minIndex = 0 // assuming min value at the index

	for i := 0; i < length; i++ {

		if input[minIndex] > input[i] {
			minIndex = i
		}
	}

	fmt.Println("Min Value:", input[minIndex])

}
