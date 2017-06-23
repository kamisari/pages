package main

import (
	"log"
	"net/http"
)

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		log.Println("req.URL: ", r.RequestURI)
		log.Println("form[key]: ", r.FormValue("key"))
		log.Println("form[key2]: ", r.FormValue("key2"))
		if s := r.FormValue("text"); s != "" {
			log.Printf("from [text]:%s\n", s)
		}
	}
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
