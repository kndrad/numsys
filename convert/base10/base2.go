package base10

import "fmt"

// Does not calculate float numbers!
// Returns number (eg. decimal) to binary number as a string.
func ToBase2(num int) (b string) {
	fmt.Printf("• %v\n", num)
	factor := 2

	digits := []int{}
	remainder := num % factor
	// append first remainder to digits
	digits = append([]int{remainder}, digits...)

	quotient := num / factor

	step := 1
	printEquastionStep(step, num, factor, quotient, remainder)

	// loop breaks when quotient is 0
	for {
		step++
		remainder = quotient % factor
		digits = append([]int{remainder}, digits...) // append next remainder at the start of digits slice

		result := quotient / factor
		printEquastionStep(step, quotient, factor, result, remainder)

		if result == 0 {
			break
		}

		// swap so quotient is now result
		quotient, result = result, 0
	}

	for _, digit := range digits {
		b += fmt.Sprintf("%v", digit)
	}
	return
}

func printEquastionStep(step, num, factor, result, remainder int) {
	fmt.Printf("%v) %v / %v = %v, reszta %v\n", step, num, factor, result, remainder)
}
