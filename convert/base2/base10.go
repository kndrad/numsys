package base2

import (
	"fmt"
	"math"
)

// Converts binary number to decimal
// Does not work with binary numbers with binary point!
func ToBase10(binary string) (num int) {
	fmt.Printf("• %v\n", binary)

	// for 11001010 msb is 7
	msb := len(binary) - 1
	lsb := 0
	i := msb 

	for {
		digit := int(binary[lsb]-'0')
		x,y := float64(2), float64(msb)
		factor := (int(math.Pow(x, y)))
		result := digit * factor // takes first digit multiples by 2^(msb)

		fmt.Printf("%v * %v^%v = %v\n", digit, x, y, result)
		num += result

		msb--
		if i == lsb { // eg. msb is 7 then if 'i' which is 7 equals lsb which is later incremented
			break
		}
		lsb++
	}
	fmt.Println("Wynik: ", num)
	return
}
