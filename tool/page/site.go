package page

import (
	"endito/files"
	"fmt"
	"net/http"
	"strings"
)

func GetPages(base string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		files, err := files.ReadDir(base, nil)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, err.Error())
			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, strings.Join(files, ","))
	}
}
