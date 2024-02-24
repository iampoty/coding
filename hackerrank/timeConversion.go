package solution

import (
	"fmt"
	"strconv"
	"strings"
)

func timeConversion(s string) string {
	// Write your code here

	//12:40:22AM
	//00:40:22
	//s = "12:40:22AM"
	meridiem := s[len(s)-2:]
	tmp := strings.Split(s[:len(s)-2], ":")
	h, _ := strconv.Atoi(tmp[0])
	m, _ := strconv.Atoi(tmp[1])
	ss, _ := strconv.Atoi(tmp[2])

	fmt.Printf("meridiem %v / h:%v / m:%v / ss:%v\n", meridiem, h, m, ss)

	if meridiem == "AM" && h == 12 {
		return fmt.Sprintf("00:%02d:%02d", m, ss)
	} else if meridiem == "PM" && h == 12 {
		return fmt.Sprintf("12:%02d:%02d", m, ss)
	} else if meridiem == "PM" {
		return fmt.Sprintf("%02d:%02d:%02d", h+12, m, ss)
	}

	// return "19:05:45"
	return fmt.Sprintf("%02d:%02d:%02d", h, m, ss)
}
