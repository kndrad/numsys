package base2

import "fmt"

// Converts binary number eg. '10010110' to octal number. In this case it's 226.
func ToBase8(binary string) int {
	decim := ToBase10(binary)
	fmt.Println(decim)
	return 1
}