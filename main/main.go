package main

import (
	"fmt"

	"github.com/gobuffalo/flect"
)

func main() {
	s := "Hello, World!"
	d := flect.Dasherize(s)
	fmt.Println(s, d)
}
