package frameworks

import (
	"bytes"
	"encoding/json"
	"fmt"

	"os"

	"github.com/Nexo/cmd/midel"
)

func searchJSON(val string) {
	filter := func(data map[string]interface{}) bool {
		if name, ok := data["libs"]; ok {
			return name == val
		}
		return false
	}
	fmt.Println(readJSON(filter)[1]["libs"])
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

func WritJSON(val midel.Blueprints) error {
	var data []midel.Blueprints
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
