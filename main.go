package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

const (
	filePath = "/mnt/secrets-store/DemoSecret"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", ExampleHandler)

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

func ExampleHandler(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Printf("Failed to read the file from the path: %s", filePath)
		w.Write([]byte("Error reading file"))
		return
	}
	w.Write(append([]byte("Hello World! Here's my secret: "), data...))
}
