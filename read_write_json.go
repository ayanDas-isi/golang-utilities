package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

//https://medium.com/@irshadhasmat/golang-simple-json-parsing-using-empty-interface-and-without-struct-in-go-language-e56d0e69968
func read_arr_and_print() {
	//Simple Employee JSON which we will parse
	empArray := `[
		{
			"id": 1,
			"name": "Mr. Boss",
			"department": "",
			"designation": "Director"
		},
		{
			"id": 11,
			"name": "Irshad",
			"department": "IT",
			"designation": "Product Manager"
		},
		{
			"id": 12,
			"name": "Pankaj",
			"department": "IT",
			"designation": "Team Lead"
		}
	]`
	fmt.Println(empArray)

	// Declared an empty interface of type Array
	var results []map[string]interface{}

	// Unmarshal or Decode the JSON to the interface.
	json.Unmarshal([]byte(read_file("emp.json")), &results)

	for key, result := range results {

		fmt.Println("Reading Value for Key :", key)
		//Reading each value by its key
		fmt.Println(
			"- Name :", result["name"],
			"- Department :", result["email"],
			"- Designation :", result["age"])
	}
	read_complex_json()
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func read_file(fname string) string {
	file, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatal(err)
		return err.Error()
	}
	return string(file)
}

func write_file(fname string, content string) {

	f, err := os.Create(fname)
	check(err)
	n3, err := f.WriteString(content)
	check(err)
	fmt.Println(n3)
	//f, err := os.Create(fname)
	//check(err)

	//defer f.Close()
}

func read_complex_json() {
	//Simple Employee JSON object which we will parse
	empArray := `[
		{
			"id": 1,
			"name": "Mr. Boss",
			"department": "",
			"designation": "Director",
			"address": {
				"city": "Mumbai",
				"state": "Maharashtra",
				"country": "India"
			}
		},
		{
			"id": 11,
			"name": "Irshad",
			"department": "IT",
			"designation": "Product Manager",
			"address": {
				"city": "Mumbai",
				"state": "Maharashtra",
				"country": "India"
			}
		},
		{
			"id": 12,
			"name": "Pankaj",
			"department": "IT",
			"designation": "Team Lead",
			"address": {
				"city": "Pune",
				"state": "Maharashtra",
				"country": "India"
			}
		}
	]`

	// Declared an empty interface of type Array
	var results []map[string]interface{}

	// Unmarshal or Decode the JSON to the interface.
	json.Unmarshal([]byte(empArray), &results)

	for key, result := range results {
		address := result["address"].(map[string]interface{})
		fmt.Println("Reading Value for Key :", key)
		//Reading each value by its key
		fmt.Println("Id :", result["id"],
			"- Name :", result["name"],
			"- Department :", result["department"],
			"- Designation :", result["designation"])
		fmt.Println("Address :", address["city"], address["state"], address["country"])
	}
}

func read_json_arr(ip string) []map[string]interface{} {
	var results []map[string]interface{}

	// Unmarshal or Decode the JSON to the interface.
	json.Unmarshal([]byte(ip), &results)

	return results
}

func read_json(ip string) map[string]interface{} {
	var results map[string]interface{}

	// Unmarshal or Decode the JSON to the interface.
	json.Unmarshal([]byte(ip), &results)

	return results
}
