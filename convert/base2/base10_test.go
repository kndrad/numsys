package base2_test

import (
	"base2"
	"testing"
)

func TestBase2Base10Conversion(t *testing.T) {
	testBinary := func(binary string, expected int) {
		if num, expected := base2.ToBase10(binary), expected; num != expected {
			t.Fatalf("%v must be %v", num, expected)
		}
	}

	testBinary("11001010", 202)
	testBinary("1101001001", 841)
	testBinary("1101001", 105)
	testBinary("11", 3)
	testBinary("111", 7)
}
