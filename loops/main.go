package main

import "log"

func main() {
	// counter()
	// rangeSlice()
	// rangeMap()
	// rangeString()
	rangeListObject()
}

func rangeListObject() {
	type User struct {
		FirstName string
		LastName  string
		Email     string
		Age       int
	}

	var users []User
	users = append(users, User{"John", "Smith", "john@smith.com", 30})
	users = append(users, User{"Mary", "Jones", "mary@jones.com", 20})
	users = append(users, User{"Sally", "Brown", "sally@brown.com", 45})
	users = append(users, User{"Alex", "Anderson", "alex@anderson", 17})

	for _, user := range users {
		log.Println(user.FirstName, user.LastName, user.Email, user.Age)
	}
}

func rangeString() {
	var firstLine = "Once upon a midnight dreary" // Edgar Allan Poe

	for i, l := range firstLine {
		log.Println(i, ":", l)
		// 0 : 79, 1 : 100, .......
	}
}

func rangeMap() {
	animals := make(map[string]string)
	animals["dog"] = "Fido"
	animals["cat"] = "Fluffy"

	for animalType, animal := range animals {
		log.Println(animalType, animal)
	}
}

func rangeSlice() {
	animals := []string{"dog", "fish", "horse", "cow", "cat"}

	for _, animal := range animals {
		log.Println(animal)
	}
}

func counter() {
	for i := 0; i <= 5; i++ {
		log.Println(i)
	}
}
