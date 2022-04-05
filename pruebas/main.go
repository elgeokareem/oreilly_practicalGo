package main

import (
	"fmt"
	"net/http"
	"os"
)

func index(w http.ResponseWriter, req *http.Request) {
	method := req.Method
	fmt.Fprintf(w, "Hola, soy el controlador y me estoy ejecutando con el método %s", method)
}

func kek(w http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	fmt.Fprintf(w, "%s", query.Get("name"))
}

func setupHandlers(mux *http.ServeMux) {
	mux.HandleFunc("/", index)
	mux.HandleFunc("/kek", kek)
}

func main() {
	listenAddr := os.Getenv("LISTEN_ADDR")

	if len(listenAddr) == 0 {
		listenAddr = ":8080"
	}

	mux := http.NewServeMux()

	setupHandlers(mux)
	http.ListenAndServe(listenAddr, mux)
}
