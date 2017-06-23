package main

import (
	"log"
	"net/http"
)

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		log.Printf("req.URL: %s\n", r.RequestURI)
		for _, s := range []string{"key1", "key2", "text"} {
			if content := r.FormValue(s); content != "" {
				log.Printf("form[%s]: %s\n", s, content)
			}
		}
	}
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
