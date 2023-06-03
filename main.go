package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

// Person is the struct of our JSONL Objects
type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"address"`
}

func peopleFromFile(file io.Reader) ([]Person, error) {
	var people []Person

	decoder := json.NewDecoder(file)
	// Each iteration will be the next JSON Object
	for decoder.More() {
		var person Person

		if err := decoder.Decode(&person); err != nil {
			return nil, err
		}

		people = append(people, person)
	}

	return people, nil
}

func main() {
	file, err := os.Open("people.jsonl")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// We can pass in any file which implements io.Reader e.g. multipart.File or os.File
	people, err := peopleFromFile(file)
	if err != nil {
		log.Fatal(err)
	}

	for _, person := range people {
		fmt.Println(person.Name, person.Age, person.Address)
	}
}
