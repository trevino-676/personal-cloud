package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/trevino-676/personal-cloud/handlers"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Gorilla!\n"))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", RootHandler)
	r.HandleFunc("/file", handlers.UploadFile).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", r))
}
