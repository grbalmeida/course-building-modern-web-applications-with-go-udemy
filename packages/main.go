package main

import (
	"log"

	"github.com/grbalmeida/building-modern-web-applications-with-go/packages/helpers"
)

func main() {
	var myVar helpers.SomeType
	myVar.TypeName = "Some name"
	log.Println(myVar.TypeName)
}
