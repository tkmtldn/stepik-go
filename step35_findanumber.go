package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

// 3.5 Работа с файлами
// https://stepik.org/lesson/351892/step/14?unit=335849

func main() {
	file, err := os.Open("./_examples2/task.data")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	rd := bufio.NewReader(file)
	result := strconv.Itoa(0)
	count := 1
	for {
		line, err := rd.ReadString(';')
		if len(line) == 0 && err != nil {
			if err == io.EOF {
				break
			}
		}
		line = strings.TrimSuffix(line, ";")
		if line == result {
			fmt.Println(count)
			break
		}
		count ++
	}
}
