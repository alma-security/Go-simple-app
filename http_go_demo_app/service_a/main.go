package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("handle func")
		http.Get("http://service-b.goapp.svc.cluster.local:80")
		log.Println("resp")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
