package frameworks

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
)

func CheckJSON(path, val, filterN string) error {
	filter := func(data map[string]interface{}) bool {
		if name, ok := data[filterN]; ok {
			return name == val
		}
		return false
	}
	res := readJSON(path, filter)
	if fmt.Sprint(res) != "[]" {
		return fmt.Errorf("This Name token pick another name Blueprint struct is Name %s Libs %s Struct %s", res[0]["name"], res[0]["libs"], res[0]["struct"])
	}
	return nil
}

func searchJSON(filterN, path, val string) []map[string]interface{} {
	filter := func(data map[string]interface{}) bool {
		if name, ok := data[filterN]; ok {
			return name == val
		}
		return false
	}
	return readJSON(path, filter)
}

func readJSON(path string, filter func(map[string]interface{}) bool) []map[string]interface{}{
	datas := []map[string]interface{}{}

	file, _ := os.ReadFile(path)
	// if err != nil {
	// 	return nil, err
	// }
	json.Unmarshal(file, &datas)

	filteredData := []map[string]interface{}{}

	for _, data := range datas {
		if filter(data) {
			filteredData = append(filteredData, data)
		}
	}

	return filteredData
}

func WritJSON[t any](val t, filePath string) error {

	var data []t
	databayte, err := os.ReadFile(filePath)
	if err != nil {
		if !os.IsNotExist(err) {
			return fmt.Errorf("failed to read file test.json: %w", err)
		}
		return err
	}
	if err := json.Unmarshal(databayte, &data); err != nil {
		return err
	}

	fmt.Print(data)
	data = append(data, val)
	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(data)

	if err := os.WriteFile(filePath, reqBodyBytes.Bytes(), os.ModePerm); err != nil {

		return err
	}
	return nil
}
