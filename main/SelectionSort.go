package main

import "fmt"

func main() {

	var input = []int{10, 0, 1, 2, 20, -1}

	var length = len(input)

	selectSort(input, length)

	for i := 0; i < length; i++ {
		fmt.Println(input[i])
	}
}

func selectSort(array []int, len int) {

	var minIndex int

	for i := 0; i < len-1; i++ {
		minIndex = i
		var minValue = array[minIndex]

		for j := i; j < len-1; j++ {
			if minValue > array[j+1] {
				minValue = array[j+1]
				minIndex = j + 1
			}
		}

		if i != minIndex {
			var temp = array[i]
			array[i] = array[minIndex]
			array[minIndex] = temp
		}
	}

}
