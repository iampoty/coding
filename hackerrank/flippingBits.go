package solution

import (
	"fmt"
	"math"
	"strconv"
)

func base2To10(str string) int64 {
	l := len(str)
	result := 0.0
	for i, n := range str {
		number, _ := strconv.Atoi(string(n))
		result += math.Exp2(float64(l-i-1)) * float64(number)
	}
	return int64(result)
}

func flippingBits(n int64) int64 {
	// Write your code here

	u32 := fmt.Sprintf("%032b", n)
	newU32 := ""

	for i := range u32 {
		if u32[i:i+1] == "1" {
			newU32 += "0"
		} else {
			newU32 += "1"
		}
	}

	x := base2To10(newU32)

	return x
}
