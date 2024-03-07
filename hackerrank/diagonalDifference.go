package solution

import (
	"math"
)

func diagonalDifference(arr [][]int32) int32 {

	// rIdx := 0
	// lIdx := len(arr) - 1

	var rVal int32
	var lVal int32

	// fmt.Printf("rIdx: %d\n", rIdx)
	// fmt.Printf("lIdx: %d\n", lIdx)

	for i := 0; i < len(arr); i++ {
		// fmt.Printf("rSide[%d][%d]: %d / lSide[%d][%d]: %d\n", i, rIdx, arr[i][rIdx], i, lIdx, arr[i][lIdx])
		j := len(arr) - 1 - i
		rVal += arr[i][i]
		lVal += arr[i][j]

		// rIdx++
		// lIdx--
	}

	// fmt.Printf("rVal: %d\n", rVal)
	// fmt.Printf("lVal: %d\n", lVal)
	// fmt.Printf("rVal - lVal: %d\n", int32(math.Abs(float64(rVal)-float64(lVal))))

	return int32(math.Abs(float64(rVal) - float64(lVal)))
}
