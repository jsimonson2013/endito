package page

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func GetPages(base string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		files, err := readDir(base, nil)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, err.Error())
			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, strings.Join(files, ","))
	}
}

func readDir(dir string, files []string) ([]string, error) {
	rdir, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	if len(rdir) == 1 && strings.Contains(rdir[0].Name(), ".html") {
		return append(files, dir+rdir[0].Name()), nil
	}

	for _, f := range rdir {
		if f.IsDir() {
			files, err = readDir(dir+f.Name()+"/", files)
			if err != nil {
				return nil, err
			}
		} else if strings.Contains(f.Name(), ".html") {
			files = append(files, dir+f.Name())
		}
	}

	return files, nil
}
