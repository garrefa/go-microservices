package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {
		log.Println("Hello world!")
		data, err := io.ReadAll(req.Body)
		if err != nil {
			http.Error(rw, "Ooooops", http.StatusBadRequest)
			return
		}
		fmt.Fprintf(rw, "Hello %s", data)
	})

	http.HandleFunc("/goodbye", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Goodbye!")
	})

	http.ListenAndServe("127.0.0.1:8080", nil)
}
