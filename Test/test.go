package main

import (
	//"bytes"
	//"encoding/json"
	"fmt"
	"log"
	"path/filepath"

	"os"
	//"github.com/Nexo/cmd/midel"
)

var (
	structF []string
	folder  []string
)

func main() {
	// test1 := midel.Blueprints{
	// 	Name: "ForCheckFunc",
	// 	Libs: "express",
	// 	StructType: "Free",
	// }
	// CheckJSON("ForCheckFunc", test1)
	// filter := func(data map[string]interface{}) bool {
	// 	if name, ok := data["name"]; ok {
	// 		return name == "achraf1"
	// 	}
	// 	return false
	// }
	// WritJSON(test1)
	// fmt.Println(readJSON(filter))
	// if fmt.Sprint(readJSON(filter)) == "[]"{
	// 	fmt.Print("empty")
	// }

	if err := BuilsStruct(&structF, &folder, "./"); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("structF:%s , folder:%s", structF, folder)

}

func BuilsStruct(valS, valF *[]string, path string) error {
	err := filepath.WalkDir(path, func(fullPath string, entry os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if entry.IsDir() {
			*valF = append(*valF, fullPath)
		} else {
			*valS = append(*valS, fullPath)
		}

		return nil
	})

	return err
}

// func CheckJSON(val string , blue midel.Blueprints) error {
// 	filter := func(data map[string]interface{}) bool {
// 		if name, ok := data["name"]; ok {
// 			return name == val
// 		}
// 		return false
// 	}
// 	if fmt.Sprint(readJSON(filter)) != "[]"{
// 		log.Fatal("This name tooken pick another name")
// 	}
// 	WritJSON(blue)
// 	return nil
// }

// func readJSON(filter func(map[string]interface{}) bool) []map[string]interface{} {
// 	datas := []map[string]interface{}{}

// 	file, _ := os.ReadFile("test1.json")
// 	json.Unmarshal(file, &datas)

// 	filteredData := []map[string]interface{}{}

// 	for _, data := range datas {
// 		// Do some filtering
// 		if filter(data) {
// 			filteredData = append(filteredData, data)
// 		}
// 	}

// 	return filteredData
// }

// func WritJSON(val midel.Blueprints) error {
// 	var data []midel.Blueprints
// 	databayte, err := os.ReadFile("test1.json")
// 	if err != nil {
// 		return err
// 	}
// 	if err := json.Unmarshal(databayte, &data); err != nil {
// 		return err
// 	}

// 	fmt.Print(data)
// 	data = append(data, val)
// 	reqBodyBytes := new(bytes.Buffer)
// 	json.NewEncoder(reqBodyBytes).Encode(data)

// 	if err := os.WriteFile("test1.json", reqBodyBytes.Bytes(), os.ModePerm); err != nil {
// 		return err
// 	}
// 	return nil
// }
