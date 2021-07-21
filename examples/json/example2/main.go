package main

import (
	"encoding/json"
	"fmt"
	"log"
)

//Student type
type Student struct {
	Name  *string `json:"name"`
	Age   int     `json:"age"`
	City  string
	Major string
}

var jsonStr = `{
   "name": "Dave",
   "age": 37,
   "city": "denver"
}`

func main() {

	var student Student
	err := json.Unmarshal([]byte(jsonStr), &student)
	if err != nil {
		log.Fatalf("json.Unmarshal failed with '%s'\n", err)
	}
	fmt.Printf("Student struct parsed from JSON: %#v\n", student)
	fmt.Printf("Name: %#v\n", *student.Name)

}
