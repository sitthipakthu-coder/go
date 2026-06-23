package httpserver

import (
	"fmt"
	"log"
	"net/http"
)

func hellohttp(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hellohttp" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supportted", http.StatusNotFound)
		return
	}

	fmt.Fprintln(w, "Hello https!")
}

func ShowHelloHTTP() {
	http.HandleFunc("/hellohttp", hellohttp)

	fmt.Println("Starting server at port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
