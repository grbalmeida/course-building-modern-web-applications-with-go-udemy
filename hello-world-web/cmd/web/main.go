package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/grbalmeida/building-modern-web-applications-with-go/pkg/config"
	"github.com/grbalmeida/building-modern-web-applications-with-go/pkg/handlers"
	"github.com/grbalmeida/building-modern-web-applications-with-go/pkg/render"
)

const portNumber = ":8080"

// main is the main application function
func main() {

	var app config.AppConfig

	tc, err := render.CreateComplexTemplateCache()

	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)

}
