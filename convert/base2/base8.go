package base2

import (
	"fmt"
	"represent"
	"strconv"
)

/*
Converts binary number eg. 10010110 to octal number. In this case it is 226.
Only works for binary numbers without point!
*/
func ToBase8(binary string) int {
	fmt.Println("z systemu binarnego na dziesietny...")
	decim := ToBase10(binary)
	fmt.Println("z systemu dziesietnego na ósemkowy...")

	factor := 8
	quotient, lsd := decim/factor, decim%factor
	digits := []int{lsd}

	step := 1

	represent.EquastionSteps(step, decim, factor, quotient, lsd)
	for {
		step++
		result, remainder := quotient/factor, quotient%factor
		digits = append([]int{remainder}, digits...)

		represent.EquastionSteps(step, quotient, factor, result, remainder)

		if result == 0 {
			break
		}

		quotient, result = result, 0
	}

	var joinedDigits string
	for _, d := range digits {
		joinedDigits += fmt.Sprint(d)
	}
	num, err := strconv.Atoi(joinedDigits)

	if err != nil {
		panic(err)
	}

	fmt.Println("Wynik:", num)
	return num
}
