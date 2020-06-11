package main

// 3.6 JSON
// https://stepik.org/lesson/353243/step/6?thread=solutions&unit=337227

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type StudentStruct struct {
	LastName    string `json:",omitempty"`
	FirstName   string `json:",omitempty"`
	MiddleName  string `json:",omitempty"`
	Birthday    string `json:",omitempty"`
	Address     string `json:",omitempty"`
	Phone       string `json:",omitempty"`
	Rating		[]int
}

type GroupStruct struct {
	ID        int `json:",omitempty"`
	Number    string `json:",omitempty"`
	Year      int `json:",omitempty"`
	Students	[]StudentStruct
}

type AnswerStruct struct {
	Average		float64
}

func main() {
	indata, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println(err)
		return
	}
	//indata := `{
	//	"Students":[
	//	{
	//		"Name":"Bob",
	//		"Rating":[1,2,3]
	//	},
	//	{
	//		"Name":"Alice",
	//		"Rating":[4,5]
	//	}
	//	]
	//}`

	var myGroupData GroupStruct
	json.Unmarshal([]byte(indata), &myGroupData)

	totalMarks := 0
	for _, student := range myGroupData.Students {
		totalMarks += len(student.Rating)
	}
	result := float64(totalMarks)/float64(len(myGroupData.Students))

	s := AnswerStruct{
		Average:   result,
	}

	data, err := json.MarshalIndent(s, "", "    ")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%s", data)
}