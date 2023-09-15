package main

import (
	"log"
	"sort"
	"unsafe"
)

type User struct {
	FirstName string
	LastName  string
}

func main() {
	// maps()
	slices()
}

func slices() {
	var mySlice []string

	log.Println(mySlice)
	// []
	log.Println(unsafe.Sizeof(mySlice))
	// 24 bytes

	mySlice = append(mySlice, "Trevor")
	log.Println(mySlice)
	// [Trevor]
	log.Println(unsafe.Sizeof(mySlice))
	// 24 bytes

	mySlice = append(mySlice, "John")
	log.Println(mySlice)
	// [Trevor John]
	log.Println(unsafe.Sizeof(mySlice))
	// 24 bytes

	mySlice = append(mySlice, "Mary")
	log.Println(mySlice)
	// [Trevor John Mary]

	var mySliceInt []int

	log.Println(unsafe.Sizeof(mySliceInt))
	// 24 bytes

	mySliceInt = append(mySliceInt, 2)
	mySliceInt = append(mySliceInt, 1)
	mySliceInt = append(mySliceInt, 3)

	log.Println(unsafe.Sizeof(mySliceInt))
	// 24 bytes

	log.Println(mySliceInt)
	// [2 1 3]

	sort.Ints(mySliceInt)
	log.Println(mySliceInt)
	// [1 2 3]

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	log.Println(numbers)
	// [1 2 3 4 5 6 7 8 9 10]
	log.Println(unsafe.Sizeof(numbers))
	// 24 bytes

	log.Println(numbers[0:2])
	// [1 2]

	log.Println(numbers[6:9])
	// [7 8 9]

	names := []string{"one", "seven", "fish", "cat"}
	log.Println(names)
	// [one seven fish cat]
	log.Println(unsafe.Sizeof(names))
	// 24 bytes
}

func maps() {
	var myString string
	var myInt int

	myString = "Hi"
	myInt = 11

	mySecondString := "another string"

	log.Println(myString, mySecondString, myInt)

	myMap := make(map[string]string)

	myMap["dog"] = "Samson"

	log.Println(myMap)
	// map[dog:Samson]
	log.Println(myMap["dog"])
	// Samson

	myMap["other-dog"] = "Cassie"

	log.Println(myMap)
	// map[dog:Samson other-dog:Cassie]
	log.Println(myMap["other-dog"])
	// Cassie

	myMap["dog"] = "fido"

	log.Println(myMap)
	// map[dog:fido other-dog:Cassie]

	myMapInt := make(map[string]int)

	log.Println(unsafe.Sizeof(myMapInt))
	// 8 bytes

	myMapInt["First"] = 1
	log.Println(unsafe.Sizeof(myMapInt))
	// 8 bytes

	myMapInt["Second"] = 2
	log.Println(unsafe.Sizeof(myMapInt))
	// 8 bytes

	myMapUser := make(map[string]User)
	log.Println(unsafe.Sizeof(myMapUser))
	// 8 bytes

	me := User{
		FirstName: "Trevor",
		LastName:  "Sawler",
	}

	myMapUser["me"] = me

	log.Println(myMapUser["me"].FirstName)
	// Trevor

	log.Println(unsafe.Sizeof(myMapUser))

	var myNewVar float32

	myNewVar = 11.1

	log.Println(unsafe.Sizeof(myNewVar))
	// 4

	var myVarInt8 int8
	myVarInt8 = 10
	log.Println(unsafe.Sizeof(myVarInt8))
	// 1 byte

	var myVarInt16 int16
	myVarInt16 = 20
	log.Println(unsafe.Sizeof(myVarInt16))
	// 2 bytes

	var myVarInt32 int32
	myVarInt32 = 30
	log.Println(unsafe.Sizeof(myVarInt32))
	// 4 bytes

	var myVarInt64 int64
	myVarInt16 = 40
	log.Println(unsafe.Sizeof(myVarInt64))
	// 8 bytes
}
