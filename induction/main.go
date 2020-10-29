package main

import (
	"fmt"
)

// I1. Prove P(1)
// I2. k = n?
// I3. Prove P(k + 1)
// I4. Increase k

func main() {
	var state string = "I1"
	var k int
	var n int = 3
	P := []int{1,2,3}
	for {
		switch state {
		case "I1":
			fmt.Println("Prove P(1)")
			k = 1
			state = "I2"
		case "I2":
			if k == n {
				state = "done"
			} else {
				state = "I3"
			}
		case "I3":
			fmt.Println(fmt.Sprintf("Prove P(%d + 1)", k))
			state = "I4"
		case "I4":
			k += 1
			state = "I2"
		case "done":
			return
		}

	}

}
