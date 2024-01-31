package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Hello from Service B")
		w.Write([]byte("Hello from Service B"))
	})

	log.Fatal(http.ListenAndServe(":8081", nil))
}
