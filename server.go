package main

import (
	"endito/compiler"
	"endito/document"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/gorilla/mux"
)

func main() {
	d := document.FromDir("./example")
	d.Print()
	c := compiler.NewPiler(d)
	c.Compile()

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	r.Post("/page", UpdatePage())

	r.Get("/page/{uri}", LoadPage())

	// r.HandleFunc("/page", UpdatePage).Methods("POST")
	// r.HandleFunc("/{id}", UpdatePage).Methods("GET")

	http.ListenAndServe(":3333", r)
}

func LoadPage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		uri := chi.URLParam(r, "uri")

		f, err := os.OpenFile(uri, os.O_RDONLY, os.ModePerm)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, err.Error())
			return
		}

		content, err := ioutil.ReadFile(f.Name())
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, err.Error())
			return
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, string(content))
	}
}

func UpdatePage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		fmt.Println(string(b))
		vars := mux.Vars(r)
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "ID: %v\n", vars["id"])
	}
}
