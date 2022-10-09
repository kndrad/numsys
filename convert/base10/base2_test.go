package base10

import "testing"

// '13' number should return '1101'
func TestBase10ToBase2(t *testing.T) {
	expected := "1101"
	if num := 13; ToBase2(num) != expected {
		t.Fatalf("num %v converted to base2 does not match %v", num, expected)
	}

	expected = "110010100001011"
	if num := 25867; ToBase2(num) != expected {
		t.Fatalf("num %v converted to base2 does not match %v", num, expected)
	}
}
