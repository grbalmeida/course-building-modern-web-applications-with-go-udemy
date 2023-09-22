package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		numberOfBytes, err := fmt.Fprintf(w, "Hello, world!")

		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(fmt.Printf("Number of bytes written: %d", numberOfBytes))
	})

	_ = http.ListenAndServe(":8080", nil)

}
