package main

import (
	"encoding/json"
//	"fmt"
	"net/http"
)

type Greeting struct {
	Hi 		string
	What	string
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		greeting := Greeting {
			Hi: "Hello, ",
			What:  "World!",
		}
		json.NewEncoder(w).Encode(greeting);
	})

	http.ListenAndServe(":8080", nil)
}
