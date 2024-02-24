package solution

func matchingStrings(strings []string, queries []string) []int32 {

	output := make([]int32, len(queries))
	for i := 0; i < len(queries); i++ {
		for j := 0; j < len(strings); j++ {
			if strings[j] == queries[i] {
				output[i]++
			}
		}
	}

	// res := map[string]int{}
	// for i := 0; i < len(queries); i++ {

	// 	if _, ok := res[queries[i]]; !ok {
	// 		res[queries[i]] = 0
	// 	}

	// 	for j := 0; j < len(strings); j++ {
	// 		if strings[j] == queries[i] {
	// 			res[strings[j]]++
	// 		}
	// 	}
	// }

	// output := []int32{}
	// for i := 0; i < len(queries); i++ {
	// 	for k, v := range res {
	// 		if k == queries[i] {
	// 			// fmt.Printf("[%d][%v] k:%v / v:%v \n", i, queries[i], k, v)
	// 			output = append(output, int32(v))
	// 		}
	// 	}
	// }
	return output
}
