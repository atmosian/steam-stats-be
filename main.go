package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	// Routes consist of a path and a handler function.
	router.HandleFunc("/", HelloHandler)
	
	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8000", router))
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Atmosian!\n"))
}
