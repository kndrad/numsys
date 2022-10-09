package base10

import "fmt"

// Does not calculate float numbers!
// Returns number (eg. decimal) to binary number as a string.
func ToBase2(num int) (b string) {
	factor := 2

	digits := []int{}
	remainder := num % factor
	// append first remainder to digits
	digits = append([]int{remainder}, digits...)

	quotient := num / factor

	step := 1
	fmt.Printf("%v) %v / %v = %v, remainder (reszta) %v\n", step, num, factor, quotient, remainder)

	// loop breaks when quotient is 0
	for {
		remainder = quotient % factor
		digits = append([]int{remainder}, digits...) // append next remainder at the start of digits slice

		quotient = quotient / factor
		fmt.Printf("%v) %v / %v = %v, remainder (reszta) %v\n", step, num, factor, quotient, remainder)
		
		if quotient == 0 {
			break
		}
		step++
	}

	for _, digit := range digits {
		b += fmt.Sprintf("%v", digit)
	}
	return
}
