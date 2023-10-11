package main

import "fmt"

func main() {
	var array = []int{2, 1, 3, 100, 99}
	var length = len(array)

	var maxValue = max(array, length)

	fmt.Printf("Max Value :%d\n", maxValue)

}

func max(arrays []int, length int) int {

	for i := 0; i < length-1; i++ {
		if arrays[i] > arrays[i+1] {
			var temp = arrays[i]
			arrays[i] = arrays[i+1]
			arrays[i+1] = temp
		}
	}

	return arrays[length-1]

}
