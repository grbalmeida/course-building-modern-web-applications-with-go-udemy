package render

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"text/template"
)

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	// get the template cache from the app config
	tc, err := CreateComplexTemplateCache()

	if err != nil {
		log.Fatal(err)
	}
	// get requested template from cache
	t, ok := tc[tmpl]

	if !ok {
		log.Fatal(err)
	}

	buf := new(bytes.Buffer)

	_ = t.Execute(buf, nil)

	// render the template
	_, err = buf.WriteTo(w)

	if err != nil {
		log.Println("Error writing template to browser", err)
	}
}

func CreateComplexTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	// get all of the files named *.page.tmpl from ./templates
	pages, err := filepath.Glob("./templates/*.page.tmpl.html")

	if err != nil {
		return myCache, err
	}

	// range through all files ending with *.page.tmpl
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)

		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl.html")

		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl.html")
		}

		myCache[name] = ts
	}

	return myCache, nil
}

var tc = make(map[string]*template.Template)

func RenderSimpleTemplate(w http.ResponseWriter, t string) {
	var tmpl *template.Template
	var err error

	// check to see if we already have the template in our cache
	_, inMap := tc[t]

	if !inMap {
		// need to create the template
		log.Println("creating template and adding to cache")
		err = createSimpleTemplateCache(t)

		if err != nil {
			log.Println(err)
		}
	} else {
		// we have the template in the cache
		log.Println("using cached template")
	}

	tmpl = tc[t]

	err = tmpl.Execute(w, nil)

	if err != nil {
		log.Println(err)
	}
}

func createSimpleTemplateCache(t string) error {
	templates := []string{
		fmt.Sprintf("./templates/%s.html", t),
		"./templates/base.layout.tmpl.html",
	}

	// parse the template
	tmpl, err := template.ParseFiles(templates...)

	if err != nil {
		return err
	}

	// add template to cache (map)
	tc[t] = tmpl

	return nil
}
