package main

// 3.7 Работа с временем
// https://stepik.org/lesson/359395/step/4?unit=343626

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

const (
	format = "2006-01-02 15:04:05"
)

func main() {
	data, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil && err != io.EOF {
		fmt.Print(err)
	}
	data = strings.TrimRight(data, "\n")

	calendar, err := time.Parse(format, data)
	if calendar.Hour() > 13 {
		calendar = calendar.AddDate(0, 0, 1)
	}

	wr := bufio.NewWriter(os.Stdout)
	wr.WriteString(calendar.Format(format))
	wr.Flush()

}