package main

// 3.4 Интерфейсы
// https://stepik.org/lesson/350788/step/10?unit=334666
import (
"fmt"
)

func main() {
	value1, value2, operation := readTask()
	// исходные данные получаются с помощью этой функции
	// все полученные значения имеют тип пустого интерфейса
	v1, ok := value1.(float64)
	if !ok {
		fmt.Printf("value=%v: %T", value1, value1)
		return
	}
	v2, ok := value2.(float64)
	if !ok {
		fmt.Printf("value=%v: %T", value2, value2)
		return
	}
	v3, ok := operation.(string)
	if !ok {
		fmt.Printf("неизвестная операция")
		return
	}
	switch v3 {
	case "+":
		fmt.Printf("%.4f", v1+v2)
	case "-":
		fmt.Printf("%.4f", v1-v2)
	case "*":
		fmt.Printf("%.4f", v1*v2)
	case "/":
		fmt.Printf("%.4f", v1/v2)
	default:
		fmt.Println("неизвестная операция")
	}
}