package main

import (
	"fmt"
	"strings"
)

// 3.4 Интерфейсы
// https://stepik.org/lesson/350788/step/13?unit=334666


type Batareika struct {
	Charge int
}

func (b Batareika) String() string {
	f_string := "[" + strings.Repeat(" ", 10-b.Charge) + strings.Repeat("X", b.Charge) + "]"
	return f_string
}

func main(){
	var s string
	fmt.Scan(&s)
	fil := strings.Count(s, "1")
	batteryForTest := Batareika{
		Charge: fil,
	}
	fmt.Println(batteryForTest)
}
