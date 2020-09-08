package main

import (
	"endito/compiler"
	"endito/document"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	d := document.FromDir("./example")
	d.Print()
	c := compiler.NewPiler(d)
	c.Compile()

	r := mux.NewRouter()
	r.HandleFunc("/{id}", HomeHandler).Methods("GET")

	http.ListenAndServe(":3333", r)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "ID: %v\n", vars["id"])
}
