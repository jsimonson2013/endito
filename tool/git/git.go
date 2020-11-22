package git

import (
	"os"
	"strings"
	"time"

	gogit "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
)

func Commit(base, msg string, files []string) error {
	repo, err := gogit.PlainOpen(base)
	if err != nil {
		return err
	}

	tree, err := repo.Worktree()
	if err != nil {
		return err
	}

	for _, file := range files {
		file = strings.Trim(file, "./")
		file = strings.Trim(file, "../")
		if _, err := tree.Add(file); err != nil {
			return err
		}
	}

	if _, err := tree.Status(); err != nil {
		return err
	}

	commit, err := tree.Commit(msg, &gogit.CommitOptions{
		Author: &object.Signature{
			Name:  os.Getenv("GIT_UNAME"),
			Email: os.Getenv("GIT_EMAIL"),
			When:  time.Now(),
		},
	})
	if err != nil {
		return err
	}

	if _, err := repo.CommitObject(commit); err != nil {
		return err
	}

	opts := &gogit.PushOptions{
		Auth: &http.BasicAuth{
			Username: os.Getenv("GITHUB_TOKEN"),
			Password: os.Getenv("GITHUB_PASS"),
		},
	}
	if err := repo.Push(opts); err != nil {
		return err
	}

	return nil
}
