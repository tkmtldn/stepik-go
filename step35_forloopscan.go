package main

// 3.5 Работа с файлами
// https://stepik.org/lesson/351892/step/9?unit=335849


import (
	"bufio"
	"io"
	"os"
	"strconv"
)

func main() {
	count := 0
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		if (input.Err() == io.EOF) || (input.Text() == "") {
			break
		}
		num, _ := strconv.Atoi(input.Text())
		count += num
		}
	io.WriteString(os.Stdout, strconv.Itoa(count))
}