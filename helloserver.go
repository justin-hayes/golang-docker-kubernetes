package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/pat"
	"github.com/urfave/negroni"
)

func main() {
	r := pat.New()
	r.Get("/", greetHandler)

	n := negroni.Classic()
	n.UseHandler(r)

	n.Run(":8080")
}

func greetHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello world!")
}
