package base8

import "fmt"

var equivalents map[string]string = map[string]string{
	"0": "000",
	"1": "001",
	"2": "010",
	"3": "011",
	"4": "100",
	"5": "101",
	"6": "110",
	"7": "111",
}

// Converts from octal to binary
func ToBase2(num int) (binary string) {
	fmt.Println("•", num)
	for _, digit := range fmt.Sprint(num) {
		binaryEquivalent := equivalents[fmt.Sprint(digit-'0')]
		fmt.Println(digit - '0', "binarnie:", binaryEquivalent)
		binary += binaryEquivalent
	}
	fmt.Println("Wynik połączenia binarnych liczb:", binary)
	fmt.Println()
	return
}
