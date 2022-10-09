package main

import (
	"fmt"
)

var _ = fmt.Print


func main() {
	fmt.Println(reprbase10posn(254))
}


// represents given number in base 10 positional notation
func reprbase10posn(num int) string {
	var repr string

	str := fmt.Sprint(num)
	len := len(str)

	for _ , d := range str {
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


