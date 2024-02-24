package solution

func miniMaxSum(arr []int32) []int64 {
	// Write your code here
	var min int64 //int64(math.Pow(10, 29)) //1000000000
	var max int64

	var curr int32
	for i := 0; i < 5; i++ {
		var sum int64

		// txt := fmt.Sprintf("%d. Sum everything except , the sum is ",(i+1))

		for j := int32(0); j < 5; j++ {
			if j == curr {
				continue
			}

			sum = sum + int64(arr[j])
			// fmt.Printf("i:%d / j:%d / sum:%d\n",i,j,sum)
		}
		curr++

		// txt = fmt.Sprintf("%s = %d", txt, sum)

		// fmt.Println(txt)
		// fmt.Printf("i:%d / sum:%d / min:%d / max:%d \n",i,sum, min, max)

		// if sum < min {
		//     min = sum
		// }

		if sum > max {
			max = sum
		}
	}

	min = max
	curr = 0
	for i := 0; i < 5; i++ {
		var sum int64
		for j := int32(0); j < 5; j++ {
			if j == curr {
				continue
			}
			sum = sum + int64(arr[j])
		}
		curr++

		if sum < min {
			min = sum
		}

	}

	// fmt.Println("10^9: ",10^9)
	return []int64{min, max}
}
