package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-git/go-git/v5"
)

// BASE is path to start searching for html
const BASE = "./"

func main() {
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

	r.Post("/update", UpdatePage())

	r.Post("/load", LoadPage())

	r.Get("/pages", GetPages())

	testGit("integrate go-git", []string{"./"})

	http.ListenAndServe(":3333", r)
}

func testGit(msg string, files []string) {
	repo, err := git.PlainOpen("./")
	if err != nil {
		return
	}

	tree, err := repo.Worktree()
	if err != nil {
		return
	}

	for _, file := range files {
		if _, err := tree.Add(file); err != nil {
			return
		}
	}

	status, err := tree.Status()
	if err != nil {
		return
	}

	fmt.Println(status)

	// commit, err := tree.Commit(msg, &git.CommitOptions{
	// 	Author: &object.Signature{
	// 		Name:  os.Getenv("GIT_UNAME"),
	// 		Email: os.Getenv("GIT_EMAIL"),
	// 		When:  time.Now(),
	// 	},
	// })
	// if err != nil {
	// 	return
	// }
	// obj, err := repo.CommitObject(commit)
	// if err != nil {
	// 	return
	// }

	// fmt.Println(obj)
}

func LoadPage() http.HandlerFunc {
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

func GetPages() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		files, err := readDir(BASE, nil)
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

func UpdatePage() http.HandlerFunc {
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

		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "wrote %d bytes", len([]byte(body["content"].(string))))
	}
}
