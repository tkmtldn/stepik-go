package main

// 3.2 Преобразование типов данных
// https://stepik.org/lesson/348563/step/14?unit=332364
// незакончено

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var s,ans string
	var first_v, second_v float64
	for i:=0; i<3; i++ {
		fmt.Scan(&s)
		ans += s
	}
	ans = strings.Replace(ans, ";", " ", 1)
	testArray := strings.Fields(ans)
	first := strings.Replace(testArray[0], ",", ".",1)
	second := strings.Replace(testArray[1], ",", ".",1)

	if s, err := strconv.ParseFloat(first, 64); err == nil {
		first_v = s
	}
	if s, err := strconv.ParseFloat(second, 64); err == nil {
		second_v = s
	}

	fmt.Printf("%.4f", first_v/second_v)
}