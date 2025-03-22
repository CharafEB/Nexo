package main

import (
	"bytes"
	"encoding/json"
	"fmt"

	"os"
)

type Blueprints struct {
	Name string `json:"name"`
	Libs string `json:"libs"`
}

func main() {
	// val := Blueprints{
	// 	Name: "test2",
	// 	Libs: "nodemon",
	// }
	filter := func(data map[string]interface{}) bool {
		if name, ok := data["name"]; ok {
			return name == "achraf"
		}
		return false
	}
	//WritJSON(val)
	fmt.Println(readJSON(filter)[1]["name"])
}
func readJSON(filter func(map[string]interface{}) bool) []map[string]interface{} {
	datas := []map[string]interface{}{}

	file, _ := os.ReadFile("test.json")
	json.Unmarshal(file, &datas)

	filteredData := []map[string]interface{}{}

	for _, data := range datas {
		// Do some filtering
		if filter(data) {
			filteredData = append(filteredData, data)
		}
	}

	return filteredData
}

func WritJSON(val Blueprints) error {
	var data []Blueprints
	databayte, err := os.ReadFile("test.json")
	if err != nil {
		return err
	}
	if err := json.Unmarshal(databayte, &data); err != nil {
		return err
	}

	fmt.Print(data)
	data = append(data, val)
	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(data)

	if err := os.WriteFile("test.json", reqBodyBytes.Bytes(), os.ModePerm); err != nil {
		return err
	}
	return nil
}
