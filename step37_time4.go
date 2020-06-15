package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// 3.7 Работа с временем
// https://stepik.org/lesson/359395/step/8?unit=343626

const now = 1589570165

func main() {
	scan := bufio.NewScanner(os.Stdin)
	if scan.Scan() {
		data := strings.Split(scan.Text(), " ")
		tMin, _ := strconv.Atoi(data[0])
		tSec, _ := strconv.Atoi(data[2])
		tPlus := now + tMin*60 + tSec
		t := time.Unix(int64(tPlus), 0)
		t = t.UTC()
		fmt.Println(t.Format(time.UnixDate))
	}
}
