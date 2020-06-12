package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// 3.6 JSON
// https://stepik.org/lesson/353243/step/9?unit=337227


type Properties struct {
	GlobalID    int `json:"global_id,omitempty"`
	SysObjID	string `json:"system_object_id,omitempty"`
	SignDate	string `json:"signature_date,omitempty"`
	Razdel		string `json:"Razdel,omitempty"`
	Kod			string `json:"Kod,omitempty"`
	Name		string `json:"Name,omitempty"`
	Idx			string `json:"Idx,omitempty"`
}

func main() {
	jsonFile, err := os.Open(`./_examples2/data-20190514T0100.json`)
	defer jsonFile.Close()
	if err != nil {
		fmt.Println(err)
	}

	var myJData []Properties
	var total int

	//var urlData = "https://raw.githubusercontent.com/semyon-dev/stepik-go/master/work_with_json/data-20190514T0100.json"
	//resp, _ := http.Get(urlData)
	//r := json.NewDecoder(resp.Body)
	//r.Decode(&myJData)

	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &myJData)

	for _,v := range myJData {
		total += v.GlobalID
	}
	fmt.Println(total)

}