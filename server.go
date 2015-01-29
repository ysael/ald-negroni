package main

import (
	//"fmt"
	"github.com/codegangsta/negroni"
	"log"
	//"github.com/gorilla/mux"
	"net/http"
)

func main() {
	n := negroni.New(
		negroni.NewRecovery(),
		negroni.HandlerFunc(MyMiddleware),
		negroni.HandlerFunc(MyHiMiddleware),
		negroni.NewLogger(),
		negroni.NewStatic(http.Dir("assets/")),
	)

	n.Run(":8080")

}

func MyMiddleware(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	log.Println("Logging on the wat there")

	if r.URL.Query().Get("password") == "jesus" {
		log.Println("yess")
		next(rw, r)
	} else {
		http.Error(rw, "Not Authorise", 401)
	}

	log.Println("logging on the way back")
}

func MyHiMiddleware(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	log.Println("HI")
}
