package base8_test

import (
	"base8"
	"testing"
)

func TestBase8Base2Conversion(t *testing.T) {
	testConversion := func(num int, expected string) {
		if b := base8.ToBase2(num); b != expected {
			t.Fatalf("%v must be %v", num, expected)
		}
	}
	testConversion(540, "101100000")
	testConversion(345, "011100101")
	testConversion(7462, "111100110010")
	testConversion(123, "001010011")
}
