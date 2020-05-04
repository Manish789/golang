package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	// router.HandleFunc tells the
	router.HandleFunc("/", HomeHandler).Methods("GET")
	router.HandleFunc("/gopher", GopherHandler).Methods("GET")
	router.HandleFunc("/quotes", QuotesHandler).Methods("GET")

	http.Handle("/", router)
	// http.ListenAndServe tells the server to listen on the TCP network address :3000.
	http.ListenAndServe(":3000", nil)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!")
}

func GopherHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello you gopher!!")
}

func QuotesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "If you cannot do great things, do small things in a great way.")
}
