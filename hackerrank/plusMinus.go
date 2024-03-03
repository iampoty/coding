package solution

import "fmt"

/*
https://www.hackerrank.com/challenges/one-month-preparation-kit-plus-minus
*/

func plusMinus(arr []int32) []string {
	n := len(arr)
	positive := 0.0
	negative := 0.0
	zero := 0.0

	for i := 0; i < n; i++ {
		if arr[i] > 0 {
			positive++
		} else if arr[i] < 0 {
			negative++
		} else {
			zero++
		}
	}

	// fmt.Printf("%06f\n", positive/float64(n))
	// fmt.Printf("%06f\n", negative/float64(n))
	// fmt.Printf("%06f\n", zero/float64(n))
	return []string{fmt.Sprintf("%06f", positive/float64(n)), fmt.Sprintf("%06f", negative/float64(n)), fmt.Sprintf("%06f", zero/float64(n))}
}
