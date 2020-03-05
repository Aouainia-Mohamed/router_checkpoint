package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func sarra(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("ahla ya a9wa instructrice fik ya GOMYCODE"))
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/sarra", sarra).Methods("GET")
	log.Fatal(http.ListenAndServe(":12345", router))
}
