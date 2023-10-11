package main

import "fmt"

func main() {
	var scores = []int{10, 20, 30, 40, 50}
	var length = len(scores)
	for i := 0; i < length; i++ {
		fmt.Println(scores[i])
	}
}
