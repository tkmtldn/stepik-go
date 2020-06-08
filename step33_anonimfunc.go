package main

// 3.3 Анонимные функции
// https://stepik.org/lesson/349644/step/7?thread=solutions&unit=333499

import (
	"fmt"
	"strconv"
)

func main() {
	fn := func(r uint) uint{
		s := strconv.FormatUint(uint64(r), 10)
		var new_s string
		for _, ui8 := range s {
			if ui8%2==0 {
				new_s += string(ui8)
			}
		}
		s1, _ := strconv.ParseUint(new_s, 10, 0)
		if s1 == 0 {
			return 100
		}
		return uint(s1)
	}
	fmt.Println(fn(727178))
	fmt.Println(fn(100))
}
