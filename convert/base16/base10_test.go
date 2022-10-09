package base16_test

import (
	"testing"
	"base16"
)

func TestBase16Base10Conversion(t *testing.T) {
	testConvert := func(hexnum string, expected int) {
		if num := base16.ToBase10(hexnum); num != expected {
			t.Fatalf("num %v converted from hex %v does not match %v", num, hexnum, expected)
		}
	}

	testConvert("123", 291)
	testConvert("ABCDEF", 11259375)
}