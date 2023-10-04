package main

import (
	"fmt"
	"log"
	"net/http"
)

func handleHello(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/hello" {
		http.Error(res, "404 not Found", http.StatusNotFound)
	}

	if req.Method != "GET" {
		http.Error(res, "403 Forbidden", http.StatusForbidden)
	}
	fmt.Fprintf(res, "Hello")
}

func handleForm(res http.ResponseWriter, req *http.Request) {
	if err := req.ParseForm(); err != nil {
		http.Error(res, "error", http.StatusBadRequest)
	}
	fmt.Fprintf(res, "POST request successful\n")

	name := req.FormValue("name")
	address := req.FormValue("address")
	fmt.Fprintf(res, "Name: %s\n", name)
	fmt.Fprintf(res, "Address: %s\n", address)
}

func main() {
	fileServer := http.FileServer(http.Dir("./views"))

	http.Handle("/", fileServer)
	http.HandleFunc("/form", handleForm)
	http.HandleFunc("/hello", handleHello)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
