package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// Student type
type Student struct {
	fullName string
	Name     string
	Age      int    `json:"age"`
	City     string `json:"city"`
}

func main() {

	student := Student{
		Name: "Dave",
		Age:  37,
		City: "Denver",
	}
	d, err := json.Marshal(&student)
	if err != nil {
		log.Fatalf("json.Marshal failed with '%s'\n", err)
	}
	fmt.Printf("Student in compact JSON: %s\n", string(d))

	d, err = json.MarshalIndent(student, "", "  ")
	if err != nil {
		log.Fatalf("json.MarshalIndent failed with '%s'\n", err)
	}
	fmt.Printf("student in pretty-printed JSON:\n%s\n", string(d))

}
