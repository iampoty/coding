package solution

func lonelyinteger(a []int32) int32 {

	tmp := map[int32]int32{}
	for i := 0; i < len(a); i++ {

		if _, ok := tmp[a[i]]; !ok {
			tmp[a[i]] = 1
		} else {
			tmp[a[i]]++
		}
	}

	// fmt.Printf("tmp :%#v \n", tmp)
	for k, v := range tmp {
		if v == 1 {
			return int32(k)
		}
	}

	return 0
}
