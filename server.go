package main

import (
	"fmt"
	//"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter().StrictSlash(false)
	r.HandleFunc("/", HomeHandler)

	fmt.Println("ca roule sur 8080")

	http.ListenAndServe(":8080", r)
}

func HomeHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("yooo")
	fmt.Fprintf(w, "we sdajkhdhave a page")
}
