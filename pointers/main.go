package main

import "log"

func main() {
	var myString string
	myString = "Green"

	log.Println("myString is set to", myString)
	changeUsingPointer(&myString)
	log.Println("after func call myString is set to", myString)
}

func changeUsingPointer(s *string) {
	log.Println("s is set to", s)
	// s is set to 0xc000028080
	log.Println("s is set to", *s)
	// s is set to Green
	newValue := "Red"
	*s = newValue
}
