package main

// 3.7 Работа с временем
// https://stepik.org/lesson/359395/step/3?unit=343626

import (
	"fmt"
	"time"
)

func main() {
	// "1986-04-16T05:20:00+06:00"
	var indata string
	fmt.Scan(&indata)

	t, err := time.Parse(time.RFC3339, indata)
	if err != nil {
		panic(err)
	}
	fmt.Println(t.Format(time.UnixDate))
}