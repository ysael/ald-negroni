package main

import (
	"fmt"
	"github.com/codegangsta/negroni"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "we have a page")

	})

	n := negroni.Classic()
	n.UseHandler(mux)
	n.Run(":3001")
}
