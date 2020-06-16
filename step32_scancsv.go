package main

// 3.2 Преобразование типов данных
// https://stepik.org/lesson/348563/step/14?unit=332364
// незакончено

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	in, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil && err != io.EOF {
		fmt.Print(err)
	}
	var first_v, second_v float64
	new_s := strings.Replace(in, " ", "", -1)
	new_s = strings.Replace(new_s, ";", " ", 1)
	new_s = strings.Replace(new_s, ",", ".", -1)

	testArray := strings.Fields(new_s)

	if s, err := strconv.ParseFloat(testArray[0], 64); err == nil {
		first_v = s
	}
	if s, err := strconv.ParseFloat(testArray[1], 64); err == nil {
		second_v = s
	}

	fmt.Printf("%.4f", first_v/second_v)
}