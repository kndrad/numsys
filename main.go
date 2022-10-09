package main

import (
	"base10"
	"fmt"
	"represent"
)

// Delete those later
var _ = base10.ToBase2
var _ = fmt.Print
var _ = represent.InBase10

func main() {
	fmt.Println(base10.ToBase2(25867))
}
