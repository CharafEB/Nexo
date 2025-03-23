package frameworks

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/Nexo/cmd/midel"
)

func CheckJSON(val string) error {
	filter := func(data map[string]interface{}) bool {
		if name, ok := data["name"]; ok {
			return name == val
		}
		return false
	}
	if fmt.Sprint(readJSON(filter)) != "[]" {
		log.Fatalf("This Name token pick another name Blueprint struct is Name %s Libs %s Struct %s", readJSON(filter)[0]["name"], readJSON(filter)[0]["libs"], readJSON(filter)[0]["struct"])
	}
	return nil
}

func searchJSON(val string) []map[string]interface{} {
	filter := func(data map[string]interface{}) bool {
		if name, ok := data["name"]; ok {
			return name == val
		}
		return false
	}
	return readJSON(filter)
}

func readJSON(filter func(map[string]interface{}) bool) []map[string]interface{} {
	datas := []map[string]interface{}{}

	file, _ := os.ReadFile("../main/test.json")
	json.Unmarshal(file, &datas)

	filteredData := []map[string]interface{}{}

	for _, data := range datas {
		if filter(data) {
			filteredData = append(filteredData, data)
		}
	}

	return filteredData
}

func WritJSON(val midel.Blueprints , file string) error {

	var data []midel.Blueprints
	databayte, err := os.ReadFile(file)
	if err != nil {
		// if _, err := os.Create(file); err != nil {
		// 	return err
		// }
		return err
	}
	if err := json.Unmarshal(databayte, &data); err != nil {
		return err
	}

	fmt.Print(data)
	data = append(data, val)
	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(data)

	if err := os.WriteFile(file, reqBodyBytes.Bytes(), os.ModePerm); err != nil {

		return err
	}
	return nil
}
