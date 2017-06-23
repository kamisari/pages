package main

import (
	"log"
	"net/http"
)

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		log.Println("req.URL: ", r.RequestURI)
		log.Println("form[key]: ", r.FormValue("key"))
		log.Printf("form[key2]: %s\n\n", r.FormValue("key2"))
	}
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
