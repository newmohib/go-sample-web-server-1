package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method is  not Supported", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "hello!")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/form" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "POST" {
		http.Error(w, "Method is  not Supported", http.StatusNotFound)
		return
	}
	if err := r.ParseForm(); err != nil {
		// http.Error(w, "Method is  not Supported", http.StatusNotFound)
		fmt.Fprintf(w, "ParseForm() erro %w", err)
		return
	}
	fmt.Fprintf(w, "POST Request Successful \n")
	name := r.FormValue("name")
	address := r.FormValue("address")

	fmt.Fprintf(w, "Name : %s\n", name)
	fmt.Fprintf(w, "Address : %s\n", address)

}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	// http.Handle("/login", http.FileServer(http.Dir("./static")+"/form.html"))
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Starting server at port 8080 \n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
