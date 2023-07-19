package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// New Instance for Router
	r := mux.NewRouter()

	//Url patterns
	r.HandleFunc("/", homeHandler).Methods("GET")
	r.HandleFunc("/hello", helloHandler).Methods("GET")
	r.HandleFunc("/goodbye", goodByeHandler).Methods("GET")

	// Registers default request handler for the url patterns
	http.Handle("/", r)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err) // This will cause our server to crash with an error message if it
	}

}

// Implementation for each url patterns
func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "This is home Page!")
}
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!")
}
func goodByeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Good Bye World!")
}
