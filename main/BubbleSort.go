package main

import "fmt"

func main() {

	var toSort = []int{8, 4, 2, 1, 10, 3}
	var length = len(toSort)

	sort(toSort, length)

	for i := 0; i < length; i++ {
		fmt.Println(toSort[i])
	}
}

func sort(array []int, length int) {

	for i := 0; i < length-1; i++ {

		for j := 0; j < length-1-i; j++ {

			if array[j] > array[j+1] {
				var temp = array[j]
				array[j] = array[j+1]
				array[j+1] = temp
			}
		}
	}
}
