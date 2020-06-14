package main

// 3.7 Работа с временем
// https://stepik.org/lesson/359395/step/7?unit=343626

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

func main() {
	format := "02.01.2006 15:04:05"
	in, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil && err != io.EOF {
			fmt.Print(err)
		}
	in = strings.TrimRight(in, "\n")
	dates := strings.Split(in, ",")
	t1, _ := time.Parse(format, dates[0])
	t2, _ := time.Parse(format, dates[1])
	if t1.After(t2) {
		fmt.Println(t1.Sub(t2))
	} else {
		fmt.Println(t2.Sub(t1))
	}
}