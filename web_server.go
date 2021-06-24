// Server1 is a minimal "echo" server.
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler) // each request calls handler
	http.HandleFunc("/hello", hello) // each request calls handler
	webServer := "localhost:8000"
	fmt.Printf("http listen on %s\n", webServer)
	log.Fatal(http.ListenAndServe(webServer, nil))
}

// handler echoes the Path component of the request URL r.
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<H1>HELLO WORLD</H1>")
}
