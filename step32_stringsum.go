package main

// 3.2 Преобразование типов данных
// https://stepik.org/lesson/348563/step/13?unit=332364

import (
	"fmt"
	"strconv"
)

func adding(a,b string) int64{
	var v1, v2 int64
	if s1, err := strconv.ParseInt(convert(a), 10, 64); err == nil {
		v1 = s1
	}
	if s2, err := strconv.ParseInt(convert(b), 10, 64); err == nil {
		v2 = s2
	}
	return  v1 + v2
}

func convert(s string) string {
	ans := ""
	for i:=0; i<len(s); i++ {
		if s[i] >47 && s[i] <58 {
			ans += string(s[i])}
	}
	return ans
}

func main() {
	fmt.Print(adding("%^80", "hhhhh20&&&&nd"))
}