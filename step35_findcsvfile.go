package main

// 3.5 Работа с файлами
// https://stepik.org/lesson/351892/step/13?unit=335849

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
)

func walkForCSV(path string, info os.FileInfo, err error) error {
	if !info.IsDir() {
		csvfile, err := os.Open(path)
		defer csvfile.Close()
		if err != nil {
			return nil
		}
		r := csv.NewReader(csvfile)
		data, err := r.ReadAll()
		if err != nil {
			fmt.Printf("Ошибка чтения: %v\n", err)
			return nil
		}
		for i, row := range data {
			if i == 4 {
				fmt.Println(row[2])
				fmt.Println(path)
				os.Exit(0)
			}
		}
	}
	return nil
}


func main() {
	const root = `C:\Users\Anthony\Desktop\task`

	if err := filepath.Walk(root, walkForCSV); err != nil {
		fmt.Printf("Ошибка: %v\n", err)
	}
}
