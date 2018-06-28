package main

import (
	"net/http"
	"log"
	"encoding/json"
	"fmt"
)

type test_struct struct {
	Test string
}

func handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
	//POST HANDLE
		var T test_struct
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&T)
		if err != nil {
			log.Println("error")
			log.Print(err)
		}
		log.Println(T.Test)
	default:
		fmt.Fprintf(w, "Ho hello there")
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":80", nil)
}
