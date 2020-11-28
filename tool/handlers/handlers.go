package handlers

import (
	"endito/page"
	"os"
	"strings"

	"fmt"
	"io/ioutil"
	"net/http"
)

// PageUpdate handles requests for updating a page
func PageUpdate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		bs, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, err.Error())
			return
		}
		defer r.Body.Close()

		if err := page.Update(bs); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, err.Error())
			return
		}

		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
	}
}

// PageLoad handles requests for loading a page
func PageLoad() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		bs, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, err.Error())
			return
		}
		defer r.Body.Close()

		content, err := page.Load(bs)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, err.Error())
			return
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, string(content))
	}
}

// GetPages handles requests for retrieving all pages
func GetPages() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fs, err := page.GetAll(os.Getenv("BASE_DIR"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, err.Error())
			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, strings.Join(fs, ","))
	}
}
