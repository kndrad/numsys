package represent

import (
	"testing"
)

// Testing '1255' as an input should return (1 x 10^4)+(2 x 10^3)+(5 x 10^2)+(5 x 10^1) string
func TestBase10Representation(t *testing.T) {
	var num uint = 1255
	result := InBase10(num)
	expected := "(1 x 10^4)+(2 x 10^3)+(5 x 10^2)+(5 x 10^1)"

	if result != expected {
		t.Fatalf("%v does not match %v", result, expected)
	}
}
