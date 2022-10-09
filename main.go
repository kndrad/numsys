package main

import (
	"base10"
	"fmt"
)

func main() {
	fmt.Println("Wynik:", base10.ToBase2(382))
	fmt.Printf("\n\n")
	fmt.Println("Wynik:", base10.ToBase2(2145))
	fmt.Printf("\n\n")
	fmt.Println("Wynik:", base10.ToBase2(0))
	fmt.Printf("\n\n")
	fmt.Println("Wynik:", base10.ToBase2(453))

}
