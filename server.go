package main

import (
	"log"
	"net/http"
)

func logging(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RemoteAddr, r.RequestURI)
		handler.ServeHTTP(w, r)
	})
}

func main() {
	http.Handle("/", logging(http.FileServer(http.Dir("./docs"))))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
