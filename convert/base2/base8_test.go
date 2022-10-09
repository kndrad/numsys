package base2_test

import (
	"testing"
	"base2"
)

func TestBase2Base8Conversion(t *testing.T) {
	testConvert := func(binary string, expected int) {
		if num := base2.ToBase8(binary); num != expected {
			t.Fatalf("%v result does not match expected num %v", num, expected)
		}
	}

	testConvert("10010110", 226)
	testConvert("111101010", 752)
	testConvert("1001001", 111)
	testConvert("101000101", 505)
}