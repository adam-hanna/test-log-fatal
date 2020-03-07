package main

import (
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	// Hello world, the web server

	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, world!\n")
	}

	http.HandleFunc("/hello", helloHandler)
	go func() {
		log.Println("in go func")
		time.Sleep(1 * time.Second)
		log.Fatal("oopsie")
	}()

	log.Println("listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
