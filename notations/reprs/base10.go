package notations

import "fmt"

// represents given number in base 10 positional notation
func ReprBase10(num int) string {
	var repr string

	str := fmt.Sprint(num)
	len := len(str)

	for _, d := range str {
		if len > -1 {
			repr += fmt.Sprintf("(%v x 10^%v)", string(d), len)
		} else {
			break
		}

		len--

		if len > 0 {
			repr += "+"
		}
	}
	return repr
}
