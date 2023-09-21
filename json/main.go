package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	HairColor string `json:"hair_color"`
	HasDog    bool   `json:"has_dog"`
}

func main() {
	myJson := `
[
	{
		"first_name": "Clark",
		"last_name": "Kent",
		"hair_color": "black",
		"has_dog": true
	},
	{
		"first_name": "Bruce",
		"last_name": "Wayne",
		"hair_color": "black",
		"has_dog": false
	}
]`

	var unmarshalled []Person

	err := json.Unmarshal([]byte(myJson), &unmarshalled)

	if err != nil {
		log.Println("Error unmarshalling json", err)
	}

	log.Printf("unmarshalled: %v", unmarshalled)

	fmt.Println()

	for _, person := range unmarshalled {
		fmt.Printf("First Name: %s\n", person.FirstName)
		fmt.Printf("Last Name: %s\n", person.LastName)
		fmt.Printf("Hair Color: %s\n", person.HairColor)
		fmt.Printf("Has Dog: %t\n", person.HasDog)
	}

	fmt.Println()

	// write json from a struct

	var mySlice []Person

	var m1 Person
	m1.FirstName = "Wally"
	m1.LastName = "West"
	m1.HairColor = "red"
	m1.HasDog = false

	mySlice = append(mySlice, m1)

	var m2 Person
	m2.FirstName = "Diana"
	m2.LastName = "Prince"
	m2.HairColor = "black"
	m2.HasDog = false

	mySlice = append(mySlice, m2)

	newJson, err := json.MarshalIndent(mySlice, "", "    ")

	if err != nil {
		log.Println("error mashalling", err)
	}

	fmt.Println(string(newJson))
	/*
		[
		    {
		        "first_name": "Wally",
		        "last_name": "West",
		        "hair_color": "red",
		        "has_dog": false
		    },
		    {
		        "first_name": "Diana",
		        "last_name": "Prince",
		        "hair_color": "black",
		        "has_dog": false
		    }
		]
	*/
}
