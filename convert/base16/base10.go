package base16

import (
	"fmt"
	"math"
	"strconv"
	"unicode"
)

var HexDecimalEquivalents map[string]string = map[string]string{
	"0": "0",
	"1": "1",
	"2": "2",
	"3": "3",
	"4": "4",
	"5": "5",
	"6": "6",
	"7": "7",
	"8": "8",
	"9": "9",
	"A": "10",
	"B": "11",
	"C": "12",
	"D": "13",
	"E": "14",
	"F": "15",
}

func ToBase10(hexdecimal string) (result int) {
	var digits []int

	for _, r := range hexdecimal {
		var key string
		if unicode.IsDigit(r) {
			key = fmt.Sprint(r - '0')
		} else {
			key = string(r)
		}

		digit, err := strconv.Atoi(HexDecimalEquivalents[key])
		if err != nil {
			panic(err)
		}
		digits = append(digits, digit)
	}

	fmt.Println(digits)
	var factor float64 = 16


	j := len(digits) - 1
	for _, digit := range digits {
		result += digit * int(math.Pow(factor, float64(j)))
		j--

		if j < 0 {
			break
		}
	}

	fmt.Println(result)
	return
}
