package main

// 3.1 Отображения (map)
// https://stepik.org/lesson/345543/step/6?unit=329288

import "fmt"

func main() {
	groupCity := map[int][]string{
		10:   []string{"ДесятьТысяч", "ДвадцатьТысяч", "СорокТысяч"}, // города с населением 10-99 тыс. человек
		100:  []string{"СтоТысяч", "ДвестиТысяч"}, // города с населением 100-999 тыс. человек
		1000: []string{"ПочтиМосква", "Москва"}, // города с населением 1000 тыс. человек и более
	}
	//fmt.Println(groupCity)
	cityPopulation := map[string]int{
		"ДесятьТысяч": 10,
		"ДвадцатьТысяч": 20,
		"СтоТысяч":   100,
		"ДвестиТысяч":  200,
		"Москва": 8000,
	}

	for _, city := range groupCity[1000] {
		delete(cityPopulation, city)
	}

	for _, city := range groupCity[10] {
		delete(cityPopulation, city)
	}

	fmt.Println(cityPopulation)
}
