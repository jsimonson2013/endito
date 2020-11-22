package page

import (
	"encoding/json"
	"endito/files"
	"endito/git"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func Load() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		bs, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, err.Error())
			return
		}
		defer r.Body.Close()

		var body map[string]interface{}
		if err := json.Unmarshal(bs, &body); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, err.Error())
			return
		}

		f, err := os.OpenFile(body["uri"].(string), os.O_RDONLY, os.ModePerm)
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

func Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		bs, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, err.Error())
			return
		}
		defer r.Body.Close()

		var body map[string]interface{}
		if err := json.Unmarshal(bs, &body); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, err.Error())
			return
		}

		if body["uname"] != os.Getenv("USERNAME") || body["pword"] != os.Getenv("PASSWORD") {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "bad username or password")
			return
		}

		if _, err = ioutil.ReadFile(body["uri"].(string)); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, err.Error())
			return
		}

		if err := ioutil.WriteFile(body["uri"].(string), []byte(body["content"].(string)), os.ModePerm); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, err.Error())
			return
		}

		rltv, err := files.GetRelativePaths(os.Getenv("RLTV_DIR"), []string{body["uri"].(string)})
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, err.Error())
			return
		}

		if err := git.Commit(os.Getenv("RLTV_DIR"), fmt.Sprintf("%s updated %s", os.Getenv("GIT_UNAME"), strings.Join(rltv, ",")), rltv); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, err.Error())
			return
		}

		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "wrote %d bytes", len([]byte(body["content"].(string))))
	}
}
