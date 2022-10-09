package represent

import "fmt"

// represents given number in base 10 positional notation
func InBase10(num int) (repr string) {
	digits := fmt.Sprint(num)
	len := len(digits)

	for _, digit := range digits {
		if len > -1 {
			repr += fmt.Sprintf("(%v x 10^%v)", string(digit), len)
		} else {
			break
		}

		len--

		if len > 0 {
			repr += "+"
		}
	}
	return
}
