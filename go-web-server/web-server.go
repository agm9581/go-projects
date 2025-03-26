package main

import (
	"fmt"
	"log"
	"net/http"
)

func noFileHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/nofile" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Forbidden method", http.StatusForbidden)
		return
	}

	fmt.Fprintf(w, "nofile route accessed")

}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() error: %v", err)
	}
	fmt.Fprintf(w, "POST request successful")
	firstInput := r.FormValue("first")
	second := r.FormValue("second")
	fmt.Fprintf(w, "First = %s", firstInput)
	fmt.Fprintf(w, "Second = %s", second)
}

func main() {
	port := 81
	fmt.Printf("Server starting in port %v \n", port)
	fileServer := http.FileServer(http.Dir("./static-files"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/nofile", noFileHandler)

	if err := http.ListenAndServe(fmt.Sprintf(":%v", port), nil); err != nil {
		log.Fatal(fmt.Sprintf("Issue occurred starting the server in port %v", port))
	}

}
